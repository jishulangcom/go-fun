package fun

import (
	"log"
	"testing"
)

type JslStruct struct{}

func (y JslStruct) Fun1(i int, str string) (int, string) {
	return i, str
}

// 调用函数
func TestCallFun(t *testing.T) {
	res := CallFun(JslStruct{}, "Fun1", 10, "test")
	log.Println(res[0], res[1])
}
