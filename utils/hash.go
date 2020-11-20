package utils

import (
	"crypto/md5"
	"crypto/sha256"
)

func Md5Hash(data []byte)[]byte{
	md5Hash := md5.New()
	md5Hash.Write(data)
	return md5Hash.Sum(nil)
}

func Sha256Hash(data []byte)[]byte{
	Sha256Hash := sha256.New()
	Sha256Hash.Write(data)
	return Sha256Hash.Sum(nil)
}