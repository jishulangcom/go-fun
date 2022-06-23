/*
	【包名:】浮点相关函数
	【作者:】技术狼
*/
package fun

import (
	"encoding/binary"
	"errors"
	"math"
	"strconv"
)


/*
	【名称:】变量是否浮点类型
	【参数:】变量(interface)
	【返回:】true/false(bool)
	【备注:】
*/
func IsFloat(variable interface{}) bool {
	return IsDouble(variable)
}


/*
	【名称:】变量是否浮点类型
	【参数:】变量(interface)
	【返回:】true/false(bool)
	【备注:】
*/
func IsDouble(variable interface{}) bool {
	variableType := GetType(variable)
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
func ToFloat64(val interface{}) (f64 float64, err error) {
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


/*
	【名称:】float64转string
	【参数:】float
	【返回:】string
	【备注:】
*/
func Float64ToStr(float64 float64) string {
	return strconv.FormatFloat(float64, 'E', -1, 64)
}


/*
	【名称:】float32转byte
	【参数:】float32
	【返回:】byte
	【备注:】
*/
func Float32ToByte(float32 float32) []byte {
	bits := math.Float32bits(float32)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}


/*
	【名称:】float64转byte
	【参数:】float32
	【返回:】byte
	【备注:】
*/
func Float64ToByte(float64 float64) []byte {
	bits := math.Float64bits(float64)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}


/*
	【名称:】float64转int64
	【参数:】float64
	【返回:】int64
	【备注:】
*/
func Float64ToInt64(float64 float64) int64 {
	return int64(float64)
}


/*
	【名称:】float64转int
	【参数:】float64
	【返回:】int
	【备注:】
*/
func Float64ToInt(float64 float64) int {
	return 0
}


/*
	【名称:】float32转int64
	【参数:】float32
	【返回:】int64
	【备注:】
*/
func Float32ToInt64(float64 float64) int64 {
	return int64(float64)
}


/*
	【名称:】float32转int
	【参数:】float32
	【返回:】int
	【备注:】
*/
func Float32ToInt(float32 float32) int {
	return 0
}