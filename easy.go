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
func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var len int = len(s)
	var sum int
	if len == 1 {
		return romanMap[s[0]]
	}
	for i := 1; i < len; i++ {
		if romanMap[s[i-1]] >= romanMap[s[i]] {
			sum += romanMap[s[i-1]]
		} else {
			sum += romanMap[s[i]] - romanMap[s[i-1]]
			i++
		}
		if i == len-1 {
			sum += romanMap[s[i]]
		}
	}
	return sum
}

// 14.
func longestCommonPrefix(strs []string) string {
	strsLen := len(strs)
	if strsLen == 0 {
		return ""
	}
	if strsLen == 1 {
		return strs[0]
	}
	same := strs[0]
	for _, v := range strs[1:] {
		var commonPrefix []byte
		var cycleLen int
		if len(same) >= len(v) {
			cycleLen = len(v)
		} else {
			cycleLen = len(same)
		}
		for i := 0; i < cycleLen; i++ {
			if v[i] == same[i] {
				commonPrefix = append(commonPrefix, same[i])
			} else {
				break
			}
		}
		same = string(commonPrefix)
	}
	return same
}

// 20.
func isValid(s string) bool {
	stack := make([]rune, 0)
	validMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	sLen := len(s)
	if sLen == 0 {
		return true
	}
	if sLen%2 == 1 {
		return false
	}

	for _, v := range s {
		switch v {
		case '(', '[', '{':
			stack = append(stack, v)
		case ')', ']', '}':
			if len(stack) == 0 || validMap[v] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// 21.
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var newList = &ListNode{}
	firstList := newList
	for list1 != nil || list2 != nil {
		if list1 == nil {
			newList.Next = list2
			break
		} else if list2 == nil {
			newList.Next = list1
			break
		} else {
			if list1.Val > list2.Val {
				newList.Next = &ListNode{Val: list2.Val}
				list2 = list2.Next
			} else {
				newList.Next = &ListNode{Val: list1.Val}
				list1 = list1.Next
			}
		}

		newList = newList.Next
	}
	if firstList == nil {
		return nil
	} else {
		return firstList.Next
	}
}

// 26.
func removeDuplicates(nums []int) int {
	var valuesMap = make(map[int]bool, len(nums))
	var count int
	for _, v := range nums {
		if _, ok := valuesMap[v]; !ok {
			nums[count] = v
			valuesMap[v] = true
			count++
		}
	}
	return len(valuesMap)
}

// 27.
func removeElement(nums []int, val int) int {
	var count int
	for _, v := range nums {
		if v != val {
			nums[count] = v
			count++
		}
	}
	return count
}

// 28.
func strStr(haystack string, needle string) int {
	hLen := len(haystack)
	nLen := len(needle)

	if nLen == 0 {
		return 0
	}

	for i := 0; i < hLen; i++ {
		if haystack[i] == needle[0] {
			if nLen > hLen-i {
				return -1
			}
			if haystack[i:nLen+i] == needle {
				return i
			}
		}
	}
	return -1
}

// 35.
func searchInsert(nums []int, target int) int {
	nLen := len(nums)
	if nLen == 0 {
		return 0
	}
	for i := 0; i < nLen; i++ {
		if nums[i] >= target {
			return i
		}
	}

	return nLen
}

// 53.
func maxSubArray(nums []int) int {
	nLen := len(nums)
	if nLen == 0 {
		return 0
	}
	if nLen == 1 {
		return nums[0]
	}

	var max, sum int = nums[0], nums[0]
	for i := 1; i < nLen; i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if max < sum {
			max = sum
		}
	}

	return max
}

// 58.
func lengthOfLastWord(s string) int {
	sLen := len(s)
	var wordCount int
	for i := sLen - 1; i >= 0; i-- {
		if s[i] != ' ' {
			wordCount++
		} else {
			if wordCount != 0 {
				return wordCount
			}
		}
	}

	return wordCount
}

// 66.
func plusOne(digits []int) []int {
	dlen := len(digits)
	var carry int
	for i := dlen - 1; i >= 0; i-- {
		var afterCarry int
		if i == dlen-1 {
			if digits[i]+1 >= 10 {
				afterCarry = 1
			} else {
				afterCarry = 0
			}
			digits[i] = (digits[i] + 1) % 10
		} else {
			if digits[i]+carry >= 10 {
				afterCarry = 1
			} else {
				afterCarry = 0
			}
			digits[i] = (digits[i] + carry) % 10
		}
		carry = afterCarry
	}
	if carry == 1 {
		newDigits := make([]int, dlen+1)
		newDigits[0] = 1
		for i := 0; i < dlen-1; i++ {
			newDigits[i+1] = digits[i]
		}
		return newDigits
	}

	return digits
}

// 67.
func addBinary(a string, b string) string {
	var s, carry int
	var result string
	la := len(a) - 1
	lb := len(b) - 1
	for la >= 0 || lb >= 0 || carry != 0 {
		s = carry
		if la >= 0 {
			s += int(a[la] - '0')
			la--
		}
		if lb >= 0 {
			s += int(b[lb] - '0')
			lb--
		}
		carry = s / 2
		result = string(s%2+'0') + result
	}
	return result
}

// 69.
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left := 0
	right := x / 2
	var pivot int

	for left <= right {
		pivot = left + (right-left)/2
		pp := pivot * pivot

		if pp == x || (pp < x && (pivot+1)*(pivot+1) > x) {
			return pivot
		} else if pp > x {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return pivot
}

// 70.
// func climbStairs(n int) int {

// }
