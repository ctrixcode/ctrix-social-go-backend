package db

// var dbInstance *gorm.DB

func DBConnection() {
	// if dbInstance != nil {
	// 	return dbInstance, nil
	// }

	// var host, username, password, dbname string = "", "", "", ""

	// currentEnv := os.Getenv("APP_ENV")
	// if currentEnv == "dev" {
	// 	host = os.Getenv("postgresHostDev")
	// 	dbname = os.Getenv("postgresDBDev")
	// 	username = os.Getenv("postgresUsernameDev")
	// 	password = os.Getenv("postgresPasswordDev")
	// }

	// if currentEnv == "production" {
	// 	host = os.Getenv("postgresHostProd")
	// 	dbname = os.Getenv("postgresDBProd")
	// 	username = os.Getenv("postgresUsernameProd")
	// 	password = os.Getenv("postgresPasswordProd")
	// }

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, username, password, dbname)
	//Open the connection

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	TranslateError: true,
	// 	NamingStrategy: schema.NamingStrategy{
	// 		SingularTable: true,
	// 	},
	// })

	// if err != nil {
	// 	return nil, err
	// }
	// dbInstance = db
	// return dbInstance, nil
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
