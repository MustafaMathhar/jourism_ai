package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"goyave.dev/goyave/v4"
	"goyave.dev/goyave/v4/config"

	"github.com/MustafaMathhar/jourism_ai/api/routes"
)

func InitializeDBConnection() *sql.DB {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[SERVER] Error loading .env file")
	}
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("[SERVER] failed to connect: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("[SERVER] failed to ping: %v", err)
	}

	log.Println("[SERVER] Successfully connected to PlanetScale!")

	return db
}

func main() {
	db := InitializeDBConnection()
	defer db.Close()

	config.Load()
	bunDb := bun.NewDB(db, mysqldialect.New())
	ds := &routes.DataStore{DB: bunDb}
	if err := goyave.Start(ds.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
	port := config.GetInt("server.port")
	log.Println("[SERVER] Started server at PORT:", port)
}
