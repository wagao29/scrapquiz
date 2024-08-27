package answer

type PostAnswerRequest struct {
	UserID    string `json:"userId" validate:"required"`
	AnswerNum int    `json:"answerNum" validate:"required"`
}
