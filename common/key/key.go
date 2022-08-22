package key

import (
	"io/ioutil"
	"log"
	"os"
)

const (
	privKeyEnv  = "JWT_PRIVATE_KEY"
	privKeyPath = "/internal/common/key/private.pem"
	pubKeyEnv   = "JWT_PUBLIC_KEY"
	pubKeyPath  = "/internal/common/key/public.pem"
)

var privKey []byte
var pubKey []byte

func mustGetKey(keyType, envVar, path string) []byte {

	key := os.Getenv(envVar)
	if key != "" {
		return []byte(key)
	}

	pwd, _ := os.Getwd()

	keyBytes, err := ioutil.ReadFile(pwd + path)
	if err != nil {
		log.Fatalf("must provide %s key", keyType)
	}

	log.Printf("WARN: using %s key from key folder", keyType)

	return keyBytes
}
func Load() {
	privKey = mustGetKey("private", privKeyEnv, privKeyPath)
	pubKey = mustGetKey("public", pubKeyEnv, pubKeyPath)
}
func Private() []byte {
	return privKey
}

func Public() []byte {
	return pubKey
}
