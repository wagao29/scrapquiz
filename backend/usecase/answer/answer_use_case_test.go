package answer

import (
	"context"
	"testing"

	answerDomain "scrapquiz/domain/answer"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
)

func TestAnswerUseCase_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAnswerRepo := answerDomain.NewMockAnswerRepository(ctrl)
	uc := NewAnswerUseCase(mockAnswerRepo)

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
