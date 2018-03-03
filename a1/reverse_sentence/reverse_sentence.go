package main

import (
	"fmt"
	"strings"
)

func main() {
	result := reverseSentence("follow the yellow brick road")
	fmt.Println(result)
}

func reverseSentence(sentence string) string {
	words := strings.Split(sentence, " ")
	reversed := make([]string, 0)
	for i := len(words) - 1; i >= 0; i-- {
		reversed = append(reversed, words[i])
	}
	return strings.Join(reversed, " ")
}
