package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	s "strings"
)

var (
	timeLimit = flag.Int("timelimit", 10, "Time limit on the quiz")
	quizFile  = flag.String("file", path.Join("data", "problems.csv"), "File path to quiz")
)

type Question struct {
	Question string
	Answer   string
}

func parseQuestions() []Question {
	bytes, err := ioutil.ReadFile(*quizFile)
	if err != nil {
		panic("Cannot read quiz file")
	}

	contents := string(bytes)
	lines := s.Split(contents, "\n")

	questions := make([]Question, len(lines))
	for i, line := range lines {
		qa := s.Split(line, ",")
		questions[i] = Question{qa[0], qa[1]}
	}

	return questions
}

func main() {
	flag.Parse()
	questions := parseQuestions()
	scanner := bufio.NewScanner(os.Stdin)

	correctAnswers := 0
	totalQuestions := len(questions)

	for _, question := range questions {
		fmt.Printf("%s: ", question.Question)
		read := scanner.Scan()
		if !read {
			fmt.Println("Couldn't read your answer. Moving on")
			continue
		}
		inputAnswer := scanner.Text()
		if inputAnswer == question.Answer {
			correctAnswers++
		}
	}

	fmt.Printf("You got %d/%d answers correct\n", correctAnswers, totalQuestions)
}
