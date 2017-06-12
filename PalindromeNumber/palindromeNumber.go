package PalindromeNumber

/**
 * @requirement:
 * 	Determine whether an integer is a palindrome.
 * 	Do this without extra space.
 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var temp = 0
	var originX = x

	for ;x >0;x=x/10  {
		temp = temp * 10 + x%10
	}
	return temp==originX
}