package dto

type CreateUserRequest struct {
	FirstName          string `json:"firstName" binding:"required"`
	LastName           string `json:"lastName" binding:"required"`
	UserName           string `json:"userName" binding:"required"`
	Age                int    `json:"age" binding:"required"`
	Email              string `json:"email" binding:"required,email"`
	Password           string `json:"password" binding:"required,min=8"`
	ChannelName        string `json:"channelName" binding:"required"`
	ChannelDescription string `json:"channelDescription" binding:"required"`
}

type GetUserResponse struct {
	UserID    int64  `json:"userId"`
	Email     string `json:"email"`
	UserName  string `json:"userName"`
	Age       int    `json:"age"`
	ChannelID int64  `json:"channelId"`
}

type UpdateUserRequest struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	UserName  *string `json:"userName,omitempty"`
	Age       *int    `json:"age,omitempty"`
	Email     *string `json:"email,omitempty"`
}

type UserSummary struct {
	UserID   int64  `json:"userId"`
	UserName string `json:"userName"`
}
