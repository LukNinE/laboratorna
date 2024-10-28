package main

import (
	"fmt"
	"strings"
)

func countVowels(s string) int {
	vowels := "а о у и і е А О У И І Е"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Print("Напишіть ваше слово: ")
	var input string
	fmt.Scanln(&input)

	vowelCount := countVowels(input)
	fmt.Printf("Число голосних в слові: \"%s\": %d\n", input, vowelCount)
}
