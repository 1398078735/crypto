package rsa

import (
	"CryptCode/utils"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"flag"
	"io/ioutil"
	"os"
)

//根据用户传入的内容,自动创建私钥和公钥并生成相应格式的文件
func GenerateKey(filename string)(*rsa.PrivateKey,error){
	//生成私钥
	private,err:= CreatePairKeys()
	if err != nil {
		return nil,err
	}
	//创建私钥文件
	err = generatePemFileByPrivateKey(private,filename)
	if err != nil {
		return nil,err
	}
	//生成公钥文件
	err = generatePemFileByPublicKey(&private.PublicKey,filename)
	if err != nil {
		return nil,err
	}
	return private,nil
}

//RSA的go语言API实现:
//私钥:type PrivateKey struct{
//    PublicKey
//    ....
//}
//私钥:
//公钥:
func CreatePairKeys() (*rsa.PrivateKey,error){
	//1,生成存私钥
	var bits int
	flag.IntVar(&bits,"b",2048,"密钥长度")
	//fmt.Println(bits)
	privateKey,err:=rsa.GenerateKey(rand.Reader,bits)
	if err != nil {
		return nil,err
	}
	//2,根据私钥生成公钥
	//publicKey := privateKey.Public()
	//3,将私钥和公钥返回
	return privateKey,nil
}

//===================证书文件的操作================
//关于pem证书文件的读取
//读取pem文件格式的密钥数据,私钥
func ReadPemPriKey(filename string)(*rsa.PrivateKey,error){
	blockBytes,err := ioutil.ReadFile("RSA_privatekey_"+filename+".pem")
	if err != nil {
		return nil,err
	}
	//decode将byte解码为内存当中的实例对象
	block,_:=pem.Decode(blockBytes)
	priBytes := block.Bytes
	priKey,err := x509.ParsePKCS1PrivateKey(priBytes)
	return priKey,err
}

//读取pem文件格式的密钥数据,公钥
func ReadPemPubKey(filename string)(*rsa.PublicKey,error){
	blockBytes,err := ioutil.ReadFile("RSA_publickey_"+filename+".pem")
	if err != nil {
		return nil,err
	}
	//decode将byte解码为内存当中的实例对象
	block,_:=pem.Decode(blockBytes)
	pubBytes := block.Bytes
	pubKey,err := x509.ParsePKCS1PublicKey(pubBytes)
	return pubKey,err
}


//根据给定的私钥数据,生成对应的pem文件
func generatePemFileByPrivateKey(private *rsa.PrivateKey,filename string)error{
	//根据pkcs1序列化后的私钥
	privateStream := x509.MarshalPKCS1PrivateKey(private)
	//pem文件,此时,privateFile文件为空
	privateFile,err := os.Create("RSA_privatekey_"+filename+".pem")//存私钥生成的文件
	if err != nil {
		return err
	}
	//pem文件中的格式，结构体
	block := &pem.Block{
		Type:    "RSA PRIVATE KEY",
		Bytes:   privateStream,
	}
	//将整理好的格式内容写入到pem文件中
	err = pem.Encode(privateFile,block)
	if err != nil {
		return err
	}
	return nil
}

////根据给定的公钥数据,生成对应的pem文件
func generatePemFileByPublicKey(public *rsa.PublicKey,filename string)error{
	//根据pkcs1序列化后的私钥
	publicPkix := x509.MarshalPKCS1PublicKey(public)
	//pem文件,此时,privateFile文件为空
	publicFile,err := os.Create("RSA_publicKey_"+filename+".pem")//存私钥生成的文件
	if err != nil {
		return err
	}
	block := &pem.Block{
		Type:    "RSA PUBLIC KEY",
		Bytes:   publicPkix,
	}
	//将整理好的格式内容写入到pem文件中
	err = pem.Encode(publicFile,block)
	if err != nil {
		return err
	}
	return nil
}

//使用RSA算法对数据进行加密,返回密文
//data不能太长,小数据量.不能超过私钥减去11个字节的长度
func RSAEncrypt(key rsa.PublicKey,data []byte)([]byte,error){
	return rsa.EncryptPKCS1v15(rand.Reader,&key,data)
}

//使用RSA算法对数据解密,返回解密后的明文
func RSADEcrypt(private *rsa.PrivateKey,cipherText []byte) ([]byte,error){
	return rsa.DecryptPKCS1v15(rand.Reader,private,cipherText)
}


//私钥签名,使用RSA算法对数据进行数字签名,并返回签名信息
func RSASign(private *rsa.PrivateKey,data []byte,hash crypto.Hash)([]byte,error){
	if hash == crypto.MD5 {
		hashed :=utils.Md5Hash(data)
		return rsa.SignPKCS1v15(rand.Reader,private,crypto.MD5,hashed)
	}else if hash == crypto.SHA256 {
		hashed :=utils.Md5Hash(data)
		return rsa.SignPKCS1v15(rand.Reader,private,crypto.SHA256,hashed)
	}else {
		return nil,errors.New("不支持你传入的类型")
	}
}

//公钥验签,使用RSA算法对数据进行签名验证,并返回签名验证结果
func RSAVerify(pub rsa.PublicKey,signText []byte,data []byte,hash crypto.Hash) (bool,error) {
	if hash == crypto.MD5 {
		hashed :=utils.Md5Hash(data)
		err := rsa.VerifyPKCS1v15(&pub,crypto.MD5,hashed,signText)
		return err == nil,err
	}else if hash == crypto.SHA256 {
		hashed :=utils.Md5Hash(data)
		err := rsa.VerifyPKCS1v15(&pub,crypto.SHA256,hashed,signText)
		return err == nil,err
	}else {
		return false,errors.New("不支持你传入的类型")
	}
}