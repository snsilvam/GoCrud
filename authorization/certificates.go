package authorization

import (
	"crypto/rsa"
	"io/ioutil"
	"sync"
)

var (
	signKey   *rsa.PrivateKey
	verifykey *rsa.PublicKey
	once      sync.Once
)

func LoadFiles(privateFile, publicFile string) error {
	var err error
	once.Do(func() {
		err = loadFiles(privateFile, publicFile)
	})
	return err
}

func loadFiles(privateFile, publicFile string) error {
	privateBytes, err := ioutil.ReadFile(privateFile)
	if err != nil {
		return err
	}
	publicBytes, err := ioutil.ReadFile(publicFile)
	if err != nil {
		return err
	}
	return parseRSA(privateBytes, publicBytes)
}
func parseRSA(privateBytes, publicBytes []byte) error {
	var err error
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return err
	}
	verifykey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return err
	}
	return nil
}
