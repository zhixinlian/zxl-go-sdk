## zxl-go-sdk
至信链go语言的sdk

## 接口说明
### NewZxlImpl(appId, appKey string) (*zxlImpl, error)
功能：通过应用id和应用密钥生成一个实例，用于用能调用  
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

### BindUserCert(pk, sk string) error 
功能： 绑定用户公钥  
参数：  
pk 用户的公钥  
sk 用户的私钥  
返回值：  
error  
绑定用户公钥是否成功，如果error为nil则绑定成功，如果不为nil则绑定失败

### UpdateUserCert(pk, sk string) error 
功能： 更新用户公钥  
参数：  
pk 用户的公钥  
sk 用户的私钥  
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

### EvidenceSave(evHash, extendInfo, sk, pk string) (interface{}, error)
功能： 保存证据  
参数：   
evHash 证据的哈希信息  
extendInfo 扩展信息  
sk 用户私钥  
pk 用户公钥  
返回值：  
interface{} 存证结果，如果用户的appId类型是腾讯，返回类型是*TencentEvidenceResp；如果用户appId类型是网安，返回类型是*CetcEvidenceResp  
error 存证过程的错误信息，如果为nil说明没有错误 

### QueryWithEvId(evId string) ([]QueryResp, error)  
功能： 通过证据id查找存证数据  
参数：  
evId 证据id  
返回值：  
[]QueryResp 查询结果  
error 查询过程发生的错误信息，如果为nil说明没有发生错误  

### QueryWithTxHash(txHash string) ([]QueryResp, error)  
功能： 通过交易哈希查找存证数据  
参数：  
txHash 交易哈希  
返回值：  
[]QueryResp 查询结果  
error 查询过程发生的错误信息，如果为nil说明没有发生错误  

### QueryWithEvHash(evHash string) ([]QueryResp, error)  
功能： 通过证据哈希查找存证数据  
参数：  
evHash 证据哈希  
返回值：  
[]QueryResp 查询结果  
error 查询过程发生的错误信息，如果为nil说明没有发生错误  

## 下载安装  
### 方法一：go get 命令  
运行命令 go get https://github.com/zhixinlian/zxl-go-sdk  

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

