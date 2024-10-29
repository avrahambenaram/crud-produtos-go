package entity

type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
