# Interfaces

go version: go1.20.5

You can create README.md

```sh
SRC={your golang source file path} go run main.go
```

## [compress/flate/inflate.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/compress/flate/inflate.go)

```go
type Resetter interface {
  Reset(r io.Reader, dict []byte) error
}
```

```go
type Reader interface {
  io.Reader
  io.ByteReader
}
```

## [compress/zlib/reader.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/compress/zlib/reader.go)

```go
type Resetter interface {
  Reset(r io.Reader, dict []byte) error
}
```

## [container/heap/heap.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/container/heap/heap.go)

```go
type Interface interface {
  sort.Interface
}
```

## [context/context.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/context/context.go)

```go
type Context interface {
  Deadline() (deadline time.Time, ok bool)

  Done() <-chan struct{}

  Err() error

  Value(key any) any
}
```

## [crypto/cipher/cipher.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/crypto/cipher/cipher.go)

```go
type Block interface {
  BlockSize() int

  Encrypt(dst, src []byte)

  Decrypt(dst, src []byte)
}
```

```go
type Stream interface {
  XORKeyStream(dst, src []byte)
}
```

```go
type BlockMode interface {
  BlockSize() int

  CryptBlocks(dst, src []byte)
}
```

## [crypto/cipher/gcm.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/crypto/cipher/gcm.go)

```go
type AEAD interface {
  NonceSize() int

  Overhead() int

  Seal(dst, nonce, plaintext, additionalData []byte) []byte

  Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error)
}
```

## [crypto/crypto.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/crypto/crypto.go)

```go
type Signer interface {
  Public() PublicKey

  Sign(rand io.Reader, digest []byte, opts SignerOpts) (signature []byte, err error)
}
```

```go
type SignerOpts interface {
  HashFunc() Hash
}
```

```go
type Decrypter interface {
  Public() PublicKey

  Decrypt(rand io.Reader, msg []byte, opts DecrypterOpts) (plaintext []byte, err error)
}
```

## [crypto/ecdh/ecdh.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/crypto/ecdh/ecdh.go)

```go
type Curve interface {
  GenerateKey(rand io.Reader) (*PrivateKey, error)

  NewPrivateKey(key []byte) (*PrivateKey, error)

  NewPublicKey(key []byte) (*PublicKey, error)

  ecdh(local *PrivateKey, remote *PublicKey) ([]byte, error)

  privateKeyToPublicKey(*PrivateKey) *PublicKey
}
```

## [crypto/elliptic/elliptic.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/crypto/elliptic/elliptic.go)

```go
type Curve interface {
  Params() *CurveParams

  IsOnCurve(x, y *big.Int) bool

  Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int)

  Double(x1, y1 *big.Int) (x, y *big.Int)

  ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int)

  ScalarBaseMult(k []byte) (x, y *big.Int)
}
```

## [crypto/tls/common.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/crypto/tls/common.go)

```go
type ClientSessionCache interface {
  Get(sessionKey string) (session *ClientSessionState, ok bool)

  Put(sessionKey string, cs *ClientSessionState)
}
```

## [database/sql/driver/driver.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/database/sql/driver/driver.go)

```go
type Driver interface {
  Open(name string) (Conn, error)
}
```

```go
type DriverContext interface {
  OpenConnector(name string) (Connector, error)
}
```

```go
type Connector interface {
  Connect(context.Context) (Conn, error)

  Driver() Driver
}
```

```go
type Pinger interface {
  Ping(ctx context.Context) error
}
```

```go
type Execer interface {
  Exec(query string, args []Value) (Result, error)
}
```

```go
type ExecerContext interface {
  ExecContext(ctx context.Context, query string, args []NamedValue) (Result, error)
}
```

```go
type Queryer interface {
  Query(query string, args []Value) (Rows, error)
}
```

```go
type QueryerContext interface {
  QueryContext(ctx context.Context, query string, args []NamedValue) (Rows, error)
}
```

```go
type Conn interface {
  Prepare(query string) (Stmt, error)

  Close() error

  Begin() (Tx, error)
}
```

```go
type ConnPrepareContext interface {
  PrepareContext(ctx context.Context, query string) (Stmt, error)
}
```

```go
type ConnBeginTx interface {
  BeginTx(ctx context.Context, opts TxOptions) (Tx, error)
}
```

```go
type SessionResetter interface {
  ResetSession(ctx context.Context) error
}
```

```go
type Validator interface {
  IsValid() bool
}
```

```go
type Result interface {
  LastInsertId() (int64, error)

  RowsAffected() (int64, error)
}
```

```go
type Stmt interface {
  Close() error

  NumInput() int

  Exec(args []Value) (Result, error)

  Query(args []Value) (Rows, error)
}
```

```go
type StmtExecContext interface {
  ExecContext(ctx context.Context, args []NamedValue) (Result, error)
}
```

```go
type StmtQueryContext interface {
  QueryContext(ctx context.Context, args []NamedValue) (Rows, error)
}
```

```go
type NamedValueChecker interface {
  CheckNamedValue(*NamedValue) error
}
```

```go
type ColumnConverter interface {
  ColumnConverter(idx int) ValueConverter
}
```

```go
type Rows interface {
  Columns() []string

  Close() error

  Next(dest []Value) error
}
```

```go
type RowsNextResultSet interface {
  Rows

  HasNextResultSet() bool

  NextResultSet() error
}
```

```go
type RowsColumnTypeScanType interface {
  Rows
  ColumnTypeScanType(index int) reflect.Type
}
```

```go
type RowsColumnTypeDatabaseTypeName interface {
  Rows
  ColumnTypeDatabaseTypeName(index int) string
}
```

```go
type RowsColumnTypeLength interface {
  Rows
  ColumnTypeLength(index int) (length int64, ok bool)
}
```

```go
type RowsColumnTypeNullable interface {
  Rows
  ColumnTypeNullable(index int) (nullable, ok bool)
}
```

```go
type RowsColumnTypePrecisionScale interface {
  Rows
  ColumnTypePrecisionScale(index int) (precision, scale int64, ok bool)
}
```

```go
type Tx interface {
  Commit() error
  Rollback() error
}
```

## [database/sql/driver/types.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/database/sql/driver/types.go)

```go
type ValueConverter interface {
  ConvertValue(v any) (Value, error)
}
```

```go
type Valuer interface {
  Value() (Value, error)
}
```

## [database/sql/sql.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/database/sql/sql.go)

```go
type Scanner interface {
  Scan(src any) error
}
```

```go
type Result interface {
  LastInsertId() (int64, error)

  RowsAffected() (int64, error)
}
```

## [debug/dwarf/type.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/debug/dwarf/type.go)

```go
type Type interface {
  Common() *CommonType
  String() string
  Size() int64
}
```

## [debug/macho/file.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/debug/macho/file.go)

```go
type Load interface {
  Raw() []byte
}
```

## [encoding/binary/binary.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/encoding/binary/binary.go)

```go
type ByteOrder interface {
  Uint16([]byte) uint16
  Uint32([]byte) uint32
  Uint64([]byte) uint64
  PutUint16([]byte, uint16)
  PutUint32([]byte, uint32)
  PutUint64([]byte, uint64)
  String() string
}
```

```go
type AppendByteOrder interface {
  AppendUint16([]byte, uint16) []byte
  AppendUint32([]byte, uint32) []byte
  AppendUint64([]byte, uint64) []byte
  String() string
}
```

## [encoding/encoding.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/encoding/encoding.go)

```go
type BinaryMarshaler interface {
  MarshalBinary() (data []byte, err error)
}
```

```go
type BinaryUnmarshaler interface {
  UnmarshalBinary(data []byte) error
}
```

```go
type TextMarshaler interface {
  MarshalText() (text []byte, err error)
}
```

```go
type TextUnmarshaler interface {
  UnmarshalText(text []byte) error
}
```

## [encoding/gob/type.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/encoding/gob/type.go)

```go
type GobEncoder interface {
  GobEncode() ([]byte, error)
}
```

```go
type GobDecoder interface {
  GobDecode([]byte) error
}
```

## [encoding/json/decode.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/encoding/json/decode.go)

```go
type Unmarshaler interface {
  UnmarshalJSON([]byte) error
}
```

## [encoding/json/encode.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/encoding/json/encode.go)

```go
type Marshaler interface {
  MarshalJSON() ([]byte, error)
}
```

## [encoding/xml/marshal.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/encoding/xml/marshal.go)

```go
type Marshaler interface {
  MarshalXML(e *Encoder, start StartElement) error
}
```

```go
type MarshalerAttr interface {
  MarshalXMLAttr(name Name) (Attr, error)
}
```

## [encoding/xml/read.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/encoding/xml/read.go)

```go
type Unmarshaler interface {
  UnmarshalXML(d *Decoder, start StartElement) error
}
```

```go
type UnmarshalerAttr interface {
  UnmarshalXMLAttr(attr Attr) error
}
```

## [encoding/xml/xml.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/encoding/xml/xml.go)

```go
type TokenReader interface {
  Token() (Token, error)
}
```

## [expvar/expvar.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/expvar/expvar.go)

```go
type Var interface {
  String() string
}
```

## [flag/flag.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/flag/flag.go)

```go
type Value interface {
  String() string
  Set(string) error
}
```

```go
type Getter interface {
  Value
  Get() any
}
```

## [fmt/print.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/fmt/print.go)

```go
type State interface {
  Write(b []byte) (n int, err error)
  Width() (wid int, ok bool)
  Precision() (prec int, ok bool)

  Flag(c int) bool
}
```

```go
type Formatter interface {
  Format(f State, verb rune)
}
```

```go
type Stringer interface {
  String() string
}
```

```go
type GoStringer interface {
  GoString() string
}
```

## [fmt/scan.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/fmt/scan.go)

```go
type ScanState interface {
  ReadRune() (r rune, size int, err error)
  UnreadRune() error
  SkipSpace()
  Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
  Width() (wid int, ok bool)
  Read(buf []byte) (n int, err error)
}
```

```go
type Scanner interface {
  Scan(state ScanState, verb rune) error
}
```

## [go/ast/ast.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/go/ast/ast.go)

```go
type Node interface {
}
```

```go
type Expr interface {
  Node
  exprNode()
}
```

```go
type Stmt interface {
  Node
  stmtNode()
}
```

```go
type Decl interface {
  Node
  declNode()
}
```

## [go/ast/walk.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/go/ast/walk.go)

```go
type Visitor interface {
  Visit(node Node) (w Visitor)
}
```

## [go/build/constraint/expr.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/go/build/constraint/expr.go)

```go
type Expr interface {
  String() string

  Eval(ok func(tag string) bool) bool

  isExpr()
}
```

## [go/constant/value.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/go/constant/value.go)

```go
type Value interface {
  Kind() Kind

  String() string

  ExactString() string

  implementsValue()
}
```

## [go/doc/comment/parse.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/go/doc/comment/parse.go)

```go
type Block interface {
  block()
}
```

```go
type Text interface {
  text()
}
```

## [go/types/api.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/go/types/api.go)

```go
type Importer interface {
  Import(path string) (*Package, error)
}
```

```go
type ImporterFrom interface {
  Importer

  ImportFrom(path, dir string, mode ImportMode) (*Package, error)
}
```

## [go/types/object.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/go/types/object.go)

```go
type Object interface {

  String() string

  order() uint32

  color() color

  setType(Type)

  setOrder(uint32)

  setColor(color color)

  setParent(*Scope)

  sameId(pkg *Package, name string) bool

  scopePos() token.Pos

  setScopePos(pos token.Pos)
}
```

## [go/types/sizes.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/go/types/sizes.go)

```go
type Sizes interface {
  Alignof(T Type) int64

  Offsetsof(fields []*Var) []int64

  Sizeof(T Type) int64
}
```

## [go/types/type.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/go/types/type.go)

```go
type Type interface {
  Underlying() Type

  String() string
}
```

## [hash/hash.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/hash/hash.go)

```go
type Hash interface {
  io.Writer

  Sum(b []byte) []byte

  Reset()

  Size() int

  BlockSize() int
}
```

```go
type Hash32 interface {
  Hash
  Sum32() uint32
}
```

```go
type Hash64 interface {
  Hash
  Sum64() uint64
}
```

## [image/color/color.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/image/color/color.go)

```go
type Color interface {
  RGBA() (r, g, b, a uint32)
}
```

```go
type Model interface {
  Convert(c Color) Color
}
```

## [image/draw/draw.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/image/draw/draw.go)

```go
type Image interface {
  image.Image
  Set(x, y int, c color.Color)
}
```

```go
type RGBA64Image interface {
  image.RGBA64Image
  Set(x, y int, c color.Color)
  SetRGBA64(x, y int, c color.RGBA64)
}
```

```go
type Quantizer interface {
  Quantize(p color.Palette, m image.Image) color.Palette
}
```

```go
type Drawer interface {
  Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point)
}
```

## [image/image.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/image/image.go)

```go
type Image interface {
  ColorModel() color.Model
  Bounds() Rectangle
  At(x, y int) color.Color
}
```

```go
type RGBA64Image interface {
  RGBA64At(x, y int) color.RGBA64
  Image
}
```

```go
type PalettedImage interface {
  ColorIndexAt(x, y int) uint8
  Image
}
```

## [image/jpeg/reader.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/image/jpeg/reader.go)

```go
type Reader interface {
  io.ByteReader
  io.Reader
}
```

## [image/png/writer.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/image/png/writer.go)

```go
type EncoderBufferPool interface {
  Get() *EncoderBuffer
  Put(*EncoderBuffer)
}
```

## [internal/coverage/encodecounter/encode.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/internal/coverage/encodecounter/encode.go)

```go
type CounterVisitor interface {
  NumFuncs() (int, error)
  VisitFuncs(f CounterVisitorFn) error
}
```

## [internal/pkgbits/codes.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/internal/pkgbits/codes.go)

```go
type Code interface {
  Marker() SyncMarker

  Value() int
}
```

## [internal/reflectlite/type.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/internal/reflectlite/type.go)

```go
type Type interface {

  Name() string

  PkgPath() string

  Size() uintptr

  Kind() Kind

  Implements(u Type) bool

  AssignableTo(u Type) bool

  Comparable() bool

  String() string

  Elem() Type

  common() *rtype
  uncommon() *uncommonType
}
```

## [internal/testlog/log.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/internal/testlog/log.go)

```go
type Interface interface {
  Getenv(key string)
  Stat(file string)
  Open(file string)
  Chdir(dir string)
}
```

## [io/fs/fs.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/io/fs/fs.go)

```go
type FS interface {
  Open(name string) (File, error)
}
```

```go
type File interface {
  Stat() (FileInfo, error)
  Read([]byte) (int, error)
  Close() error
}
```

```go
type DirEntry interface {
  Name() string

  IsDir() bool

  Type() FileMode

  Info() (FileInfo, error)
}
```

```go
type ReadDirFile interface {
  File

  ReadDir(n int) ([]DirEntry, error)
}
```

```go
type FileInfo interface {
}
```

## [io/fs/glob.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/io/fs/glob.go)

```go
type GlobFS interface {
  FS

  Glob(pattern string) ([]string, error)
}
```

## [io/fs/readdir.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/io/fs/readdir.go)

```go
type ReadDirFS interface {
  FS

  ReadDir(name string) ([]DirEntry, error)
}
```

## [io/fs/readfile.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/io/fs/readfile.go)

```go
type ReadFileFS interface {
  FS

  ReadFile(name string) ([]byte, error)
}
```

## [io/fs/stat.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/io/fs/stat.go)

```go
type StatFS interface {
  FS

  Stat(name string) (FileInfo, error)
}
```

## [io/fs/sub.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/io/fs/sub.go)

```go
type SubFS interface {
  FS

  Sub(dir string) (FS, error)
}
```

## [io/io.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/io/io.go)

```go
type Reader interface {
  Read(p []byte) (n int, err error)
}
```

```go
type Writer interface {
  Write(p []byte) (n int, err error)
}
```

```go
type Closer interface {
  Close() error
}
```

```go
type Seeker interface {
  Seek(offset int64, whence int) (int64, error)
}
```

```go
type ReadWriter interface {
  Reader
  Writer
}
```

```go
type ReadCloser interface {
  Reader
  Closer
}
```

```go
type WriteCloser interface {
  Writer
  Closer
}
```

```go
type ReadWriteCloser interface {
  Reader
  Writer
  Closer
}
```

```go
type ReadSeeker interface {
  Reader
  Seeker
}
```

```go
type ReadSeekCloser interface {
  Reader
  Seeker
  Closer
}
```

```go
type WriteSeeker interface {
  Writer
  Seeker
}
```

```go
type ReadWriteSeeker interface {
  Reader
  Writer
  Seeker
}
```

```go
type ReaderFrom interface {
  ReadFrom(r Reader) (n int64, err error)
}
```

```go
type WriterTo interface {
  WriteTo(w Writer) (n int64, err error)
}
```

```go
type ReaderAt interface {
  ReadAt(p []byte, off int64) (n int, err error)
}
```

```go
type WriterAt interface {
  WriteAt(p []byte, off int64) (n int, err error)
}
```

```go
type ByteReader interface {
  ReadByte() (byte, error)
}
```

```go
type ByteScanner interface {
  ByteReader
  UnreadByte() error
}
```

```go
type ByteWriter interface {
  WriteByte(c byte) error
}
```

```go
type RuneReader interface {
  ReadRune() (r rune, size int, err error)
}
```

```go
type RuneScanner interface {
  RuneReader
  UnreadRune() error
}
```

```go
type StringWriter interface {
  WriteString(s string) (n int, err error)
}
```

## [math/rand/rand.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/math/rand/rand.go)

```go
type Source interface {
  Int63() int64
  Seed(seed int64)
}
```

```go
type Source64 interface {
  Source
  Uint64() uint64
}
```

## [mime/multipart/formdata.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/mime/multipart/formdata.go)

```go
type File interface {
  io.Reader
  io.ReaderAt
  io.Seeker
  io.Closer
}
```

## [net/http/client.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/net/http/client.go)

```go
type RoundTripper interface {
  RoundTrip(*Request) (*Response, error)
}
```

## [net/http/cookiejar/jar.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/net/http/cookiejar/jar.go)

```go
type PublicSuffixList interface {
  PublicSuffix(domain string) string

  String() string
}
```

## [net/http/fs.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/net/http/fs.go)

```go
type FileSystem interface {
  Open(name string) (File, error)
}
```

```go
type File interface {
  io.Closer
  io.Reader
  io.Seeker
  Readdir(count int) ([]fs.FileInfo, error)
  Stat() (fs.FileInfo, error)
}
```

## [net/http/http.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/net/http/http.go)

```go
type Pusher interface {
  Push(target string, opts *PushOptions) error
}
```

## [net/http/httputil/reverseproxy.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/net/http/httputil/reverseproxy.go)

```go
type BufferPool interface {
  Get() []byte
  Put([]byte)
}
```

## [net/http/jar.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/net/http/jar.go)

```go
type CookieJar interface {
  SetCookies(u *url.URL, cookies []*Cookie)

  Cookies(u *url.URL) []*Cookie
}
```

## [net/http/server.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/net/http/server.go)

```go
type Handler interface {
  ServeHTTP(ResponseWriter, *Request)
}
```

```go
type ResponseWriter interface {
  Header() Header

  Write([]byte) (int, error)

  WriteHeader(statusCode int)
}
```

```go
type Flusher interface {
  Flush()
}
```

```go
type Hijacker interface {
  Hijack() (net.Conn, *bufio.ReadWriter, error)
}
```

```go
type CloseNotifier interface {
  CloseNotify() <-chan bool
}
```

## [net/net.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/net/net.go)

```go
type Addr interface {
}
```

```go
type Conn interface {
  Read(b []byte) (n int, err error)

  Write(b []byte) (n int, err error)

  Close() error

  LocalAddr() Addr

  RemoteAddr() Addr

  SetDeadline(t time.Time) error

  SetReadDeadline(t time.Time) error

  SetWriteDeadline(t time.Time) error
}
```

```go
type PacketConn interface {
  ReadFrom(p []byte) (n int, addr Addr, err error)

  WriteTo(p []byte, addr Addr) (n int, err error)

  Close() error

  LocalAddr() Addr

  SetDeadline(t time.Time) error

  SetReadDeadline(t time.Time) error

  SetWriteDeadline(t time.Time) error
}
```

```go
type Listener interface {
  Accept() (Conn, error)

  Close() error

  Addr() Addr
}
```

```go
type Error interface {
  error

  Temporary() bool
}
```

## [net/rpc/client.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/net/rpc/client.go)

```go
type ClientCodec interface {
  WriteRequest(*Request, any) error
  ReadResponseHeader(*Response) error
  ReadResponseBody(any) error

  Close() error
}
```

## [net/rpc/server.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/net/rpc/server.go)

```go
type ServerCodec interface {
  ReadRequestHeader(*Request) error
  ReadRequestBody(any) error
  WriteResponse(*Response, any) error

  Close() error
}
```

## [net/smtp/auth.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/net/smtp/auth.go)

```go
type Auth interface {
  Start(server *ServerInfo) (proto string, toServer []byte, err error)

  Next(fromServer []byte, more bool) (toServer []byte, err error)
}
```

## [os/exec.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/os/exec.go)

```go
type Signal interface {
  String() string
}
```

## [reflect/type.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/reflect/type.go)

```go
type Type interface {

  Align() int

  FieldAlign() int

  Method(int) Method

  MethodByName(string) (Method, bool)

  NumMethod() int

  Name() string

  PkgPath() string

  Size() uintptr

  String() string

  Kind() Kind

  Implements(u Type) bool

  AssignableTo(u Type) bool

  ConvertibleTo(u Type) bool

  Comparable() bool


  Bits() int

  ChanDir() ChanDir

  IsVariadic() bool

  Elem() Type

  Field(i int) StructField

  FieldByIndex(index []int) StructField

  FieldByName(name string) (StructField, bool)

  FieldByNameFunc(match func(string) bool) (StructField, bool)

  In(i int) Type

  Key() Type

  Len() int

  NumField() int

  NumIn() int

  NumOut() int

  Out(i int) Type

  common() *rtype
  uncommon() *uncommonType
}
```

## [runtime/error.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/runtime/error.go)

```go
type Error interface {
  error

  RuntimeError()
}
```

## [sort/sort.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/sort/sort.go)

```go
type Interface interface {
  Len() int

  Less(i, j int) bool

  Swap(i, j int)
}
```

## [sync/mutex.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/sync/mutex.go)

```go
type Locker interface {
  Lock()
  Unlock()
}
```

## [syscall/net.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/syscall/net.go)

```go
type RawConn interface {
  Control(f func(fd uintptr)) error

  Read(f func(fd uintptr) (done bool)) error

  Write(f func(fd uintptr) (done bool)) error
}
```

```go
type Conn interface {
  SyscallConn() (RawConn, error)
}
```

## [syscall/route_bsd.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/syscall/route_bsd.go)

```go
type RoutingMessage interface {
  sockaddr() ([]Sockaddr, error)
}
```

## [syscall/syscall_unix.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/syscall/syscall_unix.go)

```go
type Sockaddr interface {
}
```

## [syscall/syscall_windows.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/syscall/syscall_windows.go)

```go
type Sockaddr interface {
}
```

## [testing/quick/quick.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/testing/quick/quick.go)

```go
type Generator interface {
  Generate(rand *rand.Rand, size int) reflect.Value
}
```

## [testing/testing.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/testing/testing.go)

```go
type TB interface {
  Cleanup(func())
  Error(args ...any)
  Errorf(format string, args ...any)
  Fail()
  FailNow()
  Failed() bool
  Fatal(args ...any)
  Fatalf(format string, args ...any)
  Helper()
  Log(args ...any)
  Logf(format string, args ...any)
  Name() string
  Setenv(key, value string)
  Skip(args ...any)
  SkipNow()
  Skipf(format string, args ...any)
  Skipped() bool
  TempDir() string

  private()
}
```

## [text/template/parse/node.go](https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/text/template/parse/node.go)

```go
type Node interface {
  Type() NodeType
  String() string
  Copy() Node
  tree() *Tree
  writeTo(*strings.Builder)
}
```
