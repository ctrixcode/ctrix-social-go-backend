package database

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jmoiron/sqlx" // Changed from "database/sql"
	_ "github.com/lib/pq"
)

var dbInstance *sqlx.DB // Changed to *sqlx.DB

func DBConnection() *sqlx.DB { // Changed return type
	if dbInstance != nil {
		return dbInstance
	}
	dbConfig := getDBConfig()
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbConfig.username, dbConfig.password, dbConfig.host, dbConfig.port, dbConfig.dbname)
	db, err := sqlx.Open("postgres", connString) // Changed to sqlx.Open
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connection Established Successfully!!!")
	}

	dbInstance = db
	return dbInstance
}

func CreateInitialDBStructure() {

	db := DBConnection() // This will now return *sqlx.DB

	sqlFile, err := os.ReadFile("./sql/createTables.sql")
	if err != nil {
		log.Fatal(err)
	}
	sqlString := string(sqlFile)
	statements := strings.Split(sqlString, ";")
	for _, statement := range statements {
		statement = strings.TrimSpace(statement)
		if statement == "" {
			continue
		}
		_, err = db.Exec(statement)
		if err != nil {
			log.Fatal("Error creating initial tables: ", err)
		}
	}
}

func ResetDB() {
	dbConfig := getDBConfig()
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=postgres sslmode=disable", dbConfig.username, dbConfig.password, dbConfig.host, dbConfig.port)
	db, err := sqlx.Open("postgres", connString) // Changed to sqlx.Open
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE %s", dbConfig.dbname))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully dropped database %s\n", dbConfig.dbname)

}
func CreateDB() {
	dbConfig := getDBConfig()
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=postgres sslmode=disable", dbConfig.username, dbConfig.password, dbConfig.host, dbConfig.port)
	db, err := sqlx.Open("postgres", connString) // Changed to sqlx.Open
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbConfig.dbname))
	if err != nil {
		panic(err)
	}

	fmt.Println("DB Created Successfully!")
}

type dbConfig struct {
	host     string
	username string
	password string
	dbname   string
	port     string
}

func getDBConfig() *dbConfig {

	return &dbConfig{
		host:     os.Getenv("POSTGRES_HOST"),
		username: os.Getenv("POSTGRES_USERNAME"),
		password: os.Getenv("POSTGRES_PASSWORD"),
		dbname:   os.Getenv("POSTGRES_DB"),
		port:     os.Getenv("POSTGRES_PORT"),
	}
}