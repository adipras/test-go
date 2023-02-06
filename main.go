package main

import (
	"fmt"
	"strings"
)

func main() {
	if IsPalindromeInt(1221) {
		fmt.Println("The given number is Palindrome")
	} else {
		fmt.Println("The given number is NOT a Palindrome")
	}

	if IsPalindromeString("banana") {
		fmt.Println("The given string is Palindrome")
	} else {
		fmt.Println("The given string is NOT a Palindrome")
	}
}

func IsPalindromeInt(num int) bool {
	var reverseNumber int = 0
	var temp = num

	for temp > 0 {
		var remainder = temp % 10
		reverseNumber = (reverseNumber * 10) + remainder
		temp = temp / 10
	}

	if num == reverseNumber {
		return true
	} else {
		return false
	}
}

func IsPalindromeString(str string) bool {
	var reverseString string = ""
	var length = len(str)

	for i := length - 1; i >= 0; i-- {
		reverseString = reverseString + string(str[i])
	}

	// Case in-sensitive comparision
	if strings.ToLower(str) == strings.ToLower(reverseString) {
		return true
	} else {
		return false
	}
}
