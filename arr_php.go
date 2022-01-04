/*
	【包名:】数组(php)相关函数
	【作者:】技术狼
*/
package fun

import (
	"fmt"
	"strings"
)


/*
	【名称:】变量是否为数组(php)
	【参数:】变量(interface)
	【返回:】true/false(bool)
*/
func IsArray(val interface{}) bool {
	if Gettype(val) == "array" {
		return true
	}
	return false
}


/*
	【名称:】数组分割为字符串
	【参数:】分割符(string)，数组([]interface)，
	【返回:】字符串(string)
	【备注:】参数arr必须为字符串类型的一维数组
*/
func Implode (separator string, arr interface{}) string {
	/*
	v := reflect.ValueOf(arr)
	len := v.Len()
	if len == 0 {
		return ""
	}

	arrType := Gettype(arr)
	if(arrType != "array" && arrType != "slice" ) {
		return ""
	}

	var slice []string
	for i := 0; i < len; i++ {
		slice = append(slice, v.Index(i).String())
	}
	return strings.Join(slice, separator)
	*/
	return  strings.Replace(strings.Trim(fmt.Sprint(arr), "[]"), " ", separator, -1)
}


/*
	【名称:】数组分割为字符串
	【参数:】分割符(string)，数组([]interface)，
	【返回:】字符串(string)
*/
func Join (separator string, arr interface{}) string {
	return Implode(separator, arr)
}