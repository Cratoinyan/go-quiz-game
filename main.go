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

	defer problemsCsv.Close()

	csvReeader := csv.NewReader(problemsCsv)

	data, err := csvReeader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reflect.TypeOf(data))
}
