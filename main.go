package main

import (
	"flag"
	"fmt"
	"path"

	"kylehodgetts.com/go-quiz/quiz"
	"kylehodgetts.com/go-quiz/timer"
)

var (
	timeLimit = flag.Int("timelimit", 30, "Time limit on the quiz")
	quizFile  = flag.String("csv", path.Join("data", "problems.csv"), "File path to quiz")
)

func main() {
	flag.Parse()
	questions := quiz.ParseQuestions(quizFile)
	correctAnswers := 0
	totalQuestions := len(questions)

	endQuiz := make(chan bool)

	go quiz.StartQuiz(questions, endQuiz)
	go timer.StartTimer(*timeLimit, endQuiz)

	<-endQuiz

	fmt.Printf("You got %d/%d answers correct\n", correctAnswers, totalQuestions)
}
