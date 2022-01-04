package leetcode

// 1.
func twoSum(nums []int, target int) []int {
	var items = make(map[int]int, len(nums))
	for k, v := range nums {
		if itemKey, ok := items[target-v]; ok {
			return []int{itemKey, k}
		}
		items[v] = k
	}
	return []int{}
}

// 9.
func isPalindrome(x int) bool {
	var oldNum, newNum, position int
	oldNum = x
	for x > 0 {
		position = x % 10
		newNum = newNum*10 + position
		x = x / 10
	}

	return oldNum == newNum
}

// 13.
// func romanToInt(s string) int {

// }
