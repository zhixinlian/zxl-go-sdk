package zxl_go_sdk


var defConf = Config{
  ServerAddr: "https://sdk.zhixinblockchain.com",
  EvidenceSave: "/api/v1/spider/evidence?sdkVersion=go-v1.0.0",
  EvidenceApply: "/api/v1/spider/evidence/apply?sdkVersion=go-v1.0.0",
  EvidenceSubmit: "/api/v1/spider/evidence/submit?sdkVersion=go-v1.0.0",
  QueryWithEvId: "/api/v1/spider/sdk/evidence?sdkVersion=go-v1.0.0&evId=",
  QueryWithTxHash: "/api/v1/spider/sdk/evidence?sdkVersion=go-v1.0.0&txHash=",
  QueryWithEvHash: "/api/v1/spider/sdk/evidence?sdkVersion=go-v1.0.0&evHash=",
  UserCert: "/api/v1/spider/sdk/certificate?sdkVersion=go-v1.0.0",
}
