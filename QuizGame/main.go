package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	f, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Fail to open file: %s", *csvFileName))
	}

	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Fail to parse the csv file")
	}

	// init a timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	problems := parseLines(lines)
	corect := 0

problemLoop:
	for index, problem := range problems {
		fmt.Printf("Problem #%d :%s = \n", index, problem.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("timeout")
			break problemLoop
		case answer := <-answerCh:
			if answer == problem.answer {
				corect++
			} else {
				fmt.Printf("Wrong answer:%s, it should be:%s\n", answer, problem.answer)
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", corect, len(problems))
}

func parseLines(lines [][]string) []problem {
	// prepare the slice, in case reallocate the memory in loop
	retProblems := make([]problem, len(lines))
	for index, line := range lines {
		retProblems[index] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return retProblems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
