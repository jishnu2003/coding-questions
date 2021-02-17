package main

import "fmt"

//Problem Statement
//Given an array of sorted numbers and a target sum, find a pair in the array whose sum is equal to the given target.
//
//Example 1:
//
//Input: [1, 2, 3, 4, 6], target=6
//Output: [2, 4]
//Explanation: The numbers at index 1 and 3 add up to 6: 2 + 4 = 6
//Example 2:
//
//Input: [2, 5, 9, 11], target=11
//Output: [2, 9]
//Explanation: The numbers at index 0 and 2 add up to 11: 2 + 9 = 11

func main() {
	arr := []int{1, 2, 3, 4, 6}
	ts := 6
	res := pairWithTargetSum(arr, ts)
	fmt.Printf("Pair with target sum %d are %v\n", ts, res)

	arr = []int{2, 5, 9, 11}
	ts = 11
	res = pairWithTargetSum(arr, ts)
	fmt.Printf("Pair with target sum %d are %v\n", ts, res)
}

func pairWithTargetSum(arr []int, ts int) [2]int {
	for i, j := 0, len(arr)-1; i <= j; {
		if arr[i]+arr[j] == ts {
			return [2]int{arr[i], arr[j]}
		} else if arr[i]+arr[j] < ts {
			i += 1
		} else {
			j += 1
		}
	}
	return [2]int{-1, -1}
}
