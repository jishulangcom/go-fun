/*
	【包名:】php相关函数
	【作者:】技术狼
*/
package fun

import (
	"reflect"
)


/*
	【名称:】获取变量类型
	【参数:】变量(interface)
	【返回:】类型名(string)
	【备注:】
*/
func GetType(variable interface{}) string {
	return reflect.TypeOf(variable).Kind().String()
}


/*
	【名称:】变量是否布尔类型(php)
	【参数:】变量(interface)
	【返回:】true/false(bool)
	【备注:】
*/
func IsBool(variable interface{}) bool {
	if GetType(variable) == "bool" {
		return true
	}
	return false
}