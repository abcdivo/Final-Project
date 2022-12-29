package mysql

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitMysqlDB() *sql.DB {
	var (
		errMysql error
		dbConn   *sql.DB
	)
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := ""
	dbName := "db_gamestore"

	//root:@tcp(localhost:3306)/db_salt_academy
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, errMysql = sql.Open(`mysql`, dsn)

	if errMysql != nil {
		fmt.Println("error connect")
	}

	dbConn.SetMaxOpenConns(300)
	dbConn.SetMaxIdleConns(25)
	dbConn.SetConnMaxLifetime(5 * time.Minute)

	return dbConn

}
