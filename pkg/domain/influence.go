package domain

type Influence struct {
	CommonModel
	Name string `json:"name" gorm:"name"`
}

func (Influence) TableName() string {
	return "m_influence"
}
