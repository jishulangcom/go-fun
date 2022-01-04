/*
	【包名:】文件相关函数
	【作者:】技术狼
*/
package fun

import (
	"io/ioutil"
	"os"
)

/*
	【名称:】获取文件内容
	【参数:】路径(string)
	【返回:】文件流([]byte)
	【备注:】
*/
func FileGetContents(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
	return []byte{}, err
	}
	var ba []byte
	ba, err = ioutil.ReadAll(f)
	return ba, err
}


/*
	【名称:】文件写入内容
	【参数:】路径(string)，内容(string)
	【返回:】错误状态(error)
	【备注:】
*/
func FilePutContents(filePath string, content string) error {
	return ioutil.WriteFile(filePath, []byte(content), 0644)
}