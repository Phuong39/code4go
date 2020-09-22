package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Shelnutt2/db2struct"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DirName = "dist"
)

// 前缀
var tablePrefix string

// 后缀
var tableSuffix string

// 数据库地址
var dbHost string

// 数据库端口
var dbPort int

// 表名称 *代表所有
var dbTable string

// 数据库名称
var dbDatabase string

// 用户名
var dbUser string

// 密码
var dbPassword string

// 包名
var packageName string

// 是否生成json标签
var jsonAnnotation = true

// 是否生成gorm标签
var gormAnnotation = true

//	是否保存文件，与去掉前后缀的表名一致
var saveFile bool

var nullable bool

func handleByTable(tableName string) {
	columnDataTypes, err := db2struct.GetColumnsFromMysqlTable(dbDatabase, tableName)

	if err != nil {
		fmt.Println("Error in selecting column data information from mysql information schema")
		return
	}
	structName := tableName

	if tablePrefix != "" {
		structName = strings.TrimPrefix(structName, tablePrefix)
	}
	if tableSuffix != "" {
		structName = strings.TrimSuffix(structName, tableSuffix)
	}
	var fileName string
	if saveFile {
		fileName = structName
	}
	db2struct.HumpString(&structName)
	struc, err := db2struct.Generate(*columnDataTypes, tableName, structName, packageName, jsonAnnotation, gormAnnotation, false, nullable)

	if err != nil {
		fmt.Println("Error in creating struct from json: " + err.Error())
		return
	}
	if fileName != "" {
		file, err := os.OpenFile(DirName+"/"+fileName+".go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Open File fail: " + err.Error())
			return
		}
		length, err := file.WriteString(string(struc))
		if err != nil {
			fmt.Println("Save File fail: " + err.Error())
			return
		}
		fmt.Printf("wrote %d bytes\n", length)
	} else {
		fmt.Println(string(struc))
	}
}

func init() {
	flag.StringVar(&tablePrefix, "prefix", "t_", "")
	flag.StringVar(&tableSuffix, "suffix", "", "")
	flag.StringVar(&dbHost, "h", "127.0.0.1", "IP地址")
	flag.IntVar(&dbPort, "p", 3306, "")
	flag.StringVar(&dbTable, "table", "*", "")
	flag.StringVar(&dbDatabase, "db", "novel", "")
	flag.StringVar(&dbUser, "user", "root", "")
	flag.StringVar(&dbPassword, "pass", "root", "")
	flag.StringVar(&packageName, "packageName", "model", "")

	flag.BoolVar(&jsonAnnotation, "json", true, "")
	flag.BoolVar(&gormAnnotation, "gorm", true, "")

	flag.BoolVar(&saveFile, "save", false, "")

	flag.BoolVar(&nullable, "nullable", false, "")
	flag.Parse()

	if saveFile {
		_ = os.MkdirAll(DirName, os.ModePerm)
	}
}

func main() {
	db2struct.Init(dbUser, dbPassword, dbHost, dbPort, dbDatabase)
	if dbTable == "*" {
		for _, v := range db2struct.GetAllTable() {
			handleByTable(v)
		}
	} else {
		handleByTable(dbTable)
	}
}
