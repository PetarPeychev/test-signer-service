package api

type User struct {
	ID int `json:"id"`
}

type QuestionAnswer struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
