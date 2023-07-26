package fun

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
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
	【名称:】文件复制
	【参数:】文件路径(string)、目标路径(string)
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

/*
	【名称:】获取文件内容
	【参数:】文件路径(string)
	【返回:】[]byte，error
	【备注:】
*/
func GetFileContent(filedir string) ([]byte, error) {
	file, err := os.Open(filedir)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}


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