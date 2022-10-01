
package main

	
import (
    "fmt"
    // "os"
    "github.com/veraison/go-cose"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
    // "encoding/json"
    "encoding/base64"
    // "encoding/hex"
    "math/big"
)



type JWK struct {
    Kid string
    Kty string
    Crv string
    Alg string
    X string
    Y string
    D string
}

type PublicKey struct {
	elliptic.Curve
	X, Y *big.Int
}

type PrivateKey struct {
	PublicKey
	D *big.Int
}

func Base64UrlEncode(src []byte) []byte {
    return []byte(base64.RawURLEncoding.EncodeToString(src))
}

func Base64UrlDecode(src []byte) ([]byte, error) {
    return base64.RawURLEncoding.DecodeString(string(src))
}

func main() {
    

    var messageStr = "🌈 I wear the Armor of Mithras and the Light."

    var messageBytes = []byte(messageStr)

    privateKey, _ := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)


    // TODO: use an imported key
    // var privateKeyJwk = os.Args[1]
    // fmt.Println( privateKeyJwk  )
    // var importedPrivateKeyMap JWK
    // json.Unmarshal([]byte(privateKeyJwk), &importedPrivateKeyMap)

    // fmt.Println( importedPrivateKeyMap.X, importedPrivateKeyMap.Y, importedPrivateKeyMap.D  )
    // var d, _ = Base64UrlDecode([]byte(importedPrivateKeyMap.D))
    // fmt.Println( hex.EncodeToString(d)  )
    // fmt.Println(privateKey.D )

    signer, _ := cose.NewSigner(cose.AlgorithmES512, privateKey)

    headers := cose.Headers{
        Protected: cose.ProtectedHeader{
            cose.HeaderLabelAlgorithm: cose.AlgorithmES512,
        },
    }

    sig, _ := cose.Sign1(rand.Reader, signer, headers, messageBytes, nil)

    publicKey := privateKey.Public()
    verifier, _ := cose.NewVerifier(cose.AlgorithmES512, publicKey)

    var msg cose.Sign1Message
    _ = msg.UnmarshalCBOR(sig)
    _ = msg.Verify(nil, verifier)

    fmt.Println( string(msg.Payload) )

    
}