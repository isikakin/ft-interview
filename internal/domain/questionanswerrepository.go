package domain

import (
	"fmt"
	"ft-interview/internal/domain/entity"
	"ft-interview/pkg/cache"
)

const (
	questionAnswerKey = "question_%d"
)

type QuestionAnswerRepository interface {
	GetByQuestionId(questionId int) entity.QuestionAnswer
	GetUserAnswers() map[int]int
	Upsert(userId int, correctAnswerCount int)
}

type questionAnswerRepository struct {
	database cache.Cache
}

func (self *questionAnswerRepository) initialize() {

	answers := []entity.QuestionAnswer{
		{QuestionId: 1, CorrectAnswerId: 1},
		{QuestionId: 2, CorrectAnswerId: 1},
		{QuestionId: 3, CorrectAnswerId: 1},
		{QuestionId: 4, CorrectAnswerId: 1},
		{QuestionId: 5, CorrectAnswerId: 1},
	}

	for _, answer := range answers {
		self.database.Set(fmt.Sprintf(questionAnswerKey, answer.QuestionId), answer)
	}
}

func (self *questionAnswerRepository) GetUserAnswers() map[int]int {

	cached, found := self.database.Retrieve("answers")

	if found {

		return cached.(map[int]int)
	}

	return nil
}

func (self *questionAnswerRepository) Upsert(userId int, correctAnswerCount int) {

	userAnswers := make(map[int]int)

	cached, found := self.database.Retrieve("answers")

	if found {
		userAnswers = cached.(map[int]int)
	}

	userAnswers[userId] = correctAnswerCount

	self.database.Set("answers", userAnswers)
}

func (self *questionAnswerRepository) GetByQuestionId(questionId int) entity.QuestionAnswer {

	cached, found := self.database.Retrieve(fmt.Sprintf(questionAnswerKey, questionId))

	if found {
		return cached.(entity.QuestionAnswer)
	}

	self.initialize()

	return self.GetByQuestionId(questionId)
}

func NewQuestionAnswerRepository(database cache.Cache) QuestionAnswerRepository {
	return &questionAnswerRepository{
		database: database,
	}
}
