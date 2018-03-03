package main

import (
	"fmt"
)

func main() {
	fruits := []string{"a", "a", "a", "a", "b", "b"}
	result := unique(fruits)
	fmt.Println(result)
}

func unique(arr []string) []string {
	uniques := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		if indexOf(arr[i], uniques) == -1 {
			uniques = append(uniques, arr[i])
		}
	}
	return uniques
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
