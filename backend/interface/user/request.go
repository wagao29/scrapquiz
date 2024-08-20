package user

type PostUsersRequest struct {
	ID        string `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	AvatarURL string `json:"avatar_url" validate:"required"`
}

type PutUserRequest struct {
	Name      string `json:"name" validate:"required"`
	AvatarURL string `json:"avatar_url" validate:"required"`
}
