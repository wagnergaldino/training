package models

type User struct {
	Id      int    `gorm:"primary_key:true"`
	Usuario string `gorm:"type:varchar(50);not null"`
	Senha   string `gorm:"type:varchar(50);not null"`
}
