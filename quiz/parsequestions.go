package quiz

import (
	"io/ioutil"
	s "strings"
)

type Question struct {
	Question string
	Answer   string
}

func ParseQuestions(quizFile *string) []Question {
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
