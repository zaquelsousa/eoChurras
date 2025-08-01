package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func DbConnection(user, pssWord, dbName string) (*gorm.DB, error) {
	var connectionString = fmt.Sprintf(
		"%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pssWord, dbName,
	)

	var err error
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
		if err == nil {
			fmt.Println("Banco conectado com sucesso.")
			return db, nil
		}
		fmt.Printf("Tentativa %d de conexão falhou: %v\n", i+1, err)
		time.Sleep(3 * time.Second)
	}
	return nil, fmt.Errorf("falha ao conectar ao banco após múltiplas tentativas: %w", err)
}

func SetDB(conn *gorm.DB) {
	db = conn
}

func GetDB() *gorm.DB {
	return db
}

