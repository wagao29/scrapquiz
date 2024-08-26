package answer

import (
	"context"
	"testing"

	answerDomain "scrapquiz/domain/answer"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
)

func TestAnswerUseCase_FetchAnswerCountsByQuizID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAnswerRepo := answerDomain.NewMockAnswerRepository(ctrl)
	mockAnswerQS := NewMockAnswerQueryService(ctrl)
	uc := NewAnswerUseCase(mockAnswerRepo, mockAnswerQS)

	tests := []struct {
		name     string
		quizID   string
		mockFunc func()
		want     *AnswerQueryServiceDto
		wantErr  bool
	}{
		{
			name:   "正常系",
			quizID: "01FVSHW3SER8977QCJBYZD9HAQ",
			mockFunc: func() {
				mockAnswerQS.
					EXPECT().
					FetchAnswerCountsByQuizID(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, quizID string) (*AnswerQueryServiceDto, error) {
						return &AnswerQueryServiceDto{
							AnswerNum1Count: 6,
							AnswerNum2Count: 3,
							AnswerNum3Count: 0,
							AnswerNum4Count: 21,
						}, nil
					})
			},
			want: &AnswerQueryServiceDto{
				AnswerNum1Count: 6,
				AnswerNum2Count: 3,
				AnswerNum3Count: 0,
				AnswerNum4Count: 21,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.mockFunc()
			got, err := uc.FetchAnswerCountsByQuizID(context.Background(), tt.quizID)
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

func TestAnswerUseCase_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAnswerRepo := answerDomain.NewMockAnswerRepository(ctrl)
	mockAnswerQS := NewMockAnswerQueryService(ctrl)
	uc := NewAnswerUseCase(mockAnswerRepo, mockAnswerQS)

	tests := []struct {
		name     string
		input    AnswerUseCaseInputDto
		mockFunc func()
		want     *AnswerUseCaseOutputDto
		wantErr  bool
	}{
		{
			name: "正常系",
			input: AnswerUseCaseInputDto{
				QuizID:    "01FVSHW3SER8977QCJBYZD9HAQ",
				UserID:    "01FVSHW3SER8977QCJBYZD9HAU",
				AnswerNum: 1,
			},
			mockFunc: func() {
				mockAnswerRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
			},
			want: &AnswerUseCaseOutputDto{
				QuizID:    "01FVSHW3SER8977QCJBYZD9HAQ",
				UserID:    "01FVSHW3SER8977QCJBYZD9HAU",
				AnswerNum: 1,
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
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("diff = %v", diff)
			}
		})
	}
}
