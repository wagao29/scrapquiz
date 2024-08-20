package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	testID        = "01FVSHW3SER8977QCJBYZD9HAW"
	testName      = "太郎"
	testAvatarURL = "https://example.com/avatar.png"
)

func TestNewUser(t *testing.T) {
	type args struct {
		id        string
		name      string
		avatarURL string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				id:        testID,
				name:      testName,
				avatarURL: testAvatarURL,
			},
			want: &User{
				id:        testID,
				name:      testName,
				avatarURL: testAvatarURL,
			},
			wantErr: false,
		},
		{
			name: "異常系: name の文字数が不正",
			args: args{
				id:        testID,
				name:      "あああああああああああああああああああああ",
				avatarURL: testAvatarURL,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: avatarURL のフォーマットが不正",
			args: args{
				id:        testID,
				name:      testName,
				avatarURL: "example.com/avatar.png",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.id, tt.args.name, tt.args.avatarURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(User{}))
			if diff != "" {
				t.Errorf("NewUser() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}
