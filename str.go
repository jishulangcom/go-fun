/*
	【包名:】字符串相关函数
	【作者:】技术狼
*/
package fun

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)


/*
	【名称:】生成随机字符
	【参数:】个数(int)
	【返回:】字符串(string)
	【备注:】
*/
func Rand_string(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var src = rand.NewSource(time.Now().UnixNano())
	const (
		letterIdxBits = 6
		letterIdxMask = 1<<letterIdxBits - 1
		letterIdxMax = 63 / letterIdxBits
	)
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}


/*
	【名称:】是否手机号
	【参数:】手机号(string)
	【返回:】true/false(bool)
	【备注:】
*/
func Mobile_is(mobile string) bool{
	return Is_mobile(mobile)
}

func Is_mobile(mobile string) bool{
	result, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, mobile)
	if result {
		return true
	}
	return false
}


/*
	【名称:】手机号*号处理
	【参数:】手机号(string)
	【返回:】加*号的手机号(string)
	【备注:】非手机号则直接返回原字符
*/
func Mobile_star(mobile string) string{
	if(Mobile_is(mobile)) {
		return mobile[:3] + "****" + mobile[7:]
	}

	return mobile
}


/*
	【名称:】是否邮箱
	【参数:】邮箱(string)
	【返回:】true/false(bool)
	【备注:】
*/
func Email_is(email string) bool{
	return Is_Email(email)
}

func Is_Email(email string) bool{
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}



/*
	【名称:】string转int
	【参数:】字符串(str)
	【返回:】int
	【备注:】
*/
func Str_to_int(str string) int{
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}


/*
	【名称:】string转int64
	【参数:】字符串(str)
	【返回:】int64
	【备注:】
*/
func Str_to_int64(str string) int64{
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i
}


/*
	【名称:】string转float64
	【参数:】str
	【返回:】float64
	【备注:】
*/
func Str_to_Float64(str string) float64 {
	f, err := strconv.ParseFloat(str,64)
	if err != nil {
		return 0
	}
	return f
}


