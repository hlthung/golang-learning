package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Problem struct {
	question, answer string
}

// Ex1 From https://github.com/gophercises/quiz
// Run: go build . && ./quiz -csv=problems.csv -limit=2
func main() {
	filename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse() // parse the command line into the defined flags

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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	score := 0
	//loop: // label
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		answerChannel := make(chan string)

		// creates a gorountine to wait for userInput while timer continue
		go func() {
			var userInput string
			fmt.Scanf("%s\n", &userInput)
			answerChannel <- userInput // sending userInput to the answerChannel
		}()

		select {
		case <-timer.C:
			fmt.Println("Time out!")
			return
			//break loop
		case userInput := <-answerChannel: // if we get an answer from answerChannel
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
	}
	// we can uncomment loop and break loop, so it comes to this line
	fmt.Printf("You scored %d out of %d.\n", score, len(problems))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
