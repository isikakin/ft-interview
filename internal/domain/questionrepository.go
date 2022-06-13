package domain

import (
	"fmt"
	"ft-interview/internal/domain/entity"
	"ft-interview/pkg/cache"
)

const (
	questionKey             = "questions"
	separateQuestionsPrefix = "separatequestions_%d"
	separateQuestions       = "separatequestions"
)

type QuestionRepository interface {
	GetAll() []entity.Question
	GetAllSeparateQuestions() []entity.SeparateQuestion
	GetSeparateQuestionById(questionId int) *entity.SeparateQuestion
}

type questionRepository struct {
	database cache.Cache
}

func (self *questionRepository) GetAll() []entity.Question {

	cached, found := self.database.Retrieve(questionKey)

	if found {
		return cached.([]entity.Question)
	}

	return self.initialize()
}

func (self *questionRepository) GetAllSeparateQuestions() []entity.SeparateQuestion {

	cached, found := self.database.Retrieve(separateQuestions)

	if found {
		return cached.([]entity.SeparateQuestion)
	}

	return self.initializeSeparateQuestions(0)
}

func (self *questionRepository) GetSeparateQuestionById(questionId int) *entity.SeparateQuestion {

	cached, found := self.database.Retrieve(fmt.Sprintf(separateQuestionsPrefix, questionId))

	if found {
		founded := cached.(entity.SeparateQuestion)
		return &founded
	}

	foundedQuestion := self.initializeSeparateQuestions(questionId)

	if len(foundedQuestion) > 0 {
		return &foundedQuestion[0]
	}

	return nil
}

func (self *questionRepository) initialize() (questions []entity.Question) {

	questions = []entity.Question{
		{Id: 1, Content: "question1", Answers: []entity.Answer{
			{Id: 1, QuestionId: 1, Content: "answer1"},
			{Id: 2, QuestionId: 1, Content: "answer2"},
			{Id: 3, QuestionId: 1, Content: "answer3"},
		}},
		{Id: 2, Content: "question2", Answers: []entity.Answer{
			{Id: 4, QuestionId: 2, Content: "answer1"},
			{Id: 5, QuestionId: 2, Content: "answer2"},
			{Id: 6, QuestionId: 2, Content: "answer3"},
		}},
	}

	self.database.Set(questionKey, questions)

	return questions
}

func (self *questionRepository) initializeSeparateQuestions(questionId int) (questions []entity.SeparateQuestion) {

	questions = []entity.SeparateQuestion{
		{Id: 1, Content: "question1"},
		{Id: 2, Content: "question2"},
	}

	var searchQuestion *entity.SeparateQuestion

	for _, question := range questions {
		if question.Id == questionId {
			searchQuestion = &question
		}
		self.database.Set(fmt.Sprintf(separateQuestionsPrefix, question.Id), question)
	}

	if questionId == 0 {
		self.database.Set(separateQuestions, questions)
		return questions
	}

	if searchQuestion == nil {
		return nil
	}

	return []entity.SeparateQuestion{*searchQuestion}
}

func NewQuestionRepository(database cache.Cache) QuestionRepository {
	return &questionRepository{
		database: database,
	}
}
