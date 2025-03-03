package main

import "fmt"

func main() {
	//nums := []int{1, 2, 0, 3, 4, 5, 6, 7}
	//nums := []int{-1, -100, 3, 99}
	//nums := []int{1, 2, 3, 4}
	nums := []int{-1, 1, 0, -3, 3}
	//nums := []int{0, 0, 0, 0, 1}
	//result := moveZeroes(nums)
	//res1 := changePosition(nums)
	//fmt.Println(result)
	//fmt.Println(res1)
	//fmt.Println(rotate(nums, 2))
	fmt.Println(ProductExceptSelf(nums))
}

func changePosition(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}

	return nums
}

func moveZeroes(nums []int) []int {
	left := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}

	return nums
}

func rotate(nums []int, k int) []int {
	for j := 0; j < k; j++ {
		for i := len(nums) - 1; i > 0; i-- {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}

	return nums
}

func ProductExceptSelf(nums []int) []int {
	var res = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sum := 1
		for j := 0; j < len(nums); j++ {
			if i != j {
				sum *= nums[j]
			}
		}

		res[i] = sum
	}

	return res
}
