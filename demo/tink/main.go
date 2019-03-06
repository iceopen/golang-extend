package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/google/tink/go/subtle/aead"
	"github.com/google/tink/go/subtle/random"
	"golang.org/x/crypto/chacha20poly1305"
)

func main() {
	key := random.GetRandomBytes(chacha20poly1305.KeySize)
	aad := []byte("infoepoch2018")
	badAad := []byte{1, 2, 3}

	x, err := aead.NewXChacha20poly1305Aead(key)
	if err != nil {
		logs.Info(err)
	}
	bb, err1 := x.Decrypt(aad, badAad)
	if err1 != nil {
		logs.Info(err1)
	}
	logs.Info(string(bb))
}
