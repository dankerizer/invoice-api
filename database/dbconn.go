package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Siddheshk02/jwt-auth-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var (
	DBConn     *gorm.DB
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string


)
const (

	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "root" //Enter your password for the DB
	dbname   = "invoiceid"
)

var dsn string = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
	host, port, user, password, dbname)

var DB *gorm.DB

func ConnectDb() {
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	db.AutoMigrate(&models.User{}) // we are going to create a models.go file for the User Model.
}
