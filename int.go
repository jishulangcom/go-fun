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
func Int_to_str(int int) string {
	string := strconv.Itoa(int)
	return string
}


/*
	【名称:】int64转string
	【参数:】int64
	【返回:】string
	【备注:】
*/
func Int64_to_str(int64 int64) string {
	string := strconv.FormatInt(int64,10)
	return string
}


/*
	【名称:】int64转int
	【参数:】int64
	【返回:】int
	【备注:】
*/
func Int64_to_int(int64 int64) int {
	return *(*int)(unsafe.Pointer(&int64))
}


/*
	【名称:】int转int64
	【参数:】int
	【返回:】int64
	【备注:】
*/
func Int_to_int64(int int) int64 {
	return int64(int)
}


/*
	【名称:】int转int64
	【参数:】int
	【返回:】int64
	【备注:】
*/
func Int_to_float64(int int) float64 {
	return float64(int)
}

