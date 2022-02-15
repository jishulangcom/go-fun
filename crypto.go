package fun
/*************************************
【包名:】常用方法-加密
【作者:】技术狼
*************************************/

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

/*
	【名称:】Md5加密
	【参数:】字符串(string)
	【返回:】字符串(string)
	【备注:】
*/
func Md5(src string) string {
	//方法1
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))

	//方法2
	//return fmt.Sprintf("%x", md5.Sum([]byte(src)))
}

/*
	【名称:】sha1
	【参数:】字符串(string)
	【返回:】字符串(string)
	【备注:】
*/
func Sha1(str string) string {
	return (fmt.Sprintf("%x", sha1.Sum([]byte(str))))
}

/*
	【名称:】Sha256
	【参数:】字符串(string)
	【返回:】字符串(string)
	【备注:】
*/
func Sha256(str string) string {
	hashInBytes := sha256.Sum256([]byte(str))
	hashStr := hex.EncodeToString(hashInBytes[:])
	return hashStr
}

/*
	【名称:】Sha512加密
	【参数:】
	【返回:】
	【备注:】
*/
func Sha512(src string) string {
	h := sha512.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}