package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	csvFileName := flag.String("csv", "problems.csv", "csv file with format 'question,answer'")
	quizTime := flag.Int("time", 30, "time limit of the quiz")

	flag.Parse()

	problemsCsv, err := os.Open(*csvFileName)

	if err != nil {
		log.Fatal(err)
	}
	score := 0

	finishQuiz := func() {
		fmt.Println("your score is ", score)
		os.Exit(0)
	}

	defer problemsCsv.Close()

	csvReeader := csv.NewReader(problemsCsv)

	data, err := csvReeader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----Quiz is starting-----")

	time.AfterFunc(time.Duration(*quizTime)*time.Second, finishQuiz)

	for i := 0; i < len(data); i++ {
		fmt.Println(data[i][0], "is?")

		var solution string
		fmt.Scanf("%s", &solution)

		if solution == data[i][1] {
			score++
		}
	}

	fmt.Println("your score is ", score)
}
