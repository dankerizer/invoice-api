package database

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"os"

	"github.com/Siddheshk02/jwt-auth-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB     *gorm.DB
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

// var dsn string = fmt.Sprintf("host=%s port=%d user=%s "+
// 	"password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
// 	host, port, user, password, dbname)

// var DB *gorm.DB

func ConnectDb() {

	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", DbHost, DbPort, DbUser, DbPassword, DbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Use(tracing.NewPlugin())
	if err != nil {
		log.Error(err)
	}

	if err != nil {
		log.Printf("Gak bisa konnect ke database. %v\n", err)
		os.Exit(2)
	}
	log.Println("Terhubung ke DB")

	sqldb, err := db.DB()
	if err != nil {
		log.Printf("Error %v", err)
	}
sqldb.SetMaxIdleConns(10)
	sqldb.SetMaxOpenConns(100)
	sqldb.SetConnMaxIdleTime(time.Hour)
	sqldb.SetConnMaxLifetime(time.Hour)
	DB = db

	db.AutoMigrate(&models.User{}) // we are going to create a models.go file for the User Model.
}
