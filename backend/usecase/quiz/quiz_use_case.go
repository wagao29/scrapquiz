package quiz

import (
	"context"

	quizDomain "scrapquiz/domain/quiz"
)

type QuizUseCaseSaveInputDto struct {
	UserID      string
	Content     string
	Options     []string
	CorrectNum  int
	Explanation string
}

type QuizUseCaseSaveOutputDto struct {
	ID          string
	UserID      string
	Content     string
	Options     []string
	CorrectNum  int
	Explanation string
}

type QuizUseCase struct {
	quizRepo         quizDomain.QuizRepository
	quizQueryService QuizQueryService
}

func NewQuizUseCase(
	quizRepo quizDomain.QuizRepository,
	quizQueryService QuizQueryService,
) *QuizUseCase {
	return &QuizUseCase{
		quizRepo:         quizRepo,
		quizQueryService: quizQueryService,
	}
}

func (uc *QuizUseCase) FetchQuizByID(ctx context.Context, id string) (*QuizQueryServiceDto, error) {
	quiz, err := uc.quizQueryService.FetchQuizByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return quiz, nil
}

func (uc *QuizUseCase) FetchQuizzesByUserID(
	ctx context.Context,
	userID string,
	limit int,
	offset int,
) ([]*QuizQueryServiceDto, error) {
	quizzes, err := uc.quizQueryService.FetchQuizzesByUserID(ctx, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	return quizzes, nil
}

func (uc *QuizUseCase) FetchLatestQuizzes(
	ctx context.Context,
	limit int,
	offset int,
) ([]*QuizQueryServiceDto, error) {
	quizzes, err := uc.quizQueryService.FetchLatestQuizzes(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return quizzes, nil
}

func (uc *QuizUseCase) Save(ctx context.Context, dto QuizUseCaseSaveInputDto) (*QuizUseCaseSaveOutputDto, error) {
	quiz, err := quizDomain.NewQuiz(dto.UserID, dto.Content, dto.Options, dto.CorrectNum, dto.Explanation)
	if err != nil {
		return nil, err
	}
	if err = uc.quizRepo.Save(ctx, quiz); err != nil {
		return nil, err
	}
	return &QuizUseCaseSaveOutputDto{
		ID:          quiz.ID(),
		UserID:      quiz.UserID(),
		Content:     quiz.Content(),
		Options:     []string{quiz.Option1(), quiz.Option2(), quiz.Option3(), quiz.Option4()},
		CorrectNum:  quiz.CorrectNum(),
		Explanation: quiz.Explanation(),
	}, nil
}

func (uc *QuizUseCase) Delete(ctx context.Context, id string) error {
	if err := uc.quizRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
