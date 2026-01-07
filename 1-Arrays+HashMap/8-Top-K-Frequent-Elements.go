package main

func topKFrequent(nums []int, k int) []int {
	if len(nums) <= 1 {
		return nums
	}
	freqMap := make(map[int]int)
	countMap := make(map[int][]int)
	for i := range nums {
		freqMap[nums[i]]++
	}
	for k, v := range freqMap {
		countMap[v] = append(countMap[v], k)
	}
	out := []int{}
	for i := len(nums); i > 0; i-- {
		out = append(out, countMap[i]...)
	}
	return out[:k]
}

// func main() {
// 	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
// 	fmt.Println(topKFrequent([]int{1}, 1))
// 	fmt.Println(topKFrequent([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2))
// 	fmt.Println(topKFrequent([]int{-1, -1}, 1))
// }
