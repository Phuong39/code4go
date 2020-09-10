package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

var DB_URL = `root:fuck123@tcp(119.29.117.244:3306)/test?charset=utf8`

func init() {
	temp, err := sql.Open("mysql", DB_URL)
	if err != nil {
		panic(err)
	}
	DB = temp
}

func main() {
	open()

	//result, err := DB.Query("SHOW TABLES")
	//if err != nil {
	//	panic(err)
	//}
	//for result.Next() {
	//	var val string
	//	fmt.Println(result.Columns())
	//	result.Scan(&val)
	//	fmt.Println(val)
	//}
}
