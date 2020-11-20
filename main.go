package main

import (
	"CryptCode/aes"
	"CryptCode/des"
	"CryptCode/ecc"
	"CryptCode/rsa"
	"CryptCode/threedes"
	"crypto"
	"fmt"
)


//
func main() {
	key := []byte("12345678")//des秘钥长度：8字节
	data := "窗含西岭千秋雪呀哈喽"
	cipherText, err := des.DESEnCrypt([]byte(data),key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	originText,err :=des.DESDeCrypt(cipherText,key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("DES解密后的内容：",string(originText))

	//3DES加解密
	key1 := []byte("123456781234567812345678")//3des的秘钥长度是24字节
	//key1 := make([]byte,16)
	//key1 = append(key1,[]byte("12345678")...)
	data1 := "穷在闹市无人问，富在深山有远亲"
	cipherText1, err :=threedes.TripleDESEncrypt([]byte(data1),key1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	originText1, err :=threedes.TripleDESDecrypt(cipherText1,key1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("3DES解密后的内容:",string(originText1))

	//三,AES算法
	key2 := []byte("1234567812345678")
	data2 := "汝与曹贼何异"

	cipherText2,err:= aes.AESEnCrypt([]byte(data2),key2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	originText2,err := aes.AESDecrypt(cipherText2,key2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(originText2))





	//RSA算法的调用
	fmt.Println("=====================RSA算法=======================")
	data3 := "喜爱人妻,宠溺幼女,汝与曹贼何异之有！！！"
	privateKey,err := rsa.GenerateKey("憨憨")
	if err != nil {
		fmt.Println("生成密钥对出错",err.Error())
		return
	}

	cipherText3,err:=rsa.RSAEncrypt(privateKey.PublicKey,[]byte(data3))
	if err != nil {
		fmt.Println("加密出错",err.Error())
		return
	}
	originText3,err:=rsa.RSADEcrypt(privateKey,cipherText3)
	if err != nil {
		fmt.Println("解密出错")
		return
	}
	fmt.Println(string(originText3))

	signText,err := rsa.RSASign(privateKey,[]byte(data3),crypto.MD5)
	if err != nil {
		fmt.Println("签名出错")
		return
	}
	verifyResult,err:=rsa.RSAVerify(privateKey.PublicKey,signText,[]byte(data3),crypto.MD5)
	if err != nil {
		fmt.Println("验证出错")
		return
	}
	if verifyResult {
		fmt.Println("签名验证成功")
	}else {
		fmt.Println("签名验证失败")
	}

	//从pem文件中读取私钥和公钥
	fmt.Println("=====================Pem格式读取=================")
	priKey,_:=rsa.ReadPemPriKey("憨憨")
	pubKey,_:=rsa.ReadPemPubKey("憨憨")
	//用读取的公钥进行加密
	data5:="李白"
	cipherText4,_:=rsa.RSAEncrypt(*pubKey,[]byte(data5))
	originText4,_:=rsa.RSADEcrypt(priKey,cipherText4)
	fmt.Println(string(originText4))



	fmt.Println("===========================ECC算法===========================")
	data4 := "喜爱人妻,宠溺幼女,汝与曹贼何异之有！！！"
	pri,err:=ecc.GenerateECDSAKey()
	if err != nil {
		fmt.Println("生成ecc密钥出错")
		return
	}
	r,s,err:=ecc.ECDSASign(pri,[]byte(data4))
	fmt.Printf("%x\n",r)
	fmt.Printf("%x\n",s)
	fmt.Println(r)
	fmt.Println(s)
	verify2:=ecc.ECDSAVerify(&pri.PublicKey,r,s,[]byte(data4))
	fmt.Println("ECDS的验签结果:",verify2)
}

////将生成的私钥保存到硬盘上,持久化存储
//err = rsa.GeneratePemFileByPrivateKey(privateKey)
//if err != nil {
//	fmt.Println("生成私钥文件出错",err.Error())
//	return
//}
////将生成的公钥保存到硬盘上,持久化存储
//err = rsa.GeneratePemFileByPublicKey(&privateKey.PublicKey)
//if err != nil {
//	fmt.Println("生成公钥文件出错",err.Error())
//	return
//}