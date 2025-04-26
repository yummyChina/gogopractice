package main

import "fmt"

func main() {

	nums := []int{1, 2, 2, 3, 3, 4, 4}
	reuslt := singleNumber(nums)
	fmt.Println(reuslt)

}

func singleNumber(nums []int) int {
	countMap := map[int]int{}
	for index := range nums {
		val := nums[index]
		numberCount := countMap[val]
		numberCount++
		countMap[val] = numberCount
	}

	for key, value := range countMap {
		if value == 1 {
			return key
		}
	}

	return 0
}
