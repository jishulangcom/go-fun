/*
	【包名:】浮点相关函数
	【作者:】技术狼
*/
package fun

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
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
	【名称:】float64转string
	【参数:】float
	【返回:】string
	【备注:】
*/
func Float64ToStr(f float64) string {
	//fmt.Println("Float64ToStr:", f)
	//str := fmt.Sprintf("%f", f)
	//fmt.Println(str)
	//return fmt.Sprintf("%f", f)
	return strconv.FormatFloat(f, 'E', -1, 64)
}

func Float64Decimal(f float64) int {
	numstr := fmt.Sprint(f)
	tmp := strings.Split(numstr, ".")
	if len(tmp) <= 1 {
		return 0
	}
	return len(tmp[1])
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
	return int(float64)
}

/*
	【名称:】float32转int64
	【参数:】float32
	【返回:】int64
	【备注:】
*/
func Float32ToInt64(float32 float32) int64 {
	return int64(float32)
}

/*
	【名称:】float32转int
	【参数:】float32
	【返回:】int
	【备注:】
*/
func Float32ToInt(float32 float32) int {
	return int(float32)
}
