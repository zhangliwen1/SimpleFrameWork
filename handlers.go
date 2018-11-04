package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"test/simpleWebServer/dbops"
	"test/simpleWebServer/defs"
)

func addUser(w http.ResponseWriter,req *http.Request,p httprouter.Params) {
	reqParams,_ := ioutil.ReadAll(req.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(reqParams,ubody); err != nil {
		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUser(ubody.UserName,ubody.Pws); err != nil {
		sendErrorResponse(w,defs.ErrorDBError)
	}
	su := &defs.SignedUp{Success:true,SessionId:"xxxdfdfkdjfkdjfk"}
	if resp ,err := json.Marshal(su); err != nil {
		sendErrorResponse(w,defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w,string(resp),201)
	}

}