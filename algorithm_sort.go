package fun

import (
	"math/rand"
	"time"
)

/*
*************************************
--------------技术狼------------------
-------------常用方法------------------
-------------排序算法------------------
*************************************
*/

/*
	【名称:】快速排序
	【参数:】
	【返回:】
	【备注:】
*/
func QuickSort(nums []int, left, right int) {
	val := nums[(left+right)/2]
	i, j := left, right
	for nums[j] > val {
		j--
	}
	for nums[i] < val {
		i++
	}
	nums[i], nums[j] = nums[j], nums[i]
	i++
	j--
	if i < right {
		QuickSort(nums, i, right)
	}
	if j > left {
		QuickSort(nums, left, j)
	}
}

/*
	【名称:】冒泡排序
	【参数:】
	【返回:】
	【备注:】
*/
func BubbleSort(nums []int){
	length := len(nums)
	for i:=1;i<length;i++{
		for j:=0;j<length-1;j++{
			if nums[j]>nums[i] {
				nums[j],nums[i] = nums[i],nums[j]
			}
		}
	}
}


/*
	【名称:】洗牌算法
	【参数:】
	【返回:】
	【备注:】
*/
func Shuffle(arr []int){
	rand.Seed(time.Now().UnixNano())
	var i, j int
	var temp int
	for i = len(arr) - 1; i > 0; i-- {
		j = rand.Intn(i + 1)
		temp = arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}
}