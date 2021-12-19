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
	【名称:】变量是否int类型(php)
	【参数:】变量(interface)
	【返回:】true/false(bool)
	【备注:】
*/
func Is_int(variable interface{}) bool {
	variableType := Gettype(variable)
	if variableType == "int" || variableType == "uint" || variableType == "uint64" ||
		variableType == "int64" || variableType == "uint32" || variableType == "int32" ||
		variableType == "int8" || variableType == "uint8"{
		return true
	}
	return false
}


/*
	【名称:】变量是否数字类型
	【参数:】变量(interface)
	【返回:】true/false(bool)
	【备注:】
*/
func Is_numeric(variable interface{}) bool {
	if Is_int(variable) || Is_float(variable) {
		return true
	}

	if Is_string(variable) {
		tmpStr, err := variable.(string)
		if err {
			_, err := strconv.Atoi(tmpStr)
			if err == nil {
				return true
			}
		}
	}
	return false
}


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

