package models


import "gorm.io/gorm"

type Pedido struct {
	gorm.Model
	//Produtos []Produto
	ComandaID int
	Comanda Comanda `gorm:"foreignKey:ComandaID"`
	StatusPedido StatusPedido `json: StatusPedido`
}

type StatusPedido int
const(
	fazendo StatusPedido = iota
	pronto
	cancelado
)



