package domain

type User struct {
	CommonModel
	Username string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"`
	Token    string `json:"token" gorm:"token"`
}

func (User) TableName() string {
	return "m_user"
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
