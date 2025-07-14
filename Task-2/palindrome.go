package main

func palindrome(input string) bool {
	inputLength := len(input)

	for idx, val := range input {
		if val != rune(input[inputLength-idx-1]) {
			return false
		}
	}
	return true
}
