package fun

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

/*
	【名称:】目录是否存在
	【参数:】文件路径(string)
	【返回:】布尔
	【备注:】
*/
func DirIsExist(dirPath string) (bool, error) {
	return FileIsExist(dirPath)
}

/*
	【名称:】目录是否存在，若存在检验是否是目录
	【参数:】文件路径(string)
	【返回:】布尔
	【备注:】
*/
func DirIsExistAndChk(dirPath string) (bool, error) {
	info, err := os.Stat(dirPath)
	if err != nil {
		return false, err
	}

	if !info.IsDir() {
		return false, errors.New(dirPath + "已存在，但不是目录")
	}

	return true, nil
}

/*
	【名称:】创建目录
	【参数:】目录路径(string)，权限(os.FileMode)
	【返回:】error
	【备注:】
*/
func DirCreate(dirPath string, perm os.FileMode) error {
	return os.MkdirAll(dirPath, perm)
}

/*
	【名称:】目录复制
	【参数:】拷贝路径(string)，目标路径
	【返回:】error
	【备注:】
*/
func DirCopy(srcPath string, destPath string) error {
	var err error
	_, err = DirIsExistAndChk(srcPath)
	if err != nil {
		return err
	}

	is, _ := DirIsExistAndChk(destPath)
	if !is { // 不存在，就创建目录
		err = DirCreate(destPath, 0766)
		if err != nil {
			return err // 目录创建失败
		}
	}

	//
	err = filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if !f.IsDir() {
			path := strings.Replace(path, "\\", "/", -1)
			destNewPath := DirEndsWithSlash(destPath)
			destNewPath += f.Name()
			err = FileCopy(path, destNewPath)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

/*
	【名称:】目录路径以斜杠结尾
	【参数:】路径(string)
	【返回:】路径(string)
	【备注:】
*/
func DirEndsWithSlash(dirPath string) string {
	dirPath = strings.Replace(dirPath, "\\", "/", -1)
	if !strings.HasSuffix(dirPath, "/") {
		dirPath += "/"
	}

	return dirPath
}
