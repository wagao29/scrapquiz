package answer

type PostAnswerRequest struct {
	UserID    string `json:"user_id" validate:"required"`
	AnswerNum int    `json:"answer_num" validate:"required"`
}
