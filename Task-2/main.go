package main

import "fmt"

func main() {
	sentences := "Go function that calculates the frequency of each word in a string,"

	fmt.Println(wordFrequency(sentences))
	fmt.Println(palindrome(sentences))

	word := "aadaa"
	fmt.Println(palindrome(word))
}
