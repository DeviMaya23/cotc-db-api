package domain

type Traveller struct {
	CommonModel
	Name        string    `json:"name" gorm:"name"`
	Rarity      int       `json:"rarity" gorm:"rarity"`
	InfluenceID int       `json:"influence_id" gorm:"influence_id"`
	Influence   Influence `gorm:"foreignKey:influence_id"`
}

func (Traveller) TableName() string {
	return "tr_traveller"
}

type CreateTravellerRequest struct {
	Name      string `json:"name" validate:"required,lte=50"`
	Rarity    int    `json:"rarity" validate:"required"`
	Influence string `json:"influence" validate:"required,influence"`
}
