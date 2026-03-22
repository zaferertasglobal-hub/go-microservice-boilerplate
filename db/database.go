package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/zaferertasglobal-hub/go-microservice-boilerplate/config"
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB(cfg *config.Config) error {
	var err error
	DB, err = sql.Open("postgres", cfg.GetDSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return err
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Database is unreachable: %v", err)
		return err
	}

	log.Println("✅ Database connection established successfully")
	return nil
}

// CloseDB closes the database connection
func CloseDB() error {
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
			return err
		}
		log.Println("✅ Database connection closed successfully")
	}
	return nil
}