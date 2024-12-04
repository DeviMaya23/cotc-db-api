package domain

type Influence struct {
	ID   int64  `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name"`
}
