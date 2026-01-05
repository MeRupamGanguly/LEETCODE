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
---

