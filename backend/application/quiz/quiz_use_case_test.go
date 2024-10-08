package quiz

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/mock/gomock"

	quizDomain "scrapquiz/domain/quiz"
	utilsError "scrapquiz/utils/error"
)

var mockTime = time.Now()

func TestQuizUseCase_FetchQuizID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockQuizRepo := quizDomain.NewMockQuizRepository(ctrl)
	mockQuizQS := NewMockQuizQueryService(ctrl)
	uc := NewQuizUseCase(mockQuizRepo, mockQuizQS)

	tests := []struct {
		name     string
		inputID  string
		mockFunc func()
		want     *QuizQueryServiceDto
		wantErr  bool
	}{
		{
			name:    "正常系",
			inputID: "01FVSHW3SER8977QCJBYZD9HAW",
			mockFunc: func() {
				mockQuizQS.
					EXPECT().
					FetchQuizByID(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, id string) (*QuizQueryServiceDto, error) {
						return &QuizQueryServiceDto{
							ID:            "01FVSHW3SER8977QCJBYZD9HAW",
							Content:       "問題本文がここに入ります",
							Options:       []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
							CorrectNum:    2,
							Explanation:   "解説文がここに入ります",
							UserID:        "01FVSHW3SER8977QCJBYZD9HAU",
							UserName:      "太郎",
							UserAvatarURL: "https://example.com/avatar.png",
							CreatedAt:     mockTime,
						}, nil
					})
			},
			want: &QuizQueryServiceDto{
				ID:            "01FVSHW3SER8977QCJBYZD9HAW",
				Content:       "問題本文がここに入ります",
				Options:       []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
				CorrectNum:    2,
				Explanation:   "解説文がここに入ります",
				UserID:        "01FVSHW3SER8977QCJBYZD9HAU",
				UserName:      "太郎",
				UserAvatarURL: "https://example.com/avatar.png",
				CreatedAt:     mockTime,
			},
			wantErr: false,
		},
		{
			name: "異常系: 指定したユーザーIDに紐づくクイズが存在しない場合",
			mockFunc: func() {
				mockQuizQS.
					EXPECT().
					FetchQuizByID(gomock.Any(), gomock.Any()).
					Return(nil, utilsError.NewNotFoundError("quiz not found"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.mockFunc()
			got, err := uc.FetchQuizByID(context.Background(), tt.inputID)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("diff = %v", diff)
			}
		})
	}
}

func TestQuizUseCase_FetchQuizCounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockQuizRepo := quizDomain.NewMockQuizRepository(ctrl)
	mockQuizQS := NewMockQuizQueryService(ctrl)
	uc := NewQuizUseCase(mockQuizRepo, mockQuizQS)

	tests := []struct {
		name     string
		mockFunc func()
		want     int
		wantErr  bool
	}{
		{
			name: "正常系",
			mockFunc: func() {
				mockQuizQS.
					EXPECT().
					FetchQuizCounts(gomock.Any()).
					Return(123, nil)
			},
			want:    123,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.mockFunc()
			got, err := uc.FetchQuizCounts(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("diff = %v", diff)
			}
		})
	}
}

func TestQuizUseCase_FetchQuizCountsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockQuizRepo := quizDomain.NewMockQuizRepository(ctrl)
	mockQuizQS := NewMockQuizQueryService(ctrl)
	uc := NewQuizUseCase(mockQuizRepo, mockQuizQS)

	tests := []struct {
		name        string
		inputUserID string
		mockFunc    func()
		want        int
		wantErr     bool
	}{
		{
			name:        "正常系",
			inputUserID: "01FVSHW3SER8977QCJBYZD9HAU",
			mockFunc: func() {
				mockQuizQS.
					EXPECT().
					FetchQuizCountsByUserID(gomock.Any(), gomock.Any()).
					Return(12, nil)
			},
			want:    12,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.mockFunc()
			got, err := uc.FetchQuizCountsByUserID(context.Background(), tt.inputUserID)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("diff = %v", diff)
			}
		})
	}
}

func TestQuizUseCase_FetchQuizzesByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockQuizRepo := quizDomain.NewMockQuizRepository(ctrl)
	mockQuizQS := NewMockQuizQueryService(ctrl)
	uc := NewQuizUseCase(mockQuizRepo, mockQuizQS)

	tests := []struct {
		name        string
		inputUserID string
		inputLimit  int
		inputOffset int
		mockFunc    func()
		want        []*QuizQueryServiceDto
		wantErr     bool
	}{
		{
			name:        "正常系",
			inputUserID: "01FVSHW3SER8977QCJBYZD9HAU",
			inputLimit:  10,
			inputOffset: 0,
			mockFunc: func() {
				mockQuizQS.
					EXPECT().
					FetchQuizzesByUserID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, useID string, limit int, offset int) ([]*QuizQueryServiceDto, error) {
						var quizzes []*QuizQueryServiceDto
						for i := range 3 {
							quizzes = append(quizzes, &QuizQueryServiceDto{
								ID:            fmt.Sprintf("%d1FVSHW3SER8977QCJBYZD9HAW", i),
								Content:       "問題本文がここに入ります",
								Options:       []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
								CorrectNum:    2,
								Explanation:   "解説文がここに入ります",
								UserID:        "01FVSHW3SER8977QCJBYZD9HAU",
								UserName:      "太郎",
								UserAvatarURL: "https://example.com/avatar.png",
								CreatedAt:     mockTime,
							})
						}
						return quizzes, nil
					})
			},
			want: []*QuizQueryServiceDto{
				{
					ID:            "01FVSHW3SER8977QCJBYZD9HAW",
					Content:       "問題本文がここに入ります",
					Options:       []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
					CorrectNum:    2,
					Explanation:   "解説文がここに入ります",
					UserID:        "01FVSHW3SER8977QCJBYZD9HAU",
					UserName:      "太郎",
					UserAvatarURL: "https://example.com/avatar.png",
					CreatedAt:     mockTime,
				},
				{
					ID:            "11FVSHW3SER8977QCJBYZD9HAW",
					Content:       "問題本文がここに入ります",
					Options:       []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
					CorrectNum:    2,
					Explanation:   "解説文がここに入ります",
					UserID:        "01FVSHW3SER8977QCJBYZD9HAU",
					UserName:      "太郎",
					UserAvatarURL: "https://example.com/avatar.png",
					CreatedAt:     mockTime,
				},
				{
					ID:            "21FVSHW3SER8977QCJBYZD9HAW",
					Content:       "問題本文がここに入ります",
					Options:       []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
					CorrectNum:    2,
					Explanation:   "解説文がここに入ります",
					UserID:        "01FVSHW3SER8977QCJBYZD9HAU",
					UserName:      "太郎",
					UserAvatarURL: "https://example.com/avatar.png",
					CreatedAt:     mockTime,
				},
			},
			wantErr: false,
		},
		{
			name: "異常系: 指定したユーザーIDに紐づくクイズが存在しない場合",
			mockFunc: func() {
				mockQuizQS.
					EXPECT().
					FetchQuizzesByUserID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, utilsError.NewNotFoundError("quiz not found"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.mockFunc()
			got, err := uc.FetchQuizzesByUserID(context.Background(), tt.inputUserID, tt.inputLimit, tt.inputOffset)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("diff = %v", diff)
			}
		})
	}
}

func TestQuizUseCase_FetchLatestQuizzes(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockQuizRepo := quizDomain.NewMockQuizRepository(ctrl)
	mockQuizQS := NewMockQuizQueryService(ctrl)
	uc := NewQuizUseCase(mockQuizRepo, mockQuizQS)

	tests := []struct {
		name        string
		inputLimit  int
		inputOffset int
		mockFunc    func()
		want        []*QuizQueryServiceDto
		wantErr     bool
	}{
		{
			name:        "正常系",
			inputLimit:  10,
			inputOffset: 0,
			mockFunc: func() {
				mockQuizQS.
					EXPECT().
					FetchLatestQuizzes(gomock.Any(), gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, limit int, offset int) ([]*QuizQueryServiceDto, error) {
						var quizzes []*QuizQueryServiceDto
						for i := range 3 {
							quizzes = append(quizzes, &QuizQueryServiceDto{
								ID:            fmt.Sprintf("%d1FVSHW3SER8977QCJBYZD9HAW", i),
								Content:       "問題本文がここに入ります",
								Options:       []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
								CorrectNum:    2,
								Explanation:   "解説文がここに入ります",
								UserID:        "01FVSHW3SER8977QCJBYZD9HAU",
								UserName:      "太郎",
								UserAvatarURL: "https://example.com/avatar.png",
								CreatedAt:     mockTime,
							})
						}
						return quizzes, nil
					})
			},
			want: []*QuizQueryServiceDto{
				{
					ID:            "01FVSHW3SER8977QCJBYZD9HAW",
					Content:       "問題本文がここに入ります",
					Options:       []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
					CorrectNum:    2,
					Explanation:   "解説文がここに入ります",
					UserID:        "01FVSHW3SER8977QCJBYZD9HAU",
					UserName:      "太郎",
					UserAvatarURL: "https://example.com/avatar.png",
					CreatedAt:     mockTime,
				},
				{
					ID:            "11FVSHW3SER8977QCJBYZD9HAW",
					Content:       "問題本文がここに入ります",
					Options:       []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
					CorrectNum:    2,
					Explanation:   "解説文がここに入ります",
					UserID:        "01FVSHW3SER8977QCJBYZD9HAU",
					UserName:      "太郎",
					UserAvatarURL: "https://example.com/avatar.png",
					CreatedAt:     mockTime,
				},
				{
					ID:            "21FVSHW3SER8977QCJBYZD9HAW",
					Content:       "問題本文がここに入ります",
					Options:       []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
					CorrectNum:    2,
					Explanation:   "解説文がここに入ります",
					UserID:        "01FVSHW3SER8977QCJBYZD9HAU",
					UserName:      "太郎",
					UserAvatarURL: "https://example.com/avatar.png",
					CreatedAt:     mockTime,
				},
			},
			wantErr: false,
		},
		{
			name: "異常系: クイズが存在しない場合",
			mockFunc: func() {
				mockQuizQS.
					EXPECT().
					FetchLatestQuizzes(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, utilsError.NewNotFoundError("quiz not found"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.mockFunc()
			got, err := uc.FetchLatestQuizzes(context.Background(), tt.inputLimit, tt.inputOffset)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("diff = %v", diff)
			}
		})
	}
}

func TestQuizUseCase_FetchRandomQuizzes(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockQuizRepo := quizDomain.NewMockQuizRepository(ctrl)
	mockQuizQS := NewMockQuizQueryService(ctrl)
	uc := NewQuizUseCase(mockQuizRepo, mockQuizQS)

	tests := []struct {
		name       string
		inputLimit int
		mockFunc   func()
		want       []*QuizQueryServiceDto
		wantErr    bool
	}{
		{
			name:       "正常系",
			inputLimit: 10,
			mockFunc: func() {
				mockQuizQS.
					EXPECT().
					FetchRandomQuizzes(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, limit int) ([]*QuizQueryServiceDto, error) {
						var quizzes []*QuizQueryServiceDto
						for i := range 3 {
							quizzes = append(quizzes, &QuizQueryServiceDto{
								ID:            fmt.Sprintf("%d1FVSHW3SER8977QCJBYZD9HAW", i),
								Content:       "問題本文がここに入ります",
								Options:       []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
								CorrectNum:    2,
								Explanation:   "解説文がここに入ります",
								UserID:        "01FVSHW3SER8977QCJBYZD9HAU",
								UserName:      "太郎",
								UserAvatarURL: "https://example.com/avatar.png",
								CreatedAt:     mockTime,
							})
						}
						return quizzes, nil
					})
			},
			want: []*QuizQueryServiceDto{
				{
					ID:            "01FVSHW3SER8977QCJBYZD9HAW",
					Content:       "問題本文がここに入ります",
					Options:       []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
					CorrectNum:    2,
					Explanation:   "解説文がここに入ります",
					UserID:        "01FVSHW3SER8977QCJBYZD9HAU",
					UserName:      "太郎",
					UserAvatarURL: "https://example.com/avatar.png",
					CreatedAt:     mockTime,
				},
				{
					ID:            "11FVSHW3SER8977QCJBYZD9HAW",
					Content:       "問題本文がここに入ります",
					Options:       []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
					CorrectNum:    2,
					Explanation:   "解説文がここに入ります",
					UserID:        "01FVSHW3SER8977QCJBYZD9HAU",
					UserName:      "太郎",
					UserAvatarURL: "https://example.com/avatar.png",
					CreatedAt:     mockTime,
				},
				{
					ID:            "21FVSHW3SER8977QCJBYZD9HAW",
					Content:       "問題本文がここに入ります",
					Options:       []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
					CorrectNum:    2,
					Explanation:   "解説文がここに入ります",
					UserID:        "01FVSHW3SER8977QCJBYZD9HAU",
					UserName:      "太郎",
					UserAvatarURL: "https://example.com/avatar.png",
					CreatedAt:     mockTime,
				},
			},
			wantErr: false,
		},
		{
			name: "異常系: クイズが存在しない場合",
			mockFunc: func() {
				mockQuizQS.
					EXPECT().
					FetchRandomQuizzes(gomock.Any(), gomock.Any()).
					Return(nil, utilsError.NewNotFoundError("quiz not found"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.mockFunc()
			got, err := uc.FetchRandomQuizzes(context.Background(), tt.inputLimit)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("diff = %v", diff)
			}
		})
	}
}

func TestQuizUseCase_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockQuizRepo := quizDomain.NewMockQuizRepository(ctrl)
	mockQuizQS := NewMockQuizQueryService(ctrl)
	uc := NewQuizUseCase(mockQuizRepo, mockQuizQS)

	tests := []struct {
		name     string
		input    QuizUseCaseSaveInputDto
		mockFunc func()
		want     *QuizUseCaseSaveOutputDto
		wantErr  bool
	}{
		{
			name: "正常系",
			input: QuizUseCaseSaveInputDto{
				UserID:      "01FVSHW3SER8977QCJBYZD9HAU",
				Content:     "問題本文がここに入ります",
				Options:     []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
				CorrectNum:  1,
				Explanation: "解説文がここに入ります",
			},
			mockFunc: func() {
				mockQuizRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
			},
			want: &QuizUseCaseSaveOutputDto{
				ID:          "01FVSHW3SER8977QCJBYZD9HAW",
				UserID:      "01FVSHW3SER8977QCJBYZD9HAU",
				Content:     "問題本文がここに入ります",
				Options:     []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
				CorrectNum:  1,
				Explanation: "解説文がここに入ります",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.mockFunc()
			got, err := uc.Save(context.Background(), tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(QuizUseCaseSaveOutputDto{}, "ID")); diff != "" {
				t.Errorf("diff = %v", diff)
			}
		})
	}
}

func TestQuizUseCase_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockQuizRepo := quizDomain.NewMockQuizRepository(ctrl)
	mockQuizQS := NewMockQuizQueryService(ctrl)
	uc := NewQuizUseCase(mockQuizRepo, mockQuizQS)

	tests := []struct {
		name     string
		input    string
		mockFunc func()
		wantErr  bool
	}{
		{
			name:  "正常系",
			input: "01FVSHW3SER8977QCJBYZD9HAW",
			mockFunc: func() {
				mockQuizRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.mockFunc()
			err := uc.Delete(context.Background(), tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
