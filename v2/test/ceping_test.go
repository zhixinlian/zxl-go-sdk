/*
 Package test test-ceping_test
    - @File:  ceping_test
    - @Description: // 清洁性测评专用
    - @Author: suxiongye
    - @Date: 2024/1/18 15:53
    - @Copyright: Tencent
*/

package test

import (
	"flag"
	"fmt"
	zxl_go_sdk "github.com/zhixinlian/zxl-go-sdk/v2"
	"testing"
)

func TestGetMobileObtainResult(t *testing.T) {
	if !flag.Parsed() {
		flag.Parse()
	}
	var argLen = 3
	argList := flag.Args() // flag.Args() 返回 -args 后面的所有参数，以切片表示，每个元素代表一个参数
	if len(argList) < argLen {
		t.Fatalf("缺少参数，参数数量要求至少为：%d 个", argLen)
	}
	// 测试环境
	var serverAddr = "https://testsdk.zxinchain.com"
	var env = argList[0]
	if env == "prod" {
		serverAddr = "https://sdk.zxinchain.com"
	}
	// 代理商
	var appid = argList[1]
	var appKey = argList[2]

	var config = &zxl_go_sdk.ZxlConfig{
		AppId:      appid,
		AppKey:     appKey,
		ServerAddr: serverAddr,
	}
	zxlSdk, err := zxl_go_sdk.CreateZxlClientWithConfig(*config)

	evIdData, err := zxlSdk.GetEvidenceStatus("1703558435342809048534086", 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("evIdData is %v\n", evIdData)
}
