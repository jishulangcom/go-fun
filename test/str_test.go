package test

import (
	"fmt"
	"github.com/jishulangcom/go-fun"
	"testing"
)

// @title: 生成随机字符
// @param: 个数(int)
// @return: 字符串(string)
// @description: a-z,A-Z,0-9
// @auth: 技术狼
// @date: 2022/7/3 21:07
func TestRandString(t *testing.T) {
	str := fun.RandString(10)
	fmt.Println(str) // lIETOH7p1x
}

// @title: 是否手机号
// @param: 手机号(string)
// @return: true/false(bool)
// @description:
// @auth: 技术狼
// @date: 2022/7/3 21:10
func TestIsMobile(t *testing.T) {
	is := fun.IsMobile("136000001")
	fmt.Println(is) // true
}

// @title: 手机号*号处理
// @param: 手机号(string)
// @return: 加*号的手机号(string)
// @description: 非手机号则直接返回原字符
// @auth: 技术狼
// @date: 2022/7/3 21:13
func TestMobileStar(t *testing.T) {
	mobile := fun.MobileStar("136000001")
	fmt.Println(mobile) // 136****01
}

// @title: 是否邮箱
// @param: 邮箱(string)
// @return: true/false(bool)
// @description:
// @auth: 技术狼
// @date: 2022/7/3 21:16
func TestIsEmail(t *testing.T) {
	email := fun.IsEmail("815437877@qq.com")
	fmt.Println(email) // true
}

// @title: string转int
// @param: 字符串(str)
// @return: int, error
// @description:
// @auth: 技术狼
// @date: 2022/7/3 21:18
func TestStrToInt(t *testing.T) {
	i, err := fun.StrToInt("123456")
	fmt.Println(i, err) // 123456 <nil>
}

// @title: string转int64
// @param: 字符串(str)
// @return: int64，error
// @description:
// @auth: 技术狼
// @date: 2022/7/3 21:25
func TestStrToInt64(t *testing.T) {
	i, err := fun.StrToInt64("123456")
	fmt.Println(i, err) // 123456 <nil>
}

// @title: string转uint64
// @param: 字符串(str)
// @return: uint64，error
// @description:
// @auth: 技术狼
// @date: 2022/7/3 21:28
func TestStrToUint64(t *testing.T) {
	i, err := fun.StrToUint64("123456")
	fmt.Println(i, err) // 123456 <nil>
}

// @title: string转float64
// @param: 字符串(str)
// @return: float64, error
// @description:
// @auth: 技术狼
// @date: 2022/7/3 21:29
func TestStrToFloat64(t *testing.T) {
	i, err := fun.StrToFloat64("123456.789")
	fmt.Println(i, err) // 123456 <nil>
}

// @title: string转*big.float
// @param: 字符串(str)
// @return: *big.Float, bool
// @description:
// @auth: 技术狼
// @date: 2022/7/3 21:31
func TestStrToBigFloat(t *testing.T) {
	i, err := fun.StrToBigFloat("123456.789")
	fmt.Println(i, err) // 123456 <nil>
}

// @title: string转*big.Int
// @param:
// @return:
// @description:
// @auth: 技术狼
// @date: 2022/7/3 21:03
func TestStrToBigInt(t *testing.T) {
	bigInt, ok := fun.StrToBigInt("0.123", 3)
	fmt.Println(bigInt.String(), ok) // 123 true
}

// @title: 某字符串是否存在数组中
// @param: 字符(string)，数组(array)
// @return: 布尔(bool)
// @description:
// @auth: 技术狼
// @date: 2022/7/3 21:33
func TestStrIsInArr(t *testing.T) {
	arr := []string{"技术狼", "jishulang.com"}
	is := fun.StrIsInArr("技术狼", arr)
	fmt.Println(is) // true
}
