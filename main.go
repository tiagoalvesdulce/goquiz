package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func countTime(limit int, done chan<- bool) {
	time.Sleep(time.Duration(limit) * time.Second)
	done <- true
}

func main() {
	quizFilename := flag.String("quizFilename", "problems.csv", "The name of the file where the quiz is")
	limit := flag.Int("limit", 30, "The time limit for a quiz in seconds")
	flag.Parse()

	var rightAnswers int
	var done = make(chan bool, 1)

	fmt.Println("---> Loading quiz: ", *quizFilename)
	quizFile, err := os.Open(*quizFilename)
	if err != nil {
		log.Fatalf("Error openning quiz file. Err: %s", err)
	}
	reader := csv.NewReader(quizFile)

	go countTime(*limit, done)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Could not read. Err: %s", err)
	}

	go func() {
		for i, problem := range lines {
			question := problem[0]
			answer := strings.Trim(strings.ToLower(problem[1]), " ")
			var userAnswer string

			fmt.Printf("Problem #%d: %s = ", i+1, question)

			_, err = fmt.Scan(&userAnswer)
			if err != nil {
				log.Fatalf("Error getting answer from user. Err: %s", err)
			}

			userAnswer = strings.Trim(strings.ToLower(userAnswer), " ")

			if userAnswer == answer {
				rightAnswers++
			}
		}
		done <- true
	}()
	<-done
	fmt.Printf("\nYou got %d/%d!\n", rightAnswers, len(lines))
}
