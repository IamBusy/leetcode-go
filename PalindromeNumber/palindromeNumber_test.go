package PalindromeNumber

import (
	"testing"
	"fmt"
)

func TestIsPalindrome(t *testing.T) {
	rtn := isPalindrome(1)
	rtn2 := isPalindrome(1001)
	fmt.Println(rtn)
	fmt.Println(rtn2)
}
