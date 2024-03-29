package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

type interfaceInfoList []interfaceInfo

func (l *interfaceInfoList) append(file, text string) {
	for i, ii := range *l {
		if ii.file == file {
			l := []interfaceInfo(*l)
			l[i].texts = append(l[i].texts, text)
			return
		}
	}
	*l = append(*l, interfaceInfo{
		file:  file,
		texts: []string{text},
	})
}

type interfaceInfo struct {
	file  string
	texts []string
}

func main() {
	files := []string{}
	src := os.Getenv("GO_SRC")
	if err := filepath.WalkDir(src, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			if strings.HasSuffix(entry.Name(), "cmd") {
				return fs.SkipDir
			}

			if strings.HasSuffix(entry.Name(), "internal") {
				return fs.SkipDir
			}

			if strings.HasSuffix(entry.Name(), "testdata") {
				return fs.SkipDir
			}

			if strings.HasSuffix(entry.Name(), "vendor") {
				return fs.SkipDir
			}

			return nil
		}

		if strings.HasSuffix(entry.Name(), ".go") && !strings.HasSuffix(entry.Name(), "_test.go") {
			files = append(files, path)
		}

		return nil
	}); err != nil {
		log.Fatalln(err)
	}

	iInfo := make(interfaceInfoList, 0)

	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			log.Fatalln(err)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)
		name := strings.ReplaceAll(file.Name(), src+"/", "")
		var text []string
		for scanner.Scan() {
			line := scanner.Text()
			if regexp.MustCompile("^type [A-Z].+interface.+{.+}$").MatchString(line) { // it means interface written by one-line.
				iInfo.append(name, line)
				text = []string{}
				continue
			}

			if regexp.MustCompile("^type [A-Z].+ interface {").MatchString(line) { // it means interface written by multi-lines.
				text = append(text, line)
				continue
			}

			if len(text) > 0 {
				if strings.Contains(line, "//") {
					continue
				}

				text = append(text, line)
				if strings.HasPrefix(line, "}") && strings.HasSuffix(line, "}") { // it means the line is "}"
					iInfo.append(name, strings.Join(text, "\n"))
					text = []string{}
				}
			}
		}
	}

	var builder strings.Builder

	version := runtime.Version()
	if _, err := builder.WriteString(fmt.Sprintf("# Interfaces\n\ngo version: %s\n\nYou can create README.md\n\n```sh\nGO_SRC={GO SOURCE} go run main.go\n```\n", version)); err != nil {
		log.Println(err)
	}

	for _, ii := range iInfo {
		link := fmt.Sprintf("https://cs.opensource.google/go/go/+/refs/tags/%s:src/%s", version, ii.file)
		_, err := builder.WriteString(fmt.Sprintf("\n## [%s](%s)\n", ii.file, link))
		if err != nil {
			log.Fatalln(err)
		}
		for _, text := range ii.texts {
			v := strings.ReplaceAll(text, "\t", "  ")
			_, err := builder.WriteString(fmt.Sprintf("\n```go\n%s\n```\n", v))
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	file, err := os.Create("README.md")
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	if _, err = file.WriteString(builder.String()); err != nil {
		log.Fatalln(err)
	}
}
