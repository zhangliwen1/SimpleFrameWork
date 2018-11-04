package dbops

import "database/sql"

var (
	dbConn *sql.DB
	err error
)

func init(){
	db,err := sql.Open("mysql","root:199999@tcp(127.0.0.1:3306)/beego?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	dbConn = db
}
