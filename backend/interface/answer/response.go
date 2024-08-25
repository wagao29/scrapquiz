package answer

type PostAnswerResponse struct {
	QuizID    string `json:"quiz_id"`
	UserID    string `json:"user_id"`
	AnswerNum int    `json:"answer_num"`
}

type FetchAnswerCountsResponse struct {
	AnswerNum1Count int `json:"answer_num_1_count"`
	AnswerNum2Count int `json:"answer_num_2_count"`
	AnswerNum3Count int `json:"answer_num_3_count"`
	AnswerNum4Count int `json:"answer_num_4_count"`
}
