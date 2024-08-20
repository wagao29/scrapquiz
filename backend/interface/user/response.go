package user

type (
	GetUserResponse  = userResponseModel
	PostUserResponse = userResponseModel
	PutUserResponse  = userResponseModel
)

type GetUsersResponse = struct {
	User []userResponseModel `json:"users"`
}

type userResponseModel struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}
