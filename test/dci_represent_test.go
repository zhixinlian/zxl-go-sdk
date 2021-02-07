package test

import (
	"fmt"
	"testing"
	"time"
	"github.com/zhixinlian/zxl-go-sdk"
	"github.com/zhixinlian/zxl-go-sdk/constants"
)

// 代理商确权
func TestDciClaimAgent(t *testing.T) {

	agentAppId := ""
	agentAppKey := ""
	//agentSK := ""
	//agentPK :=
	//	""

	appId = ""
	appKey = ""
	pk = ""
	sk = ""

	zxlSDK, err := zxl_go_sdk.NewZxlImpl(agentAppId, agentAppKey)
	if err != nil {
		fmt.Println("初始化 SDK 错误")
	}

	righter := zxl_go_sdk.DciRighter{
		RighterName: "",
		RighterIdCard: "",
		RighterEmail: "",
		RighterGainedWay: constants.GAINED_WAY_ORIGINAL,
		Sk: sk,
		RighterType: constants.RIGHTER_TYPE_LEGAL,
	}

	right := zxl_go_sdk.DciRight{
		Type: constants.RIGHT_TYPE_ALL,
		RighterInfoList: []zxl_go_sdk.DciRighter{righter},
	}

	author := zxl_go_sdk.DciAuthor{
		AuthorName: "",
		AuthorIdCard: "",
		AuthorType: constants.AUTHOR_TYPE_LEGAL,
	}

	dciClaim := zxl_go_sdk.DciClaim{
		RepresentAppId: appId,
		DciName: "图片作品1",
		DciUrl: "https://news.sina.com.cn/w/2021-02-14/doc-ikftpnny6679806.shtml",
		ProposerEmail: "",
		DciType: constants.DCI_TYPE_FILMING,
		DciCreateProperty: constants.DCI_CREATE_PROPERTY_ADAPT,
		DciCreateTime: "2021-02-10 12:59:59",
		RightInfoList: []zxl_go_sdk.DciRight{right},
		AuthorList: []zxl_go_sdk.DciAuthor{author},
	}
	resp, err := zxlSDK.SubmitDciClaim(dciClaim, sk, 10 * time.Second)

	if err != nil {
		fmt.Printf("提交确权请求出错 %+v", err)
		return
	}

	fmt.Printf("确权结果为 %+v", resp)
}

func TestQueryRepresentResult(t *testing.T) {

	agentAppId := ""
	agentAppKey := ""

	zxlSDK, err := zxl_go_sdk.NewZxlImpl(agentAppId, agentAppKey)
	if err != nil {
		fmt.Println("初始化 SDK 错误")
	}

	dciQuery := zxl_go_sdk.DciClaimQuery{
		TaskId: "",
	}

	resp, err := zxlSDK.QueryDciClaimResult(dciQuery, 10 * time.Second)
	if err != nil {
		fmt.Printf("确权查询出错 %+v", err)
		return
	}
	fmt.Printf("确权查询 %+v", resp)
}
