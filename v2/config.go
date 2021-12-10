package zxl_go_sdk

var defConf = Config{
	ServerAddr: "https://sdk.zxinchain.com",
	//ServerAddr:      "http://127.0.0.1:7082",
	EvidenceSave:    "/api/v1/spider/evidence?sdkVersion=go-v2.0.4",
	EvidenceApply:   "/api/v1/spider/evidence/apply?sdkVersion=go-v2.0.4",
	EvidenceSubmit:  "/api/v1/spider/evidence/submit?sdkVersion=go-v2.0.4",
	QueryWithEvId:   "/api/v1/spider/sdk/evidence?sdkVersion=go-v2.0.4&evId=",
	QueryWithTxHash: "/api/v1/spider/sdk/evidence?sdkVersion=go-v2.0.4&txHash=",
	QueryWithEvHash: "/api/v1/spider/sdk/evidence?sdkVersion=go-v2.0.4&evHash=",
	QueryWithHash:   "/api/v1/spider/sdk/evidence?sdkVersion=go-v2.0.4&hash=",
	ProxyHost:       "172.16.2.16",
	ProxyPort:       "10800",
	IsProxy:         false,
	ServerCrtPath:   "",
	//ServerCrtPath:   "./vendor/cert/access-test.crt",
	//ServerCrtPath:   "d:/certificate/wlpt-private-key.pem",
	UserCert:       "/api/v1/spider/sdk/certificate?sdkVersion=go-v2.0.4",
	ContentCapture: "/api/v1/spider/sdk/req/forward?sdkVersion=go-v2.0.4",
	ReqFilePath:    "/api/v1/spider/sdk/req/file/forward?sdkVersion=go-v2.0.4",
	DefaultHttpTimeout: 10,
}
