package controller

import (
	"churras/database"
	"churras/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getProdutos(w http.ResponseWriter, r *http.Request){
	var produtos []models.Produto

	db := database.GetDB()
	result := db.Find(&produtos)

	if result.Error != nil{
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(produtos)
}
func getProduto(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var produto models.Produto
	db := database.GetDB()
	db.Find(&produto, id)

	if produto.ID != 0 {
		json.NewEncoder(w). Encode(produto)
	}
	http.NotFound(w, r)
}

func createProduto(w http.ResponseWriter, r *http.Request){
	var produto models.Produto

	err := json.NewDecoder(r.Body).Decode(&produto)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	//TODO some validations ne pai kkkkkkkkk se nao so cara bota ate a mae no DB

	db := database.GetDB()
	db.Create(&produto)
}

func updateProduto(w http.ResponseWriter, r *http.Request){}
func deleteProduto(w http.ResponseWriter, r *http.Request){}
