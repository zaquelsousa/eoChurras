package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Comanda struct {
	gorm.Model
	Identificaçao string
	//Pedidos []Pedido `gorm:"many2many:pedidos;"`
	EstaFechada bool
	UserID int
	User Users `gorm:"foreignKey:UserID"`
	Valor decimal.Decimal `gorm:"type:decimal(10,2)" json:"Preco"`

}


type ComandaPedido struct {
	gorm.Model
	ComandaID   uint	`json:"ComandaID"`
	Comanda Comanda		`gorm:"foreignKey:ComandaID" json:"-"`
	PedidoID  uint     `json:"PedidoID"`
	Pedido Pedido       `gorm:"foreignKey:PedidoID" json:"-"`
}


