package dto

import "github.com/shopspring/decimal"

//struct auxiliar pra essa porra DE REQUEST SIMPLISMENTE OUIDEIO SQL SERIO SE FGUDER
// COMO PODE ESSA PORRA NAO ACEIDTAR UM ARRAYU MANO SERIO E UMA ESTRUTURA BASICA SE
//FUDER LIXO DO CARAI NA EU DEVERIA TROCAR PARA MONGODB

type ComandaRequest struct {
	Identificacao string
	Pedidos []ComandaPedidoRequest `gorm:"many2many:pedidos;"`
	EstaFechada bool
	UserID int
	Valor decimal.Decimal `gorm:"type:decimal(10,2)" json:"Preco"`

}

type ComandaPedidoRequest struct {
	PedidoID  uint            `json:"PedidoID"`
}


