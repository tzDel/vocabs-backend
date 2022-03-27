package commons

import (
	"fmt"
	"log"
	"os"
	"vocabs-backend/src/vocabs"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func MigrateDatabase() {
	db.AutoMigrate(&vocabs.VocabEntity{})
	// db.AutoMigrate(&UserEntity{})
}

func CloseDatabase() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}

func InitDatabaseConnection() {
	loadEnvFile()
	dialect := os.Getenv("DIALECT")
	connectToDatabase(dialect, buildDSN())
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

func connectToDatabase(dialect string, dsn string) {
	pgDialector := buildPostgresDialector(dsn)
	db, err = gorm.Open(pgDialector, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database connection established")
	}
}

func buildPostgresDialector(dsn string) gorm.Dialector {
	return postgres.New(
		postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		})
}
