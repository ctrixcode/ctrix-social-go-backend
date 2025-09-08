package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var dbInstance *sql.DB

func DBConnection() *sql.DB {
	if dbInstance != nil {
		return dbInstance
	}
	dbConfig := getDBConfig()
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbConfig.username, dbConfig.password, dbConfig.host, dbConfig.port, dbConfig.dbname)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	dbInstance = db
	return dbInstance
}

func CreateInitialDBStructure() {

	// db, err := DBConnection()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// sqlFile, err := os.ReadFile("./sql/createTables.sql")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = db.Exec(string(sqlFile)).Error
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Initial Tables are Created Successfully!!!")
	// }

}
func ResetDB() {
	// db, err := DBConnection()
	// if err != nil {
	// 	fmt.Println("error here!1")
	// 	log.Fatal("Error connecting to db: ", err)
	// }
	// if err := db.Exec("DROP DATABASE Ctrix_Social_DB"); err != nil {
	// 	fmt.Println("error here!2")

	// 	log.Fatal(err)
	// }
	// if err := db.Exec("CREATE DATABASE IF NOT EXISTS Ctrix_Social_DB"); err != nil {
	// 	fmt.Println("error here!3")

	// 	log.Fatal(err)
	// }

	// sqlFile, err := os.ReadFile("./sql/resetDB.sql")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = db.Exec(string(sqlFile)).Error
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("DB Resetted Successfully!!!")
	// }

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
		dbname:   os.Getenv("POSTGRES_DBNAME"),
		port:     os.Getenv("POSTGRES_PORT"),
	}
}
