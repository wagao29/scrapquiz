package quiz

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

const testUserID = "01FVSHW3SER8977QCJBYZD9HAW"

func TestNewQuiz(t *testing.T) {
	type args struct {
		userID      string
		content     string
		options     []string
		correctNum  int
		explanation string
	}
	tests := []struct {
		name    string
		args    args
		want    *Quiz
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				userID:      testUserID,
				content:     "問題本文がここに入ります",
				options:     []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
				correctNum:  1,
				explanation: "解説文がここに入ります",
			},
			want: &Quiz{
				userID:  testUserID,
				content: "問題本文がここに入ります",
				options: Options{
					option1: "選択肢1",
					option2: "選択肢2",
					option3: "選択肢3",
					option4: "選択肢4",
				},
				correctNum:  1,
				explanation: "解説文がここに入ります",
			},
			wantErr: false,
		},
		{
			name: "異常系: content の文字数が不正",
			args: args{
				userID:      testUserID,
				content:     strings.Repeat("あ", 1200),
				options:     []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
				correctNum:  1,
				explanation: "解説文がここに入ります",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: explanation の文字数が不正",
			args: args{
				userID:      testUserID,
				content:     "問題本文がここに入ります",
				options:     []string{"選択肢1", "選択肢2", "選択肢3", "選択肢4"},
				correctNum:  1,
				explanation: strings.Repeat("あ", 600),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: option の文字数が不正",
			args: args{
				userID:      testUserID,
				content:     "問題本文がここに入ります",
				options:     []string{"選択肢1", strings.Repeat("あ", 230), "選択肢3"},
				correctNum:  1,
				explanation: "解説文がここに入ります",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: options の数より correctNum が大きい",
			args: args{
				userID:      testUserID,
				content:     "問題本文がここに入ります",
				options:     []string{"選択肢1", "選択肢2"},
				correctNum:  3,
				explanation: "解説文がここに入ります",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewQuiz(
				tt.args.userID,
				tt.args.content,
				tt.args.options,
				tt.args.correctNum,
				tt.args.explanation,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(Quiz{}, Options{}), cmpopts.IgnoreFields(Quiz{}, "id"))
			if diff != "" {
				t.Errorf("NewQuiz() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}
