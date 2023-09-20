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
	tt, _ := time.Parse(time.RFC3339, str)
	return tt.Format(ORIGIN_TIME)
}

/*
	【名称:】标准时间转化为时间戳
	【参数:】标准时间(string)
	【返回:】时间戳（秒）(int64)
	【备注:】如2018-01-14 21:45:54 转为 1515937554
*/
func DatetimeToTimestamp(str string) int64 {
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(ORIGIN_TIME, str, loc)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64
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
	【名称:】获取当前日期时间
	【参数:】
	【返回:】标准时间(string)
	【备注:】时间格式：2021-12-30 23:14:07
*/
func CurrentDateTime() string {
	return time.Now().Format(ORIGIN_TIME)
}

/*
	【名称:】获取时间戳（秒）
	【参数:】
	【返回:】时间戳(int64)
	【备注:】
*/
func TimeStampSecond() int64 {
	return time.Now().Unix()
}

/*
	【名称:】获取时间戳（毫秒）
	【参数:】
	【返回:】时间戳(int64)
	【备注:】
*/
func Millisecond() int64 {
	return time.Now().UnixNano() / 1e6
}


/*
	【名称:】时间戳转日期
	【参数:】int64，*time.Location
	【返回:】string
	【备注:】
*/
func TimestampToDate(timestamp int64, loc *time.Location) string {
	tm := time.Unix(timestamp, 0).In(loc) // time.UTC
	return tm.Format("2006-01-02")
}

/*
	【名称:】时间戳转日期数字
	【参数:】int64，*time.Location
	【返回:】string
	【备注:】
*/
func TimestampToDateNumber(timestamp int64, loc *time.Location) string {
	tm := time.Unix(timestamp, 0).In(time.UTC) // time.UTC
	return tm.Format("20060102")
}

/*
	【名称:】时区-东八区(中国上海)
	【参数:】
	【返回:】*time.Location
	【备注:】
*/
func LocAsiaShanghai() *time.Location {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return loc
}

/*
	【名称:】时区-0时区
	【参数:】
	【返回:】*time.Location
	【备注:】
*/
func LocUTC() *time.Location {
	return time.UTC
}

/*
	【名称:】时区-西八区(洛杉矶)
	【参数:】
	【返回:】*time.Location
	【备注:】
*/
func LocAmericaLosAngeles() *time.Location {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	return loc
}

/*
	【名称:】获取某个时间相隔N天的连续时间
	【参数:】t:某个时间
			days:相隔N天
			dayN:取多少天
			SameDay:当天是否在内
	【返回:】[]time.Time
	【备注:】
*/
func GetSometimeApartNDaysTimes(t time.Time, days int, dayN int, SameDay bool) []time.Time {
	list := make([]time.Time, 0, dayN)
	I := 0
	if SameDay {
		I = 1
		list = append(list, t)
	}
	for i := I; i < dayN; i++ {
		newTime := t.AddDate(0, 0, days)
		list = append(list, newTime)
		t = newTime
	}
	return list
}

/*
	【名称:】获取某个时间相隔N天的连续时间(指定日期格式)
	【参数:】t:某个时间
			days:相隔N天
			dayN:取多少天
			SameDay:当天是否在内
			format:日期格式（如:2006-01-02 15:04:05）
	【返回:】[]time.Time
	【备注:】
*/
func GetSometimeApartNDaysTimesFormat(t time.Time, days int, dayN int, SameDay bool, format string) []string {
	list := GetSometimeApartNDaysTimes(t, days, dayN, SameDay)
	newList := make([]string, 0, len(list))
	for _, t := range list {
		date := t.Format(format)
		newList = append(newList, date)
	}

	return newList
}


/*
	【名称:】获取某个时间相隔N天的时间(指定日期格式)
	【参数:】t:某个时间
			days:相隔N天
			dayN:取多少天
			SameDay:当天是否在内
			format:日期格式（如:2006-01-02 15:04:05）
	【返回:】[]time.Time
	【备注:】 可以用来获取明天、昨天日期
*/
func GetSometimeApartNDaysTimeFormat(t time.Time, days int, format string) string {
	newT := t.AddDate(0, 0, days)
	str := newT.Format(format)
	return str
}


/*
	【名称:】时间差（当前时间-某个时间）
	【参数:】dateTimeStr:某个时间（格式：2006-01-02 15:04:05）
			loc:时区（如:time.UTC）
	【返回:】 秒
	【备注:】 now-dateTimeStr
*/
func TimeDifferenceNowSubDateTime(dateTimeStr string, loc *time.Location) (int, error) {
	now, err := time.ParseInLocation(ORIGIN_TIME, time.Now().Format("2006-01-02 15:04:05"), loc)
	if err != nil {
		return 0, err
	}

	t, err := time.ParseInLocation(ORIGIN_TIME, dateTimeStr, loc)
	if err != nil {
		return 0, err
	}

	reverseTime := now.Sub(t)
	return int(reverseTime.Seconds()), nil
}


/*
	【名称:】时间差（两个指定时间）
	【参数:】dateTimeStr1:某个时间， dateTimeStr2:某个时间（格式：2006-01-02 15:04:05）
			loc:时区（如:time.UTC）
	【返回:】 秒
	【备注:】 dateTimeStr1-dateTimeStr2
*/
func TimeDifferenceTwoDateTime(dateTimeStr1 string, dateTimeStr2 string, loc *time.Location) (int, error) {
	t1, err := time.ParseInLocation(ORIGIN_TIME, dateTimeStr1, loc)
	if err != nil {
		return 0, err
	}

	t2, err := time.ParseInLocation(ORIGIN_TIME, dateTimeStr2, loc)
	if err != nil {
		return 0, err
	}

	reverseTime := t1.Sub(t2)
	return int(reverseTime.Seconds()), nil
}


