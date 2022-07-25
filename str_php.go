/*
	【包名:】字符(php)串相关函数
	【作者:】技术狼
*/
package fun

import (
	"strings"
)


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