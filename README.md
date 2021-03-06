

#### bool

```go
package test

import (
	"fmt"
	"github.com/jishulangcom/go-fun"
	"testing"
)

// @title: 变量是否布尔类型
func TestIsBool(t *testing.T) {
	v := true
	is := fun.IsBool(v)
	fmt.Println(is) // true
}
```



#### string

```go
package test

import (
	"fmt"
	"github.com/jishulangcom/go-fun"
	"testing"
)

// @title: 生成随机字符
func TestRandString(t *testing.T) {
	str := fun.RandString(10)
	fmt.Println(str) // lIETOH7p1x
}

// @title: 是否手机号
func TestIsMobile(t *testing.T) {
	is := fun.IsMobile("136000001")
	fmt.Println(is) // true
}

// @title: 手机号*号处理
func TestMobileStar(t *testing.T) {
	mobile := fun.MobileStar("136000001")
	fmt.Println(mobile) // 136****01
}

// @title: 是否邮箱
func TestIsEmail(t *testing.T) {
	email := fun.IsEmail("815437877@qq.com")
	fmt.Println(email) // true
}

// @title: string转int
func TestStrToInt(t *testing.T) {
	i, err := fun.StrToInt("123456")
	fmt.Println(i, err) // 123456 <nil>
}

// @title: string转int64
func TestStrToInt64(t *testing.T) {
	i, err := fun.StrToInt64("123456")
	fmt.Println(i, err) // 123456 <nil>
}

// @title: string转uint64
func TestStrToUint64(t *testing.T) {
	i, err := fun.StrToUint64("123456")
	fmt.Println(i, err) // 123456 <nil>
}

// @title: string转float64
func TestStrToFloat64(t *testing.T) {
	i, err := fun.StrToFloat64("123456.789")
	fmt.Println(i, err) // 123456 <nil>
}

// @title: string转*big.float
func TestStrToBigFloat(t *testing.T) {
	i, err := fun.StrToBigFloat("123456.789")
	fmt.Println(i, err) // 123456 <nil>
}

// @title: string转*big.Int
func TestStrToBigInt(t *testing.T) {
	bigInt, ok := fun.StrToBigInt("0.123", 3)
	fmt.Println(bigInt.String(), ok) // 123 true
}

// @title: 某字符串是否存在数组中
func TestStrIsInArr(t *testing.T) {
	arr := []string{"技术狼", "jishulang.com"}
	is := fun.StrIsInArr("技术狼", arr)
	fmt.Println(is) // true
}
```

