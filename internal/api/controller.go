package api

import (
	"fmt"
	"ft-interview/internal/api/modal"
	"ft-interview/internal/domain"
	"ft-interview/internal/domain/entity"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//GetQuestionsWithAnswers simple in memory cache question with answers
func GetQuestionsWithAnswers(e *echo.Echo, questionRepository domain.QuestionRepository) {

	e.GET("questions", func(c echo.Context) error {
		var (
			questions []entity.Question
		)

		questions = questionRepository.GetAll()

		if len(questions) == 0 {
			return c.NoContent(http.StatusNoContent)
		}

		return c.JSON(http.StatusOK, questions)
	})
}

//GetQuestionsWithSeparateAnswers get all question with answers by question id like database struct
func GetQuestionsWithSeparateAnswers(e *echo.Echo, questionRepository domain.QuestionRepository, answerRepository domain.AnswerRepository) {
	e.GET("v2/questions", func(c echo.Context) error {

		var (
			questions []entity.SeparateQuestion
			response  []modal.SeparateQuestionResponse
		)

		questions = questionRepository.GetAllSeparateQuestions()

		if len(questions) == 0 {
			return c.NoContent(http.StatusNoContent)
		}

		for _, question := range questions {
			var separateQuestion modal.SeparateQuestionResponse

			separateQuestion.Id = question.Id
			separateQuestion.Content = question.Content

			var responseItem modal.SepareteQuestionResponseItem

			answers := answerRepository.GetByQuestionId(question.Id)

			for _, answer := range answers {
				responseItem.QuestionId = question.Id
				responseItem.Id = answer.Id
				responseItem.Content = answer.Content

				separateQuestion.Answers = append(separateQuestion.Answers, responseItem)
			}

			response = append(response, separateQuestion)
		}

		return c.JSON(http.StatusOK, response)
	})
}

//GetQuestionsWithSeparateAnswersByQuestionId get one question with answer by questionId
func GetQuestionsWithSeparateAnswersByQuestionId(e *echo.Echo, questionRepository domain.QuestionRepository, answerRepository domain.AnswerRepository) {
	e.GET("questions/:questionId", func(c echo.Context) error {

		var (
			question *entity.SeparateQuestion
		)

		questionId, err := strconv.Atoi(c.Param("questionId"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid questionId")
		}

		question = questionRepository.GetSeparateQuestionById(questionId)

		if question == nil {
			return c.NoContent(http.StatusNoContent)
		}

		var separateQuestion modal.SeparateQuestionResponse

		separateQuestion.Id = question.Id
		separateQuestion.Content = question.Content

		var responseItem modal.SepareteQuestionResponseItem

		answers := answerRepository.GetByQuestionId(question.Id)

		for _, answer := range answers {
			responseItem.QuestionId = question.Id
			responseItem.Id = answer.Id
			responseItem.Content = answer.Content

			separateQuestion.Answers = append(separateQuestion.Answers, responseItem)
		}

		return c.JSON(http.StatusOK, separateQuestion)
	})
}

func Finish(e *echo.Echo, questionAnswerRepository domain.QuestionAnswerRepository) {
	e.POST("questions/complete", func(c echo.Context) error {
		var (
			request = new(modal.FinishTestRequest)
			err     error
		)

		if err = c.Bind(request); err != nil {
			return c.JSON(http.StatusBadRequest, "Json deserialize error")
		}

		if IsDuplicateAnswers(request.Data) {
			return c.JSON(http.StatusBadRequest, "You can select just one answer per question!")
		}

		wrongAnswerCount, correctAnswerCount := 0, 0

		for _, data := range request.Data {

			answer := questionAnswerRepository.GetByQuestionId(data.QuestionId)

			if answer.CorrectAnswerId != data.AnswerId {
				wrongAnswerCount += 1
				continue
			}

			correctAnswerCount += 1
		}

		questionAnswerRepository.Upsert(request.UserId, correctAnswerCount)

		return c.JSON(http.StatusOK, modal.FinishTestResponse{
			CorrectCount: correctAnswerCount,
			WrongCount:   wrongAnswerCount,
		})
	})
}

func CompareToOtherUsers(e *echo.Echo, questionAnswerRepository domain.QuestionAnswerRepository) {

	e.GET("questions/compare/:userId", func(c echo.Context) error {

		userId, err := strconv.Atoi(c.Param("userId"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid user id")
		}

		userAnswers := questionAnswerRepository.GetUserAnswers()

		if _, ok := userAnswers[userId]; !ok {
			return c.JSON(http.StatusBadRequest, "soru çözmemişsin xd")
		}

		userIndex := FindUserOrderByCorrectAnswerCount(userId, userAnswers)

		percentage := (userIndex / float64(len(userAnswers))) * 100

		return c.String(http.StatusOK, fmt.Sprintf("You were better than  %.0f%% of all quizzers", percentage))
	})
}
