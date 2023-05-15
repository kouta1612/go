package word

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsParindrome("detartrated") {
		t.Error(`IsParindrome("detartrated") = false`)
	}
	if !IsParindrome("kayak") {
		t.Error(`IsParindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsParindrome("palindrome") {
		t.Error(`IsParindrome("palindrome") = true`)
	}
}
