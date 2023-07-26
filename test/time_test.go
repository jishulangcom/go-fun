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

// @title: 获取某个时间相隔N天的连续时间
// @auth: 技术狼(jishulang.com)
// @date: 2023/7/26 21:07
func TestGetSometimeApartNDaysTimes(t *testing.T){
	list := fun.GetSometimeApartNDaysTimesFormat(time.Now().UTC(), +1, 7, true, "20060102")
	fmt.Println(list)
}

// @title: 获取某个时间相隔N天的时间
// @auth: 技术狼(jishulang.com)
// @date: 2023/7/26 21:07
func TestGetSometimeApartNDaysTimeFormat(t *testing.T){
	date := fun.GetSometimeApartNDaysTimeFormat(time.Now().UTC(), +1, "2006-01-02")
	fmt.Println(date) // 明天日期
}
