package models


import (
	"gorm.io/gorm"
	"github.com/shopspring/decimal"
)


type Produto struct {
	gorm.Model
	Name string          `json:"Name"`
	Preco decimal.Decimal `gorm:"type:decimal(10,2)" json:"Preco"`
	Qtd  uint16           `json:"Qtd"`
}



type PedidoProduto struct {
	gorm.Model
	PedidoID   uint            `json:"PedidoID"`
	Pedido     Pedido          `gorm:"foreignKey:PedidoID" json:"-"`
	ProdutoID  uint            `json:"ProdutoID"`
	Produto    Produto         `gorm:"foreignKey:ProdutoID" json:"-"`
	Quantidade uint16          `json:"Quantidade"`
	Preco      decimal.Decimal `gorm:"type:decimal(10,2)" json:"Preco"` 
}


