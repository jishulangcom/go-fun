/*
	【包名:】int相关函数
	【作者:】技术狼
*/
package fun

import (
	"strconv"
)


/*
	【名称:】变量是否int类型(php)
	【参数:】变量(interface)
	【返回:】true/false(bool)
	【备注:】
*/
func IsInt(variable interface{}) bool {
	variableType := GetType(variable)
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
func IsNumeric(variable interface{}) bool {
	if IsInt(variable) || IsFloat(variable) {
		return true
	}

	if IsString(variable) {
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