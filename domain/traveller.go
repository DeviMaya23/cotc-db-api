package domain

type Traveller struct {
	ID        int64     `json:"id" gorm:"id"`
	Name      string    `json:"name" gorm:"name"`
	Rarity    int       `json:"rarity" gorm:"rarity"`
	Influence Influence `gorm:"foreignKey:influence_id"`
}
