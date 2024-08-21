package quiz

import (
	"fmt"
	"scrapquiz/utils/ulid"
	"unicode/utf8"

	utilsError "scrapquiz/utils/error"
)

const (
	maxOptionLength      = 200
	maxContentLength     = 1000
	maxExplanationLength = 500
)

type Options struct {
	option1 string
	option2 string
	option3 string
	option4 string
}

type Quiz struct {
	id          string
	userID      string
	content     string
	options     Options
	correctNum  int
	explanation string
}

func Reconstruct(
	id string,
	userID string,
	content string,
	options []string,
	correctNum int,
	explanation string,
) (*Quiz, error) {
	return newQuiz(
		id, userID, content, options, correctNum, explanation,
	)
}

func NewQuiz(
	userID string,
	content string,
	options []string,
	correctNum int,
	explanation string,
) (*Quiz, error) {
	return newQuiz(
		ulid.NewULID(), userID, content, options, correctNum, explanation,
	)
}

func (u *Quiz) ID() string {
	return u.id
}

func (u *Quiz) UserID() string {
	return u.userID
}

func (u *Quiz) Content() string {
	return u.content
}

func (u *Quiz) Options() Options {
	return u.options
}

func (u *Quiz) CorrectNum() int {
	return u.correctNum
}

func (u *Quiz) Explanation() string {
	return u.explanation
}

func newQuiz(
	id string,
	userID string,
	content string,
	options []string,
	correctNum int,
	explanation string,
) (*Quiz, error) {

	if !ulid.IsValid(id) {
		return nil, utilsError.NewBadRequestError("quiz id is invalid")
	}

	if utf8.RuneCountInString(content) > maxContentLength {
		return nil, utilsError.NewBadRequestError("content is too long")
	}

	if utf8.RuneCountInString(explanation) > maxExplanationLength {
		return nil, utilsError.NewBadRequestError("explanation is too long")
	}

	if correctNum > len(options)-1 {
		return nil, utilsError.NewBadRequestError("correctNum is too large for the number of options")
	}

	var opts Options
	for i, option := range options {
		if utf8.RuneCountInString(option) > maxOptionLength {
			return nil, utilsError.NewBadRequestError(fmt.Sprintf("option %d is too long", i+1))
		}
		switch i + 1 {
		case 1:
			opts.option1 = option
		case 2:
			opts.option2 = option
		case 3:
			opts.option3 = option
		case 4:
			opts.option4 = option
		}
	}

	return &Quiz{
		id:          id,
		userID:      userID,
		content:     content,
		options:     opts,
		correctNum:  correctNum,
		explanation: explanation,
	}, nil
}
