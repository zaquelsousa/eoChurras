package models

import (
	"time"

	"gorm.io/gorm"
)

type Comanda struct {
	gorm.Model
	Identificaçao string
	Pedidos []Pedido
	EstaFechada bool
	Data time.Time
}


