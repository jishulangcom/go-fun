package fun

import (
	"os"
	"path/filepath"
)


/*
	【名称:】获取导入包的绝对路径
	【参数:】包路由(string)，版本号
	【返回:】包的绝对地址
	【备注:】
*/
func GetImportPackagePath(packagePath string, version string) string {
	path := filepath.Join(os.Getenv("GOPATH"), "pkg", "mod")
	arr := Explode("/", packagePath)

	for _, v := range arr {
		path = filepath.Join(path, v)
	}
	path += "@"+ version

	return path
}
