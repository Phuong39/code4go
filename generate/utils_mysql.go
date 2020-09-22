package db2struct

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var db *sql.DB

type ColumnBean struct {
	ColumnName string
	Sort       int
}

func Init(mariadbUser string, mariadbPassword string, mariadbHost string, mariadbPort int, mariadbDatabase string) {
	var err error
	if mariadbPassword != "" {
		db, err = sql.Open("mysql", mariadbUser+":"+mariadbPassword+"@tcp("+mariadbHost+":"+strconv.Itoa(mariadbPort)+")/"+mariadbDatabase+"?&parseTime=True")
	} else {
		db, err = sql.Open("mysql", mariadbUser+"@tcp("+mariadbHost+":"+strconv.Itoa(mariadbPort)+")/"+mariadbDatabase+"?&parseTime=True")
	}
	if err != nil {
		panic(err)
	}
}

func GetColumnsFromMysqlTable(databaseName, tableName string) (*map[*ColumnBean]map[string]string, error) {
	// Store colum as map of maps
	columnDataTypes := make(map[*ColumnBean]map[string]string)
	// Select columnd data from INFORMATION_SCHEMA
	columnDataTypeQuery := "SELECT COLUMN_NAME, COLUMN_KEY, DATA_TYPE, IS_NULLABLE FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = ? AND table_name = ?"

	if Debug {
		fmt.Println("running: " + columnDataTypeQuery)
	}
	rows, err := db.Query(columnDataTypeQuery, databaseName, tableName)

	if err != nil {
		fmt.Println("Error selecting from db: " + err.Error())
		return nil, err
	}
	if rows != nil {
		defer rows.Close()
	} else {
		return nil, errors.New("No results returned for table")
	}
	var sort int
	for rows.Next() {
		var column string
		var columnKey string
		var dataType string
		var nullable string
		rows.Scan(&column, &columnKey, &dataType, &nullable)
		columnDataTypes[&ColumnBean{column, sort}] = map[string]string{"value": dataType, "nullable": nullable, "primary": columnKey}
		sort++
	}
	return &columnDataTypes, err
}

// Generate go struct entries for a map[string]interface{} structure
func generateMysqlTypes(obj map[*ColumnBean]map[string]string, depth int, jsonAnnotation bool, gormAnnotation bool, gureguTypes bool, nullableC bool) string {
	structure := "struct {"

	keys := make([]*ColumnBean, len(obj), len(obj))

	for i := 0; i < len(obj); i++ {
		keys[i] = getBySort(i, obj)
	}
	//for index,key := range obj {
	//	keys[index] = obj[index]
	//	keys = append(keys, key)
	//}
	// sort.Strings(keys)

	for _, key := range keys {
		mysqlType := obj[key]
		nullable := false
		if mysqlType["nullable"] == "YES" {
			nullable = true
		}

		primary := ""
		if mysqlType["primary"] == "PRI" {
			primary = ";primary_key"
		}

		// Get the corresponding go value type for this mysql type
		var valueType string
		// If the guregu (https://github.com/guregu/null) CLI option is passed use its types, otherwise use go's sql.NullX

		valueType = mysqlTypeToGoType(mysqlType["value"], nullable && nullableC, gureguTypes)

		fieldName := fmtFieldName(stringifyFirstChar(key.ColumnName))
		var annotations []string
		if gormAnnotation == true {
			annotations = append(annotations, fmt.Sprintf("gorm:\"column:%s%s\"", key.ColumnName, primary))
		}
		if jsonAnnotation == true {
			annotations = append(annotations, fmt.Sprintf("json:\"%s\"", key.ColumnName))
		}
		if len(annotations) > 0 {
			structure += fmt.Sprintf("\n%s %s `%s`",
				fieldName,
				valueType,
				strings.Join(annotations, " "))

		} else {
			structure += fmt.Sprintf("\n%s %s",
				fieldName,
				valueType)
		}
	}
	return structure
}

func getBySort(sort int, obj map[*ColumnBean]map[string]string) *ColumnBean {
	for k := range obj {
		if k.Sort == sort {
			return k
		}
	}
	return nil
}

// mysqlTypeToGoType converts the mysql types to go compatible sql.Nullable (https://golang.org/pkg/database/sql/) types
func mysqlTypeToGoType(mysqlType string, nullable bool, gureguTypes bool) string {
	switch mysqlType {
	case "tinyint", "int", "smallint", "mediumint":
		if nullable {
			if gureguTypes {
				return gureguNullInt
			}
			return sqlNullInt
		}
		return golangInt
	case "bigint":
		if nullable {
			if gureguTypes {
				return gureguNullInt
			}
			return sqlNullInt
		}
		return golangInt64
	case "char", "enum", "varchar", "longtext", "mediumtext", "text", "tinytext", "json":
		if nullable {
			if gureguTypes {
				return gureguNullString
			}
			return sqlNullString
		}
		return "string"
	case "date", "datetime", "time", "timestamp":
		if nullable && gureguTypes {
			return gureguNullTime
		}
		return golangTime
	case "decimal", "double":
		if nullable {
			if gureguTypes {
				return gureguNullFloat
			}
			return sqlNullFloat
		}
		return golangFloat64
	case "float":
		if nullable {
			if gureguTypes {
				return gureguNullFloat
			}
			return sqlNullFloat
		}
		return golangFloat32
	case "binary", "blob", "longblob", "mediumblob", "varbinary":
		return golangByteArray
	}
	return ""
}

func GetAllTable() []string {
	rows, err := db.Query("show tables")
	if err != nil {
		panic(err)
	}
	result := make([]string, 0, 10)
	for rows.Next() {
		var tableName string
		_ = rows.Scan(&tableName)
		result = append(result, tableName)
	}
	return result
}
