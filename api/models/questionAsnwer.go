package models

type QuestionAnswer map[string][]int

func (q QuestionAnswer) IsQuestionAnswered(username string, questionID int) bool {
	if answeredQuestions, ok := q[username]; ok {
		for _, q := range answeredQuestions {
			if q == questionID {
				return true
			}
		}
	}
	return false
}
