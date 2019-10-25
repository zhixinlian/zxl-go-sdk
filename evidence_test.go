package zxl_go_sdk

import (
	"encoding/hex"
	"fmt"
	"github.com/zhixinlian/zxl-go-sdk/sm/sm3"
	"testing"
	"time"
)

func TestCetc(t *testing.T) {
	zxl, err := NewZxlImpl("xxxxxxxxxx1xxxx", "appTest77")
	if err != nil {
		t.Error(err.Error())
	}
	pk, sk, err := zxl.GenerateKeyPair()
	if err != nil {
		t.Error(err.Error())
	}
	//err = zxl.BindUserCert(pk, sk)
	//if err != nil {
	//	t.Error(err)
	//}
	err = zxl.UpdateUserCert(pk, sk)
	if err != nil {
		t.Error(err.Error())
	}

	hashData := sm3.SumSM3([]byte("123123123"))
	evHash := hex.EncodeToString(hashData)
	result, err := zxl.EvidenceSave(evHash, "abc", sk, pk)
	if err != nil {
		t.Error(err.Error())
	}
	formatData, _ := result.(*CetcEvidenceResp)
	fmt.Println(formatData.EvId, formatData.TxHash, formatData.CreateTime)

	time.Sleep(time.Second * 10)
	queryResult1, err := zxl.QueryWithTxHash(formatData.TxHash)
	if err != nil || queryResult1[0].EvHash != evHash{
		t.Error("QueryWithTxHash error")
	}

	queryResult2, err := zxl.QueryWithEvId(queryResult1[0].EvId)
	if err != nil || queryResult2[0].EvHash != evHash{
		t.Error("QueryWithEvId error")
	}
	queryResult3, err := zxl.QueryWithEvHash(evHash)
	if err != nil || queryResult3[0].EvHash != evHash{
		t.Error("QueryWithEvHash error")
	}
}


func TestTencent(t *testing.T) {
	zxl, err := NewZxlImpl("xxxxxxxxxx0xxxx", "appTestTX")
	if err != nil {
		t.Error(err)
	}
	pk, sk, err := zxl.GenerateKeyPair()
	if err != nil {
		t.Error(err)
	}
	err = zxl.UpdateUserCert(pk, sk)
	if err != nil {
		t.Error(err.Error())
	}

	hashData := sm3.SumSM3([]byte("123123123"))
	evHash := hex.EncodeToString(hashData)
	result, err := zxl.EvidenceSave(evHash, "abc", sk, pk)
	if err != nil {
		t.Error(err.Error())
	}
	formatData, _ := result.(*TencentEvidenceResp)
	fmt.Println(formatData.TxHash, formatData.CreateTime, formatData.BlockHeight)

	time.Sleep(time.Second * 10)
	queryResult1, err := zxl.QueryWithTxHash(result.GetTxHash())
	if err != nil || queryResult1[0].EvHash != evHash{
		t.Error("QueryWithTxHash error")
	}

	queryResult2, err := zxl.QueryWithEvId(result.GetEvId())
	if err != nil || queryResult2[0].EvHash != evHash{
		t.Error("QueryWithEvId error")
	}
	queryResult3, err := zxl.QueryWithEvHash(evHash)
	if err != nil || queryResult3[0].EvHash != evHash{
		t.Error("QueryWithEvHash error")
	}
}

func TestSign(t *testing.T) {
	zxl, err := NewZxlImpl("xxxxxxxxxx0xxxx", "appTestTX")
	if err != nil {
		t.Error(err)
	}
	pk, sk, err := zxl.GenerateKeyPair()
	if err != nil {
		t.Error(err)
	}

	rawData := []byte("abcdefghijklmnopqrstuvwxyz")
	signStr, err := zxl.Sign(sk, rawData)
	fmt.Println(signStr, err)

	success, err := zxl.Verify(pk, signStr, rawData)
	fmt.Println(success, err)
}

func TestCipher(t *testing.T) {
	zxl, err := NewZxlImpl("xxxxxxxxxx0xxxx", "appTestTX")
	if err != nil {
		t.Error(err)
	}
	rawData := []byte("123abc")
	pwd := "111"
	encryptedData, err := zxl.EncryptData(pwd, rawData)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(encryptedData, err)
	decryptedData, err := zxl.DecryptData(pwd, encryptedData)
	fmt.Println(string(decryptedData), err)
}

func TestHash(t *testing.T) {
	zxl, err := NewZxlImpl("xxxxxxxxxx1xxxx", "appTestTX")
	if err != nil {
		t.Error(err)
	}
	hashStr, err := zxl.CalculateHash("G:\\channel.zip")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(hashStr)
}


