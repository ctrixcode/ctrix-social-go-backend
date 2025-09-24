package database

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/config"
	"github.com/jmoiron/sqlx" // Changed from "database/sql"
	_ "github.com/lib/pq"
)

var dbInstance *sqlx.DB // Changed to *sqlx.DB

func NewDBConnection(cfg config.DatabaseConfig) *sqlx.DB {
	if dbInstance != nil {
		return dbInstance
	}
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		slog.Error("Failed to open database connection", "error", err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	} else {
		slog.Info("DB Connection Established Successfully!!!")
	}

	dbInstance = db
	return dbInstance
}

func CreateDB(cfg config.DatabaseConfig) {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=postgres sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.SSLMode)
	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		slog.Error("Failed to open database connection for create", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", cfg.DBName))
	if err != nil {
		slog.Error("Failed to create database", "db_name", cfg.DBName, "error", err)
		os.Exit(1)
	}

	slog.Info("DB Created Successfully!")
}
