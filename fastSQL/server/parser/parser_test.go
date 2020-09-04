package parser

import (
	"fmt"
	"testing"
)

func TestTokenizer(t *testing.T) {
	list := Tokenizer("CREATE TABLE Persons (Id int not null primary key, LastName varchar(255) NOT NULL, FirstName varchar(255), Address varchar(255), City varchar(255) );")

	for _, v := range list {
		fmt.Println(v)
	}
}
