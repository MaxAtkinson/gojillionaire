package questions

import (
	"testing"
)

func TestReadQuestions(t *testing.T) {
	questions := ReadQuestions("test_data/questions.json")
	
	questionCount := len(questions)
	if questionCount == 0 {
		t.Errorf("got %q, wanted %q", questionCount, "> 0")
	}
}