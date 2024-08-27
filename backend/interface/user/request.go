package user

type PostUserRequest struct {
	ID        string `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	AvatarURL string `json:"avatarUrl" validate:"required"`
}

type PutUserRequest struct {
	Name      string `json:"name" validate:"required"`
	AvatarURL string `json:"avatarUrl" validate:"required"`
}
