package utils

func GetAnswerIndexFromLetter(answer string) int {
	answerMap := map[string]int{"A": 0, "B": 1, "C": 2, "D": 3}
	return answerMap[answer]
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
