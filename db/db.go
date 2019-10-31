package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
type DB struct {
	ReadInstance *sqlx.DB
	WriteInstance *sqlx.DB
}
var dbConfig *DB
func GetInstance() *DB{
	if(dbConfig == nil){
		dbConnection,err := sqlx.Connect("mysql","root:@/todo?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic("failed to connect to db")
		}else{
			dbConfig = &DB{dbConnection,dbConnection}
			fmt.Println("connected to db")
		}
		return dbConfig
	}else{
		return dbConfig
	}
}

func init(){
	dbConnection,err := sqlx.Connect("mysql","root:@/todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to db")
	}else{
		dbConfig = &DB{dbConnection,dbConnection}
		fmt.Println("connected to db")
	}
}
