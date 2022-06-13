package entity

type Question struct {
	Id      int
	Content string
	Answers []Answer
}

type SeparateQuestion struct {
	Id      int
	Content string
}

