package main

func intersection(nums1 []int, nums2 []int) []int {
	seenMap := make(map[int]bool)
	uniqMap := make(map[int]bool)
	out := []int{}
	for i := range nums1 {
		seenMap[nums1[i]] = true
	}
	for i := range nums2 {
		if seenMap[nums2[i]] {
			if !uniqMap[nums2[i]] {
				out = append(out, nums2[i])
			}
			uniqMap[nums2[i]] = true
		}
	}
	return out
}

// func main() {
// 	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
// 	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
// 	fmt.Println(intersection([]int{1, 2, 3, 4}, []int{4}))
// 	fmt.Println(intersection([]int{1, 2}, []int{4, 5, 6, 2}))
// }
