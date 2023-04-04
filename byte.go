package fun

import (
	"encoding/base32"
	"encoding/binary"
	"io"
	"math"
	"reflect"
	"unsafe"
)


/*
	【名称:】[]byte转float32
	【参数:】[]byte
	【返回:】float32
	【备注:】
*/
func BytesToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}


/*
	【名称:】[]byte转float64
	【参数:】[]byte
	【返回:】float64
	【备注:】
*/
func BytesToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

/*
	【名称:】不安全的[]byte转string
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

/*
	【名称:】[]byte转io.Reader
	【参数:】[]byte
	【返回:】io.Reader
	【备注:】
*/
func BytesToIoReader(b []byte) io.Reader {
	return IoReader(b)
}

/*
	【名称:】base32加密
	【参数:】[]byte
	【返回:】string
	【备注:】
*/
func Base32Encode(src []byte) string {
	return base32.StdEncoding.EncodeToString(src)
}

/*
	【名称:】base32解密
	【参数:】string
	【返回:】[]byte, error
	【备注:】
*/
func Base32Decode(s string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(s)
}


