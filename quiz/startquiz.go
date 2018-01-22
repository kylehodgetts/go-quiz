package quiz

import (
	"bufio"
	"fmt"
	"os"
)

func StartQuiz(questions []Question, stop chan<- bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for i, question := range questions {
		fmt.Printf("#%d: %s = ", i+1, question.Question)

		if !scanner.Scan() {
			fmt.Println("Couldn't read your answer. Moving on")
			continue
		}

		inputAnswer := scanner.Text()
		if inputAnswer == question.Answer {
			
		}
	}
	stop <- true
}
