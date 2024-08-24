package quiz

type PostQuizRequest struct {
	UserID      string   `json:"user_id" validate:"required"`
	Content     string   `json:"content" validate:"required"`
	Options     []string `json:"options" validate:"required"`
	CorrectNum  int      `json:"correct_num" validate:"required"`
	Explanation string   `json:"explanation" validate:"required"`
}
