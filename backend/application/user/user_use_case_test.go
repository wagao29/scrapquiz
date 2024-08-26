package user

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"

	userDomain "scrapquiz/domain/user"
	utilsError "scrapquiz/utils/error"
)

const (
	testID        = "01FVSHW3SER8977QCJBYZD9HAW"
	testName      = "太郎"
	testAvatarURL = "https://example.com/avatar.png"
)

func TestUserUseCase_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := userDomain.NewMockUserRepository(ctrl)
	uc := NewUserUseCase(mockUserRepo)

	tests := []struct {
		name     string
		input    string
		mockFunc func()
		want     *UserUseCaseOutputDto
		wantErr  bool
	}{
		{
			name:  "正常系",
			input: "01FVSHW3SER8977QCJBYZD9HAW",
			mockFunc: func() {
				mockUserRepo.
					EXPECT().
					FindByID(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, id string) (*userDomain.User, error) {
						return userDomain.NewUser(
							testID,
							testName,
							testAvatarURL,
						)
					})
			},
			want: &UserUseCaseOutputDto{
				ID:        testID,
				Name:      testName,
				AvatarURL: testAvatarURL,
			},
			wantErr: false,
		},
		{
			name:  "異常系: idを持つユーザーが存在しない場合",
			input: "01FVSHW3SER8977QCJBYZD9HAX",
			mockFunc: func() {
				mockUserRepo.
					EXPECT().
					FindByID(gomock.Any(), gomock.Any()).
					Return(nil, utilsError.NewNotFoundError("user not found"))
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
			got, err := uc.FindByID(context.Background(), tt.input)
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

func TestUserUseCase_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := userDomain.NewMockUserRepository(ctrl)
	uc := NewUserUseCase(mockUserRepo)

	tests := []struct {
		name     string
		mockFunc func()
		want     []*UserUseCaseOutputDto
		wantErr  bool
	}{
		{
			name: "正常系",
			mockFunc: func() {
				mockUserRepo.
					EXPECT().
					FindAll(gomock.Any()).
					DoAndReturn(func(ctx context.Context) ([]*userDomain.User, error) {
						var users []*userDomain.User
						for i := range 3 {
							user, _ := userDomain.NewUser(
								fmt.Sprintf("%d1FVSHW3SER8977QCJBYZD9HAW", i),
								testName,
								testAvatarURL,
							)
							users = append(users, user)
						}
						return users, nil
					})
			},
			want: []*UserUseCaseOutputDto{
				{
					ID:        "01FVSHW3SER8977QCJBYZD9HAW",
					Name:      testName,
					AvatarURL: testAvatarURL,
				},
				{
					ID:        "11FVSHW3SER8977QCJBYZD9HAW",
					Name:      testName,
					AvatarURL: testAvatarURL,
				},
				{
					ID:        "21FVSHW3SER8977QCJBYZD9HAW",
					Name:      testName,
					AvatarURL: testAvatarURL,
				},
			},
			wantErr: false,
		},
		{
			name: "異常系: ユーザーが存在しない場合",
			mockFunc: func() {
				mockUserRepo.
					EXPECT().
					FindAll(gomock.Any()).
					Return(nil, utilsError.NewNotFoundError("user not found"))
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
			got, err := uc.FindAll(context.Background())
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

func TestUserUseCase_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := userDomain.NewMockUserRepository(ctrl)
	uc := NewUserUseCase(mockUserRepo)

	tests := []struct {
		name     string
		input    UserUseCaseInputDto
		mockFunc func()
		want     *UserUseCaseOutputDto
		wantErr  bool
	}{
		{
			name: "正常系",
			input: UserUseCaseInputDto{
				ID:        testID,
				Name:      testName,
				AvatarURL: testAvatarURL,
			},
			mockFunc: func() {
				mockUserRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
			},
			want: &UserUseCaseOutputDto{
				ID:        testID,
				Name:      testName,
				AvatarURL: testAvatarURL,
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

func TestUserUseCase_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := userDomain.NewMockUserRepository(ctrl)
	uc := NewUserUseCase(mockUserRepo)

	tests := []struct {
		name     string
		input    UserUseCaseInputDto
		mockFunc func()
		wantErr  bool
	}{
		{
			name: "正常系",
			input: UserUseCaseInputDto{
				ID:        testID,
				Name:      testName,
				AvatarURL: testAvatarURL,
			},
			mockFunc: func() {
				mockUserRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.mockFunc()
			err := uc.Update(context.Background(), tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserUseCase_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := userDomain.NewMockUserRepository(ctrl)
	uc := NewUserUseCase(mockUserRepo)

	tests := []struct {
		name     string
		input    string
		mockFunc func()
		wantErr  bool
	}{
		{
			name:  "正常系",
			input: testID,
			mockFunc: func() {
				mockUserRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
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
