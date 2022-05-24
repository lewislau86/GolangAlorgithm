package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return j
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
}
