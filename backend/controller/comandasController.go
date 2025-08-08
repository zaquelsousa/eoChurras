package controller

import (
	"churras/database"
	"churras/dto"
	"churras/models"
	"churras/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
)

func getComandas(w http.ResponseWriter, r *http.Request){
	
	comandas, err := services.GetAllComandas()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comandas)
}


func getComanda(w http.ResponseWriter, r *http.Request){}


func createComanda(w http.ResponseWriter, r *http.Request){
	var comanda dto.ComandaRequest
	if err := json.NewDecoder(r.Body).Decode(&comanda); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := services.NewTab(comanda)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	BroadcastMensagem("comanda", comanda)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comanda)

}

//DTO sei la como chama isso  eum objeto pra virar o json pro front
type Produto struct {
    Name          string          `json:"name"`
    Qtd           uint            `json:"qtd"`
    ValorUnitario decimal.Decimal `json:"valor_unitario"`
    ValorTT       decimal.Decimal `json:"valor_total"`
}

type ProdutoList struct {
    Produtos         []Produto      `json:"produtos"`
    ValorTotalPedido decimal.Decimal `json:"valor_total_pedido"`
}

//func for when client close the bill
func closeTab(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid comanda ID", http.StatusBadRequest)
		return
	}

	var comandaPedido []models.ComandaPedido
	db := database.GetDB()
	result := db.Where("comanda_id = ?", id).Find(&comandaPedido)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	
	//vo usar uma hashmap pq ai eu consigo saber se os rpduitos se repetem, ai eu concateno pra a lista de produtos nao ficar comfusa
	produtoMap := make(map[uint]Produto)
	var valorTotal decimal.Decimal

	for _, cp := range comandaPedido {
		var pedidoProduto []models.PedidoProduto
		result := db.Where("pedido_id = ?", cp.PedidoID).Find(&pedidoProduto)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		for _, pp := range pedidoProduto {
			var produtoFromDb models.Produto
			result := db.First(&produtoFromDb, pp.ProdutoID)
			if result.Error != nil {
				http.Error(w, result.Error.Error(), http.StatusInternalServerError)
				return
			}
			
			//verifica se tem itens iguais e concatrena
			//legal do go é que qudn usa um valor de map pode ter dopis valores o chave
			//eoo esse found que é um bool onde ele e true se a chave existe tropzera ne nao
			key := produtoFromDb.ID
			if existing, found := produtoMap[key]; found {
				existing.Qtd += uint(pp.Quantidade)
				existing.ValorTT = existing.ValorTT.Add(pp.Preco)
				produtoMap[key] = existing
			} else {
				produtoMap[key] = Produto{
					Name:          produtoFromDb.Name,
					Qtd:           uint(pp.Quantidade),
					ValorUnitario: produtoFromDb.Preco,
					ValorTT:       pp.Preco,
				}
			}

			valorTotal = valorTotal.Add(pp.Preco)
		}
	}

	//converte o map em slice
	var produtos []Produto
	for _, p := range produtoMap {
		produtos = append(produtos, p)
	}

	allProds := ProdutoList{
		Produtos:         produtos,
		ValorTotalPedido: valorTotal,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allProds)
}


func updateComanda(w http.ResponseWriter, r *http.Request){}
func deleteComanda(w http.ResponseWriter, r *http.Request){}
