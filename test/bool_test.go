package test

import (
	"fmt"
	"github.com/jishulangcom/go-fun"
	"testing"
)

// @title: 变量是否布尔类型
// @param: 变量(interface)
// @return: true/false(bool)
// @description:
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/3 21:40
func TestIsBool(t *testing.T) {
	v := true
	is := fun.IsBool(v)
	fmt.Println(is) // true
}