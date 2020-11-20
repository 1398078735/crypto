package base

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
)

var base58Alphabets = []byte("123456789ABCDEFGHJKLMNPQSTUVWXYZabcdefghijkmnopqrstuvwxyz")
//base58编码
func Base58Encode(data []byte) []byte {
	var result []byte
	x := big.NewInt(0).SetBytes(data)
	base :=big.NewInt(int64(len(base58Alphabets)))
	zero :=big.NewInt(0)
	mod := &big.Int{}

	for x.Cmp(zero) != 0 {
		x.DivMod(x,base,mod)
		result = append(result,base58Alphabets[mod.Int64()])
	}

	if data[0] == 0x00 {
		result = append(result,base58Alphabets[0])
	}
	ReverseBytes(result)
	return result
}

//将字节数组使用base58解码
func Base58Decode(data []byte)[]byte{
	result := big.NewInt(0)
	for _,b:=range data{
		charIndex := bytes.IndexByte(base58Alphabets,b)
		result.Mul(result,big.NewInt(58))
		result.Add(result,big.NewInt(int64(charIndex)))
	}
	decoded := result.Bytes()
	if data[0] == base58Alphabets[0] {
		decoded = append([]byte{0x00},decoded...)
	}
	return decoded
}


//转换成十六进制字符串
func Base58EncodeHexString(data string) string{
	arr,_ := hex.DecodeString(data)
	res:=Base58Encode(arr)
	return fmt.Sprintf("%s",res)
}

//字节逆转
func ReverseBytes(data []byte){
	for i, j := 0,len(data) - 1;i < j ; i,j=i+1,j-1 {
		data[i],data[j]=data[j],data[i]
	}
}