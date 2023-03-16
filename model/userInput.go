package model

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Admin    bool   `json:"admin" `
}

type UserUpdateInput struct {
	Password string `json:"password"`
	Admin    bool   `json:"admin" `
}
