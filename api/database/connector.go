package database

import (
	"fmt"
	"log"
	"os"
	"vocabs-backend/api/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func GetDatabase() *gorm.DB {
	return db
}

func Migrate() {
	db.AutoMigrate(&model.Vocab{})
	// db.AutoMigrate(&UserEntity{})
}

func CloseConnection() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}

func InitConnection() *gorm.DB {
	loadEnvFile()
	dialect := os.Getenv("DIALECT")
	return connectToDatabase(dialect, buildDSN())
}

func loadEnvFile() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal(err)
	}
}

func buildDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
		os.Getenv("HOST"),
		os.Getenv("USER"),
		os.Getenv("NAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBPORT"))
}

func connectToDatabase(dialect string, dsn string) *gorm.DB {
	pgDialector := buildPostgresDialector(dsn)
	db, err = gorm.Open(pgDialector, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database connection established")
	}
	return db
}

func buildPostgresDialector(dsn string) gorm.Dialector {
	return postgres.New(
		postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		})
}
