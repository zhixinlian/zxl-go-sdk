package zxl_go_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
	"github.com/zhixinlian/zxl-go-sdk/constants"
)

/**
行业相关信息
*/
const (
	//农业
	Cid_one = 1
	//林业
	Cid_two = 2
	//畜牧业
	Cid_three = 3
	//渔业
	Cid_four = 4
	//农、林、牧、渔专业及辅助性活动
	Cid_five = 5
	//煤炭开采和洗选业
	Cid_six = 6
	//石油和天然气开采业
	Cid_seven = 7
	//黑色金属矿采选业
	Cid_eight = 8
	//有色金属矿采选业
	Cid_nine = 9
	//非金属矿采选业
	Cid_ten = 10
	//开采专业及辅助性活动
	Cid_eleven = 11
	//其他采矿业
	Cid_twelve = 12
	//农副食品加工业
	Cid_thirteen = 13
	//食品制造业
	Cid_fourteen = 14
	//酒、饮料和精制茶制造业
	Cid_fifteen = 15
	//烟草制品业
	Cid_sixteen = 16
	//纺织业
	Cid_seventeen = 17
	//纺织服装、服饰业
	Cid_eighteen = 18
	//皮革、毛皮、羽毛及其制品和制鞋业
	Cid_nineteen = 19
	//木材加工和木、竹、藤、棕、草制品业
	Cid_twenty = 20
	//家具制造业
	Cid_twenty_one = 21
	//造纸和纸制品业
	Cid_twenty_two = 22
	//印刷和记录媒介复制业
	Cid_twenty_three = 23
	//文教、工美、体育和娱乐用品制造业
	Cid_twenty_four = 24
	//石油、煤炭及其他燃料加工业
	Cid_twenty_five = 25
	//化学原料和化学制品制造业
	Cid_twenty_six = 26
	//医药制造业
	Cid_twenty_seven = 27
	//化学纤维制造业
	Cid_twenty_eight = 28
	//橡胶和塑料制品业
	Cid_twenty_nine = 29
	//非金属矿物制品业
	Cid_thirty = 30
	//黑色金属冶炼和压延加工业
	Cid_thirty_one = 31
	//有色金属冶炼和压延加工业
	Cid_thirty_two = 32
	//金属制品业
	Cid_thirty_three = 33
	//通用设备制造业
	Cid_thirty_four = 34
	//专用设备制造业
	Cid_thirty_five = 35
	//汽车制造业
	Cid_thirty_six = 36
	//铁路、船舶、航空航天和其他运输设备制造业
	Cid_thirty_seven = 37
	//电气机械和器材制造业
	Cid_thirty_eight = 38
	//计算机、通信和其他电子设备制造业
	Cid_thirty_nine = 39
	//仪器仪表制造业
	Cid_forty = 40
	//其他制造业
	Cid_forty_one = 41
	//废弃资源综合利用业
	Cid_forty_two = 42
	//金属制品、机械和设备修理业
	Cid_forty_three = 43
	//电力、热力生产和供应业
	Cid_forty_four = 44
	//燃气生产和供应业
	Cid_forty_five = 45
	//水的生产和供应业
	Cid_forty_six = 46
	//房屋建筑业
	Cid_forty_seven = 47
	//土木工程建筑业
	Cid_forty_eight = 48
	//建筑安装业
	Cid_forty_nine = 49
	//建筑装饰、装修和其他建筑业
	Cid_fifty = 50
	//批发业
	Cid_fifty_one = 51
	//零售业
	Cid_fifty_two = 52
	//铁路运输业
	Cid_fifty_three = 53
	//道路运输业
	Cid_fifty_four = 54
	//水上运输业
	Cid_fifty_five = 55
	//航空运输业
	Cid_fifty_six = 56
	//管道运输业
	Cid_fifty_seven = 57
	//多式联运和运输代理业
	Cid_fifty_eight = 58
	//装卸搬运和仓储业
	Cid_fifty_nine = 59
	//邮政业
	Cid_sixty = 60
	//住宿业
	Cid_sixty_one = 61
	//餐饮业
	Cid_sixty_two = 62
	//电信、广播电视和卫星传输服务
	Cid_sixty_three = 63
	//互联网和相关服务
	Cid_sixty_four = 64
	//软件和信息技术服务业
	Cid_sixty_five = 65
	//货币金融服务
	Cid_sixty_six = 66
	//资本市场服务
	Cid_sixty_seven = 67
	//保险业
	Cid_sixty_eight = 68
	//其他金融业
	Cid_sixty_nine = 69
	//房地产业
	Cid_seventy = 70
	//租赁业
	Cid_seventy_one = 71
	//商务服务业
	Cid_seventy_two = 72
	//研究和试验发展
	Cid_seventy_three = 73
	//专业技术服务业
	Cid_seventy_four = 74
	//科技推广和应用服务业
	Cid_seventy_five = 75
	//水利管理业
	Cid_seventy_six = 76
	//生态保护和环境治理业
	Cid_seventy_seven = 77
	//公共设施管理业
	Cid_seventy_eight = 78
	//土地管理业
	Cid_seventy_nine = 79
	//居民服务业
	Cid_eighty = 80
	//机动车、电子产品和日用产品修理业
	Cid_eighty_one = 81
	//其他服务业
	Cid_eighty_two = 82
	//教育
	Cid_eighty_three = 83
	//卫生
	Cid_eighty_four = 84
	//社会工作
	Cid_eighty_five = 85
	//新闻和出版业
	Cid_eighty_six = 86
	//广播、电视、电影和录音制作业
	Cid_eighty_seven = 87
	//3","18","文化艺术业
	Cid_eighty_eight = 88
	//体育
	Cid_eighty_nine = 89
	//娱乐业
	Cid_ninety = 90
	//中国共产党机关
	Cid_ninety_one = 91
	//国家机构
	Cid_ninety_two = 92
	//人民政协、民主党派
	Cid_ninety_three = 93
	//社会保障
	Cid_ninety_four = 94
	//群众团体、社会团体和其他成员组织
	Cid_ninety_five = 95
	//基层群众自治组织
	Cid_ninety_six = 96
	//国际组织
	Cid_ninety_seven = 97
)

/**
初始化及选择行业相关信息
*/
func initCategory(categoryId int) (cateInfo Category) {
	var cMap = make(map[int]Category)
	cMap[1] = Category{1, 1, "农业"}
	cMap[2] = Category{2, 1, "林业"}
	cMap[3] = Category{3, 1, "畜牧业"}
	cMap[4] = Category{4, 1, "渔业"}
	cMap[5] = Category{5, 1, "农、林、牧、渔专业及辅助性活动"}
	cMap[6] = Category{1, 2, "煤炭开采和洗选业"}
	cMap[7] = Category{2, 2, "石油和天然气开采业"}
	cMap[8] = Category{3, 2, "黑色金属矿采选业"}
	cMap[9] = Category{4, 2, "有色金属矿采选业"}
	cMap[10] = Category{5, 2, "非金属矿采选业"}
	cMap[11] = Category{6, 2, "开采专业及辅助性活动"}
	cMap[12] = Category{7, 2, "其他采矿业"}
	cMap[13] = Category{1, 3, "农副食品加工业"}
	cMap[14] = Category{2, 3, "食品制造业"}
	cMap[15] = Category{3, 3, "酒、饮料和精制茶制造业"}
	cMap[16] = Category{4, 3, "烟草制品业"}
	cMap[17] = Category{5, 3, "纺织业"}
	cMap[18] = Category{6, 3, "纺织服装、服饰业"}
	cMap[19] = Category{7, 3, "皮革、毛皮、羽毛及其制品和制鞋业"}
	cMap[20] = Category{8, 3, "木材加工和木、竹、藤、棕、草制品业"}
	cMap[21] = Category{9, 3, "家具制造业"}
	cMap[22] = Category{10, 3, "造纸和纸制品业"}
	cMap[23] = Category{11, 3, "印刷和记录媒介复制业"}
	cMap[24] = Category{12, 3, "文教、工美、体育和娱乐用品制造业"}
	cMap[25] = Category{13, 3, "石油、煤炭及其他燃料加工业"}
	cMap[26] = Category{14, 3, "化学原料和化学制品制造业"}
	cMap[27] = Category{15, 3, "医药制造业"}
	cMap[28] = Category{16, 3, "化学纤维制造业"}
	cMap[29] = Category{17, 3, "橡胶和塑料制品业"}
	cMap[30] = Category{18, 3, "非金属矿物制品业"}
	cMap[31] = Category{19, 3, "黑色金属冶炼和压延加工业"}
	cMap[32] = Category{20, 3, "有色金属冶炼和压延加工业"}
	cMap[33] = Category{21, 3, "金属制品业"}
	cMap[34] = Category{22, 3, "通用设备制造业"}
	cMap[35] = Category{23, 3, "专用设备制造业"}
	cMap[36] = Category{24, 3, "汽车制造业"}
	cMap[37] = Category{25, 3, "铁路、船舶、航空航天和其他运输设备制造业"}
	cMap[38] = Category{26, 3, "电气机械和器材制造业"}
	cMap[39] = Category{27, 3, "计算机、通信和其他电子设备制造业"}
	cMap[40] = Category{28, 3, "仪器仪表制造业"}
	cMap[41] = Category{29, 3, "其他制造业"}
	cMap[42] = Category{30, 3, "废弃资源综合利用业"}
	cMap[43] = Category{31, 3, "金属制品、机械和设备修理业"}
	cMap[44] = Category{1, 4, "电力、热力生产和供应业"}
	cMap[45] = Category{2, 4, "燃气生产和供应业"}
	cMap[46] = Category{3, 4, "水的生产和供应业"}
	cMap[47] = Category{1, 5, "房屋建筑业"}
	cMap[48] = Category{2, 5, "土木工程建筑业"}
	cMap[49] = Category{3, 5, "建筑安装业"}
	cMap[50] = Category{4, 5, "建筑装饰、装修和其他建筑业"}
	cMap[51] = Category{1, 6, "批发业"}
	cMap[52] = Category{2, 6, "零售业"}
	cMap[53] = Category{1, 7, "铁路运输业"}
	cMap[54] = Category{2, 7, "道路运输业"}
	cMap[55] = Category{3, 7, "水上运输业"}
	cMap[56] = Category{4, 7, "航空运输业"}
	cMap[57] = Category{5, 7, "管道运输业"}
	cMap[58] = Category{6, 7, "多式联运和运输代理业"}
	cMap[59] = Category{7, 7, "装卸搬运和仓储业"}
	cMap[60] = Category{8, 7, "邮政业"}
	cMap[61] = Category{1, 8, "住宿业"}
	cMap[62] = Category{2, 8, "餐饮业"}
	cMap[63] = Category{1, 9, "电信、广播电视和卫星传输服务"}
	cMap[64] = Category{2, 9, "互联网和相关服务"}
	cMap[65] = Category{3, 9, "软件和信息技术服务业"}
	cMap[66] = Category{1, 10, "货币金融服务"}
	cMap[67] = Category{2, 10, "资本市场服务"}
	cMap[68] = Category{3, 10, "保险业"}
	cMap[69] = Category{4, 10, "其他金融业"}
	cMap[70] = Category{1, 11, "房地产业"}
	cMap[71] = Category{1, 12, "租赁业"}
	cMap[72] = Category{2, 12, "商务服务业"}
	cMap[73] = Category{1, 13, "研究和试验发展"}
	cMap[74] = Category{2, 13, "专业技术服务业"}
	cMap[75] = Category{3, 13, "科技推广和应用服务业"}
	cMap[76] = Category{1, 14, "水利管理业"}
	cMap[77] = Category{2, 14, "生态保护和环境治理业"}
	cMap[78] = Category{3, 14, "公共设施管理业"}
	cMap[79] = Category{4, 14, "土地管理业"}
	cMap[80] = Category{1, 15, "居民服务业"}
	cMap[81] = Category{2, 15, "机动车、电子产品和日用产品修理业"}
	cMap[82] = Category{3, 15, "其他服务业"}
	cMap[83] = Category{1, 16, "教育"}
	cMap[84] = Category{1, 17, "卫生"}
	cMap[85] = Category{2, 17, "社会工作"}
	cMap[86] = Category{1, 18, "新闻和出版业"}
	cMap[87] = Category{2, 18, "广播、电视、电影和录音制作业"}
	cMap[88] = Category{3, 18, "文化艺术业"}
	cMap[89] = Category{4, 18, "体育"}
	cMap[90] = Category{5, 18, "娱乐业"}
	cMap[91] = Category{1, 19, "中国共产党机关"}
	cMap[92] = Category{2, 19, "国家机构"}
	cMap[93] = Category{3, 19, "人民政协、民主党派"}
	cMap[94] = Category{4, 19, "社会保障"}
	cMap[95] = Category{5, 19, "群众团体、社会团体和其他成员组织"}
	cMap[96] = Category{6, 19, "基层群众自治组织"}
	cMap[97] = Category{1, 20, "国际组织"}
	if val, ok := cMap[categoryId]; ok {
		return val
	}
	return Category{0, 0, "其他"}
}

/**
行业相关信息
*/
type Category struct {
	//行业子id
	Cid int
	//行业父id
	Pid int
	//行业描述
	Msg string
}

/**上传附件类型*/
const (
	UploadLetter    = "letterId"
	UploadCardFront = "id"
	UploadBus       = "busiLicenseId"
)

/**
代理用户注册所需参数
*/
type AgentUser struct {
	//代理用户注册邮箱
	RepresentEmail string `json:"representEmail"`
	// 密码
	Pwd string `json:"pwd"`
	// 公函
	LetterFile string `json:"letterFile"`
	// 身份证正面
	CardFrontFile string `json:"cardFrontFile"`
	// 身份证反面
	CardBackendFile string `json:"cardBackendFile"`
	// 上午执照
	LicenseFile string `json:"licenseFile"`
	/**法人代表*/
	Representative string `json:"representative"`
	/**企业名称*/
	EpName string `json:"epName"`
	/**企业信用代码*/
	CreditCode string `json:"creditCode"`
	/**企业执照ID*/
	BusiLicenseId int `json:"busiLicenseId"`
	/**公函ID*/
	OfficialLetterId int `json:"officialLetterId"`
	/**身份证正面文件id*/
	IdcardFrontId int `json:"idcardFrontId"`
	/**身份证背面文件id*/
	IdcardBackId int `json:"idcardBackId"`
	/**身份证号码*/
	Idcard string `json:"idcard"`
	/**申请人姓名*/
	Contact string `json:"contact"`
	/**申请人职位*/
	Title string `json:"title"`
	/**申请人手机号*/
	Mobile string `json:"mobile"`
	/**行业id信息**/
	Category int `json:"category"`
	/**自然人姓名**/
	PersonName string `json:"personName"`
	/**用户类型，默认为1：法人主体， 2：自然人主体**/
	UserType constants.UserType `json:"userType"`
	/**接入平台名称：数字+字母，不超过20字**/
	PlatformName string `json:"platformName"`
	/**接入平台地址**/
	PlatformUrl string `json:"platformUrl"`
	/**平台业务类型：
	  1--金融类
	  2--版权类
	  3--其他类
	  4--未填写 default*/
	BusinessType constants.BusinessType `json:"businessType"`
}

/**通用转发请求结构体*/
type CommonReq struct {
	RequestType string `json:"requestType"`
	RedirectUrl string `json:"redirectUrl"`
}

/****
判断字段值是否为空
*/
func judgeIsNull(structName interface{}) bool {
	t := reflect.TypeOf(structName)
	fieldNum := t.NumField()
	msg := ""
	v := reflect.ValueOf(structName)
	for i := 0; i < fieldNum; i++ {
		val := v.FieldByName(t.Field(i).Name)
		msg = t.Field(i).Name
		var errMsg = "the field " + msg + " must be not nil!"
		if t.Field(i).Type.Name() == "string" && len(val.String()) == 0 {
			panic(errMsg)
		}
		if t.Field(i).Type.Name() == "int" && val.Int() == 0 {
			panic(errMsg)
		}
	}
	return true
}

const (
	/**============================================相关请求url===============================================*/
	//代理用户注册url
	SubmitRegisterUser = "sdk/zhixin-api/v2/ump/user/sign_up"
	//上传公函
	UploadLetterAddr = "sdk/zhixin-api/v2/ump/ep/upload_official_letter"
	//上传身份证正面
	UploadCardFrontAddr = "sdk/zhixin-api/v2/ump/ep/upload_idcard_front"
	//上传身份证反面
	UploadCardBackAddr = "sdk/zhixin-api/v2/ump/ep/upload_idcard_back"
	//上传商务执照
	UploadBusAddr = "sdk/zhixin-api/v2/ump/ep/upload_busi_license"
	//提交商务信息
	SubmitBusAddr = "sdk/zhixin-api/v2/ump/ep/apply"
	// 提交个人实名
	SubmitPersonAddr = "sdk/zhixin-api/v2/ump/ep/person_apply"
	//查询企业信息
	GetEpInfoAddr = "sdk/zhixin-api/v1/ump/ep/get"
	/**===========================================相关请求类型================================================*/
	ReqTypeGet    = "GET"
	ReqTypePost   = "POST"
	ReqTypePut    = "PUT"
	ReqTypeDelete = "DELETE"
)

/**代理用户信息注册*/
type AgentUserSubmitInfo struct {
	CommonReq
	Email      string `json:"email"`
	Pwd        string `json:"pwd"`
	AgentAppId string `json:"agentAppId"`
}

/**上传附件**/
type AgentUserUploadFile struct {
	CommonReq
	AgentAppId        string `json:"Agent-App-Id"`
	RepresentAppEmail string `json:"Represent-App-Email"`
}

/**提交商务信息*/
type LicenseUser struct {
	AgentAppId       string `json:"agentAppId"`
	RepresentEmail   string `json:"representEmail"`
	Representative   string `json:"representative"`
	EpName           string `json:"epName"`
	BusiLicenseId    int    `json:"busiLicenseId"`
	OfficialLetterId int    `json:"officialLetterId"`
	IdcardFrontId    int    `json:"idcardFrontId"`
	IdcardBackId     int    `json:"idcardBackId"`
	Idcard           string `json:"idcard"`
	CreditCode       string `json:"creditCode"`
	Contact          string `json:"contact"`
	Title            string `json:"title"`
	Mobile           string `json:"mobile"`
	CategoryCId      int    `json:"categoryCId"`
	CategoryPId      int    `json:"categoryPId"`
	PersonName       string `json:"personName"`
	UserType         constants.UserType    `json:"userType"`
	PlatformName     string `json:"platformName"`
	PlatformUrl      string `json:"platformUrl"`
	BusinessType     constants.BusinessType    `json:"businessType"`
	CommonReq
}

/**审核结果结构体*/
type ReviewData struct {
	AppId     string `json:"appId"`
	AppKey    string `json:"appKey"`
	State     int    `json:"state"`
	Reason    string `json:"reason"`
	AgentCode string `json:"agentCode"`
}

/**
代理用户注册，其包含以下几步
1、账号注册（邮箱+代理商appId+用户设置的密码）
2、上传公函（可选）
3、身份证正面
4、身份证反面
5、上传执照
6、提交整个注册信息
*/
func (zxl *zxlImpl) RegisterUser(info AgentUser, timeout time.Duration) (bool, error) {
	//提交注册信息
	pwd, err := zxl.CalculateStrHash(info.Pwd)
	if err != nil {
		return false, errors.New("calculate password abstract fail:" + err.Error())
	}
	info.Pwd = pwd
	err = submitRegister(zxl.appId, zxl.appKey, info, timeout)
	if err != nil {
		return false, err
	}
	//上传公函
	if len(info.LetterFile) != 0 {
		letterId, err := uploadFile(info, zxl.appId, info.LetterFile, UploadLetterAddr, UploadLetter)
		if err != nil {
			return false, errors.New("upload letter fail : " + err.Error())
		}
		info.OfficialLetterId = letterId
	}
	//上传身份证正面
	frontId, err := uploadFile(info, zxl.appId, info.CardFrontFile, UploadCardFrontAddr, UploadCardFront)
	if err != nil {
		return false, errors.New("upload card front fail : " + err.Error())
	}
	info.IdcardFrontId = frontId
	//上传身份证反面
	backId, err := uploadFile(info, zxl.appId, info.CardBackendFile, UploadCardBackAddr, UploadCardFront)
	if err != nil {
		return false, errors.New("upload card back fail : " + err.Error())
	}
	info.IdcardBackId = backId
	//上传商务执照
	if len(info.LicenseFile) != 0 {
		busiId, err := uploadFile(info, zxl.appId, info.LicenseFile, UploadBusAddr, UploadBus)
		if err != nil {
			return false, errors.New("upload busi fail : " + err.Error())
		}
		info.BusiLicenseId = busiId
    }
	//提交商务信息
	return submitBusInfo(info, zxl.appId, zxl.appKey)
}

/**查询代理用户的审核状态信息*/
func (zxl *zxlImpl) SelectEpInfo(email string, timeout time.Duration) (ReviewData, error) {
	var ret ReviewData
	if len(email) <= 0 {
		return ret, errors.New("please enter the correct email address")
	}
	var param = make(map[string]string)
	param["requestType"] = "POST"
	param["redirectUrl"] = GetEpInfoAddr
	param["representEmail"] = email
	param["agentAppId"] = zxl.appId
	bytes, _ := json.Marshal(param)
	url := defConf.ServerAddr + defConf.ContentCapture
	retBytes, err := sendTxMidRequest(zxl.appId, zxl.appKey, "POST", url, bytes, timeout)
	if err != nil {
		return ret, errors.New(err.Error())
	}
	json.Unmarshal(retBytes, &ret)
	return ret, nil
}

/**绑定代理用户的公私钥*/
func (zxl *zxlImpl) BindRepresentUserCert(representAppId, representAppKey, representPk, representSk string) (bool, error) {
	if len(representSk) <= 0 {
		panic("")
	}
	rawData := strings.Join([]string{representAppId, representPk}, ",")
	signedStr, err := zxl.Sign(representSk, []byte(rawData))
	if err != nil {
		return false, errors.New("BindUserCertError (Sign): " + err.Error())
	}
	certReq := UserCertReq{Pk: representPk, Sign: signedStr, RepresentAppId: representAppId, RepresentAppKey: representAppKey}
	dataBytes, err := json.Marshal(&certReq)
	if err != nil {
		return false, errors.New("BindUserCertError (Marshal): " + err.Error())
	}
	_, err = sendRequest(zxl.appId, zxl.appKey, "POST", defConf.ServerAddr+defConf.UserCert, dataBytes, 0)
	if err != nil {
		return false, errors.New("BindUserCertError (sendRequest): " + err.Error())
	}
	return true, nil
}

/**更新代理用户公私钥*/
func (zxl *zxlImpl) UpdateRepresentUserCert(representAppId, representAppKey, representPk, representSk string) (bool, error) {
	if len(representSk) <= 0 {
		panic("")
	}
	rawData := strings.Join([]string{representAppId, representPk}, ",")
	signedStr, err := zxl.Sign(representSk, []byte(rawData))
	if err != nil {
		return false, errors.New("UpdateUserCertError (Sign): " + err.Error())
	}
	certReq := UserCertReq{Pk: representPk, Sign: signedStr, RepresentAppId: representAppId, RepresentAppKey: representAppKey}
	dataBytes, err := json.Marshal(&certReq)
	if err != nil {
		return false, errors.New("UpdateUserCertError (Marshal): " + err.Error())
	}
	_, err = sendRequest(zxl.appId, zxl.appKey, "PUT", defConf.ServerAddr+defConf.UserCert, dataBytes, 0)
	if err != nil {
		return false, errors.New("UpdateUserCertError (sendRequest): " + err.Error())
	}
	return true, nil
}

/**代理用户上链**/
func (zxl *zxlImpl) RepresentSave(evHash, extendInfo, representSk, representAppId string, timeout time.Duration) (*EvSaveResult, error) {
	uid, err := generateUid()
	if err != nil {
		return nil, errors.New("EvidenceSave (cetc generateUid) error:" + err.Error())
	}
	ed := CetcEvidenceReq{EvId: uid, EvHash: evHash, ExtendInfo: extendInfo, AepresentAppId: representAppId}
	rawStr := []byte(strings.Join([]string{representAppId, ed.EvHash, ed.ExtendInfo, ed.EvId}, ","))
	signStr, err := zxl.Sign(representSk, rawStr)
	if err != nil {
		return nil, errors.New("EvidenceSave (cetc Sign) error:" + err.Error())
	}
	ed.Sign = signStr
	bodyData, _ := json.Marshal(&ed)
	respBytes, err := sendRequest(zxl.appId, zxl.appKey, "POST", defConf.ServerAddr+defConf.EvidenceSave, bodyData, timeout)
	if err != nil {
		return nil, errors.New("EvidenceSave (cetc sendRequest) error:" + err.Error())
	}
	var saveResp EvSaveResult
	err = json.Unmarshal(respBytes, &saveResp)
	if err != nil {
		return nil, errors.New("EvidenceSave (cetc Unmarshal) error:" + err.Error())
	}
	saveResp.EvHash = evHash
	saveResp.EvId = uid
	return &saveResp, nil
}

/**代理用户的注册*/
func submitRegister(appId, appKey string, info AgentUser, timeout time.Duration) error {
	var req = CommonReq{ReqTypePost, SubmitRegisterUser}
	var send = AgentUserSubmitInfo{req, info.RepresentEmail, info.Pwd, appId}
	paramBytes, _ := json.Marshal(send)
	url := defConf.ServerAddr + defConf.ContentCapture
	_, err := sendTxMidRequest(appId, appKey, "POST", url, paramBytes, timeout)
	if err != nil {
		return errors.New("register agent user fail: " + err.Error())
	}
	return nil
}

/**上传附件(返回对应的附件id)**/
func uploadFile(info AgentUser, appId, filename, urlConst, uploadType string) (int, error) {
	var param = make(map[string]string)
	param["requestType"] = "POST"
	param["redirectUrl"] = urlConst
	param["Agent-App-Id"] = appId
	param["Represent-App-Email"] = info.RepresentEmail
	url := defConf.ServerAddr + defConf.ReqFilePath
	upRet, err := postFile(filename, url, param)
	if err != nil {
		return 0, errors.New("send file fail: " + err.Error())
	}
	if retCode, ok := upRet["retCode"]; ok {
		rCode := int(retCode.(float64))
		if rCode != 0 {
			return 0, errors.New("upload file error : " + upRet["retMsg"].(string))
		}
	}
	//TODO 考虑是否需判断else情况
	if _, ok := upRet[uploadType]; ok {
		return int(upRet[uploadType].(float64)), nil
	}
	return 0, errors.New("input right key, the key '" + uploadType + "' was err")
}

/**提交商务信息*/
func submitBusInfo(user AgentUser, appId, appKey string) (bool, error) {
	bytes, _ := json.Marshal(user)
	var a = &LicenseUser{}
	json.Unmarshal(bytes, a)
	a.AgentAppId = appId
	c := initCategory(user.Category)
	a.CategoryCId = c.Cid
	a.CategoryPId = c.Pid
	a.RequestType = "POST"
	// 商务类型为 0 表示没有设置
	if a.BusinessType == 0 {
		a.BusinessType = constants.BUSINESS_DEFAULT
	}
	if user.UserType == constants.USER_NATURAL_PERSON {
		a.RedirectUrl = SubmitPersonAddr
	} else {
		a.RedirectUrl = SubmitBusAddr
	}
	bytes, _ = json.Marshal(a)
	url := defConf.ServerAddr + defConf.ContentCapture
	_, err := sendTxMidRequest(appId, appKey, "POST", url, bytes, 0)
	if err != nil {
		return false, errors.New("submit busi info fail :" + err.Error())
	}
	return true, nil
}

/**附件上传**/
func postFile(filename string, targetUrl string, param map[string]string) (map[string]interface{}, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		return nil, errors.New("create file from fail : " + err.Error())
	}
	//打开文件操作句柄
	fh, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("Open file fail: " + err.Error())
	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, errors.New("copy file fail: " + err.Error())
	}
	if len(param) != 0 {
		for k, v := range param {
			bodyWriter.WriteField(k, v)
		}
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return nil, errors.New("send http fail: " + err.Error())
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("read response fail : " + err.Error())
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("cli.Do error bad status : " + resp.Status)
	}
	var retMap = make(map[string]interface{})
	json.Unmarshal(respBody, &retMap)
	return retMap, nil
}
