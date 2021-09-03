package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DatabaseInit() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:rakamin@tcp(localhost:3306)/coba")

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	fmt.Println("Connected to database")

	return db, nil
}

func GormDatabaseConn() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot open file .env")
	}

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_pass, db_host, db_port, db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Errorf("Cannot connect db", err)
		return nil, err
	}

	return db, nil
}
