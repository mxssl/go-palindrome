package main

import "testing"

func TestPalindrome(t *testing.T) {
	resultRight := isPalindrome("racecar")
	resultWrong := isPalindrome("hello")

	if !resultRight || resultWrong {
		t.Error("Palindrome func is incorrect")
	}
}

func TestReverse(t *testing.T) {
	word := "hello"
	reverseWord := "olleh"
	result := reverse(word)

	if result != reverseWord {
		t.Errorf("Reverse func is incorrect. Got: %v, Want: %v",
			result, reverseWord)
	}
}
