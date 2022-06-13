package api

import (
	"ft-interview/internal/api/modal"
	"sort"
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

func FindUserOrderByCorrectAnswerCount(userId int, userAnswers map[int]int) (index float64) {

	keys := make([]int, 0, len(userAnswers))
	for key, _ := range userAnswers {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return userAnswers[keys[i]] < userAnswers[keys[j]]
	})

	for _, k := range keys {
		if k == userId {
			break
		}
		index++
	}

	return index
}
