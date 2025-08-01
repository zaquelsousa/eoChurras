package models


import "gorm.io/gorm"

type Pedido struct {
	gorm.Model
	Produtos    []PedidoProduto `gorm:"foreignKey:PedidoID"`
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



