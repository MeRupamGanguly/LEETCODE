package main

func frequencySort(s string) string {
	freqMap := make(map[rune]int)
	countMap := make(map[int][]rune)
	for _, r := range s {
		freqMap[r]++
	}
	for k, v := range freqMap {
		countMap[v] = append(countMap[v], k)
	}
	out := []rune{}
	for i := len(s); i > 0; i-- {
		if len(countMap[i]) < 1 {
			continue
		}
		aR := countMap[i]
		for j := 0; j < len(countMap[i]); j++ {
			r := aR[j]
			for k := 0; k < i; k++ {
				out = append(out, r)
			}
		}
	}
	return string(out)
}

// func main() {
// 	fmt.Println(frequencySort("tree"))
// 	fmt.Println(frequencySort("cccaaa"))
// 	fmt.Println(frequencySort("Aabb"))
// 	fmt.Println(frequencySort("dApqqbh7pNtDplFTLnSf"))
// 	fmt.Println(frequencySort("oMd6DBkNgnXEAYxN0KBUZi2t7HOEZy77UOtZo3l119bKJOgaD3"))
// }
