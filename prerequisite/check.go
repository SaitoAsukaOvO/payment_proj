package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"payment_proj/config"
)

func main()  {
	conf := config.LoadAppConfig()
	log.Printf("Connectting to mysql...")
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/",
		conf.Mysql.User,
		conf.Mysql.Password,
		conf.Mysql.Host,
		conf.Mysql.Port,
	))
	if err != nil {
		fmt.Printf("Error - Failed to initialize the mysql client, %s\n", err.Error())
		log.Panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS WorldWallet;")
	log.Printf("Database WorldWallet Created")
	if err != nil {
		log.Panic(err.Error())
	}
	_, err = db.Exec("USE WorldWallet;")
	if err != nil {
		log.Panic(err.Error())
	}
	log.Printf("Database checked")
}