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

/*
	【名称:】获取io.Writer
	【参数:】
	【返回:】*bytes.Buffer
	【备注:】
*/
func IoWriter() *bytes.Buffer {
	return new(bytes.Buffer)
}

/*
	【名称:】获取io.Reader
	【参数:】
	【返回:】*bytes.Buffer
	【备注:】
*/
func IoReader(b []byte) io.Reader {
	return bytes.NewReader(b)
}
