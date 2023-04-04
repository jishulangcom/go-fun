package test

import (
	"fmt"
	"github.com/jishulangcom/go-fun"
	"testing"
	"time"
)

// @title: 时间戳转日期
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/3 21:07
func TestTimestampToDate(t *testing.T) {
	str := fun.TimestampToDate(time.Now().Unix(), time.UTC)
	fmt.Println(str) // 2023-04-04
}

// @title: 时间戳转日期数字
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/3 21:07
func TestTimestampToDateNumber(t *testing.T) {
	str := fun.TimestampToDateNumber(time.Now().Unix(), fun.LocAsiaShanghai())
	fmt.Println(str) // 20230404
}