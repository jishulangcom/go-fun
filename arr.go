/*
	【包名:】数组相关函数
	【作者:】技术狼
*/
package fun

import (
	"context"
	"errors"
	"reflect"
	"time"
)

/*
	【名称:】一维数组求和
	【参数:】arr(interface)
	【返回:】和(float64)，err(error)
	【备注:】可用于数组、切片；小数点计算会有精度问题
*/
func ArraySum(arr interface{}) (sum float64, err error) {
	list := reflect.ValueOf(arr)
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < list.Len(); i++ {
			val := list.Index(i)
			v, err := InterfaceToFloat64(val.Interface())
			if err != nil {
				return 0, err
			}
			sum += v
		}
	default:
		return 0, errors.New("arr参数必须是数组或切片")
	}
	return
}

/*
	【名称:】变量是否为切片
	【参数:】变量(interface)
	【返回:】true/false(bool)
*/
func IsSlice(val interface{}) bool {
	if GetType(val) == "slice" {
		return true
	}
	return false
}

// @title: 查找一维数组中某元素的全部index
// @auth: 技术狼(jishulang.com)
// @date: 2022/9/22 21:51
func ArrIndexOfIntAll(arr []int, coNub int, target int, timeout uint) []int {
	l := len(arr)
	sp := l / coNub
	ctx, _ := context.WithCancel(context.Background())
	ch := make(chan int)

	for i := 0; i < coNub; i++ {
		go func(i int, ctx context.Context) {
			start := i * sp
			curArr := arr[i*sp : (i+1)*sp]
			for k, v := range curArr {
				select {
				case <-ctx.Done():
					return
				default:
				}

				if v == target {
					index := start + k
					ch <- index
				}
			}
		}(i, ctx)
	}

	indexArr := []int{}
	select {
	case index := <-ch:
		indexArr = append(indexArr, index)
	case <-time.After(time.Duration(timeout/1000) * time.Second):

	}

	return indexArr
}

// @title: 查找一维数组中某元素的一个index
// @auth: 技术狼(jishulang.com)
// @date: 2022/9/22 21:51
func ArrIndexOfIntOne(arr []int, coNub int, target int, timeout uint) *int {
	l := len(arr)
	sp := l / coNub
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)

	for i := 0; i < coNub; i++ {
		go func(i int, ctx context.Context) {
			start := i * sp
			curArr := arr[i*sp : (i+1)*sp]
			for k, v := range curArr {
				select {
				case <-ctx.Done():
					return
				default:
				}

				if v == target {
					index := start + k
					ch <- index
				}
			}
		}(i, ctx)
	}

	select {
	case index := <-ch:
		cancel()
		return &index
	case <-time.After(time.Duration(timeout) * time.Second):

	}

	return nil
}

// @title: 查找一维数组中某元素的全部index
// @auth: 技术狼(jishulang.com)
// @date: 2022/9/22 21:51
func ArrIndexOfStringAll(arr []int, coNub int, target int, timeout uint) []int {
	l := len(arr)
	sp := l / coNub
	ctx, _ := context.WithCancel(context.Background())
	ch := make(chan int)

	for i := 0; i < coNub; i++ {
		go func(i int, ctx context.Context) {
			start := i * sp
			curArr := arr[i*sp : (i+1)*sp]
			for k, v := range curArr {
				select {
				case <-ctx.Done():
					return
				default:
				}

				if v == target {
					index := start + k
					ch <- index
				}
			}
		}(i, ctx)
	}

	indexArr := []int{}
	select {
	case index := <-ch:
		indexArr = append(indexArr, index)
	case <-time.After(time.Duration(timeout/1000) * time.Second):

	}

	return indexArr
}

// @title: 查找一维数组中某元素的一个index
// @auth: 技术狼(jishulang.com)
// @date: 2022/9/22 21:51
func ArrIndexOfStringOne(arr []int, coNub int, target int, timeout uint) *int {
	l := len(arr)
	sp := l / coNub
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)

	for i := 0; i < coNub; i++ {
		go func(i int, ctx context.Context) {
			start := i * sp
			curArr := arr[i*sp : (i+1)*sp]
			for k, v := range curArr {
				select {
				case <-ctx.Done():
					return
				default:
				}

				if v == target {
					index := start + k
					ch <- index
				}
			}
		}(i, ctx)
	}

	select {
	case index := <-ch:
		cancel()
		return &index
	case <-time.After(time.Duration(timeout) * time.Second):

	}

	return nil
}
