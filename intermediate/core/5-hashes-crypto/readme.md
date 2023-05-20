# Hashes and Cryptography

## Crypto package

- Cryptography is the practice of ensuring secure communication in the
  presence of third parties.
- It makes use of several encryption, decryption and cryptic techniques
  to ensure that digital data is not exploited b unwanted and harmful
  entities.
- Go provides good support for cryptography and hashing through the package
  `crypto`

## Crypto package

- aes
- cipher
- rsa
- sha
- md5

## docs

```go
$ go doc encoding/hex
package hex // import "encoding/hex"

Package hex implements hexadecimal encoding and decoding.

var ErrLength = errors.New("encoding/hex: odd length hex string")
func Decode(dst, src []byte) (int, error)
func DecodeString(s string) ([]byte, error)
func DecodedLen(x int) int
func Dump(data []byte) string
func Dumper(w io.Writer) io.WriteCloser
func Encode(dst, src []byte) int
func EncodeToString(src []byte) string
func EncodedLen(n int) int
func NewDecoder(r io.Reader) io.Reader
func NewEncoder(w io.Writer) io.Writer
type InvalidByteError byte
```

```go
$ go doc crypto/md5
package md5 // import "crypto/md5"

Package md5 implements the MD5 hash algorithm as defined in RFC 1321.

MD5 is cryptographically broken and should not be used for secure
applications.

const BlockSize = 64
const Size = 16
func New() hash.Hash
func Sum(data []byte) [Size]byte
```

```go
$ go doc crypto/rsa
package rsa // import "crypto/rsa"

Package rsa implements RSA encryption as specified in PKCS #1 and RFC 8017.

RSA is a single, fundamental operation that is used in this package to
implement either public-key encryption or public-key signatures.

The original specification for encryption and signatures with RSA is PKCS #1
and the terms "RSA encryption" and "RSA signatures" by default refer to PKCS
#1 version 1.5. However, that specification has flaws and new designs should
use version 2, usually called by just OAEP and PSS, where possible.

Two sets of interfaces are included in this package. When a more abstract
interface is not necessary, there are functions for encrypting/decrypting
with v1.5/OAEP and signing/verifying with v1.5/PSS. If one needs to abstract
over the public key primitive, the PrivateKey type implements the Decrypter
and Signer interfaces from the crypto package.

The RSA operations in this package are not implemented using constant-time
algorithms.

const PSSSaltLengthAuto = 0 ...
var ErrDecryption = errors.New("crypto/rsa: decryption error")
var ErrMessageTooLong = errors.New("crypto/rsa: message too long for RSA public key size")
var ErrVerification = errors.New("crypto/rsa: verification error")
func DecryptOAEP(hash hash.Hash, random io.Reader, priv *PrivateKey, ciphertext []byte, ...) ([]byte, error)
func DecryptPKCS1v15(rand io.Reader, priv *PrivateKey, ciphertext []byte) ([]byte, error)
func DecryptPKCS1v15SessionKey(rand io.Reader, priv *PrivateKey, ciphertext []byte, key []byte) error
func EncryptOAEP(hash hash.Hash, random io.Reader, pub *PublicKey, msg []byte, label []byte) ([]byte, error)
func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) ([]byte, error)
func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error)
func SignPSS(rand io.Reader, priv *PrivateKey, hash crypto.Hash, digest []byte, ...) ([]byte, error)
func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error
func VerifyPSS(pub *PublicKey, hash crypto.Hash, digest []byte, sig []byte, opts *PSSOptions) error
type CRTValue struct{ ... }
type OAEPOptions struct{ ... }
type PKCS1v15DecryptOptions struct{ ... }
type PSSOptions struct{ ... }
type PrecomputedValues struct{ ... }
type PrivateKey struct{ ... }
    func GenerateKey(random io.Reader, bits int) (*PrivateKey, error)
    func GenerateMultiPrimeKey(random io.Reader, nprimes int, bits int) (*PrivateKey, error)
type PublicKey struct{ ... }
```

```go
$ go doc crypto/aes
package aes // import "crypto/aes"

Package aes implements AES encryption (formerly Rijndael), as defined in
U.S. Federal Information Processing Standards Publication 197.

The AES operations in this package are not implemented using constant-time
algorithms. An exception is when running on systems with enabled hardware
support for AES that makes these operations constant-time. Examples include
amd64 systems using AES-NI extensions and s390x systems using
Message-Security-Assist extensions. On such systems, when the result of
NewCipher is passed to cipher.NewGCM, the GHASH operation used by GCM is
also constant-time.

const BlockSize = 16
func NewCipher(key []byte) (cipher.Block, error)
type KeySizeError int
```

```go
$ go doc crypto/cipher
package cipher // import "crypto/cipher"

Package cipher implements standard block cipher modes that can be wrapped
around low-level block cipher implementations. See
https://csrc.nist.gov/groups/ST/toolkit/BCM/current_modes.html and NIST
Special Publication 800-38A.

type AEAD interface{ ... }
    func NewGCM(cipher Block) (AEAD, error)
    func NewGCMWithNonceSize(cipher Block, size int) (AEAD, error)
    func NewGCMWithTagSize(cipher Block, tagSize int) (AEAD, error)
type Block interface{ ... }
type BlockMode interface{ ... }
    func NewCBCDecrypter(b Block, iv []byte) BlockMode
    func NewCBCEncrypter(b Block, iv []byte) BlockMode
type Stream interface{ ... }
    func NewCFBDecrypter(block Block, iv []byte) Stream
    func NewCFBEncrypter(block Block, iv []byte) Stream
    func NewCTR(block Block, iv []byte) Stream
    func NewOFB(b Block, iv []byte) Stream
type StreamReader struct{ ... }
type StreamWriter struct{ ... }
```
