package fun
/*************************************
--------------技术狼------------------
-------------常用方法------------------
-------------查找算法------------------
*************************************/

/*
	【名称:】二分查找法
	【参数:】
	【返回:】
	【备注:】二分查找的基础是先做排序，然后做二分查找
*/
func BinarySearch(nums []int,left,right,val int) int{
	k := (left+right)/2
	if nums[k]>val {
		return BinarySearch(nums,left,k,val)
	}else if nums[k] < val {
		return BinarySearch(nums,k,right,val)
	}else{
		return k
	}
}

/*
	【名称:】斐波拉契数列-递归方法
	【参数:】
	【返回:】
	【备注:】 找出第n位数是什么，时间复杂度是n方
*/
func FibonacciRecursion(n int)int{
	if n==0 {
		return 0
	}else if n==1 {
		return 1
	}else{
		return FibonacciRecursion(n-1)+FibonacciRecursion(n-2)
	}
}

/*
	【名称:】斐波拉契数列-迭代法
	【参数:】
	【返回:】
	【备注:】 找出第n位数是什么，时间复杂度是n方
*/
func FibonacciFind(n int)int{
	x,y,fib := 0,1,0
	for i:=0;i<=n;i++{
		if i==0 {
			fib=0
		}else if i== 1{
			fib = x+y
		}else{
			fib=x+y
			x,y = y,fib
		}
	}
	return fib
}