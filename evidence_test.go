package main

import (
	"encoding/hex"
	"fmt"
	"github.com/zhixinlian/zxl-go-sdk/sm/sm3"
	"testing"
	"time"
)

func TestCetc(t *testing.T) {
	//zxl, err := NewZxlImpl("200515000110001","0e4bce1b0ef8471fb9140b849e776f48")
	//if err != nil {
	//	fmt.Println("错误")
	//}
	//result ,err := zxl.EvidenceSave("123213123123123321312312","test","04115802b333e51625853f3f8c0d6117041615689e3a1318a3e8fc05dc50dc44255cc33d2304012cf79f0b0fccf150c40a5af5aca4417d441ae4ec95094d1e0c3a",
	//	"bf8ebcee5542976b7b8a0df6c33862aa09bc274d036bc2007c6190b8a6d20ad4",time.Second*10);
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(result)

	/** 2020-07-24 add start */
	zxl, err := NewZxlImpl("200514000200001","9d7f4ba445a54ed2b041e142d5ea12f3")
	if err != nil {
		fmt.Println("错误")
	}
	//图片取证
	result, err := zxl.EvidenceObtainPic("https://www.baidu.com","图片","go_sdk_test",time.Second*2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
	//根据orderNo查询取证任务状态及结果
	result1, err := zxl.GetEvidenceStatus("1595578460075584005313706",time.Second*2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result1)
	zxl1, err := NewZxlImpl("200515000110001","0e4bce1b0ef8471fb9140b849e776f48")
	if err != nil {
		fmt.Println("错误")
	}
	//下发截屏任务
	result2, err := zxl1.ContentCapturePic("https://www.baidu.com",time.Second*2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result2)
	//根据订单号查询截屏任务的状态及结果
	result4, err:= zxl1.GetContentStatus("1595578461244000782838012",time.Second*2)
	if err != nil {
		fmt.Println("321321",err.Error())
	}
	fmt.Println(result4)
	/** 2020-07-24 add end */
	//zxl, err := NewZxlImpl("190725000110077", "appTest77")
	//if err != nil {
	//	t.Error(err.Error())
	//}
	//pk, sk, err := zxl.GenerateKeyPair()
	//if err != nil {
	//	t.Error(err.Error())
	//}
	////err = zxl.BindUserCert(pk, sk)
	////if err != nil {
	////	t.Error(err)
	////}
	//err = zxl.UpdateUserCert(pk, sk, 0)
	//if err != nil {
	//	t.Error(err.Error())
	//	return
	//}
	//
	//hashData := sm3.SumSM3([]byte("123123123"))
	//evHash := hex.EncodeToString(hashData)
	//fmt.Print("111111")
	//result, err := zxl.EvidenceSave(evHash, "abc", sk, pk, 0)
	//fmt.Print("2222222")
	//if err != nil {
	//	t.Error(err.Error())
	//}
	//
	//fmt.Println(result.EvId, result.TxHash, result.CreateTime)
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
