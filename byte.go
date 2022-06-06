package fun

import (
	"encoding/binary"
	"math"
	"reflect"
	"unsafe"
)


/*
	【名称:】byte转float32
	【参数:】[]byte
	【返回:】float32
	【备注:】
*/
func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}


/*
	【名称:】byte转float64
	【参数:】[]byte
	【返回:】float64
	【备注:】
*/
func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

/*
	【名称:】不安全字节转string
	【参数:】[]byte
	【返回:】string
	【备注:】
*/
func UnsafeBytesToString(bytes []byte) string {
	if len(bytes) == 0 {
		return ""
	}

	hdr := &reflect.StringHeader{
		Data: uintptr(unsafe.Pointer(&bytes[0])),
		Len:  len(bytes),
	}
	return *(*string)(unsafe.Pointer(hdr))
}

