package main

import "fmt"

// Given an array of sorted numbers, remove all duplicates from it.
// You should not use any extra space; after removing the duplicates in-place
// return the length of the subarray that has no duplicate in it.

//Example 1:
//
//Input: [2, 3, 3, 3, 6, 9, 9]
//Output: 4
//Explanation: The first four elements after removing the duplicates will be [2, 3, 6, 9].
//Example 2:
//
//Input: [2, 2, 2, 11]
//Output: 2
//Explanation: The first two elements after removing the duplicates will be [2, 11].

func main() {
	arr := []int{2, 3, 3, 3, 6, 9, 9}
	arr = removeDuplicates(arr)
	fmt.Println(arr)

	arr = []int{2, 2, 2, 11}
	arr = removeDuplicates(arr)
	fmt.Println(arr)
}

func removeDuplicates(arr []int) []int {
	pos := 1 // Position to insert next non-duplicate element
	for i := 1; i < len(arr); i++ {
		//fmt.Printf("Comparing %d and %d\n", arr[pos-1], arr[i])
		if arr[pos-1] != arr[i] {
			//fmt.Printf("Moving %d to position %d\n", arr[i], pos)
			arr[pos] = arr[i]
			pos += 1
		}
	}
	arr = arr[:pos]
	//fmt.Printf("Total non duplicate elements are : %d\n", pos)
	return arr
}
