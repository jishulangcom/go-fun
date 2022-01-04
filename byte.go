package fun

import (
	"encoding/binary"
	"math"
)


/*
	【名称:】byte转float32
	【参数:】byte
	【返回:】float32
	【备注:】
*/
func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}


/*
	【名称:】byte转float64
	【参数:】byte
	【返回:】float64
	【备注:】
*/
func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

