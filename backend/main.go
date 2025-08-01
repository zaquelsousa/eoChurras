package main

import (
	"churras/controller"
	"churras/database"
	"churras/models"
	"log"
)


func main() {
	//db confs tenho que passar essas credenciaais para um lugar mais seruhro tipo um .env
	user, pssWord, dbName := "zaquel", "1234", "churrasquinho"

	db, err := database.DbConnection(user, pssWord, dbName)

	if err != nil {
		log.Fatalln(err)
	}
	
	database.SetDB(db)
	db.AutoMigrate(&models.Users{}, &models.Produto{}, models.Pedido{}, models.Comanda{}, models.PedidoProduto{}, models.ComandaPedido{})

	controller.Router()
}
