package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"test/simpleWebServer/dbops"
	"test/simpleWebServer/defs"
	"test/simpleWebServer/session"
)

func addUser(w http.ResponseWriter,req *http.Request,p httprouter.Params) {
	reqParams,_ := ioutil.ReadAll(req.Body) //获取参数
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(reqParams,ubody); err != nil { //json还原到struct
		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUser(ubody.UserName,ubody.Pws); err != nil { //入库
		sendErrorResponse(w,defs.ErrorDBError)
	}
	id := session.GenerateNewSessionId(ubody.UserName) // 生产sessionid
	su := &defs.SignedUp{Success: true, SessionId: id} //返回格式

	if resp, err := json.Marshal(su); err != nil { // json序列化
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}

}

func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w,uname)
}