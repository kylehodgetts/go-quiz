package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"time"

	"kylehodgetts.com/go-quiz/quiz"
)

var (
	timeLimit      = flag.Int("timelimit", 30, "Time limit on the quiz")
	quizFile       = flag.String("csv", path.Join("data", "problems.csv"), "File path to quiz")
	correctAnswers = 0
)

const (
	welcomeMessage = "You have %d seconds to answer %d questions. Press any key to begin.\n"
	endQuizMessage = "You got %d/%d answers correct\n"
)

func StartTimer(timeLimit time.Duration, stop chan<- bool) {
	time.Sleep(time.Duration(timeLimit * time.Second))
	stop <- true
}

func StartQuiz(questions []quiz.Question, scanner *bufio.Scanner, stop chan<- bool) {
	for i, question := range questions {
		fmt.Printf("#%d: %s = ", i+1, question.Question)

		if !scanner.Scan() {
			fmt.Println("Couldn't read your answer. Moving on")
			continue
		}

		inputAnswer := scanner.Text()
		if inputAnswer == question.Answer {
			correctAnswers++
		}
	}
	stop <- true
}

func main() {
	flag.Parse()
	questions := quiz.ParseQuestions(quizFile)
	endQuiz := make(chan bool)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(welcomeMessage, *timeLimit, len(questions))

	go StartQuiz(questions, scanner, endQuiz)
	go StartTimer(time.Duration(*timeLimit), endQuiz)

	<-endQuiz

	fmt.Printf(endQuizMessage, correctAnswers, len(questions))
}
