package zxl_go_sdk

var defConf = Config{
	//ServerAddr: "https://sdk.zxinchain.com",
	//ServerAddr: "https://testsdk.zxchain.net:9087",
	ServerAddr: "https://access.com:9086",
	//ServerAddr:      "https://127.0.0.1:11123",
	EvidenceSave:    "/api/v1/spider/evidence?sdkVersion=go-v2.0.0",
	EvidenceApply:   "/api/v1/spider/evidence/apply?sdkVersion=go-v2.0.0",
	EvidenceSubmit:  "/api/v1/spider/evidence/submit?sdkVersion=go-v2.0.0",
	QueryWithEvId:   "/api/v1/spider/sdk/evidence?sdkVersion=go-v2.0.0&evId=",
	QueryWithTxHash: "/api/v1/spider/sdk/evidence?sdkVersion=go-v2.0.0&txHash=",
	QueryWithEvHash: "/api/v1/spider/sdk/evidence?sdkVersion=go-v2.0.0&evHash=",
	QueryWithHash:   "/api/v1/spider/sdk/evidence?sdkVersion=go-v2.0.0&hash=",
	ProxyHost:       "172.16.2.16",
	ProxyPort:       "10800",
	IsProxy:         false,
	ServerCrtPath:   "",
	//ServerCrtPath:   "d:/certificate/wlpt-private-key.pem",
	UserCert:       "/api/v1/spider/sdk/certificate?sdkVersion=go-v2.0.0",
	ContentCapture: "/api/v1/spider/sdk/req/forward?sdkVersion=go-v2.0.0",
}
