package dbops

import (
	_ "github.com/go-sql-driver/mysql"
)

func AddUser(loginName string,pwd string) error {
	stmtIns, err := dbConn.Prepare("insert into user (name,pwd) values (?,?)")
	if err != nil {
		return err
	}

	_,err = stmtIns.Exec(loginName,pwd)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}
