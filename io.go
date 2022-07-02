package fun

import (
	"bytes"
	"io"
)

/*
	【名称:】io.Reader转[]byte
	【参数:】io.Reader
	【返回:】[]byte
	【备注:】
*/
func IoReaderToBytes(reader io.Reader) []byte {
	buf := &bytes.Buffer{}
	buf.ReadFrom(reader)
	data := buf.Bytes()
	return data
}