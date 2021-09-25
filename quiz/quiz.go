package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type question struct {
	question, correctAnswer, userAnswer string
}

func (q question) isCorrect() bool {
	return q.correctAnswer == q.userAnswer
}

func main() {
	file := flag.String("csv", "problems.csv", "Name of the CSV file that contains quiz questions and answer")
	flag.Parse()
	data, err := os.ReadFile(*file)
	if err != nil {
		log.Fatal("Failed to read file: " + err.Error())
	}
	reader := csv.NewReader(strings.NewReader(string(data)))
	var correct, questions int
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var q question
		q.question = record[0]
		q.correctAnswer = record[1]
		fmt.Print(q.question + " = ")
		fmt.Scan(&q.userAnswer)
		if q.isCorrect() {
			correct += 1
		}
		questions += 1
	}
	fmt.Println(correct, "out of", questions, "correct")
}
