package user

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	testUserID        = "123456789012345678901"
	testUserName      = "太郎"
	testUserAvatarURL = "https://example.com/avatar.png"
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
				id:        testUserID,
				name:      testUserName,
				avatarURL: testUserAvatarURL,
			},
			want: &User{
				id:        testUserID,
				name:      testUserName,
				avatarURL: testUserAvatarURL,
			},
			wantErr: false,
		},
		{
			name: "異常系: id の長さが不正",
			args: args{
				id:        "12345678901234567890123",
				name:      testUserName,
				avatarURL: testUserAvatarURL,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: name の文字数が不正",
			args: args{
				id:        testUserID,
				name:      strings.Repeat("あ", 31),
				avatarURL: testUserAvatarURL,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: avatarURL のフォーマットが不正",
			args: args{
				id:        testUserID,
				name:      testUserName,
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
