// https://leetcode.com/problems/two-sum/description/
package leetcode

import (
	"fmt"
	"havecompilerwillgo/test"
)

func RunTwoSum() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	expectedOut1 := []int{0, 1}

	nums2 := []int{3, 2, 4}
	target2 := 6
	expectedOut2 := []int{1, 2}

	nums3 := []int{2, 7, 11, 15}
	target3 := 9
	expectedOut3 := []int{0, 1}

	fmt.Println("==========Running Two Sum==========")
	test.AssertEqual(twoSum(nums1, target1), expectedOut1)
	test.AssertEqual(twoSum(nums2, target2), expectedOut2)
	test.AssertEqual(twoSum(nums3, target3), expectedOut3)
	fmt.Println("Success")
}

// Task: given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
// solution intuition:
//   - best possible runtime of an array of numbers is o(n). This is achievable easily if you keep track of the numbers
//     you've already seen or keep track of their complement. By using an o(1) storage structure like a hash, you do the
//     complement lookup quickly
func twoSum(nums []int, target int) []int {
	visitedMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		// if the complement is contained in the map
		if _, exists := visitedMap[complement]; exists {
			return []int{visitedMap[complement], i}
		}

		visitedMap[num] = i
	}

	// no pair adding up to the target was found
	return []int{}
}
