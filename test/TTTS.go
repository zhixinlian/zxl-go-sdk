package main

import (
	"crypto/rand"
	"fmt"
	"github.com/zhixinlian/zxl-go-sdk/sm/sm2"
	"os"
)

func main() {
	msg := []byte{1,2,3,4}
	for i := 0; i < 1; i++ {
		privateKey, _ := sm2.GenerateKey(rand.Reader)
		epk := sm2.EncodePubKey(&privateKey.PublicKey)
		esk := sm2.EncodePrivKey(privateKey)

		fmt.Println("epk:",epk)
		fmt.Println("spk",esk)
		pk, _ :=sm2.DecodePubKey(epk)
		sk, _ :=sm2.DecodePrivKey(esk)

		sign,err:=sk.Sign(rand.Reader,msg,nil)

		if err!=nil{
			fmt.Println("err",i)
			os.Exit(0)
		}
		b:=pk.Verify(msg,sign)
		if !b {
			fmt.Println(b,i)
			os.Exit(0)
		}

	}
	fmt.Println("over")

}

