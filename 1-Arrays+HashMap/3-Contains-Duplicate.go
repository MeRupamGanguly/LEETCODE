package main

func containsDuplicate(nums []int) bool {
	freqMap := make(map[int]int)
	for i := range nums {
		freqMap[nums[i]]++
		if freqMap[nums[i]] == 2 {
			return true
		}
	}
	return false
}

// func main() {
// 	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
// 	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
// 	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
// 	fmt.Println(containsDuplicate([]int{7, 10, 5, 5, 6, 6, 4, 10, 5, 4, 9, 4, 9, 6, 5, 9, 6, 3, 6, 5, 6, 7, 7, 4, 9, 9, 10, 5, 8, 1, 8, 3, 2, 7, 5, 10, 1, 8, 5, 8, 4, 3, 6, 4, 9, 4, 2, 8, 3, 2, 2, 1, 5, 6, 3, 2, 6, 1, 8, 6, 2, 9, 1, 4, 5, 10, 8, 5, 10, 5, 10, 1, 4, 8, 3, 6, 4, 10, 9, 1, 1, 1, 2, 2, 9, 6, 6, 8, 1, 9, 2, 5, 5, 2, 1, 8, 5, 2, 3, 10}))
// 	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
// }
