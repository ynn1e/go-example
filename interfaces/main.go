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
	for _, ii := range *l {
		if ii.file == file {
			ii.texts = append(ii.texts, text)
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
	if err := filepath.WalkDir(os.Getenv("SRC"), func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			if strings.HasSuffix(entry.Name(), "cmd") {
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
		name := filepath.Base(file.Name())
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
				text = append(text, line)
				if strings.HasPrefix(line, "}") && strings.HasSuffix(line, "}") { // it means the line is "}"
					iInfo.append(name, strings.Join(text, "\n"))
					text = []string{}
				}
			}
		}
	}

	var builder strings.Builder
	if _, err := builder.WriteString(fmt.Sprintf("# Interfaces\n\ngo version: %s\n\nYou can create README.md\n\n```sh\nSRC={your golang source file path} go run main.go\n```\n", runtime.Version())); err != nil {
		log.Println(err)
	}

	for _, ii := range iInfo {
		_, err := builder.WriteString(fmt.Sprintf("\n## %s\n", ii.file))
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
