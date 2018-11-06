package dbops

import (
	"database/sql"
	"strconv"
	"sync"
	"test/simpleWebServer/defs"
)

/**
session的回话机制是web关键、不论何种介子、所有的原理都是对session crud
 */
func InserSession(sid string,ttl int64,uname string) error {
	ttlstr := strconv.FormatInt(ttl,10)
	stmtIns,err := dbConn.Prepare("insert into session (session_id,TTl,login_name) values (?,?,?)")
	if err != nil {
		return err
	}

	_,err = stmtIns.Exec(sid,ttlstr,uname)
	if err != nil {
		return err
	}

	//release db connecting resource
	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession,error) {
	ss := &defs.SimpleSession{}
	stmtOut,err := dbConn.Prepare("select ttl,login_name from sessions where session_id = ?")
	if err != nil {
		return nil, err
	}
	var ttl,uname string

	stmtOut.QueryRow(sid).Scan(&ttl,&uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if res,err := strconv.ParseInt(ttl,10,64); err == nil {
		ss.TTl = res
		ss.Username = uname
	}else {
		return nil,err
	}

	defer stmtOut.Close()
	return ss,nil
}

func RetrieveAllSessions()(*sync.Map,error) {
	m := &sync.Map{} // 同步map，原子操作
	stmtOut,err := dbConn.Prepare("select * from sessions")
	if err != nil {
		return nil,err
	}

	rows,err := stmtOut.Query()
	if err != nil {
		return nil,err
	}

	for rows.Next() {
		var id,ttlstr,login_name string
		if er := rows.Scan(&id,&ttlstr,&login_name);er != nil {
			break
		}
		if ttl, err1 := strconv.ParseInt(ttlstr,10,64);err1 != nil {
			ss := &defs.SimpleSession{Username:login_name,TTl:ttl}
			m.Store(id,ss) // 入map,以session_id为key
		}
	}
	return m,nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil {
		return err
	}
	if _,err := stmtOut.Query(sid);err != nil {
		return err
	}
	return nil
}



