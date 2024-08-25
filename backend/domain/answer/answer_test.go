package answer

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewAnswer(t *testing.T) {
	type args struct {
		quizID    string
		userID    string
		answerNum int
	}
	tests := []struct {
		name    string
		args    args
		want    *Answer
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				quizID:    "01FVSHW3SER8977QCJBYZD9HAQ",
				userID:    "01FVSHW3SER8977QCJBYZD9HAU",
				answerNum: 1,
			},
			want: &Answer{
				quizID:    "01FVSHW3SER8977QCJBYZD9HAQ",
				userID:    "01FVSHW3SER8977QCJBYZD9HAU",
				answerNum: 1,
			},
			wantErr: false,
		},
		{
			name: "異常系: quizID のフォーマットが不正",
			args: args{
				quizID:    "jf93&f328*2@()df",
				userID:    "01FVSHW3SER8977QCJBYZD9HAU",
				answerNum: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: answerNum が 1 より小さい",
			args: args{
				quizID:    "01FVSHW3SER8977QCJBYZD9HAQ",
				userID:    "01FVSHW3SER8977QCJBYZD9HAU",
				answerNum: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: answerNum が 4 より大きい",
			args: args{
				quizID:    "01FVSHW3SER8977QCJBYZD9HAQ",
				userID:    "01FVSHW3SER8977QCJBYZD9HAU",
				answerNum: 5,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAnswer(tt.args.quizID, tt.args.userID, tt.args.answerNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(Answer{}))
			if diff != "" {
				t.Errorf("NewAnswer() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}
