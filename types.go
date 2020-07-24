package zxl_go_sdk

type SaveResult struct {
	//链上存证的hash
	Hash string `json:"hash"`
	//存证区块的高度
	Height int64 `json:"height"`
	//存证区块的时间
	Time string `json:"time"`
}

type EvidenceData struct {
	AppId       string `json:"appId"`
	EvId        string `json:"evId"`
	EvHash      string `json:"evHash"`
	ExtendInfo  string `json:"extendInfo"`
	TxHash      string `json:"txHash"`
	BlockHeight int64  `json:"blockHeight"`
	CreateTime  string `json:"createTime"`
}

type Config struct {
	ServerAddr      string `json:"serverAddr"`
	EvidenceApply   string `json:"evidenceApply"`
	EvidenceSubmit  string `json:"evidenceSubmit"`
	EvidenceSave    string
	QueryWithEvId   string `json:"queryWithEvId"`
	QueryWithEvHash string `json:"queryWithEvHash"`
	QueryWithTxHash string `json:"queryWithTxHash"`
	UserCert        string `json:"userCert"`
	ContentCapture  string `json:"contentCapture"`
}

type CommonRet struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserCertReq struct {
	Pk   string `json:"pk"`
	Sign string `json:"sign"`
}

type UserCertResp struct {
	AppId       string `json:"appId"`
	BlockHeight string `json:"blockHeight"`
	TxHash      string `json:"txHash"`
	CreateTime  string `json:"createTime"`
}

type CetcEvidenceReq struct {
	EvId       string `json:"evId"`
	EvHash     string `json:"evHash"`
	ExtendInfo string `json:"extendInfo"`
	Time       string `json:"time"`
	Sign       string `json:"sign"`
}

type TencentEvidenceReq struct {
	BodyData string `json:"bodyData"`
	BodySign string `json:"bodySign"`
}

type SystemParams struct {
	MchType   int    `json:"mch_type"`
	MchId     string `json:"mch_id"`
	SignType  string `json:"sign_type"`
	TimeStamp int64  `json:"timestamp"`
	Version   string `json:"version"`
}

type ApplyContent struct {
	Content    string `json:"content"`
	EvId       string `json:"ev_id"`
	ExtendInfo string `json:"extend_info"`
}

type EvidenceApplyReq struct {
	SysParams SystemParams `json:"sys_params"`
	Content   ApplyContent `json:"content"`
	PublicKey string       `json:"public_key"`
}

type EvidenceApplyResp struct {
	ToBeSign string `json:"toBeSign"`
	Session0 string `json:"session0"`
	Session1 string `json:"session1"`
	TxHash   string `json:"txHash"`
}

type EvidenceSubmitReq struct {
	SysParams SystemParams `json:"sys_params"`
	Sign      string       `json:"sign"`
	Session0  string       `json:"session0"`
	Session1  string       `json:"session1"`
}

type QueryReq struct {
	EvId   string `json:"evId"`
	TxHash string `json:"txHash"`
	EvHash string `json:"evHash"`
}

type QueryResp struct {
	AppId       string `json:"appId"`
	EvId        string `json:"evId"`
	EvHash      string `json:"evHash"`
	ExtendInfo  string `json:"extendInfo"`
	TxHash      string `json:"txHash"`
	BlockHeight int64  `json:"blockHeight"`
	CreateTime  string `json:"createTime"`
}

type EvSaveResult struct {
	BlockHeight int64  `json:"blockHeight"`
	EvId        string `json:"evId"`
	TxHash      string `json:"txHash"`
	EvHash      string `json:"evHash"`
	CreateTime  string `json:"createTime"`
}

//下发截屏/录屏任务
type EvObtainTask struct {
	WebUrls     string `json:"webUrls"`
	Type        int    `json:"type"`
	AppId       string `json:"appId"`
	RequestType string `json:"requestType"`
	RedirectUrl string `json:"redirectUrl"`
	OrderNo     string `json:"orderNo"`
	Title       string `json:"title"`
	Remark      string `json:"remark"`
}

//任务Data
type TaskEvData struct {
	status    int    `json:"status"`
	statusMsg string `json:"statusMsg"`
	url       string `json:"url"`
	hash      string `json:"hash"`
}

//证书状态及结果
type EvIdData struct {
	status     int    `json:"status"`
	evidUrl    string `json:"evidUrl"`
	voucherUrl string `json:"voucherUrl"`
}

//定义tx接口返回结构体，最外层
type TxRetCommonData struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Detail  TxRetDetail `json:"detail"`
}

//定义tx接口返回结构体,最内层
type TxRetDetail struct {
	OrderNo    string `json:"orderNo"`
	Status     int    `json:"status"`
	StatusMsg  string `json:"statusMsg"`
	Url        string `json:"url"`
	Hash       string `json:"hash"`
	EvIdUrl    string `json:"evidUrl"`
	VoucherUrl string `json:"voucherUrl"`
}

func (evData *EvSaveResult) GetEvId() string {
	return evData.EvId
}

func (evData *EvSaveResult) GetEvHash() string {
	return evData.EvHash
}

func (evData *EvSaveResult) GetTxHash() string {
	return evData.TxHash
}

func (evData *EvSaveResult) GetCreateTime() string {
	return evData.CreateTime
}
func (evData *EvSaveResult) GetBlockHeight() int64 {
	return evData.BlockHeight
}