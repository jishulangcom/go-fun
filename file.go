package fun

import "os"

/*
	【名称:】文件是否存在
	【参数:】文件路径(string)
	【返回:】布尔
	【备注:】
*/
func FileIsExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}