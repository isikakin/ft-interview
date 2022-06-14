package api

import (
	"fmt"
	"ft-interview/internal/api"
	"ft-interview/internal/domain"
	"ft-interview/pkg/cache"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

func Init(cmd *cobra.Command, args []string) error {

	e := echo.New()
	e.Debug = false
	e.HideBanner = true
	e.HidePort = true

	var database = cache.New(5*time.Minute, 10*time.Minute)

	var questionRepository = domain.NewQuestionRepository(database)
	var answerRepository = domain.NewAnswerRepository(database)
	var questionAnswerRepository = domain.NewQuestionAnswerRepository(database)

	e.GET("/v1/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "Service Up")
	})

	api.GetQuestionsWithAnswers(e, questionRepository)
	api.GetQuestionsWithSeparateAnswers(e, questionRepository, answerRepository)
	api.GetQuestionsWithSeparateAnswersByQuestionId(e, questionRepository, answerRepository)
	api.Finish(e, questionRepository, questionAnswerRepository)
	api.CompareToOtherUsers(e, questionAnswerRepository)

	if err := e.Start(fmt.Sprintf(":%s", "5000")); err != nil {
		panic(err)
	}

	return nil
}
