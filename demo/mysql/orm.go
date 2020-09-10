/*
* 作者：刘时明
* 时间：2020/9/10 0010-15:54
* 作用：
 */
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func open() {
	db, err := gorm.Open("mysql", DB_URL)
	if err != nil {
		panic(err)
	}

	query, err := db.DB().Query("SHOW TABLES")
	columns, err := query.Columns()

	fmt.Println(columns)
}
