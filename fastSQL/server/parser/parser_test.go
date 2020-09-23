package parser

import (
	"fmt"
	"testing"
)

func TestTokenizer(t *testing.T) {
	sql1 := "select * from Persons where id=1"

	sql2 := "CREATE TABLE Persons (Id int not null primary key, LastName varchar(255) NOT NULL, FirstName varchar(255), Address varchar(255), City varchar(255) );"

	list1 := Tokenizer(sql1)
	for _, v := range list1 {
		fmt.Println(v)
	}
	list2 := Tokenizer(sql2)
	for _, v := range list2 {
		fmt.Println(v)
	}
}
