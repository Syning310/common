package main

import "fmt"

// 二分查找：递归法
func BinaryFind(arr []int, left int, right int, findVal int) {

	if left > right {
		fmt.Println("找不到...")
		return
	}

	// 取中间下标
	mid := (left + right) / 2

	if arr[mid] > findVal {
		BinaryFind(arr, left, mid-1, findVal)
	} else if arr[mid] < findVal {
		BinaryFind(arr, mid+1, right, findVal)
	} else {
		fmt.Println("找到了，下标为 ", mid)
	}
}


// 循环法
func PoolBinarySearch(arr []int, findVal int) int {

	left := 0
	right := len(arr) - 1

	for left <= right {
    
		mid := (left + right) / 2
    
		if arr[mid] == findVal {
			fmt.Println("找到了")
			return mid
		} else if arr[mid] > findVal {
			fmt.Println("值在左边")
			right = mid - 1
		} else if arr[mid] < findVal {
			fmt.Println("值在右边")
			left = mid + 1
		}

	}
	return -1
}




func main() {

	var arr []int = []int{1, 10, 30, 500, 1200, 3200, 9800, 12900}

	fmt.Println("arr = ", arr)

	BinaryFind(arr, 0, len(arr), 1200)

	index := PoolBinarySearch(arr, 500)

	fmt.Println("index = ", index)

}
