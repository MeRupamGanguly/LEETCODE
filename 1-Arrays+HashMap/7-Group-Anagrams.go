package main

func groupAnagrams(strs []string) [][]string {
	if len(strs) < 1 {
		return [][]string{}
	}
	if len(strs) == 1 {
		return [][]string{strs}
	}
	anagramMap := make(map[string][]string)
	out := [][]string{}
	for _, s := range strs {
		charIntArr := make([]rune, 26)
		for _, r := range s {
			charIntArr[r-'a']++
		}
		anagramMap[string(charIntArr)] = append(anagramMap[string(charIntArr)], s)
	}
	for _, v := range anagramMap {
		out = append(out, v)
	}
	return out
}

// func main() {
// 	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
// 	fmt.Println(groupAnagrams([]string{""}))
// 	fmt.Println(groupAnagrams([]string{"a"}))
// 	fmt.Println(groupAnagrams([]string{"ddddddddddg", "dgggggggggg"}))
// }
