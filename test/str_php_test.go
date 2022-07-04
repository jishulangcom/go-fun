package test

import (
	"fmt"
	"github.com/jishulangcom/go-fun"
	"testing"
)

// @title: 移除字符串左侧的字符
// @param: 字符串(string)，要移除的字符(string)
// @return: 字符串(string)
// @description:
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/4 22:07
func TestLtrim(t *testing.T) {
	str := fun.Ltrim("技术狼jishulang.com", "技术狼")
	fmt.Println(str)
}

// @title: 字符替换
// @param: 要替换的字符串(string)，旧字符(string)，新字符(string)
// @return: string
// @description:
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/4 22:14
func TestStrReplace(t *testing.T) {
	str := fun.StrReplace("技术狼jishulang.com", "jishulang.com", "jishulang.cn")
	fmt.Println(str)
}