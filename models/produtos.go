package models


import "gorm.io/gorm"

type Produto struct {
	gorm.Model
	Name string
	Preco float32
	Qtd uint16
}


