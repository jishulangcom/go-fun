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

// @title: 生成随机字符
// @param: 个数(int)
// @return: 字符串(string)
// @description: a-z,A-Z,0-9
// @date: 2022/7/3 21:07
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

// @title: 是否手机号
// @param: 手机号(string)
// @return: true/false(bool)
// @description:
// @date: 2022/7/3 21:10
func IsMobile(mobile string) bool {
	result, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, mobile)
	if result {
		return true
	}
	return false
}

// @title: 手机号*号处理
// @param: 手机号(string)
// @return: 加*号的手机号(string)
// @description: 非手机号则直接返回原字符
// @date: 2022/7/3 21:13
func MobileStar(mobile string) string {
	if IsMobile(mobile) {
		return mobile[:3] + "****" + mobile[7:]
	}

	return mobile
}

// @title: 是否邮箱
// @param: 邮箱(string)
// @return: true/false(bool)
// @description:
// @date: 2022/7/3 21:16
func IsEmail(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// @title: string转int
// @param: 字符串(str)
// @return: int, error
// @description:
// @date: 2022/7/3 21:18
func StrToInt(str string) (int, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// @title: string转int64
// @param: 字符串(str)
// @return: int64，error
// @description:
// @date: 2022/7/3 21:25
func StrToInt64(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// @title: string转uint64
// @param: 字符串(str)
// @return: uint64，error
// @description:
// @date: 2022/7/3 21:28
func StrToUint64(str string) (uint64, error) {
	i, err := StrToInt64(str)
	if err != nil {
		return 0, err
	}

	ui := uint64(i)

	return ui, nil
}

// @title: string转float64
// @param: 字符串(str)
// @return: float64, error
// @description:
// @date: 2022/7/3 21:29
func StrToFloat64(str string) (float64, error) {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// @title: string转*big.float
// @param: 字符串(str)
// @return: *big.Float, bool
// @description:
// @date: 2022/7/3 21:31
func StrToBigFloat(str string) (*big.Float, bool) {
	bigval, ok := new(big.Float).SetString(str)
	if !ok {
		return nil, false
	}
	return bigval, true
}

// @title: string转*big.Int
// @param: (字符)string，(精度位)int64
// @return: *big.Int, bool
// @description:
// @auth: 技术狼
// @date: 2022/7/3 21:00
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

// @title: 某字符串是否存在数组中
// @param: 字符(string)，数组(array)
// @return: 布尔(bool)
// @description:
// @date: 2022/7/3 21:33
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
