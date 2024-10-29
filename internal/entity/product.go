package entity

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Description string
	Price       uint
}
