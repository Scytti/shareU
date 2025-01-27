package main

import (
	"database/sql"
	"embed"
	"fmt"
	"os"
	"shareU/internal/config"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	var db *sql.DB
	cfg := config.MustLoad()
	dbURL := cfg.DBConfig.ConnectionURL()
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	path, _ := os.Getwd()
	println(path)
	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
}
