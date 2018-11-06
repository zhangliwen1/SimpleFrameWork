package session

import (
	"test/simpleWebServer/utils"
	"sync"
	"test/simpleWebServer/dbops"
	"test/simpleWebServer/defs"
	"time"
)

var sessionMap *sync.Map

func init(){
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().Unix() / 1000000
}

func deleteExpiredSession(sid string){
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func LoadSessionsFromDB() {
	r,err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(key, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(key,ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id,_:= utils.NewUUID()
	ct := nowInMilli()
	ttl := ct + 30 * 60 * 1000 //30min

	ss := &defs.SimpleSession{Username:un,TTl:ttl}
	sessionMap.Store(id,ss)
	dbops.InserSession(id,ttl,un)
	return id
}

func IsSessionExpired(sid string) (string,bool){
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTl < ct { // ducking type 断言
			deleteExpiredSession(sid) // 清理过期
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, true
	}
	return "", true
}
