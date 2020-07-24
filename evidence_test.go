package zxl_go_sdk

import (
	"encoding/hex"
	"fmt"
	"github.com/zhixinlian/zxl-go-sdk/sm/sm3"
	"testing"
	"time"
)

func TestCetc(t *testing.T) {
	//zxl, err := NewZxlImpl("200514000200001","9d7f4ba445a54ed2b041e142d5ea12f3")
	//if err != nil {
	//	fmt.Println("错误")
	//}
	////图片取证
	//result, err := zxl.evidenceObtainPic("https://www.baidu.com","图片","go_sdk_test",time.Second*2)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(result)
	////根据orderNo查询取证任务状态及结果
	//result1, err := zxl.getEvidenceStatus("1595503953240243004615586",time.Second*2)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(result1)
	//zxl1, err := NewZxlImpl("200515000110001","0e4bce1b0ef8471fb9140b849e776f48")
	//if err != nil {
	//	fmt.Println("错误")
	//}
	////下发截屏任务
	//result2, err := zxl1.ContentCapturePic("https://www.baidu.com",time.Second*2)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(result2)
	////根据订单号查询截屏任务的状态及结果
	//result4, err:= zxl1.getContentStatus("1595503953240243004615586",time.Second*2)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(result4)

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
	err = zxl.UpdateUserCert(pk, sk, 0)
	if err != nil {
		t.Error(err.Error())
		return
	}

	hashData := sm3.SumSM3([]byte("123123123"))
	evHash := hex.EncodeToString(hashData)
	result, err := zxl.EvidenceSave(evHash, "abc", sk, pk, 0)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(result.EvId, result.TxHash, result.CreateTime)

	time.Sleep(time.Second * 10)
	queryResult1, err := zxl.QueryWithTxHash(result.TxHash, 0)
	if err != nil || queryResult1[0].EvHash != evHash {
		t.Error("QueryWithTxHash error")
	}

	queryResult2, err := zxl.QueryWithEvId(queryResult1[0].EvId, 0)
	if err != nil || queryResult2[0].EvHash != evHash {
		t.Error("QueryWithEvId error")
	}
	queryResult3, err := zxl.QueryWithEvHash(evHash, time.Second*2)
	if err != nil || queryResult3[0].EvHash != evHash {
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
