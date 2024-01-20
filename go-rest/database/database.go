package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq" // don't forget to add it. It doesn't be added automatically
)

var Db *sql.DB //created outside to make it global.

// make sure your function start with uppercase to call outside of the directory.
func ConnectDatabase() {
	//
	//err := godotenv.Load() //by default, it is .env, so we don't have to write
	//if err != nil {
	//	log.Println("Error is occurred  on .env file please check", err)
	//}
	//we read our .env file
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT")) // don't forget to convert int since port is int type.
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASSWORD")
	maxIdleConn, _ := strconv.Atoi(os.Getenv("POOL_MAX_IDLE_CONN"))
	maxOpenConn, _ := strconv.Atoi(os.Getenv("POOL_MAX_OPEN_CONN"))

	// set up postgres sql to open it.
	//psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
	//	host, port, user, dbname, pass)
	psqlSetup := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user,
		pass,
		host,
		port,
		dbname,
	)
	log.Println("Connecting with: ", psqlSetup)
	db, err := sql.Open("postgres", psqlSetup)

	//defer db.Close()

	if err != nil {
		log.Println("There is an error while connecting to the database ", err)
		log.Fatal(err)
	}

	Db = db
	Db.SetMaxIdleConns(maxIdleConn)
	Db.SetMaxOpenConns(maxOpenConn)
	err = Db.Ping()

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully connected to database!")
	}

}
