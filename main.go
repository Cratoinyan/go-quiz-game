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
	//add flags for our cli
	csvFileName := flag.String("csv", "problems.csv", "csv file with format 'question,answer'")
	quizTime := flag.Int("time", 30, "time limit of the quiz")
	flag.Parse()

	//open csv file and control if it is succesfull
	problemsCsv, err := os.Open(*csvFileName)

	if err != nil {
		log.Fatal(err)
	}
	//init score
	score := 0

	//function to pass to time.AfterFunc so the quiz will be ended after the set amount given with the time flag has passed
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
	//set the timer
	time.AfterFunc(time.Duration(*quizTime)*time.Second, finishQuiz)
	//ask every question read from the quiz
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
