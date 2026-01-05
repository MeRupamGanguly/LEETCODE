package main

func twoSum(nums []int, target int) []int {
	seenMap := make(map[int]int)
	for i := range nums {
		comp := target - nums[i]
		v, ok := seenMap[comp]
		if ok {
			return []int{v, i}
		} else {
			seenMap[nums[i]] = i
		}
	}
	return []int{-1, -1}
}

// func main() {
// 	arr1 := []int{2, 7, 11, 15}
// 	arr2 := []int{3, 2, 4}
// 	arr3 := []int{-1, -2, -3, -4, -5}
// 	target1 := 9
// 	target2 := 6
// 	target3 := -8
// 	fmt.Println(twoSum(arr1, target1))
// 	fmt.Println(twoSum(arr2, target2))
// 	fmt.Println(twoSum(arr3, target3))
// }
