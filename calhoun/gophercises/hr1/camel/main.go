package main

import (
	"fmt"
	"strings"
)

// go run main.go < camel.in
func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	//fmt.Println("Input is:", input)
	count := countCamel(input)
	fmt.Println(count)
}

func countCamel(input string) int {
	ans := 1
	for _, char := range input {
		str := string(char)
		if strings.ToUpper(str) == str {
			ans++
		}
	}
	return ans
}
