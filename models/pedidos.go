package models


import "gorm.io/gorm"

type Pedido struct {
	gorm.Model
	Name string
	Produtos []Produto
	StatusPedido StatusPedido
}

type StatusPedido int
const(
	fazendo StatusPedido = iota
	pronto
	cancelado
)



