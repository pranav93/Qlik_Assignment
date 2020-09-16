package utils

// CheckPalindrome CheckPalindrome
func CheckPalindrome(input string) bool {
	runeList := []rune(input)

	length := len(runeList)
	for i := 0; i < length/2; i++ {
		if runeList[i] != runeList[length-1-i] {
			return false
		}
	}
	return true
}
