package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	fmt.Println(max(1, 2, -3), min(1, 2, -3))
	// fmt.Println(min())

	fmt.Println(max2(2, 3, 1), min2(2, 3, 1))
}

func max(nums ...int) int {
	if len(nums) == 0 {
		fmt.Fprintln(os.Stderr, "引数は少なくとも１つは必要です。")
		os.Exit(1)
	}
	sort.Ints(nums)
	return nums[len(nums)-1]
}

func min(nums ...int) int {
	if len(nums) == 0 {
		fmt.Fprintln(os.Stderr, "引数は少なくとも１つは必要です。")
		os.Exit(1)
	}
	sort.Ints(nums)
	return nums[0]
}

func max2(n int, nums ...int) int {
	slice := append(nums, n)
	sort.Ints(slice)
	return slice[len(slice)-1]
}

func min2(n int, nums ...int) int {
	slice := append(nums, n)
	sort.Ints(slice)
	return slice[0]
}
