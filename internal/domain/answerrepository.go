package domain

import (
	"fmt"
	"ft-interview/internal/domain/entity"
	"ft-interview/pkg/cache"
)

const (
	answerTableName = "answer_%d"
)

type AnswerRepository interface {
	GetByQuestionId(questionId int) []entity.Answer
}

type answerRepository struct {
	database cache.Cache
}

func (self *answerRepository) GetByQuestionId(questionId int) []entity.Answer {

	cached, found := self.database.Retrieve(fmt.Sprintf(answerTableName, questionId))

	if found {
		return cached.([]entity.Answer)
	}

	return self.initialize()
}

func (self *answerRepository) initialize() (answers []entity.Answer) {

	answers = []entity.Answer{
		{Id: 1, Content: "answer1", QuestionId: 1},
		{Id: 2, Content: "answer2", QuestionId: 1},
		{Id: 3, Content: "answer3", QuestionId: 1},
		{Id: 4, Content: "answer1", QuestionId: 2},
		{Id: 5, Content: "answer2", QuestionId: 2},
		{Id: 6, Content: "answer3", QuestionId: 2},
	}

	groupingAnswer := groupingAnswerByQuestionId(answers)

	for key, value := range groupingAnswer {
		self.database.Set(fmt.Sprintf(answerTableName, key), value)
	}

	return answers
}

func groupingAnswerByQuestionId(answers []entity.Answer) map[int][]entity.Answer {

	questionAnswers := make(map[int][]entity.Answer)

	for _, answer := range answers {
		questionAnswers[answer.QuestionId] = append(questionAnswers[answer.QuestionId], entity.Answer{
			Id:         answer.Id,
			QuestionId: answer.QuestionId,
			Content:    answer.Content,
		})
	}

	return questionAnswers
}

func NewAnswerRepository(database cache.Cache) AnswerRepository {
	return &answerRepository{
		database: database,
	}
}
