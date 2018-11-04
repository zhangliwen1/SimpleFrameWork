package main

import (
	"encoding/json"
	"io"
	"net/http"
	"test/simpleWebServer/defs"
)

func sendErrorResponse(w http.ResponseWriter,errResp defs.ErrResponse) {
	w.WriteHeader(errResp.HttpSC) // 设置状态码

	resStr,_ := json.Marshal(&errResp.Error)
	io.WriteString(w,string(resStr))
}

func sendNormalResponse(w http.ResponseWriter,resp string,sc int) {
	w.WriteHeader(sc)
	io.WriteString(w,resp)
}
