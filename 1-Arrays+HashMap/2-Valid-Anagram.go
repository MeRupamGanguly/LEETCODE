package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	freqMap := make(map[rune]int)
	for _, r := range s {
		freqMap[r]++
	}
	for _, r := range t {
		v, ok := freqMap[r]
		if !ok {
			return false
		} else {
			if v == 0 {
				return false
			} else {
				freqMap[r]--
			}
		}
	}
	return true
}

// func main() {
// 	fmt.Println(isAnagram("anagram", "nagaram"))
// 	fmt.Println(isAnagram("rat", "car"))
// 	fmt.Println(isAnagram("ab", "abc"))
// 	fmt.Println(isAnagram("aacc", "ccac"))
// 	fmt.Println(isAnagram("aba", "aab"))
// }
