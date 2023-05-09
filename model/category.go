package model

type Category struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"type:varchar(20);not null; unique"`
	CreatedAt Time   `json:"created_at" gorm:"type:timestamp"`
	UpdateAt  Time   `json:"update_at" gorm:"type:timestamp"`
}
