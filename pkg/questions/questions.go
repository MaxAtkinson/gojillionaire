package questions

import (
	"encoding/json"
	"fmt"
	"github.com/MaxAtkinson/gojillionaire/pkg/utils"
	"io/ioutil"
	"os"
	"strings"
)

var VALID_ANSWERS = []string{"A", "B", "C", "D"}

type Answer struct {
	Text      string `json:text`
	IsCorrect bool   `json:isCorrect`
}

func (a *Answer) String() string {
	return a.Text
}

type Question struct {
	Text    string   `json:text`
	Answers []Answer `json:answers`
}

func (q *Question) String() string {
	var sb strings.Builder

	sb.WriteString(q.Text)
	sb.WriteString("\n\n")
	sb.WriteString(fmt.Sprintf("A) %s\n", q.Answers[0].Text))
	sb.WriteString(fmt.Sprintf("B) %s\n", q.Answers[1].Text))
	sb.WriteString(fmt.Sprintf("C) %s\n", q.Answers[2].Text))
	sb.WriteString(fmt.Sprintf("D) %s\n", q.Answers[3].Text))

	return sb.String()
}

func (q *Question) Ask() string {
	var answer string

	fmt.Println(q.String())
	fmt.Println("Enter your answer:")
	fmt.Scanf("%s", &answer)
	answer = strings.ToUpper(strings.TrimSpace(answer))

	if !utils.StringInSlice(strings.ToUpper(answer), VALID_ANSWERS) {
		fmt.Println(fmt.Sprintf("\nInvalid answer \"%s\" supplied. Must be one of %s\n",
			answer,
			VALID_ANSWERS))
		q.Ask()
	}
	return answer
}

func (q *Question) IsCorrectAnswer(answer string) bool {
	answerIndex := utils.GetAnswerIndexFromLetter(answer)
	return q.Answers[answerIndex].IsCorrect
}

func ReadQuestions(params ...string) []Question {
	var questions []Question
    var fileName string

    if len(params) == 0 {
        fileName = "data/questions.json"
    } else if len(params) == 1 {
        fileName = params[0]
    } else {
        panic("Invalid arguments supplied")
    }

	jsonFile, err := os.Open(fileName)
	defer jsonFile.Close()
	if err != nil {
		panic(err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(byteValue, &questions)

	return questions
}
