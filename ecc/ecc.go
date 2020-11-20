package ecc

import (
	"CryptCode/utils"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
)

//ECC elliptic curve cryptography椭圆曲线加密算法
//ecc elliptic curve digital signature algorithm椭圆数字签名算法

//=====================私钥签名,公钥验签=============================
//生成私钥和公钥对
func GenerateECDSAKey()(*ecdsa.PrivateKey,error){
	//1,实例化一个椭圆曲线方程实例
	curve :=elliptic.P256()
	pri,err:=ecdsa.GenerateKey(curve,rand.Reader)
	if err != nil {
		return nil,err
	}
	return pri,nil
}

//使用ecdsa算法中的私钥对数据进行签名
func ECDSASign(pri *ecdsa.PrivateKey,data []byte)(*big.Int,*big.Int,error){
	hash := utils.Sha256Hash(data)
	r,s,err:=ecdsa.Sign(rand.Reader,pri,hash)
	if err != nil {
		return nil,nil,err
	}
	return r,s,nil
}

func ECDSAVerify(pub *ecdsa.PublicKey,r,s *big.Int,data []byte)bool{
	hash := utils.Sha256Hash(data)
	return ecdsa.Verify(pub,hash,r,s)
}


//der格式：将内存中的实例（比如r，s *big。int）进行序列化编码，使之可以在网络中进行长距离传输
//或者可以持久化存储起来
//der是一种编码格式，有具体的编码规范，只需要按照其规定好的格式进行数据拼接就可以
func MakeSignatureDerString(r,s string) string{
	lenSigR := len(r) / 2
	lenSigS := len(s) / 2
	lenSequence := lenSigR + lenSigS + 4
	strLenSigR := DecimalToHex(int64(lenSigR))
	strLenSigS := DecimalToHex(int64(lenSigS))
	strLenSequence := DecimalToHex(int64(lenSequence))
	derString := "30" + strLenSequence
	derString  =  derString + "02" + strLenSigR + r
	derString  =  derString + "02" + strLenSigS + s
	derString  =  derString + "01"
	return derString
}


//十进制转十六进制
func DecimalToHex(n int64)string{
	if n < 0 {
		fmt.Println("不符合要求！")
		return ""
	}
	if n == 0 {
		return "0"
	}
	hex:=map[int64]int64{10:65,11:66,12:67,13:68,14:69,15:70}
	s:=""
	for q := n; q > 0; q = q / 16 {
		m:=q % 16
		if m > 9 && m < 16 {
			m = hex[m]
			s = fmt.Sprintf("%v%v",string(m),s)
			continue
		}
		s = fmt.Sprintf("%v%v",m,s)
	}
	return s
}