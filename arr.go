/*
	【包名:】数组相关函数
	【作者:】技术狼
*/
package fun

import (
	"errors"
	"reflect"
)


/*
	【名称:】一维数组求和
	【参数:】arr(interface)
	【返回:】和(float64)，err(error)
	【备注:】可用于数组、切片；小数点计算会有精度问题
*/
func ArraySum(arr interface{}) (sum float64, err error) {
	list := reflect.ValueOf(arr)
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < list.Len(); i++ {
			val := list.Index(i)
			v, err := ToFloat64(val.Interface())
			if err != nil {
				return 0, err
			}
			sum += v
		}
	default:
		return 0, errors.New("arr参数必须是数组或切片")
	}
	return
}


/*
	【名称:】变量是否为切片
	【参数:】变量(interface)
	【返回:】true/false(bool)
*/
func IsSlice(val interface{}) bool {
	if GetType(val) == "slice" {
		return true
	}
	return false
}