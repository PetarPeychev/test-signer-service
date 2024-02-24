package api

import "time"

type User struct {
	ID int `json:"id"`
}

type QuestionAnswer struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Signature struct {
	ID               int              `json:"id"`
	UserID           int              `json:"userId"`
	Timestamp        time.Time        `json:"timestamp"`
	QuestionsAnswers []QuestionAnswer `json:"questionsAnswers"`
}
