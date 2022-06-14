package api

import (
	"ft-interview/internal/api/modal"
)

func IsDuplicateAnswers(answers []modal.FinishTestRequestItem) bool {

	questions := make(map[int]int)
	for _, data := range answers {
		if _, ok := questions[data.QuestionId]; ok {
			return true
		}
		questions[data.QuestionId] = 1
	}

	return false
}

func FindUserOrderByCorrectAnswerCount(userId int, userAnswers map[int]int) (count int) {

	for key, value := range userAnswers {
		if key == userId {
			continue
		}

		if userAnswers[userId] > value {
			count++
		}
	}

	return count
}
