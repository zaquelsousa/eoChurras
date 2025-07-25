package controller

import (
	"churras/database"
	"churras/models"
	"encoding/json"
	"github.com/shopspring/decimal"
	"net/http"
)

func getPedidos(w http.ResponseWriter, r *http.Request){
	var pedidos []models.Pedido

	db := database.GetDB()
	result := db.Find(&pedidos)

	if result.Error != nil{
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pedidos)
}
func getPedido(w http.ResponseWriter, r *http.Request){}


//struct auxiliar pra essa porra DE REQUEST SIMPLISMENTE OUIDEIO SQL SERIO SE FGUDER
// COMO PODE ESSA PORRA NAO ACEIDTAR UM ARRAYU MANO SERIO E UMA ESTRUTURA BASICA SE 
//FUDER LIXO DO CARAI NA EU DEVERIA TROCAR PARA MONGODB

type PedidoRequest struct {
	StatusPedido   models.StatusPedido           `json:"StatusPedido"`
	Produtos       []PedidoProdutoRequest        `json:"Produtos"`
	ComandaID int
}

type PedidoProdutoRequest struct {
	ProdutoID  uint            `json:"ProdutoID"`
	Quantidade uint16          `json:"Quantidade"`
	Preco      decimal.Decimal `json:"Preco"`
}

func createPedido(w http.ResponseWriter, r *http.Request) {
	var pedidoReq PedidoRequest
	if err := json.NewDecoder(r.Body).Decode(&pedidoReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := database.GetDB()

	// Cria o Pedido
	pedido := models.Pedido{
		StatusPedido: pedidoReq.StatusPedido,
		ComandaID: pedidoReq.ComandaID,
	}
	if err := db.Create(&pedido).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Cria os PedidoProduto associados
	for _, p := range pedidoReq.Produtos {
		pp := models.PedidoProduto{
			PedidoID:   pedido.ID,
			ProdutoID:  p.ProdutoID,
			Quantidade: p.Quantidade,
			Preco:      p.Preco,
		}
		if err := db.Create(&pp).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	addOrderOnBill(uint(pedidoReq.ComandaID), pedido.ID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pedido)
}


func updatePedido(w http.ResponseWriter, r *http.Request){}
func deletePedido(w http.ResponseWriter, r *http.Request){}
