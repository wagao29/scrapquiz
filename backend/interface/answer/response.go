package answer

type PostAnswerResponse struct {
	QuizID    string `json:"quizId"`
	UserID    string `json:"userId"`
	AnswerNum int    `json:"answerNum"`
}

type FetchAnswerCountsResponse struct {
	AnswerNum1Count int `json:"answerNum1Count"`
	AnswerNum2Count int `json:"answerNum2Count"`
	AnswerNum3Count int `json:"answerNum3Count"`
	AnswerNum4Count int `json:"answerNum4Count"`
}
