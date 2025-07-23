package main

import (
	"churras/database"
	"log"
)

func main() {
	//db confs tenho que passar essas credenciaais para um lugar mais seruhro tipo um .env
	user, pssWord, dbName := "zaquel", "1234", "churrasquinho"

	db, err := database.DbConnection(user, pssWord, dbName)

	if err != nil {
		log.Fatalln(err)
	}
	slqDB, err := db.DB()
	err = slqDB.Ping()
	if err != nil {
		log.Fatalln("deu ruim no ping com a DB", err)
	}
}
