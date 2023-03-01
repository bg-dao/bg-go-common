package baseUtil

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	var aesKey = []byte("7894abu9y8d2a8f7")
	pass := []byte("17612345678")
	xPass, err := AesEncrypt(pass, aesKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	pass64 := base64.StdEncoding.EncodeToString(xPass)
	fmt.Printf("加密后:%v\n", pass64)

	bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	if err != nil {
		fmt.Println(err)
		return
	}

	tPass, err := AesDecrypt(bytesPass, aesKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("解密后:%s\n", tPass)
}
