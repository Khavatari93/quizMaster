package main

import (
	"encoding/csv"
	f "fmt"
	"log"
	"os"
)

// struct for Questions
type Questions struct {
	question string
	answer   string
}

func readCSV(path string) [][]string {

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func createQuestions(problems [][]string) []Questions {
	var questions []Questions
	for i, line := range problems {
		if i > 0 {
			var rec Questions
			for j, field := range line {
				if j == 0 {
					rec.question = field
				} else if j == 1 {
					rec.answer = field
				}
			}
			questions = append(questions, rec)
		}
	}
	return questions
}

func startQuiz(questions []Questions) {
	var counter int
	var answer string
	for index, element := range questions {
		f.Println("Question No.", index, "\n", element.question)
		f.Scanln(&answer)
		if answer == element.answer {
			counter++
		}
	}
	f.Println("Your Score", counter)
}

func main() {
	f.Println("Quiz Master!")
	problems := readCSV("questions/problems.csv")
	quiz := createQuestions(problems)
	startQuiz(quiz)
}
