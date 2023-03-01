package baseUtil

import (
	"testing"
)

func TestRsaSignWithMd5(t *testing.T) {
	data := ""
	privateKeyHex := ""
	result, err := RsaSignWithMd5(data, privateKeyHex)
	if err != nil {
		t.Errorf("RsaSignWithMd5 error")
		return
	}
	t.Logf("result:%v", result)
}

func TestRsaVerifySignWithMd5(t *testing.T) {
	originalData := ""
	signData := ""
	publicKeyHex := ""
	resultFlag := RsaVerifySignWithMd5(originalData, signData, publicKeyHex)
	t.Logf("resultFlag:%v", resultFlag)
}
