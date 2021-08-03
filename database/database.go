package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func StartDB(){
	str := "user=postgres dbname=sergio password=123456 host=localhost sslmode=disable"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil{
		log.Fatal("error: ", err) // verificando se tem erro
	}

	db = database // setando com o database

	config, _ := db.DB()

	config.SetMaxIdleConns(10) // conexoes maximas
	config.SetMaxOpenConns(100) // conexoes maximas abertas
	config.SetConnMaxLifetime(time.Hour) // tempo total da conexao (1 hora)
}

func GetDatabase() *gorm.DB{
	return db
}



