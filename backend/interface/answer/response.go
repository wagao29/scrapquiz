package answer

type PostAnswerResponse struct {
	QuizID    string `json:"quizId"`
	UserID    string `json:"userId"`
	AnswerNum int    `json:"answerNum"`
}

type FetchAnswerCountsResponse struct {
	AnswerCounts []int `json:"answerCounts"`
}
