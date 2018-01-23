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

func StartTimer(timeLimit time.Duration, stop chan<- bool) {
	time.Sleep(time.Duration(timeLimit * time.Second))
	stop <- true
}

func StartQuiz(questions []quiz.Question, stop chan<- bool) {
	scanner := bufio.NewScanner(os.Stdin)
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

	fmt.Printf("You have %d seconds to answer %d questions. GO!\n", *timeLimit, len(questions))
	go StartQuiz(questions, endQuiz)
	go StartTimer(time.Duration(*timeLimit), endQuiz)

	<-endQuiz

	fmt.Printf("You got %d/%d answers correct\n", correctAnswers, len(questions))
}
