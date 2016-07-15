package models

type Address struct {
	Cep         string `gorm:"type:varchar(10);primary_key:true"`
	Logradouro  string `gorm:"type:varchar(110)"`
	Bairro      string `gorm:"type:varchar(110)"`
	Complemento string `gorm:"type:varchar(110)"`
	Nome        string `gorm:"type:varchar(110)"`
	Localidade  string `gorm:"type:varchar(70)"`
	Uf          string `gorm:"type:varchar(5)"`
	Restricao   string `gorm:"type:varchar(110)"`
	Latitude    string `gorm:"type:varchar(20)"`
	Longitude   string `gorm:"type:varchar(20)"`
}
