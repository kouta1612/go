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

func TestFrenchPalindrome(t *testing.T) {
	if !IsParindrome("été") {
		t.Error(`IsParindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsParindrome(input) {
		t.Errorf(`IsParindrome(%q) = false`, input)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsParindrome("palindrome") {
		t.Error(`IsParindrome("palindrome") = true`)
	}
}
