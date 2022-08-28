package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	question, answer string
}

// run: go build . && ./quiz -csv=problems.csv
func main() {
	filename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *filename))
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	var problems []Problem
	for _, line := range lines {
		problems = append(problems, Problem{line[0], strings.TrimSpace(line[1])})
	}

	score := 0
	for i, p := range problems {
		var userInput string
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		fmt.Scanf("%s\n", &userInput)

		if userInput == "break" {
			exit("Exiting the program.")
		} else {
			if _, err := strconv.Atoi(userInput); err != nil {
				fmt.Printf("%q is not a number.\n", userInput)
			}
			if userInput == p.answer {
				score++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", score, len(problems))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
