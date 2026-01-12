package main

func firstUniqChar(s string) int {
	freqMap := make(map[rune]int)
	for _, r := range s {
		freqMap[r]++
	}
	for i, r := range s {
		if freqMap[r] == 1 {
			return i
		}
	}
	return -1
}

// func main() {
// 	fmt.Println(firstUniqChar("leetcode"))
// 	fmt.Println(firstUniqChar("loveleetcode"))
// 	fmt.Println(firstUniqChar("aabb"))
// 	fmt.Println(firstUniqChar("aaaaaaaaaaavvvvvvvvvvvvvvvvvvvbbbbbbbbbbbbbbbbbbbbcde"))
// }
