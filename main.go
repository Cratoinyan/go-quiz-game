package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	problemsCsv, err := os.Open("problems.csv")

	if err != nil {
		log.Fatal(err)
	}

	score := 0

	defer problemsCsv.Close()

	csvReeader := csv.NewReader(problemsCsv)

	data, err := csvReeader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reflect.TypeOf(data))

	fmt.Println("-----Quiz is starting-----")

	for i := 0; i < len(data); i++ {
		fmt.Println(data[i][0], "is?")

		var solution string
		fmt.Scanf("%s", &solution)

		if solution == data[i][1] {
			fmt.Println("correct")
			score++
		} else {
			fmt.Println("wrong")
		}
	}

	fmt.Println("your score is ", score)
}
