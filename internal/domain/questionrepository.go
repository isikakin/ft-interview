package domain

import (
	"ft-interview/internal/domain/entity"
	"ft-interview/pkg/cache"
)

const (
	questionKey       = "questions"
	separateQuestions = "separatequestions"
)

type QuestionRepository interface {
	GetAll() []entity.Question
	GetAllSeparateQuestions() []entity.SeparateQuestion
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

	return self.initializeSeparateQuestions()
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

func (self *questionRepository) initializeSeparateQuestions() (questions []entity.SeparateQuestion) {

	questions = []entity.SeparateQuestion{
		{Id: 1, Content: "question1"},
		{Id: 2, Content: "question2"},
	}

	self.database.Set(separateQuestions, questions)

	return questions
}

func NewQuestionRepository(database cache.Cache) QuestionRepository {
	return &questionRepository{
		database: database,
	}
}
