package model

type Question struct {
	ID      string
	Text    string
	Type    string
	Points  float64
	QuizID  string
	Answers []Answer
}

type Answer struct {
	AnswerID  string
	Text      string
	IsCorrect bool
}
