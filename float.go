/*
	【包名:】浮点相关函数
	【作者:】技术狼
*/
package fun

import (
	"errors"
	"strconv"
)


/*
	【名称:】变量是否浮点类型
	【参数:】变量(interface)
	【返回:】true/false(bool)
	【备注:】
*/
func Is_float(variable interface{}) bool {
	return Is_double(variable)
}


/*
	【名称:】变量是否浮点类型
	【参数:】变量(interface)
	【返回:】true/false(bool)
	【备注:】
*/
func Is_double(variable interface{}) bool {
	variableType := Gettype(variable)
	if variableType == "float64" || variableType == "float32" {
		return true
	}
	return false
}


/*
	【名称:】转换为 Float64
	【参数:】val(interface)
	【返回:】和(float64)，err(error)
	【备注:】
*/
func To_float64(val interface{}) (f64 float64, err error) {
	switch value := val.(type) {
	case float64:
		return value, nil
	case float32:
		return float64(value), nil
	case uint8:
		return float64(value), nil
	case uint16:
		return float64(value), nil
	case uint32:
		return float64(value), nil
	case uint64:
		return float64(value), nil
	case uint:
		if strconv.IntSize == 32 || strconv.IntSize == 64 {
			return float64(value), nil
		}
		return 0, errors.New("uint转换为float64失败")
	case int8:
		return float64(value), nil
	case int16:
		return float64(value), nil
	case int32:
		return float64(value), nil
	case int64:
		return float64(value), nil
	case int:
		if strconv.IntSize == 32 || strconv.IntSize == 64 {
			return float64(value), nil
		}
		return 0, errors.New("int转换为float64失败")
	case bool:
		if value {
			f64 = 1
		}
		return
	case string:
		f64, err = strconv.ParseFloat(value, 64)
		if err == nil {
			return
		}
		return 0, errors.New("字符串转换为float64失败")
	default:
		return 0, errors.New("转换为float64失败")
	}
}