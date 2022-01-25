package test

import (
	"fmt"
	"testing"
	"time"
	zxl_go_sdk "github.com/zhixinlian/zxl-go-sdk/v2"
	"github.com/zhixinlian/zxl-go-sdk/v2/constants"
)


func TestDciClaim(t *testing.T) {

	appId = ""
	appKey = ""
	pk = ""
	sk = ""

	zxlSDK, err := zxl_go_sdk.NewZxlImpl(appId, appKey)
	if err != nil {
		fmt.Println("初始化 SDK 错误")
	}

	righter := zxl_go_sdk.DciRighter{
		RighterName: "",
		RighterIdCard: "",
		RighterEmail: "",
		RighterGainedWay: constants.GAINED_WAY_ORIGINAL,
		RighterSk: sk,
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
		DciName: "图片作品1",
		DciUrl: "https://sports.sina.com.cn/basketball/nba/2021-02-19/doc-ikftssap7045257.shtml",
		ProposerEmail: "",
		ProposerSk: "",
		DciType: constants.DCI_TYPE_FILMING,
		DciCreateProperty: constants.DCI_CREATE_PROPERTY_ADAPT,
		RightInfoList: []zxl_go_sdk.DciRight{right},
		AuthorList: []zxl_go_sdk.DciAuthor{author},
	}
	resp, err := zxlSDK.SubmitDciClaim(dciClaim, 10 * time.Second)

	if err != nil {
		fmt.Printf("提交确权请求出错 %+v", err)
		return
	}

	fmt.Printf("确权结果为 %+v", resp)
}




func TestDciClaimTwoAuthor(t *testing.T) {

	appId = ""
	appKey = ""
	pk = ""
	sk = ""

	zxlSDK, err := zxl_go_sdk.NewZxlImpl(appId, appKey)
	if err != nil {
		fmt.Println("初始化 SDK 错误")
	}

	righter := zxl_go_sdk.DciRighter{
		RighterName: "",
		RighterIdCard: "",
		RighterEmail: "",
		RighterGainedWay: constants.GAINED_WAY_ORIGINAL,
		RighterSk: sk,
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

	author1 := zxl_go_sdk.DciAuthor{
		AuthorName: "",
		AuthorIdCard: "",
		AuthorType: constants.AUTHOR_TYPE_LEGAL,
	}

	dciClaim := zxl_go_sdk.DciClaim{
		DciName: "图片作品1",
		DciUrl: "https://k.sina.com.cn/article_1887344341_707e96d50200110wn.html?from=news&subch=onews",
		ProposerEmail: "",
		ProposerSk: "",
		DciType: constants.DCI_TYPE_FILMING,
		DciCreateProperty: constants.DCI_CREATE_PROPERTY_ADAPT,
		RightInfoList: []zxl_go_sdk.DciRight{right},
		AuthorList: []zxl_go_sdk.DciAuthor{author, author1},
	}
	resp, err := zxlSDK.SubmitDciClaim(dciClaim, 10 * time.Second)

	if err != nil {
		fmt.Printf("提交确权请求出错 %+v", err)
		return
	}

	fmt.Printf("确权结果为 %+v", resp)
}



func TestDciClaimTwoRighter(t *testing.T) {

	appId = ""
	appKey = ""
	pk = ""
	sk = ""

	zxlSDK, err := zxl_go_sdk.NewZxlImpl(appId, appKey)
	if err != nil {
		fmt.Println("初始化 SDK 错误")
	}

	righter := zxl_go_sdk.DciRighter{
		RighterName: "",
		RighterIdCard: "",
		RighterEmail: "",
		RighterGainedWay: constants.GAINED_WAY_ORIGINAL,
		RighterSk: sk,
		RighterType: constants.RIGHTER_TYPE_LEGAL,
	}

	righter2 := zxl_go_sdk.DciRighter{
		RighterName: "",
		RighterIdCard: "",
		RighterEmail: "",
		RighterGainedWay: constants.GAINED_WAY_ORIGINAL,
		RighterSk: sk,
		RighterType: constants.RIGHTER_TYPE_LEGAL,
	}

	right := zxl_go_sdk.DciRight{
		Type: constants.RIGHT_TYPE_ALL,
		RighterInfoList: []zxl_go_sdk.DciRighter{righter, righter2},
	}

	author := zxl_go_sdk.DciAuthor{
		AuthorName: "",
		AuthorIdCard: "",
		AuthorType: constants.AUTHOR_TYPE_LEGAL,
	}

	dciClaim := zxl_go_sdk.DciClaim{
		DciName: "图片作品1",
		DciUrl: "https://news.sina.com.cn/o/2021-02-13/doc-ikftssap5618045.shtml",
		ProposerEmail: "",
		ProposerSk: "",
		DciType: constants.DCI_TYPE_FILMING,
		DciCreateProperty: constants.DCI_CREATE_PROPERTY_ADAPT,
		RightInfoList: []zxl_go_sdk.DciRight{right},
		AuthorList: []zxl_go_sdk.DciAuthor{author},
	}
	resp, err := zxlSDK.SubmitDciClaim(dciClaim, 10 * time.Second)

	if err != nil {
		fmt.Printf("提交确权请求出错 %+v", err)
		return
	}

	fmt.Printf("确权结果为 %+v", resp)
}

func TestDciClaimTortSearchNormal(t *testing.T) {

	appId = ""
	appKey = ""
	pk = ""
	sk = ""

	zxlSDK, err := zxl_go_sdk.NewZxlImpl(appId, appKey)
	if err != nil {
		fmt.Println("初始化 SDK 错误")
	}

	righter := zxl_go_sdk.DciRighter{
		RighterName: "",
		RighterIdCard: "",
		RighterEmail: "",
		RighterGainedWay: constants.GAINED_WAY_ORIGINAL,
		RighterSk: sk,
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
		DciName: "图片作品1",
		DciUrl: "https://pic2.zhimg.com/80/v2-73c0a03f0aead285983ade7764e32225_1440w.jpg",
		ProposerEmail: "",
		ProposerSk: "",
		DciType: constants.DCI_TYPE_PIC_MODEL,
		DciCreateProperty: constants.DCI_CREATE_PROPERTY_ORIGINAL,
		RightInfoList: []zxl_go_sdk.DciRight{right},
		AuthorList: []zxl_go_sdk.DciAuthor{author},
	}
	resp, err := zxlSDK.SubmitDciClaim(dciClaim, 10 * time.Second)

	if err != nil {
		fmt.Printf("提交确权请求出错 %+v", err)
		return
	}

	fmt.Printf("确权结果为 %+v", resp)
}

func TestDciClaimTortSearch(t *testing.T) {

	appId = ""
	appKey = ""
	pk = ""
	sk = ""

	zxlSDK, err := zxl_go_sdk.NewZxlImpl(appId, appKey)
	if err != nil {
		fmt.Println("初始化 SDK 错误")
	}

	righter := zxl_go_sdk.DciRighter{
		RighterName: "",
		RighterIdCard: "",
		RighterEmail: "",
		RighterGainedWay: constants.GAINED_WAY_ORIGINAL,
		RighterSk: sk,
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
		DciName: "图片作品1",
		DciUrl: "https://inews.gtimg.com/newsapp_bt/0/5001rcns97nr04er/1000?appid=ee22ce76657290e1",
		ProposerEmail: "",
		ProposerSk: "",
		DciType: constants.DCI_TYPE_PIC_MODEL,
		DciCreateProperty: constants.DCI_CREATE_PROPERTY_ORIGINAL,
		RightInfoList: []zxl_go_sdk.DciRight{right},
		AuthorList: []zxl_go_sdk.DciAuthor{author},
	}
	resp, err := zxlSDK.SubmitDciClaim(dciClaim, 10 * time.Second)

	if err != nil {
		fmt.Printf("提交确权请求出错 %+v", err)
		return
	}

	fmt.Printf("确权结果为 %+v", resp)
}



func TestDciClaimParamNotEnough(t *testing.T) {

	appId = ""
	appKey = ""
	pk = ""
	sk = ""

	zxlSDK, err := zxl_go_sdk.NewZxlImpl(appId, appKey)
	if err != nil {
		fmt.Println("初始化 SDK 错误")
	}

	righter := zxl_go_sdk.DciRighter{
		RighterName: "",
		RighterIdCard: "",
		RighterEmail: "",
		RighterGainedWay: constants.GAINED_WAY_ORIGINAL,
		RighterSk: sk,
		RighterType: constants.RIGHTER_TYPE_LEGAL,
	}

	right := zxl_go_sdk.DciRight{
		Type: constants.RIGHT_TYPE_ALL, // 当前只能使用这个权利项
		RighterInfoList: []zxl_go_sdk.DciRighter{righter},
	}

	dciClaim := zxl_go_sdk.DciClaim{
		DciName: "图片作品1",
		DciUrl: "https://inews.gtimg.com/newsapp_bt/0/5001rcns97nr04er/1000?appid=ee22ce76657290e1",
		ProposerEmail: "",
		ProposerSk: "",
		DciType: constants.DCI_TYPE_PIC_MODEL,
		DciCreateProperty: constants.DCI_CREATE_PROPERTY_ORIGINAL,
		RightInfoList: []zxl_go_sdk.DciRight{right},
	}
	resp, err := zxlSDK.SubmitDciClaim(dciClaim, 10 * time.Second)

	if err != nil {
		fmt.Printf("提交确权请求出错 %+v", err)
		return
	}

	fmt.Printf("确权结果为 %+v", resp)
}

func TestQueryResult(t *testing.T) {

	appId = ""
	appKey = ""

	zxlSDK, err := zxl_go_sdk.NewZxlImpl(appId, appKey)
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
