package main

import (
	"fmt"
	"time"
)
/*
解法一
两层循环，时间复杂度O(n^2)
*/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

/*
解法二
一层循环，
简历一个快查表，如果
*/
func twoSum2(nums []int, target int) []int {  
	m := make(map[int]int)
	for i := 0; i < len(nums); i++{
		result := target-nums[i]
		if _,ok:=m[result];ok {
			return []int{m[result],i}
		}
		m[nums[i]]=i
	}	
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	// 解法1
	start := time.Now() // 获取当前时间
	fmt.Println(twoSum(nums, target))
	elapsed := time.Now().Sub(start)
	fmt.Println("解法1:", elapsed)

	// 解法2
	start = time.Now() // 获取当前时间
	fmt.Println(twoSum2(nums, target))
	elapsed = time.Now().Sub(start)
	fmt.Println("解法2:", elapsed)
}
