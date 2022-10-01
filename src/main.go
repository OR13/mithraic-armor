
package main

	
import (
    "fmt"
    "github.com/veraison/go-cose"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)


func main() {
    
    var messageStr = "🌈 I wear the Armor of Mithras and the Light."

    var messageBytes = []byte(messageStr)

    // fmt.Println( messageStr )

    privateKey, _ := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
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