package main

import (
	"fmt"
)

func isValidNumber(s string, index int) bool {
	if index >= len(s) {
		return true
	}
	if s[index] < '0' || s[index] > '9' {
		return false
	}
	return isValidNumber(s, index+1)
}

// count total different pair
func countMismatch(s []rune, left, right int) int {
	if left >= right {
		return 0
	}
	if s[left] != s[right] {
		return 1 + countMismatch(s, left+1, right-1)
	}
	return countMismatch(s, left+1, right-1)
}

// Make Minimal Palindrome Before Upgrading
func makePalindrome(original []rune, left, right int, result []rune, changed []bool) ([]rune, []bool) {
	if left > right {
		return result, changed
	}
	if left == right {
		result[left] = original[left]
		return result, changed
	}

	if original[left] == original[right] {
		result[left] = original[left]
		result[right] = original[right]
		return makePalindrome(original, left+1, right-1, result, changed)
	}

	// beda → samakan ke digit lebih besar
	if original[left] > original[right] {
		result[left] = original[left]
		result[right] = original[left]
		changed[left] = true
	} else {
		result[left] = original[right]
		result[right] = original[right]
		changed[right] = true
	}
	return makePalindrome(original, left+1, right-1, result, changed)
}

// Upgrading Polindrome To Highest Number
func maximizePalindrome(result []rune, left, right int, kLeft int, changed []bool) ([]rune, int) {
	if left > right {
		return result, kLeft
	}
	if left == right {
		if kLeft > 0 && result[left] != '9' {
			result[left] = '9'
			kLeft--
		}
		return result, kLeft
	}

	if result[left] == '9' && result[right] == '9' {
		return maximizePalindrome(result, left+1, right-1, kLeft, changed)
	}

	if (changed[left] || changed[right]) && kLeft >= 1 {
		result[left] = '9'
		result[right] = '9'
		kLeft--
		return maximizePalindrome(result, left+1, right-1, kLeft, changed)
	}

	if kLeft >= 2 {
		result[left] = '9'
		result[right] = '9'
		kLeft -= 2
		return maximizePalindrome(result, left+1, right-1, kLeft, changed)
	}

	return maximizePalindrome(result, left+1, right-1, kLeft, changed)
}

func HighestPalindrome(s string, k int) string {
	if len(s) == 0 {
		return "-1"
	}
	if !isValidNumber(s, 0) {
		return "-1"
	}

	original := []rune(s)
	n := len(original)

	mismatch := countMismatch(original, 0, n-1)
	if mismatch > k {
		return "-1"
	}

	result := make([]rune, n)
	changed := make([]bool, n)
	result, changed = makePalindrome(original, 0, n-1, result, changed)

	kLeft := k - mismatch
	result, _ = maximizePalindrome(result, 0, n-1, kLeft, changed)

	return string(result)
}

func main() {
	var s string
	var k int

	fmt.Print("Number: ")
	_, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Println("Error reading input number:", err)
		return
	}

	fmt.Print("Limit: ")
	_, err = fmt.Scanln(&k)
	if err != nil {
		fmt.Println("Error reading change limit:", err)
		return
	}

	result := HighestPalindrome(s, k)
	fmt.Println("Highest Palindrome:", result)
}
