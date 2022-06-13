package modal

type SeparateQuestionResponse struct {
	Id      int
	Content string
	Answers []SepareteQuestionResponseItem
}

type SepareteQuestionResponseItem struct {
	Id         int
	QuestionId int
	Content    string
}

type FinishTestResponse struct {
	CorrectCount int
	WrongCount   int
}
