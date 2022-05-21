/*
	【包名:】字符串相关函数
	【作者:】技术狼
*/
package fun

import (
	"math/big"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
	【名称:】生成随机字符
	【参数:】个数(int)
	【返回:】字符串(string)
	【备注:】
*/
func RandString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var src = rand.NewSource(time.Now().UnixNano())
	const (
		letterIdxBits = 6
		letterIdxMask = 1<<letterIdxBits - 1
		letterIdxMax  = 63 / letterIdxBits
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
func IsMobile(mobile string) bool {
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
func MobileStar(mobile string) string {
	if IsMobile(mobile) {
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
func IsEmail(email string) bool {
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
func StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

/*
	【名称:】string转int64
	【参数:】字符串(str)
	【返回:】int64，error
	【备注:】
*/
func StrToInt64(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

/*
	【名称:】string转uint64
	【参数:】字符串(str)
	【返回:】uint64，error
	【备注:】
*/
func StrToUint64(str string) (uint64, error) {
	i, err := StrToInt64(str)
	if err != nil {
		return 0, err
	}

	ui := uint64(i)

	return ui, nil
}

/*
	【名称:】string转float64
	【参数:】str
	【返回:】float64
	【备注:】
*/
func StrToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return f
}

/*
	【名称:】string转*big.float
	【参数:】str
	【返回:】*big.Float
	【备注:】
*/
func StrToBigFloat(str string) (*big.Float, bool) {
	bigval, ok := new(big.Float).SetString(str)
	if !ok {
		return nil, false
	}
	return bigval, true
}

/*
	【名称:】string转*big.Int
	【参数:】str，exp精度位
	【返回:】*big.Int
	【备注:】
*/
func StrToBigInt(str string, exp int64) (*big.Int, bool) {
	bigval, ok := StrToBigFloat(str)
	if !ok {
		return nil, false
	}

	coin := new(big.Float)
	b := new(big.Int).Exp(big.NewInt(10), big.NewInt(exp), nil)
	coin.SetInt(b)

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result)

	return result, true
}

/*
	【名称:】某字符串是否存在数组中
	【参数:】字符(string)，数组(array)
	【返回:】布尔(bool)
	【备注:】
*/
func StrIsInArr(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	//index的取值：[0,len(str_array)]
	if index < len(str_array) && str_array[index] == target { //需要注意此处的判断，先判断 &&左侧的条件，如果不满足则结束此处判断，不会再进行右侧的判断
		return true
	}
	return false
}

/*
	【名称:】某字符是否存在字符串中
	【参数:】字符(string)，字符串(string)
	【返回:】布尔(bool)
	【备注:】
*/
func StrIsInString(target string, strs string) bool {
	if find := strings.Contains(strs, target); find {
		return true
	}
	return false
}
