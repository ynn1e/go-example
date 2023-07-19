[Handler]: https://pkg.go.dev/net/http#Handler

# HTTP server

- ServeMux
  - ServeMux.Handle
  - ServeMux.Handler
- HandleFunc
- Handler

## ServeMux

> ServeMux is an HTTP request multiplexer.

様々な URL パターンから最もマッチする1つの handler(Entry) へルーティングするという理解。
その役割は `m` が担っている。

```go
type ServeMux struct {
  mu    sync.RWMutex
  m     map[string]muxEntry
  es    []muxEntry
  hosts bool
}
```

```go
var DefaultServeMux = &defaultServeMux

var defaultServeMux ServeMux
```

## ServeMux.Handle

パターンにエントリを関連づける。
上記ルーティングの理解の箇所。

```go
func (mux *ServeMux) Handle(pattern string, handler Handler) {
  // 省略
  e := muxEntry{h: handler, pattern: pattern}
  mux.m[pattern] = e
  // 省略
}
```

```go
type muxEntry struct {
  h       Handler
  pattern string
}
```

## ServeMux.Handler

```go
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string) {
  // 省略
  return mux.handler(host, r.URL.Path)
}
```

最もマッチする(より固有な)パターンを返す。

```go
func (mux *ServeMux) handler(host, path string) (h Handler, pattern string) {
  mux.mu.RLock()
  defer mux.mu.RUnlock()

  // Host-specific pattern takes precedence over generic ones
  if mux.hosts {
    h, pattern = mux.match(host + path)
  }
  if h == nil {
    h, pattern = mux.match(path)
  }
  if h == nil {
    h, pattern = NotFoundHandler(), ""
  }
  return
}
```

> Find a handler on a handler map given a path string.
> Most-specific (longest) pattern wins.

```go
func (mux *ServeMux) match(path string) (h Handler, pattern string) {
  // Check for exact match first.
  v, ok := mux.m[path]
  if ok {
    return v.h, v.pattern
  }

  // Check for longest valid match.  mux.es contains all patterns
  // that end in / sorted from longest to shortest.
  for _, e := range mux.es {
    if strings.HasPrefix(path, e.pattern) {
      return e.h, e.pattern
    }
  }
  return nil, ""
}
```

## HandleFunc

```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
  DefaultServeMux.HandleFunc(pattern, handler)
}
```

## [Handler]

`ServeHTTP(ResponseWriter, *Request)` を実装すれば `Handler` interface を満たす。

```go
type Handler interface {
  ServeHTTP(ResponseWriter, *Request)
}
```
