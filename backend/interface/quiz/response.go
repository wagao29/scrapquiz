package quiz

type FetchQuizzesResponse = struct {
	Quizzes []FetchQuizResponse `json:"quizzes"`
}

type FetchQuizCountsResponse struct {
	QuizCounts int `json:"quizCounts"`
}

type PostQuizResponse struct {
	ID          string   `json:"id"`
	UserID      string   `json:"userId"`
	Content     string   `json:"content"`
	Options     []string `json:"options"`
	CorrectNum  int      `json:"correctNum"`
	Explanation string   `json:"explanation"`
}

type FetchQuizResponse struct {
	ID            string   `json:"id"`
	Content       string   `json:"content"`
	Options       []string `json:"options"`
	CorrectNum    int      `json:"correctNum"`
	Explanation   string   `json:"explanation"`
	UserID        string   `json:"userId"`
	UserName      string   `json:"userName"`
	UserAvatarURL string   `json:"userAvatarUrl"`
}
