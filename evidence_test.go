package zxl_go_sdk

import (
	"encoding/hex"
	"fmt"
	"github.com/zhixinlian/zxl-go-sdk/sm/sm3"
	"testing"
	"time"
)

func TestCetc(t *testing.T) {
	zxl, err := NewZxlImpl("190725000110077", "appTest77")
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
	err = zxl.UpdateUserCert(pk, sk, 0)
	if err != nil {
		t.Error(err.Error())
		return
	}

	hashData := sm3.SumSM3([]byte("123123123"))
	evHash := hex.EncodeToString(hashData)
	fmt.Print("111111")
	result, err := zxl.EvidenceSave(evHash, "abc", sk, pk, 0)
	fmt.Print("2222222")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(result.EvId, result.TxHash, result.CreateTime)
	//
	//time.Sleep(time.Second * 10)
	//queryResult1, err := zxl.QueryWithTxHash(result.TxHash,0)
	//if err != nil || queryResult1[0].EvHash != evHash{
	//	t.Error("QueryWithTxHash error")
	//}
	//
	//queryResult2, err := zxl.QueryWithEvId(queryResult1[0].EvId,0)
	//if err != nil || queryResult2[0].EvHash != evHash{
	//	t.Error("QueryWithEvId error")
	//}
	//queryResult3, err := zxl.QueryWithEvHash(evHash, time.Second*2)
	//if err != nil || queryResult3[0].EvHash != evHash{
	//	t.Error("QueryWithEvHash error")
	//}
}

func TestTencent(t *testing.T) {
	zxl, err := NewZxlImpl("190725000110077", "appTest77")
	if err != nil {
		t.Error(err)
	}
	pk, sk, err := zxl.GenerateKeyPair()
	if err != nil {
		t.Error(err)
	}
	err = zxl.UpdateUserCert(pk, sk, time.Second*2)
	if err != nil {
		t.Error(err.Error())
	}

	hashData := sm3.SumSM3([]byte("123123123"))
	evHash := hex.EncodeToString(hashData)
	result, err := zxl.EvidenceSave(evHash, "abc", sk, pk, time.Second*2)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(result.TxHash, result.CreateTime, result.BlockHeight)

	time.Sleep(time.Second * 10)
	queryResult1, err := zxl.QueryWithTxHash(result.GetTxHash(), time.Second*2)
	if err != nil || queryResult1[0].EvHash != evHash {
		t.Error("QueryWithTxHash error")
	}

	queryResult2, err := zxl.QueryWithEvId(result.GetEvId(), time.Second*2)
	if err != nil || queryResult2[0].EvHash != evHash {
		t.Error("QueryWithEvId error")
	}
	queryResult3, err := zxl.QueryWithEvHash(evHash, time.Second*2)
	if err != nil || queryResult3[0].EvHash != evHash {
		t.Error("QueryWithEvHash error")
	}
}

func TestSign(t *testing.T) {
	zxl, err := NewZxlImpl("xxxxxxxxxx0xxxx", "appTestTX")
	if err != nil {
		t.Error(err)
	}
	pk, sk, err := zxl.GenerateKeyPair()
	fmt.Println(pk)
	fmt.Println(sk)
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
	decryptedData, err := zxl.DecryptData("asdf", encryptedData)
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
func TestStrHash(t *testing.T) {
	zxl, err := NewZxlImpl("190725000110077", "appTestTX")
	hashStr, err := zxl.CalculateStrHash("fwejfoiwfjoweifjowf")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(hashStr)
}

func TestHelloWorld(t *testing.T) {
	fmt.Println("ha")
}

func TestVerify(t *testing.T) {
	zxl, err := NewZxlImpl("xxxxxxxxxx0xxxx", "appTestTX")
	if err != nil {
		t.Error(err)
	}

	var pk = "046bde58352d3360228f71e014d832b00df19c49725eb7ec13986bfce10c54d03e6789cef641961b4a9f69d7bc30c736b60ce62b9674902cff9eecbdb57380341c"
	//var sk = "c31c4b9e1582fe3fa1c9b9ee8be831c410ce243ffb94c263905c91f40d11cc7a"
	//rawData := []byte("3045022100f86e124bc0e8063c62078a8ca1f57a047173b90822133890b000adb6a5f414ee0220262dc4043dffd74d5624d08b3def08bbea1f09d584991c8507754f3084395299")
	rawData := []byte("")
	var signStr = "304502206dc7a6903269024bc0c42ba6a35493f75015c096be6aded8a4e1e46964f6c165022100c020dc22f8059f69f3447fb76190d157f92890bad5a21ffd174f461155eb75af"
	//signStr, err := zxl.Sign(sk, rawData)
	//公钥 签名后数据 签名原数据
	success, err := zxl.Verify(pk, signStr, rawData)
	fmt.Println("begin")
	fmt.Println(success, err)
}
