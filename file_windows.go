// +build windows

/*
	【包名:】文件操作相关方法-windows
	【作者:】技术狼
*/
package fun

import (
	"syscall"
)

/*
	【名称:】文件隐藏
	【参数:】文件路径(string)
	【返回:】是否成功(bool)，错误信息(error)
	【备注:】
*/
func FileHide(filePath string) (bool, error) {
	namep, err := syscall.UTF16PtrFromString(filePath)
	if err != nil {
		return false, err
	}

	if err := syscall.SetFileAttributes(namep, syscall.FILE_ATTRIBUTE_HIDDEN); err != nil {
		return false, err
	}

	return true, nil
}
