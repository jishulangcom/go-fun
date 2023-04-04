package fun

import (
	"os"
	"path/filepath"
	"reflect"
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
	path += "@" + version

	return path
}

/*
	【名称:】调用函数（魔法函数）
	【参数:】包路由(string)，版本号
	【返回:】包的绝对地址
	【备注:】
*/
func CallFun(any interface{}, funName string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	return reflect.ValueOf(any).MethodByName(funName).Call(inputs)
}

/*
	【名称:】获取环境变量值
	【参数:】
	【返回:】
	【备注:】
*/
func GetEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

/*
	【名称:】设置环境变量值
	【参数:】
	【返回:】
	【备注:】
*/
func SetEnv(key string, val string) error {
	return os.Setenv(key, val)
}
