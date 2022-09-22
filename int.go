/*
	【包名:】int相关函数
	【作者:】技术狼
*/
package fun

import (
	"strconv"
	"unsafe"
)

/*
	【名称:】int转字符
	【参数:】int
	【返回:】string
	【备注:】
*/
func IntToStr(int int) string {
	string := strconv.Itoa(int)
	return string
}

/*
	【名称:】int64转string
	【参数:】int64
	【返回:】string
	【备注:】
*/
func Int64ToStr(int64 int64) string {
	string := strconv.FormatInt(int64, 10)
	return string
}

/*
	【名称:】int64转int
	【参数:】int64
	【返回:】int
	【备注:】
*/
func Int64ToInt(int64 int64) int {
	return *(*int)(unsafe.Pointer(&int64))
}

/*
	【名称:】int转int64
	【参数:】int
	【返回:】int64
	【备注:】
*/
func IntToInt64(int int) int64 {
	return int64(int)
}

/*
	【名称:】int转int64
	【参数:】int
	【返回:】int64
	【备注:】
*/
func IntToFloat64(int int) float64 {
	return float64(int)
}

// 判断是奇数
func IntIsOdd(num int) bool {
	return Int64IsOdd(int64(num))
}

// 判断是偶数
func IntIsEven(num int) bool {
	return Int64IsEven(int64(num))
}

// 判断是奇数
func Int64IsOdd(num int64) bool {
	if num%2 == 0 {
		return false
	}
	return true
}

// 判断是偶数
func Int64IsEven(num int64) bool {
	if num%2 == 0 {
		return true
	}
	return false
}
