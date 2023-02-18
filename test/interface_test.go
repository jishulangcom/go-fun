package test

import (
	"fmt"
	"github.com/jishulangcom/go-fun"
	"testing"
)

func TestInterfaceToInt64(t *testing.T) {
	var val int64
	var err error

	fmt.Println("======浮点=======")
	val, err = fun.InterfaceToInt64(1.111111)
	fmt.Println(val, err)

	fmt.Println("======浮点（科学计数）=======")
	val, err = fun.InterfaceToInt64("1.111111e+06")
	fmt.Println(val, err)
}
