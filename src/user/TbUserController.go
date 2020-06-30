package user

import (
	"net/http"
	"encoding/json"
	"commons"
)

//所有user模块的handler
func UserHandler() {
	commons.Router.HandleFunc("/login", loginController)
}

//登录
func loginController(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	er := LoginService(username, password)
	//把结构体转换为json数据
	b, _ := json.Marshal(er)
	//设置响应内容为json
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}
