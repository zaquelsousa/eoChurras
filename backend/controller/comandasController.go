package controller

import (
	"churras/database"
	"churras/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shopspring/decimal"
)

func getComandas(w http.ResponseWriter, r *http.Request){
	var comandas []models.Comanda

	db := database.GetDB()
	result := db.Find(&comandas)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comandas)
}

func getComanda(w http.ResponseWriter, r *http.Request){}

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

func createComanda(w http.ResponseWriter, r *http.Request){
	var comandaReq ComandaRequest
	if err := json.NewDecoder(r.Body).Decode(&comandaReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

db := database.GetDB()

	comanda := models.Comanda{
		Identificacao: comandaReq.Identificacao,
		EstaFechada: comandaReq.EstaFechada,
		UserID: comandaReq.UserID,
		Valor: comandaReq.Valor,
	}
	if err := db.Create(&comanda).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, c := range comandaReq.Pedidos {
		cp := models.ComandaPedido{
			ComandaID: comanda.ID,
			PedidoID: c.PedidoID,
		}
		if err := db.Create(&cp).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	
	BroadcastMensagem("comanda", comanda)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comanda)

}

func addOrderOnBill(comandaID uint , pedidoID uint){
	cp := models.ComandaPedido{
		ComandaID: comandaID,
		PedidoID: pedidoID,
	}

	db := database.GetDB()

	if err := db.Create(&cp).Error; err != nil {
		fmt.Println("deu merda ", err)
		return
	}

}
func updateComanda(w http.ResponseWriter, r *http.Request){}
func deleteComanda(w http.ResponseWriter, r *http.Request){}
