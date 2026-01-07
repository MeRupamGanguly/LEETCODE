## These notes summarize the logic for the Go implementations of Leetcode problems.
---

## 1. Two Sum

**File:** `1-Two-Sum.go`

* **Goal:** Find two indices such that the numbers at those indices add up to a specific target.
* **Logic:** The code uses a hash map to keep track of numbers it has already seen and their corresponding indices.
* **Process:** For every number in the array, it calculates a "complement" by subtracting the current number from the target. It then checks the map to see if that complement has already been encountered; if so, it returns the current index and the stored index. If not, it saves the current number and its index into the map for future checks.
```go
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
```
## 2. Valid Anagram

**File:** `2-Valid-Anagram.go`

* **Goal:** Determine if one string is a rearrangement of another.
* **Logic:** The code verifies if both strings have the exact same characters with the exact same frequencies.
* **Process:** It first checks if the strings are of equal length, returning false if they are not. It then populates a frequency map using the characters of the first string. Finally, it iterates through the second string, decrementing the counts in the map. If a character is missing from the map or its count is already zero, the strings are not anagrams.
```go
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
```
## 3. Contains Duplicate

**File:** `3-Contains-Duplicate.go`

* **Goal:** Check if any value appears at least twice in an array.
* **Logic:** The code uses a map to record the frequency of each integer as it iterates through the slice.
* **Process:** For every element, it increments a counter in the `freqMap`. Immediately after incrementing, it checks if the count has reached **2**. If it has, the function returns `true`. If the loop completes without this condition being met, it returns `false`.
```go
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
```
## 4. Single Number

**File:** `4-Single-Number.go`

* **Goal:** In an array where every element appears more than once except for one, find that unique element.
* **Logic:** The code performs a two-pass approach using a frequency map to count occurrences.
* **Process:** The first pass iterates through the entire array to count how many times each number appears. The second pass iterates through the resulting map to identify and return the key that has a frequency value of **1**.
```go
func singleNumber(nums []int) int {
	freqMap := make(map[int]int)
	for i := range nums {
		freqMap[nums[i]]++
	}
	for k, v := range freqMap {
		if v == 1 {
			return k
		}
	}
	return -1
}
```
## 5. Intersection of Two Arrays

**File:** `5-Intersection-of-Arrays.go`

* **Goal:** Find the unique common elements present in two different arrays.
* **Logic:** The code uses two maps: one to track elements seen in the first array and another to ensure the output contains only unique values.
* **Process:** It first stores every element from `nums1` into `seenMap`. Then, it iterates through `nums2`; if an element exists in `seenMap` and has not yet been added to the result (tracked by `uniqMap`), it appends that element to the output list.
```go
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
```
## 6. Majority Element

**File:** `6-Majority-Element.go`

* **Goal:** Identify the element that appears most frequently in an array.
* **Logic:** The code utilizes a frequency map to count all elements and then finds the maximum value among those counts.
* **Process:** It first loops through the array to build a complete map of element frequencies. After the map is built, it iterates through the map keys and values, updating a `max` variable and a `majorityElem` variable whenever it finds a frequency higher than the current maximum.
```go
func majorityElement(nums []int) int {

	majorityElem := nums[0]
	max := 0
	freqMap := make(map[int]int)
	for i := range nums {
		freqMap[nums[i]]++
	}
	for k, v := range freqMap {
		if v > max {
			max = v
			majorityElem = k
		}
	}
	return majorityElem
}
```
## 7. Group Anagrams
**File:** 7-Group-Anagrams.go  
**Goal:** Group words that contain the same characters into sub-lists.  

**Logic:**  
The code creates a unique character-frequency "signature" for each string to use as a map key.  

**Process:**  
- Initialize a map where the key is a string representation of character counts and the value is a slice of strings.  
- For each word, create a rune array of size 26 (for 'a'-'z') to count occurrences.  
- Convert this count array into a string so all anagrams produce the same key.  
- Append the original word to the map at that specific key.  
- Return all map values as a 2D slice.  

```go
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
```
---

## 8. Top K Frequent Elements
**File:** 8-Top-K-Frequent-Elements.go  
**Goal:** Return the *k* most frequent numbers in an array.  

**Logic:**  
It uses a frequency map followed by a "bucket" mapping where frequencies act as indices.  

**Process:**  
- Populate `freqMap` to count how many times each integer appears.  
- Reverse this relationship in `countMap`, where the key is the frequency and the value is a slice of numbers with that frequency.  
- Iterate backward from the highest possible frequency (the array length) to find the most frequent items.  
- Collect these items into an output slice until the top *k* elements are gathered.  

```go

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
```
---

## 9. Sort Characters By Frequency
**File:** 9-Sort-Characters-By-Frequency.go  
**Goal:** Sort a string so characters appearing more often come first.  

**Logic:**  
Similar to bucket sort, characters are grouped by their counts and then rebuilt into a string in descending order of those counts.  

**Process:**  
- Count character occurrences using `freqMap`.  
- Build `countMap` to group all characters that share the same frequency.  
- Starting from the largest possible frequency index, retrieve the characters from the map.  
- For each character found, append it to the result rune slice *i* times (where *i* is the current frequency).  
- Return the final string.  
```go
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
```
---

## 10. First Unique Character in a String
**File:** 10-First-Unique-Character.go  
**Goal:** Find the index of the first character that appears only once in the string.  

**Logic:**  
A two-pass approach is used to count character frequencies and then check them in the original sequence.  

**Process:**  
- **Pass 1:** Iterate through the string once, incrementing counts for each character in `freqMap`.  
- **Pass 2:** Iterate through the string a second time from start to finish.  
  - For each character, check its value in the map.  
  - If the value is 1, immediately return the current index.  
- If the loop completes without finding a unique character, return `-1`.  
```go
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
```
---

