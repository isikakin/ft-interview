package modal

type FinishTestRequest struct {
	UserId int
	Data   []FinishTestRequestItem
}

type FinishTestRequestItem struct {
	QuestionId int
	AnswerId   int
}
