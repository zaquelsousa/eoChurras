package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db  *gorm.DB

func DbConnection(user string, pssWord string, dbName string) (*gorm.DB, error){
	var connectionString = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pssWord, dbName,
	)
	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}

func SetDB(conn *gorm.DB){
	db = conn
}

func GetDB() *gorm.DB{
	return db
}
