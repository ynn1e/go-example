# Interfaces

go version: go1.20.5

## inflate.go

```go
type Resetter interface {
  // Reset discards any buffered data and resets the Resetter as if it was
  // newly initialized with the given reader.
  Reset(r io.Reader, dict []byte) error
}
```

## reader.go

```go
type Resetter interface {
  // Reset discards any buffered data and resets the Resetter as if it was
  // newly initialized with the given reader.
  Reset(r io.Reader, dict []byte) error
}
```

## heap.go

```go
type Interface interface {
  sort.Interface
  Push(x any) // add x as element Len()
  Pop() any   // remove and return element Len() - 1.
}
```

## context.go

```go
type Context interface {
  // Deadline returns the time when work done on behalf of this context
  // should be canceled. Deadline returns ok==false when no deadline is
  // set. Successive calls to Deadline return the same results.
  Deadline() (deadline time.Time, ok bool)

  // Done returns a channel that's closed when work done on behalf of this
  // context should be canceled. Done may return nil if this context can
  // never be canceled. Successive calls to Done return the same value.
  // The close of the Done channel may happen asynchronously,
  // after the cancel function returns.
  //
  // WithCancel arranges for Done to be closed when cancel is called;
  // WithDeadline arranges for Done to be closed when the deadline
  // expires; WithTimeout arranges for Done to be closed when the timeout
  // elapses.
  //
  // Done is provided for use in select statements:
  //
  //  // Stream generates values with DoSomething and sends them to out
  //  // until DoSomething returns an error or ctx.Done is closed.
  //  func Stream(ctx context.Context, out chan<- Value) error {
  //    for {
  //      v, err := DoSomething(ctx)
  //      if err != nil {
  //        return err
  //      }
  //      select {
  //      case <-ctx.Done():
  //        return ctx.Err()
  //      case out <- v:
  //      }
  //    }
  //  }
  //
  // See https://blog.golang.org/pipelines for more examples of how to use
  // a Done channel for cancellation.
  Done() <-chan struct{}

  // If Done is not yet closed, Err returns nil.
  // If Done is closed, Err returns a non-nil error explaining why:
  // Canceled if the context was canceled
  // or DeadlineExceeded if the context's deadline passed.
  // After Err returns a non-nil error, successive calls to Err return the same error.
  Err() error

  // Value returns the value associated with this context for key, or nil
  // if no value is associated with key. Successive calls to Value with
  // the same key returns the same result.
  //
  // Use context values only for request-scoped data that transits
  // processes and API boundaries, not for passing optional parameters to
  // functions.
  //
  // A key identifies a specific value in a Context. Functions that wish
  // to store values in Context typically allocate a key in a global
  // variable then use that key as the argument to context.WithValue and
  // Context.Value. A key can be any type that supports equality;
  // packages should define keys as an unexported type to avoid
  // collisions.
  //
  // Packages that define a Context key should provide type-safe accessors
  // for the values stored using that key:
  //
  //   // Package user defines a User type that's stored in Contexts.
  //   package user
  //
  //   import "context"
  //
  //   // User is the type of value stored in the Contexts.
  //   type User struct {...}
  //
  //   // key is an unexported type for keys defined in this package.
  //   // This prevents collisions with keys defined in other packages.
  //   type key int
  //
  //   // userKey is the key for user.User values in Contexts. It is
  //   // unexported; clients use user.NewContext and user.FromContext
  //   // instead of using this key directly.
  //   var userKey key
  //
  //   // NewContext returns a new Context that carries value u.
  //   func NewContext(ctx context.Context, u *User) context.Context {
  //     return context.WithValue(ctx, userKey, u)
  //   }
  //
  //   // FromContext returns the User value stored in ctx, if any.
  //   func FromContext(ctx context.Context) (*User, bool) {
  //     u, ok := ctx.Value(userKey).(*User)
  //     return u, ok
  //   }
  Value(key any) any
}
```

## cipher.go

```go
type Block interface {
  // BlockSize returns the cipher's block size.
  BlockSize() int

  // Encrypt encrypts the first block in src into dst.
  // Dst and src must overlap entirely or not at all.
  Encrypt(dst, src []byte)

  // Decrypt decrypts the first block in src into dst.
  // Dst and src must overlap entirely or not at all.
  Decrypt(dst, src []byte)
}
```

## gcm.go

```go
type AEAD interface {
  // NonceSize returns the size of the nonce that must be passed to Seal
  // and Open.
  NonceSize() int

  // Overhead returns the maximum difference between the lengths of a
  // plaintext and its ciphertext.
  Overhead() int

  // Seal encrypts and authenticates plaintext, authenticates the
  // additional data and appends the result to dst, returning the updated
  // slice. The nonce must be NonceSize() bytes long and unique for all
  // time, for a given key.
  //
  // To reuse plaintext's storage for the encrypted output, use plaintext[:0]
  // as dst. Otherwise, the remaining capacity of dst must not overlap plaintext.
  Seal(dst, nonce, plaintext, additionalData []byte) []byte

  // Open decrypts and authenticates ciphertext, authenticates the
  // additional data and, if successful, appends the resulting plaintext
  // to dst, returning the updated slice. The nonce must be NonceSize()
  // bytes long and both it and the additional data must match the
  // value passed to Seal.
  //
  // To reuse ciphertext's storage for the decrypted output, use ciphertext[:0]
  // as dst. Otherwise, the remaining capacity of dst must not overlap plaintext.
  //
  // Even if the function fails, the contents of dst, up to its capacity,
  // may be overwritten.
  Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error)
}
```

## crypto.go

```go
type Signer interface {
  // Public returns the public key corresponding to the opaque,
  // private key.
  Public() PublicKey

  // Sign signs digest with the private key, possibly using entropy from
  // rand. For an RSA key, the resulting signature should be either a
  // PKCS #1 v1.5 or PSS signature (as indicated by opts). For an (EC)DSA
  // key, it should be a DER-serialised, ASN.1 signature structure.
  //
  // Hash implements the SignerOpts interface and, in most cases, one can
  // simply pass in the hash function used as opts. Sign may also attempt
  // to type assert opts to other types in order to obtain algorithm
  // specific values. See the documentation in each package for details.
  //
  // Note that when a signature of a hash of a larger message is needed,
  // the caller is responsible for hashing the larger message and passing
  // the hash (as digest) and the hash function (as opts) to Sign.
  Sign(rand io.Reader, digest []byte, opts SignerOpts) (signature []byte, err error)
}
```

## ecdh.go

```go
type Curve interface {
  // GenerateKey generates a new PrivateKey from rand.
  GenerateKey(rand io.Reader) (*PrivateKey, error)

  // NewPrivateKey checks that key is valid and returns a PrivateKey.
  //
  // For NIST curves, this follows SEC 1, Version 2.0, Section 2.3.6, which
  // amounts to decoding the bytes as a fixed length big endian integer and
  // checking that the result is lower than the order of the curve. The zero
  // private key is also rejected, as the encoding of the corresponding public
  // key would be irregular.
  //
  // For X25519, this only checks the scalar length.
  NewPrivateKey(key []byte) (*PrivateKey, error)

  // NewPublicKey checks that key is valid and returns a PublicKey.
  //
  // For NIST curves, this decodes an uncompressed point according to SEC 1,
  // Version 2.0, Section 2.3.4. Compressed encodings and the point at
  // infinity are rejected.
  //
  // For X25519, this only checks the u-coordinate length. Adversarially
  // selected public keys can cause ECDH to return an error.
  NewPublicKey(key []byte) (*PublicKey, error)

  // ecdh performs a ECDH exchange and returns the shared secret. It's exposed
  // as the PrivateKey.ECDH method.
  //
  // The private method also allow us to expand the ECDH interface with more
  // methods in the future without breaking backwards compatibility.
  ecdh(local *PrivateKey, remote *PublicKey) ([]byte, error)

  // privateKeyToPublicKey converts a PrivateKey to a PublicKey. It's exposed
  // as the PrivateKey.PublicKey method.
  //
  // This method always succeeds: for X25519, the zero key can't be
  // constructed due to clamping; for NIST curves, it is rejected by
  // NewPrivateKey.
  privateKeyToPublicKey(*PrivateKey) *PublicKey
}
```

## elliptic.go

```go
type Curve interface {
  // Params returns the parameters for the curve.
  Params() *CurveParams

  // IsOnCurve reports whether the given (x,y) lies on the curve.
  //
  // Note: this is a low-level unsafe API. For ECDH, use the crypto/ecdh
  // package. The NewPublicKey methods of NIST curves in crypto/ecdh accept
  // the same encoding as the Unmarshal function, and perform on-curve checks.
  IsOnCurve(x, y *big.Int) bool

  // Add returns the sum of (x1,y1) and (x2,y2).
  //
  // Note: this is a low-level unsafe API.
  Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int)

  // Double returns 2*(x,y).
  //
  // Note: this is a low-level unsafe API.
  Double(x1, y1 *big.Int) (x, y *big.Int)

  // ScalarMult returns k*(x,y) where k is an integer in big-endian form.
  //
  // Note: this is a low-level unsafe API. For ECDH, use the crypto/ecdh
  // package. Most uses of ScalarMult can be replaced by a call to the ECDH
  // methods of NIST curves in crypto/ecdh.
  ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int)

  // ScalarBaseMult returns k*G, where G is the base point of the group
  // and k is an integer in big-endian form.
  //
  // Note: this is a low-level unsafe API. For ECDH, use the crypto/ecdh
  // package. Most uses of ScalarBaseMult can be replaced by a call to the
  // PrivateKey.PublicKey method in crypto/ecdh.
  ScalarBaseMult(k []byte) (x, y *big.Int)
}
```

## common.go

```go
type ClientSessionCache interface {
  // Get searches for a ClientSessionState associated with the given key.
  // On return, ok is true if one was found.
  Get(sessionKey string) (session *ClientSessionState, ok bool)

  // Put adds the ClientSessionState to the cache with the given key. It might
  // get called multiple times in a connection if a TLS 1.3 server provides
  // more than one session ticket. If called with a nil *ClientSessionState,
  // it should remove the cache entry.
  Put(sessionKey string, cs *ClientSessionState)
}
```

## driver.go

```go
type Driver interface {
  // Open returns a new connection to the database.
  // The name is a string in a driver-specific format.
  //
  // Open may return a cached connection (one previously
  // closed), but doing so is unnecessary; the sql package
  // maintains a pool of idle connections for efficient re-use.
  //
  // The returned connection is only used by one goroutine at a
  // time.
  Open(name string) (Conn, error)
}
```

## types.go

```go
type ValueConverter interface {
  // ConvertValue converts a value to a driver Value.
  ConvertValue(v any) (Value, error)
}
```

## sql.go

```go
type Scanner interface {
  // Scan assigns a value from a database driver.
  //
  // The src value will be of one of the following types:
  //
  //    int64
  //    float64
  //    bool
  //    []byte
  //    string
  //    time.Time
  //    nil - for NULL values
  //
  // An error should be returned if the value cannot be stored
  // without loss of information.
  //
  // Reference types such as []byte are only valid until the next call to Scan
  // and should not be retained. Their underlying memory is owned by the driver.
  // If retention is necessary, copy their values before the next call to Scan.
  Scan(src any) error
}
```

## type.go

```go
type Type interface {
  Common() *CommonType
  String() string
  Size() int64
}
```

## file.go

```go
type Load interface {
  Raw() []byte
}
```

## binary.go

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

## encoding.go

```go
type BinaryMarshaler interface {
  MarshalBinary() (data []byte, err error)
}
```

## decode.go

```go
type Unmarshaler interface {
  UnmarshalJSON([]byte) error
}
```

## encode.go

```go
type Marshaler interface {
  MarshalJSON() ([]byte, error)
}
```

## marshal.go

```go
type Marshaler interface {
  MarshalXML(e *Encoder, start StartElement) error
}
```

## read.go

```go
type Unmarshaler interface {
  UnmarshalXML(d *Decoder, start StartElement) error
}
```

## xml.go

```go
type TokenReader interface {
  Token() (Token, error)
}
```

## expvar.go

```go
type Var interface {
  // String returns a valid JSON value for the variable.
  // Types with String methods that do not return valid JSON
  // (such as time.Time) must not be used as a Var.
  String() string
}
```

## flag.go

```go
type Value interface {
  String() string
  Set(string) error
}
```

## print.go

```go
type State interface {
  // Write is the function to call to emit formatted output to be printed.
  Write(b []byte) (n int, err error)
  // Width returns the value of the width option and whether it has been set.
  Width() (wid int, ok bool)
  // Precision returns the value of the precision option and whether it has been set.
  Precision() (prec int, ok bool)

  // Flag reports whether the flag c, a character, has been set.
  Flag(c int) bool
}
```

## scan.go

```go
type ScanState interface {
  // ReadRune reads the next rune (Unicode code point) from the input.
  // If invoked during Scanln, Fscanln, or Sscanln, ReadRune() will
  // return EOF after returning the first '\n' or when reading beyond
  // the specified width.
  ReadRune() (r rune, size int, err error)
  // UnreadRune causes the next call to ReadRune to return the same rune.
  UnreadRune() error
  // SkipSpace skips space in the input. Newlines are treated appropriately
  // for the operation being performed; see the package documentation
  // for more information.
  SkipSpace()
  // Token skips space in the input if skipSpace is true, then returns the
  // run of Unicode code points c satisfying f(c).  If f is nil,
  // !unicode.IsSpace(c) is used; that is, the token will hold non-space
  // characters. Newlines are treated appropriately for the operation being
  // performed; see the package documentation for more information.
  // The returned slice points to shared data that may be overwritten
  // by the next call to Token, a call to a Scan function using the ScanState
  // as input, or when the calling Scan method returns.
  Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
  // Width returns the value of the width option and whether it has been set.
  // The unit is Unicode code points.
  Width() (wid int, ok bool)
  // Because ReadRune is implemented by the interface, Read should never be
  // called by the scanning routines and a valid implementation of
  // ScanState may choose always to return an error from Read.
  Read(buf []byte) (n int, err error)
}
```

## ast.go

```go
type Node interface {
  Pos() token.Pos // position of first character belonging to the node
  End() token.Pos // position of first character immediately after the node
}
```

## walk.go

```go
type Visitor interface {
  Visit(node Node) (w Visitor)
}
```

## expr.go

```go
type Expr interface {
  // String returns the string form of the expression,
  // using the boolean syntax used in //go:build lines.
  String() string

  // Eval reports whether the expression evaluates to true.
  // It calls ok(tag) as needed to find out whether a given build tag
  // is satisfied by the current build configuration.
  Eval(ok func(tag string) bool) bool

  // The presence of an isExpr method explicitly marks the type as an Expr.
  // Only implementations in this package should be used as Exprs.
  isExpr()
}
```

## value.go

```go
type Value interface {
  // Kind returns the value kind.
  Kind() Kind

  // String returns a short, quoted (human-readable) form of the value.
  // For numeric values, the result may be an approximation;
  // for String values the result may be a shortened string.
  // Use ExactString for a string representing a value exactly.
  String() string

  // ExactString returns an exact, quoted (human-readable) form of the value.
  // If the Value is of Kind String, use StringVal to obtain the unquoted string.
  ExactString() string

  // Prevent external implementations.
  implementsValue()
}
```

## parse.go

```go
type Block interface {
  block()
}
```

## api.go

```go
type Importer interface {
  // Import returns the imported package for the given import path.
  // The semantics is like for ImporterFrom.ImportFrom except that
  // dir and mode are ignored (since they are not present).
  Import(path string) (*Package, error)
}
```

## object.go

```go
type Object interface {
  Parent() *Scope // scope in which this object is declared; nil for methods and struct fields
  Pos() token.Pos // position of object identifier in declaration
  Pkg() *Package  // package to which this object belongs; nil for labels and objects in the Universe scope
  Name() string   // package local object name
  Type() Type     // object type
  Exported() bool // reports whether the name starts with a capital letter
  Id() string     // object name if exported, qualified name if not exported (see func Id)

  // String returns a human-readable string of the object.
  String() string

  // order reflects a package-level object's source order: if object
  // a is before object b in the source, then a.order() < b.order().
  // order returns a value > 0 for package-level objects; it returns
  // 0 for all other objects (including objects in file scopes).
  order() uint32

  // color returns the object's color.
  color() color

  // setType sets the type of the object.
  setType(Type)

  // setOrder sets the order number of the object. It must be > 0.
  setOrder(uint32)

  // setColor sets the object's color. It must not be white.
  setColor(color color)

  // setParent sets the parent scope of the object.
  setParent(*Scope)

  // sameId reports whether obj.Id() and Id(pkg, name) are the same.
  sameId(pkg *Package, name string) bool

  // scopePos returns the start position of the scope of this Object
  scopePos() token.Pos

  // setScopePos sets the start position of the scope for this Object.
  setScopePos(pos token.Pos)
}
```

## sizes.go

```go
type Sizes interface {
  // Alignof returns the alignment of a variable of type T.
  // Alignof must implement the alignment guarantees required by the spec.
  Alignof(T Type) int64

  // Offsetsof returns the offsets of the given struct fields, in bytes.
  // Offsetsof must implement the offset guarantees required by the spec.
  Offsetsof(fields []*Var) []int64

  // Sizeof returns the size of a variable of type T.
  // Sizeof must implement the size guarantees required by the spec.
  Sizeof(T Type) int64
}
```

## hash.go

```go
type Hash interface {
  // Write (via the embedded io.Writer interface) adds more data to the running hash.
  // It never returns an error.
  io.Writer

  // Sum appends the current hash to b and returns the resulting slice.
  // It does not change the underlying hash state.
  Sum(b []byte) []byte

  // Reset resets the Hash to its initial state.
  Reset()

  // Size returns the number of bytes Sum will return.
  Size() int

  // BlockSize returns the hash's underlying block size.
  // The Write method must be able to accept any amount
  // of data, but it may operate more efficiently if all writes
  // are a multiple of the block size.
  BlockSize() int
}
```

## color.go

```go
type Color interface {
  // RGBA returns the alpha-premultiplied red, green, blue and alpha values
  // for the color. Each value ranges within [0, 0xffff], but is represented
  // by a uint32 so that multiplying by a blend factor up to 0xffff will not
  // overflow.
  //
  // An alpha-premultiplied color component c has been scaled by alpha (a),
  // so has valid values 0 <= c <= a.
  RGBA() (r, g, b, a uint32)
}
```

## draw.go

```go
type Image interface {
  image.Image
  Set(x, y int, c color.Color)
}
```

## image.go

```go
type Image interface {
  // ColorModel returns the Image's color model.
  ColorModel() color.Model
  // Bounds returns the domain for which At can return non-zero color.
  // The bounds do not necessarily contain the point (0, 0).
  Bounds() Rectangle
  // At returns the color of the pixel at (x, y).
  // At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
  // At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
  At(x, y int) color.Color
}
```

## writer.go

```go
type EncoderBufferPool interface {
  Get() *EncoderBuffer
  Put(*EncoderBuffer)
}
```

## codes.go

```go
type Code interface {
  // Marker returns the SyncMarker for the Code's dynamic type.
  Marker() SyncMarker

  // Value returns the Code's ordinal value.
  Value() int
}
```

## log.go

```go
type Interface interface {
  Getenv(key string)
  Stat(file string)
  Open(file string)
  Chdir(dir string)
}
```

## fs.go

```go
type FS interface {
  // Open opens the named file.
  //
  // When Open returns an error, it should be of type *PathError
  // with the Op field set to "open", the Path field set to name,
  // and the Err field describing the problem.
  //
  // Open should reject attempts to open names that do not satisfy
  // ValidPath(name), returning a *PathError with Err set to
  // ErrInvalid or ErrNotExist.
  Open(name string) (File, error)
}
```

## glob.go

```go
type GlobFS interface {
  FS

  // Glob returns the names of all files matching pattern,
  // providing an implementation of the top-level
  // Glob function.
  Glob(pattern string) ([]string, error)
}
```

## readdir.go

```go
type ReadDirFS interface {
  FS

  // ReadDir reads the named directory
  // and returns a list of directory entries sorted by filename.
  ReadDir(name string) ([]DirEntry, error)
}
```

## readfile.go

```go
type ReadFileFS interface {
  FS

  // ReadFile reads the named file and returns its contents.
  // A successful call returns a nil error, not io.EOF.
  // (Because ReadFile reads the whole file, the expected EOF
  // from the final Read is not treated as an error to be reported.)
  //
  // The caller is permitted to modify the returned byte slice.
  // This method should return a copy of the underlying data.
  ReadFile(name string) ([]byte, error)
}
```

## stat.go

```go
type StatFS interface {
  FS

  // Stat returns a FileInfo describing the file.
  // If there is an error, it should be of type *PathError.
  Stat(name string) (FileInfo, error)
}
```

## sub.go

```go
type SubFS interface {
  FS

  // Sub returns an FS corresponding to the subtree rooted at dir.
  Sub(dir string) (FS, error)
}
```

## io.go

```go
type Reader interface {
  Read(p []byte) (n int, err error)
}
```

## rand.go

```go
type Source interface {
  Int63() int64
  Seed(seed int64)
}
```

## formdata.go

```go
type File interface {
  io.Reader
  io.ReaderAt
  io.Seeker
  io.Closer
}
```

## client.go

```go
type RoundTripper interface {
  // RoundTrip executes a single HTTP transaction, returning
  // a Response for the provided Request.
  //
  // RoundTrip should not attempt to interpret the response. In
  // particular, RoundTrip must return err == nil if it obtained
  // a response, regardless of the response's HTTP status code.
  // A non-nil err should be reserved for failure to obtain a
  // response. Similarly, RoundTrip should not attempt to
  // handle higher-level protocol details such as redirects,
  // authentication, or cookies.
  //
  // RoundTrip should not modify the request, except for
  // consuming and closing the Request's Body. RoundTrip may
  // read fields of the request in a separate goroutine. Callers
  // should not mutate or reuse the request until the Response's
  // Body has been closed.
  //
  // RoundTrip must always close the body, including on errors,
  // but depending on the implementation may do so in a separate
  // goroutine even after RoundTrip returns. This means that
  // callers wanting to reuse the body for subsequent requests
  // must arrange to wait for the Close call before doing so.
  //
  // The Request's URL and Header fields must be initialized.
  RoundTrip(*Request) (*Response, error)
}
```

## jar.go

```go
type PublicSuffixList interface {
  // PublicSuffix returns the public suffix of domain.
  //
  // TODO: specify which of the caller and callee is responsible for IP
  // addresses, for leading and trailing dots, for case sensitivity, and
  // for IDN/Punycode.
  PublicSuffix(domain string) string

  // String returns a description of the source of this public suffix
  // list. The description will typically contain something like a time
  // stamp or version number.
  String() string
}
```

## http.go

```go
type Pusher interface {
  // Push initiates an HTTP/2 server push. This constructs a synthetic
  // request using the given target and options, serializes that request
  // into a PUSH_PROMISE frame, then dispatches that request using the
  // server's request handler. If opts is nil, default options are used.
  //
  // The target must either be an absolute path (like "/path") or an absolute
  // URL that contains a valid host and the same scheme as the parent request.
  // If the target is a path, it will inherit the scheme and host of the
  // parent request.
  //
  // The HTTP/2 spec disallows recursive pushes and cross-authority pushes.
  // Push may or may not detect these invalid pushes; however, invalid
  // pushes will be detected and canceled by conforming clients.
  //
  // Handlers that wish to push URL X should call Push before sending any
  // data that may trigger a request for URL X. This avoids a race where the
  // client issues requests for X before receiving the PUSH_PROMISE for X.
  //
  // Push will run in a separate goroutine making the order of arrival
  // non-deterministic. Any required synchronization needs to be implemented
  // by the caller.
  //
  // Push returns ErrNotSupported if the client has disabled push or if push
  // is not supported on the underlying connection.
  Push(target string, opts *PushOptions) error
}
```

## reverseproxy.go

```go
type BufferPool interface {
  Get() []byte
  Put([]byte)
}
```

## server.go

```go
type Handler interface {
  ServeHTTP(ResponseWriter, *Request)
}
```

## net.go

```go
type Addr interface {
  Network() string // name of the network (for example, "tcp", "udp")
  String() string  // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")
}
```

## auth.go

```go
type Auth interface {
  // Start begins an authentication with a server.
  // It returns the name of the authentication protocol
  // and optionally data to include in the initial AUTH message
  // sent to the server.
  // If it returns a non-nil error, the SMTP client aborts
  // the authentication attempt and closes the connection.
  Start(server *ServerInfo) (proto string, toServer []byte, err error)

  // Next continues the authentication. The server has just sent
  // the fromServer data. If more is true, the server expects a
  // response, which Next should return as toServer; otherwise
  // Next should return toServer == nil.
  // If Next returns a non-nil error, the SMTP client aborts
  // the authentication attempt and closes the connection.
  Next(fromServer []byte, more bool) (toServer []byte, err error)
}
```

## exec.go

```go
type Signal interface {
  String() string
  Signal() // to distinguish from other Stringers
}
```

## error.go

```go
type Error interface {
  error

  // RuntimeError is a no-op function but
  // serves to distinguish types that are run time
  // errors from ordinary errors: a type is a
  // run time error if it has a RuntimeError method.
  RuntimeError()
}
```

## sort.go

```go
type Interface interface {
  // Len is the number of elements in the collection.
  Len() int

  // Less reports whether the element with index i
  // must sort before the element with index j.
  //
  // If both Less(i, j) and Less(j, i) are false,
  // then the elements at index i and j are considered equal.
  // Sort may place equal elements in any order in the final result,
  // while Stable preserves the original input order of equal elements.
  //
  // Less must describe a transitive ordering:
  //  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
  //  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
  //
  // Note that floating-point comparison (the < operator on float32 or float64 values)
  // is not a transitive ordering when not-a-number (NaN) values are involved.
  // See Float64Slice.Less for a correct implementation for floating-point values.
  Less(i, j int) bool

  // Swap swaps the elements with indexes i and j.
  Swap(i, j int)
}
```

## mutex.go

```go
type Locker interface {
  Lock()
  Unlock()
}
```

## route_bsd.go

```go
type RoutingMessage interface {
  sockaddr() ([]Sockaddr, error)
}
```

## syscall_unix.go

```go
type Sockaddr interface {
  sockaddr() (ptr unsafe.Pointer, len _Socklen, err error) // lowercase; only we can define Sockaddrs
}
```

## syscall_windows.go

```go
type Sockaddr interface {
  sockaddr() (ptr unsafe.Pointer, len int32, err error) // lowercase; only we can define Sockaddrs
}
```

## quick.go

```go
type Generator interface {
  // Generate returns a random instance of the type on which it is a
  // method using the size as a size hint.
  Generate(rand *rand.Rand, size int) reflect.Value
}
```

## testing.go

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

  // A private method to prevent users implementing the
  // interface and so future additions to it will not
  // violate Go 1 compatibility.
  private()
}
```

## node.go

```go
type Node interface {
  Type() NodeType
  String() string
  // Copy does a deep copy of the Node and all its components.
  // To avoid type assertions, some XxxNodes also have specialized
  // CopyXxx methods that return *XxxNode.
  Copy() Node
  Position() Pos // byte position of start of node in full original input string
  // tree returns the containing *Tree.
  // It is unexported so all implementations of Node are in this package.
  tree() *Tree
  // writeTo writes the String output to the builder.
  writeTo(*strings.Builder)
}
```
