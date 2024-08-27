package quiz

type PostQuizRequest struct {
	UserID      string   `json:"userId" validate:"required"`
	Content     string   `json:"content" validate:"required"`
	Options     []string `json:"options" validate:"required"`
	CorrectNum  int      `json:"correctNum" validate:"required"`
	Explanation string   `json:"explanation" validate:"required"`
}
