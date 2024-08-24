package quiz

type FetchQuizzesResponse = struct {
	Quizzes []FetchQuizResponse `json:"quizzes"`
}

type PostQuizResponse struct {
	ID          string   `json:"id"`
	UserID      string   `json:"user_id"`
	Content     string   `json:"content"`
	Options     []string `json:"options"`
	CorrectNum  int      `json:"correct_num"`
	Explanation string   `json:"explanation"`
}

type FetchQuizResponse struct {
	ID            string   `json:"id"`
	Content       string   `json:"content"`
	Options       []string `json:"options"`
	CorrectNum    int      `json:"correct_num"`
	Explanation   string   `json:"explanation"`
	UserID        string   `json:"user_id"`
	UserName      string   `json:"user_name"`
	UserAvatarURL string   `json:"user_avatar_url"`
}
