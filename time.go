package fun

import (
	"time"
)

const ORIGIN_TIME = "2006-01-02 15:04:05" //go原始时间

/*
	【名称:】Rfc3339格式时间转为正常时间
	【参数:】Rfc3339时间字符(string)
	【返回:】时间字符串(string)
	【备注:】如2018-01-14T21:45:54+08:00 转为 2018-01-14 21:45:54
*/
func Rfc3339_to_datetime(str string) string {
	tt, _:=time.Parse(time.RFC3339,str)
	return tt.Format(ORIGIN_TIME)
}


/*
	【名称:】标准时间转化为时间戳
	【参数:】标准时间(string)
	【返回:】时间戳（秒）(int64)
	【备注:】如2018-01-14 21:45:54 转为 1515937554
*/
func DatetimeToTimestamp(str string) int64 {
	loc, _ := time.LoadLocation("Local")    //获取时区
	tmp, _ := time.ParseInLocation(ORIGIN_TIME, str, loc)
	timestamp := tmp.Unix()    //转化为时间戳 类型是int64
	return timestamp
}


/*
	【名称:】时间戳转化为标准时间
	【参数:】时间戳(int64)
	【返回:】标准时间(string)
	【备注:】如2018-01-14 21:45:54 转为 1515937554
*/
func TimestampToDatetime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(ORIGIN_TIME)
}


/*
	【名称:】获取当前时间
	【参数:】
	【返回:】标准时间(string)
	【备注:】时间格式：2021-12-30 23:14:07
*/
func CurrentTime() string {
	return time.Now().Format(ORIGIN_TIME)
}


/*
	【名称:】获取时间戳
	【参数:】
	【返回:】时间戳(int64)
	【备注:】秒
*/
func TimeStamp() int64 {
	return time.Now().Unix()
}

