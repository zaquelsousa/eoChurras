package controller

import (
	"churras/database"
	"churras/models"
	"churras/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
)

func getPedidos(w http.ResponseWriter, r *http.Request){
	var pedidos []models.Pedido

	db := database.GetDB()
	result := db.Preload("Produtos").Find(&pedidos)

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
	if err := db.Preload("Produtos").First(&pedido, pedido.ID).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := services.AddOrderOnbill(uint(pedidoReq.ComandaID), pedido.ID)

	if err != nil {
		http.Error(w, fmt.Sprintf("Erro: %s", err.Error()), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
	BroadcastMensagem("pedido", pedido)
	json.NewEncoder(w).Encode(pedido)
}


type Notificacao struct {
	Tipo     string `json:"tipo"`
    Mensagem string `json:"mensagem"`
}


func pedidoPronto(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id := params["id"]

	var pedido  models.Pedido
	db := database.GetDB()
	db.Find(&pedido, id)

	if pedido.StatusPedido == 0 {
		var pedido_produto models.PedidoProduto
		db.Where("pedido_id = ?", id).First(&pedido_produto)
		
		var comanda models.Comanda
		db.Find(&comanda, pedido.ComandaID)
		valor := pedido_produto.Preco
		comanda.Valor = comanda.Valor.Add(valor)
		db.Save(&comanda)

		pedido.StatusPedido = 1
		db.Save(&pedido)

		notificacao := Notificacao{
			Tipo:     "pedido_pronto",
			Mensagem: "O pedido da comanda "+ comanda.Identificacao + " está pronto para ser entregue.",
		}
		BroadcastMensagem("notificacao", notificacao)
	}



}

func updatePedido(w http.ResponseWriter, r *http.Request){}
func deletePedido(w http.ResponseWriter, r *http.Request){}
