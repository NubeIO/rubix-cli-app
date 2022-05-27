package model

type User struct {
	UUID      string `json:"uuid" gorm:"primary_key" get:"true" delete:"true"`
	Username  string `json:"username" get:"true" post:"true" patch:"true"`
	IsAdmin   bool   `json:"is_admin" get:"true" post:"true" patch:"true"`
	UserGroup bool   `json:"user_group" get:"true" post:"true" patch:"true"`
	Email     string `json:"email" get:"true" post:"true" patch:"true"`
	Password  string `json:"password" get:"false" post:"true" patch:"true"`
	Hash      string `json:"-"`
	UID       string `json:"-"`
	Role      string `json:"-"`
}

type NewUser struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
