/*
	// @title: 字符串相关函数
	// @auth: 技术狼(jishulang.com)
*/
package fun

import (
	"fmt"
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
	expStr := fmt.Sprintf("1e%d", exp) // "1e" + "8"

	amount := new(big.Rat)
	_, ok := amount.SetString(str)
	if !ok {
		return nil, false
	}
	bigexp := new(big.Rat)
	_, ok2 := bigexp.SetString(expStr)
	if !ok2 {
		return nil, false
	}

	result := new(big.Rat).Mul(amount, bigexp)

	result.FloatString(0)

	return new(big.Int).SetString(result.FloatString(0), 10)
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


/*
	【名称:】变量是否字符类型
	【参数:】变量(interface)
	【返回:】true/false(bool)
	【备注:】
*/
func IsString(variable interface{}) bool {
	if GetType(variable) == "string" {
		return true
	}
	return false
}


/*
	【名称:】字符串打散为数组
	【参数:】分割字符串(string)，要分割的字符串(string)
	【返回:】数组(map[int]string)
	【备注:】
*/
func Explode(separator, str string) []string {
	return strings.Split(str, separator)
}


/*
	【名称:】获取字符串的长度
	【参数:】字符串(string)
	【返回:】长度(int)
	【备注:】
*/
func MbStrlen(str string) int {
	return len([]rune(str))
}


/*
	【名称:】字符首字母转成小写
	【参数:】字符串(string)
	【返回:】字符串(string)
	【备注:】
*/
func Lcfirst(str string) string {
	if len(str) == 0 {
		return ""
	}

	s := []rune(str)
	if s[0] >= 65 && s[0] <= 90 {
		s[0] += 32
	}
	return string(s)
}


/*
	【名称:】字符首字母转成大写
	【参数:】字符串(string)
	【返回:】字符串(string)
	【备注:】
*/
func Ucfirst(str string) string {
	if len(str) == 0 {
		return ""
	}

	s := []rune(str)
	if s[0] >= 97 && s[0] <= 122 {
		s[0] -= 32
	}
	return string(s)
}


/*
	【名称:】字符转大写
	【参数:】字符串(string)
	【返回:】字符串(string)
	【备注:】
*/
func StrToUpper(str string) string {
	return strings.ToUpper(str)
}


/*
	【名称:】字符转小写
	【参数:】字符串(string)
	【返回:】字符串(string)
	【备注:】
*/
func StrToLower(str string) string {
	return strings.ToLower(str)
}


/*
	【名称:】每个单词的首字符转换为大写
	【参数:】字符串(string)
	【返回:】字符串(string)
	【备注:】
*/
func Ucword(str string) string {
	return strings.Title(str)
}


/*
	【名称:】移除字符串两侧的字符
	【参数:】字符串(string)，要移除的字符(string)
	【返回:】字符串(string)
	【备注:】
*/
func Trim(str, charlist string) string {
	if charlist == "" {
		charlist = " \r\n\t\x0B"
	}
	return strings.Trim(str, charlist)
}


/*
	【名称:】去除全部空格
	【参数:】字符串(string)
	【返回:】字符串(string)
	【备注:】
*/
func TrimAll(str string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.Join(strings.Fields(str),"")
}


/*
	【名称:】移除字符串左侧的字符
	【参数:】字符串(string)，要移除的字符(string)
	【返回:】字符串(string)
	【备注:】
*/
func Ltrim(str, charlist string) string {
	if charlist == "" {
		charlist = " \r\n\t\x0B"
	}
	return strings.TrimLeft(str, charlist)
}

// @title: 移除字符串右侧的字符
// @param: 字符串(string)，要移除的字符(string)
// @return: 字符串(string)
// @description:
// @date: 2022/7/4 22:07
func Rtrim(str, charlist string) string {
	if charlist == "" {
		charlist = " \r\n\t\x0B"
	}
	return strings.TrimRight(str, charlist)
}

/*
	【名称:】从字符串右端移除字符
	【参数:】字符串(string)，要移除的字符(string)
	【返回:】字符串(string)
	【备注:】
*/
func Chop(str, charlist string) string {
	return Rtrim(str, charlist)
}

/*
	【名称:】返回字符串中第一个字符的 ASCII 值
	【参数:】字符串(string)
	【返回:】ASCII值(int)
	【备注:】
*/
func Ord(str string) int {
	if str == "" {
		return 0
	}
	s := []rune(str)
	return int(s[0])
}


/*
	【名称:】ASCII值返回字符
	【参数:】ASCII值(int32)
	【返回:】字符(string)
	【备注:】
*/
func Chr(ascii int32) string {
	return string(ascii)
}


/*
	【名称:】把字符串重复指定的次数
	【参数:】字符串(string)，次数(int)
	【返回:】字符(string)
	【备注:】
*/
func StrRepeat(input string, multiplier int) string {
	return strings.Repeat(input, multiplier)
}

// @title: 字符替换
// @param: 要替换的字符串(string)，旧字符(string)，新字符(string)
// @return: string
// @description:
// @date: 2022/7/4 22:14
func StrReplace(strs string, oldStr string, newStr string) string {
	return strings.Replace(strs, oldStr, newStr, -1)
}