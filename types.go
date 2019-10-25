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
	AppId string `json:"appId"`
	EvId string `json:"evId"`
	EvHash string `json:"evHash"`
	ExtendInfo string `json:"extendInfo"`
	TxHash string `json:"txHash"`
	BlockHeight int64 `json:"blockHeight"`
	CreateTime string `json:"createTime"`
}

type Config struct {
	ServerAddr string `json:"serverAddr"`
	EvidenceApply string `json:"evidenceApply"`
	EvidenceSubmit string `json:"evidenceSubmit"`
	EvidenceSave string
	QueryWithEvId string `json:"queryWithEvId"`
	QueryWithEvHash string `json:"queryWithEvHash"`
	QueryWithTxHash string `json:"queryWithTxHash"`
	UserCert string `json:"userCert"`
}

type CommonRet struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type UserCertReq struct {
	Pk string `json:"pk"`
	Sign string `json:"sign"`
}

type UserCertResp struct {
	AppId string `json:"appId"`
	BlockHeight string `json:"blockHeight"`
	TxHash string `json:"txHash"`
	CreateTime string `json:"createTime"`
}

type CetcEvidenceReq struct {
	EvId string `json:"evId"`
	EvHash string `json:"evHash"`
	ExtendInfo string `json:"extendInfo"`
	Time string `json:"time"`
	Sign string `json:"sign"`
}

type TencentEvidenceReq struct {
	BodyData string `json:"bodyData"`
	BodySign string `json:"bodySign"`
}

type SystemParams struct {
	MchType int `json:"mch_type"`
	MchId string `json:"mch_id"`
	SignType string `json:"sign_type"`
	TimeStamp int64 `json:"timestamp"`
	Version string `json:"version"`
}

type ApplyContent struct {
	Content string `json:"content"`
	EvId string `json:"ev_id"`
	ExtendInfo string `json:"extend_info"`
}

type EvidenceApplyReq struct {
	SysParams SystemParams `json:"sys_params"`
	Content ApplyContent `json:"content"`
	PublicKey string `json:"public_key"`
}

type EvidenceApplyResp struct {
	ToBeSign string `json:"toBeSign"`
	Session0 string `json:"session0"`
	Session1 string `json:"session1"`
	TxHash string `json:"txHash"`
}

type EvidenceSubmitReq struct {
	SysParams SystemParams `json:"sys_params"`
	Sign string `json:"sign"`
	Session0 string `json:"session0"`
	Session1 string `json:"session1"`
}

type QueryReq struct {
	EvId string `json:"evId"`
	TxHash string `json:"txHash"`
	EvHash string `json:"evHash"`
}

type QueryResp struct {
	AppId string `json:"appId"`
	EvId string `json:"evId"`
	EvHash string `json:"evHash"`
	ExtendInfo string `json:"extendInfo"`
	TxHash string `json:"txHash"`
	BlockHeight int64 `json:"blockHeight"`
	CreateTime string `json:"createTime"`
}

type EvSaveResult interface {
	GetEvId() string
	GetEvHash() string
	GetTxHash() string
	GetCreateTime() string
	GetBlockHeight() int64
}

type TencentEvidenceResp struct {
	BlockHeight int64 `json:"blockHeight"`
	TxHash string `json:"txHash"`
	CreateTime string `json:"createTime"`
	EvId string `json:"evId"`
	EvHash string `json:"evHash"`
}

func (evData *TencentEvidenceResp) GetEvId() string {
	return evData.EvId
}

func (evData *TencentEvidenceResp) GetEvHash() string{
	return evData.EvHash
}

func (evData *TencentEvidenceResp) GetTxHash() string{
	return evData.TxHash
}

func (evData *TencentEvidenceResp) GetCreateTime() string {
	return evData.CreateTime
}
func (evData *TencentEvidenceResp) GetBlockHeight() int64 {
	return evData.BlockHeight
}

type CetcEvidenceResp struct {
	EvId string `json:"evId"`
	TxHash string `json:"txHash"`
	EvHash string `json:""`
	CreateTime string `json:"createTime"`
}

func (evData *CetcEvidenceResp) GetEvId() string {
	return evData.EvId
}

func (evData *CetcEvidenceResp) GetEvHash() string{
	return evData.EvHash
}

func (evData *CetcEvidenceResp) GetTxHash() string{
	return evData.TxHash
}

func (evData *CetcEvidenceResp) GetCreateTime() string {
	return evData.CreateTime
}
func (evData *CetcEvidenceResp) GetBlockHeight() int64 {
	return 0
}
