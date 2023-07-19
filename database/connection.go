package database

import (
	"fmt"
	"time"
	
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//Connect to sql db
func DbConn() (db *sql.DB) {
	db, err := sql.Open("mysql","root:password@tcp(localhost:3306)/database")  
	if err != nil {
	log.Printf("Setup MySQL connect error %+v\n", err)
	}
	return db
}

//Mysql db setup
func Setup() {
	// init mysql.
	db := DbConn()

	defer db.Close()
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println("MySQL connection successful")
}


