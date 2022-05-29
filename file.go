package fun

import (
	"io"
	"os"
	"strings"
	"syscall"
)

/*
	【名称:】文件是否存在
	【参数:】文件路径(string)
	【返回:】布尔
	【备注:】
*/
func FileIsExist(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

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

/*
	【名称:】文件复制
	【参数:】文件路径(string)、目标路径
	【返回:】error
	【备注:】
*/
func FileCopy(srcPath, destPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	//
	destSplitPathDirs := strings.Split(destPath, "/") //分割path目录
	destSplitPath := ""                               //检测时候存在目录
	for index, dir := range destSplitPathDirs {
		if index < len(destSplitPathDirs)-1 {
			destSplitPath = destSplitPath + dir + "/"
			is, _ := DirIsExist(destSplitPath)
			if is == false {
				err := os.Mkdir(destSplitPath, os.ModePerm)
				if err != nil {
					return err
				}
			}
		}
	}

	//
	dstFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}
