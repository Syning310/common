package main

import "fmt"

func BubbleSort(arr *[]int) {

	tmp := 0

	for i := 0; i < len(*arr); i++ {

		for j := 0; j < len(*arr)-1-i; j++ {

			if (*arr)[j] > (*arr)[j+1] {
				// 交换
				tmp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = tmp
			}

		}

	}
}

func main() {

	var arr []int = []int{12, 5, 23, 66, 99, 10, 28, 77, 18}

	fmt.Printf("排序前： %v\n", arr)

	BubbleSort(&arr)

	fmt.Printf("排序后： %v\n", arr)

}
