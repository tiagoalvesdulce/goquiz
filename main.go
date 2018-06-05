package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	quizFilename := flag.String("quizFilename", "problems.csv", "The name of the file where the quiz is")
	flag.Parse()
	fmt.Println("Loading quiz: ", *quizFilename)
	quizFile, err := os.Open(*quizFilename)
	if err != nil {
		log.Fatalf("Error openning quiz file. Err: %s", err)
	}
	reader := csv.NewReader(bufio.NewReader(quizFile))
	var totalQuestions, rightAnswers int
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not read. Err: %s", err)
		}
		question := line[0]
		answer := line[1]

		var userAnswer string
		fmt.Printf("%s = ", question)
		_, err = fmt.Scan(&userAnswer)
		if err != nil {
			log.Fatalf("Error getting answer from user. Err: %s", err)
		}

		if userAnswer == answer {
			rightAnswers++
		}
		totalQuestions++
	}
	fmt.Printf("You got %d/%d\n", rightAnswers, totalQuestions)
}
