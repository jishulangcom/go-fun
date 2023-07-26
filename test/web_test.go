package test

import (
	"fmt"
	"github.com/jishulangcom/go-fun"
	"testing"
)

// @title: 获取URL主域名
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/3 21:07
func TestUrlDomain(t *testing.T) {
	str := fun.UrlDomain("https://www.jishulang.com/index.html?source=gofun")
	fmt.Println(str) // jishulang.com
}
