package controller

import (
	"churras/database"
	"churras/helper"
	"churras/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getUser(w http.ResponseWriter, r *http.Request){
	var users []models.Users

	db := database.GetDB()
	result := db.Find(&users)

	if result.Error != nil{
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func getUsers(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var user models.Users
	db := database.GetDB()
	db.Find(&user, id)

	if user.ID != 0 {
		json.NewEncoder(w).Encode(user)
	}
	http.NotFound(w, r)
}

func createUser(w http.ResponseWriter, r *http.Request){
	//TODO: implementar JWT para create, update e delete ne paizao, somente o gerente deve ser capaz de fazer sas porra
	var user models.Users

	err := json.NewDecoder(r. Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user.PassWorld, err = helper.HashPassword(user.PassWorld)
	
	if err != nil {
		panic("deu merda na hora de fazer o hashing da senha")
	}
	
	db := database.GetDB()
	db.Create(&user)
	
}

func updateUser(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var existeingUser models.Users
	db := database.GetDB()

	err = db.First(&existeingUser, id).Error
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)

	}

	var updatedUser models.Users
	json.NewDecoder(r.Body).Decode(&updatedUser)

	existeingUser.Name = updatedUser.Name
	existeingUser.PassWorld, err = helper.HashPassword(existeingUser.PassWorld)
	existeingUser.Role = updatedUser.Role

	db.Save(&existeingUser)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existeingUser)

}

func deleteUser(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	
	db := database.GetDB()
	var user models.Users

	err = db.First(&user, id).Error

	if err != nil{
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	db.Delete(&user)
	w.WriteHeader(http.StatusNoContent)
}
