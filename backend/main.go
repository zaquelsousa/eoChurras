package main

import (
	"churras/controller"
	"churras/database"
	"churras/models"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func main() {
	//laod env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("deu merda ao carregar env file %s", err)
	}

	//db confs tenho que passar essas credenciaais para um lugar mais seruhro tipo um .env

	user := os.Getenv("DB_USER")
	pssWord := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	db, err := database.DbConnection(user, pssWord, dbName)

	if err != nil {
		log.Fatalln(err)
	}
	
	database.SetDB(db)
	db.AutoMigrate(&models.Users{}, &models.Produto{}, models.Pedido{}, models.Comanda{}, models.PedidoProduto{}, models.ComandaPedido{})

	controller.Router()
}
