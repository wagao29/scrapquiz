package answer

import (
	utilsError "scrapquiz/utils/error"
	"scrapquiz/utils/ulid"
)

const (
	minAnswerNum = 1
	maxAnswerNum = 4
)

type Answer struct {
	quizID    string
	userID    string
	answerNum int
}

func NewAnswer(quizID string, userID string, answerNum int) (*Answer, error) {
	if !ulid.IsValid(quizID) {
		return nil, utilsError.NewBadRequestError("quiz id is invalid")
	}

	if answerNum < minAnswerNum || answerNum > maxAnswerNum {
		return nil, utilsError.NewBadRequestError("answerNum is out of range")
	}

	return &Answer{
		quizID:    quizID,
		userID:    userID,
		answerNum: answerNum,
	}, nil
}

func (u *Answer) QuizID() string {
	return u.quizID
}

func (u *Answer) UserID() string {
	return u.userID
}

func (u *Answer) AnswerNum() int {
	return u.answerNum
}
