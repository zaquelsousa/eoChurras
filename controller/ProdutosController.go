package controller

import (
	"churras/database"
	"churras/models"
	"encoding/json"
	"net/http"
)

func getProdutos(w http.ResponseWriter, r *http.Request){}
func getProduto(w http.ResponseWriter, r *http.Request){}

func createProduto(w http.ResponseWriter, r *http.Request){
	var produto models.Produtos

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
