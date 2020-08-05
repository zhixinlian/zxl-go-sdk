## zxl-go-sdk
至信链go语言的sdk

## 接口说明
### NewZxlImpl(appId, appKey string) (*zxlImpl, error)
功能：通过应用id和应用密钥生成一个实例，用于功能调用  
参数：  
appId 应用id  
appKey 应用密钥  
返回值：  
*zxlImpl 生成的调用实例指针  
error 错误信息  

### GenerateKeyPair() (pk string, sk string, err error)  
功能： 生成公钥和私钥  
参数：无  
返回值：  
pk 公钥字符串  
sk 私钥字符串  
error 生成是否出错，如果不为nil则生成成功，否则生成失败  

### BindUserCert(pk, sk string, timeout time.Duration) error  
功能： 绑定用户公钥  
参数：  
pk 用户的公钥  
sk 用户的私钥  
timeout 调用超时时间，0为不设置  
返回值：  
error  
绑定用户公钥是否成功，如果error为nil则绑定成功，如果不为nil则绑定失败  

### UpdateUserCert(pk, sk string, timeout time.Duration) error 
功能： 更新用户公钥  
参数：  
pk 用户的公钥  
sk 用户的私钥  
timeout 调用超时时间，0为不设置  
返回值：  
error 更新用户公钥是否成功，如果error为nil则绑定成功，如果不为nil则绑定失败  

### EncryptData(pwd string, rawData []byte) (string, error)  
功能： 加密信息    
参数：  
pwd 加密密码  
rawData 需要加密的数据  
返回值：  
string 加密后的字符串  
error 加密的出错信息，如果为nil说明没有错误    

### DecryptData(pwd string, encryptedData string) ([]byte, error)  
功能： 解密信息  
参数：  
pwd 加密密码  
encrypedData 加密后的字符串  
返回值：  
[]byte 解密后的数据内容  
error 解密的出错信息，如果为nil说明没有错误  

### CalculateHash(path string) (string, error)
功能： 通过文件路径，计算文件的哈希值  
参数：  
path 需要计算哈希的文件路径  
返回值：  
string 文件的哈希值  
error 计算过程中的错误信息，如果为nil说明没有错误  

### Sign(sk string, rawData []byte) (string, error)  
功能： 对数据进行签名  
参数：  
sk 私钥字符串  
rawData 待签名数据  
返回值：  
string 签名结果字符串  
error 签名过程中的错误信息，如果为nil说明没有错误  

### Verify(pk, signedStr string, rawData []byte) (bool, error)  
功能： 验证数据的签名是否正确  
参数：  
pk 用户公钥字符串  
signedStr 签名字符串  
rawData 数据内容  
返回值：  
bool 验证是否通过  
error 验证过程是否发生错误，如果为nil说明没有错误  

### EvidenceSave(evHash, extendInfo, sk, pk string, timeout time.Duration) (*EvSaveResult, error)  
功能： 保存证据  
参数：   
evHash 证据的哈希信息  
extendInfo 扩展信息  
sk 用户私钥  
pk 用户公钥  
timeout 调用超时时间，0为不设置  
返回值：  
*EvSaveResult 存证结果
error 存证过程的错误信息，如果为nil说明没有错误 

### QueryWithEvId(evId string, timeout time.Duration) ([]QueryResp, error)  
功能： 通过证据id查找存证数据  
参数：  
evId 证据id  
timeout 调用超时时间，0为不设置  
返回值：  
[]QueryResp 查询结果  
error 查询过程发生的错误信息，如果为nil说明没有发生错误  

### QueryWithTxHash(txHash string, timeout time.Duration) ([]QueryResp, error)  
功能： 通过交易哈希查找存证数据  
参数：  
txHash 交易哈希  
timeout 调用超时时间，0为不设置  
返回值：  
[]QueryResp 查询结果  
error 查询过程发生的错误信息，如果为nil说明没有发生错误  

### QueryWithEvHash(evHash string, timeout time.Duration) ([]QueryResp, error)  
功能： 通过证据哈希查找存证数据  
参数：  
evHash 证据哈希  
timeout 调用超时时间，0为不设置  
返回值：  
[]QueryResp 查询结果  
error 查询过程发生的错误信息，如果为nil说明没有发生错误  

### ContentCaptureVideo(webUrls string, timeout time.Duration) (string, error)

功能： 下发录屏任务  
参数：  
webUrls 需要录屏的url  
timeout 调用超时时间，0为不设置  
返回值：  
string 订单号  
error 查询过程发生的错误信息，如果为nil说明没有发生错误 

### ContentCapturePic(webUrls string, timeout time.Duration) (string, error)

功能： 下发截屏任务  
参数：  
webUrls 需要截屏的url  
timeout 调用超时时间，0为不设置  
返回值：  
string 订单号  
error 查询过程发生的错误信息，如果为nil说明没有发生错误 

### GetContentStatus(orderNo string, timeout time.Duration) (*TaskEvData, error)

功能： 获取截屏/录屏任务状态及结果  
参数：  
orderNo 订单号  
timeout 调用超时时间，0为不设置  
返回值：  
*TaskEvData 任务执行状态信息	{status:0[运行中]|2[成功]|10[失败],statusMsg:任务状态解读[运行中]>>[运行成功]>>[运行失败],url:状态成功时,对应的cosurl,hash:截图成功时,对应的存证hash}  
error 查询过程发生的错误信息，如果为nil说明没有发生错误 

### EvidenceObtainVideo(webUrls, title, remark string, timeout time.Duration) (string, error)

功能： 视频取证接口  
参数：  
webUrls 需要视频取证的，必填  
title 标题，必填  
remark 备注，  
timeout 调用超时时间，0为不设置  
返回值：  
string 订单号  
error 查询过程发生的错误信息，如果为nil说明没有发生错误 

### EvidenceObtainPic(webUrls, title, remark string, timeout time.Duration) (string, error)

功能： 图片取证接口  
参数：  
webUrls 需要视频取证的，必填  
title 标题，必填  
remark 备注，  
timeout 调用超时时间，0为不设置  
返回值：  
string 订单号  
error 查询过程发生的错误信息，如果为nil说明没有发生错误 

### GetEvidenceStatus(orderNo string, timeout time.Duration) (*EvIdData, error)

功能： 获取取证证书任务状态及结果接口  
参数：  
orderNo 取证订单号，必填  
timeout 调用超时时间，0为不设置  
返回值：  
*EvIdData 取证任务状态及结果 	｛status:0[运行中]|2[任务完成]|10[失败],evidUrl:取证证据下载地址[当状态为1时],voucherUrl:取证证书下载地址[[当状态为1时]｝  
error 查询过程发生的错误信息，如果为nil说明没有发生错误 



## 下载安装  
### 方法一：go get 命令  
运行命令 go get [github.com/zhixinlian/zxl-go-sdk](https://github.com/zhixinlian/zxl-go-sdk)  

### 方法二：git 命令  
1. 进入到go根目录或者当前项目的vendor目录中的github.com/zhixinlian目录下  
2. 运行命令 git clone https://github.com/zhixinlian/zxl-go-sdk.git  

### 方法三：手动下载
1. 进入网页https://github.com/zhixinlian/zxl-go-sdk  
2. 下载zip压缩版  
3. 解压到go根目录或者当前项目vendor目录中的github.com/zhixinlian/zxl-go-sdk目录中  

## 使用示例
```
package main

import (
	"fmt"
	zxl "github.com/zhixinlian/zxl-go-sdk"
)

func main() {
	//初始化应用
	zxl, err := zxl.NewZxlImpl("123456000110000", "appkey")
	if err != nil {
		panic(err)
	}

	//生成公私钥对
	pk, sk, err := zxl.GenerateKeyPair()
	if err != nil {
		panic(err)
	}
	fmt.Println("公钥:", pk)
	fmt.Println("私钥:", sk)

	//计算文件hash
	hashStr, err := zxl.CalculateHash("G:\\channel.zip")
	if err != nil {
		panic(err.Error())
	}
	fmt.Print(hashStr)
}

```

