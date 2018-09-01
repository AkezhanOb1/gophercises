package main

import (
	"os"
	"encoding/csv"
	"io"
	"fmt"
	"time"
)

type Problem struct {
	Question string
	Answer string
}

func main() {

	csvFile,_ := os.Open("problems.csv")
	reader := csv.NewReader(csvFile)
	var problems []Problem
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		problems = append(problems, Problem{
			Question: line[0],
			Answer:   line[1],
		})
	}

	var userGuess string
	var correctAnswers int
	var wrongAnswers int

	fmt.Println("Press Any key to start up the timer ")
	fmt.Scanln( &userGuess)
	timer := time.NewTimer(30 * time.Second)

	go func() {
		for i := 0; i < len(problems); i++ {

			fmt.Print("Guess the result of this problem ", problems[i].Question, " ")
			fmt.Scanln( &userGuess)
			if userGuess == problems[i].Answer {
				correctAnswers ++
			}else {
				wrongAnswers ++
			}

		}
	}()


	<-timer.C
	fmt.Println("Timer  expired")


	fmt.Println("CORRECT ANSWERS ", correctAnswers)
	fmt.Println("WRONG ANSWERS ", wrongAnswers)

}
