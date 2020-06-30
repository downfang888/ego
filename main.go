package main

import (
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"commons"
	"item"
	"item/cat"
	"user"
	"item/param"
	"item/paramitem"
)

//显示登录页面
func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/login.html")
	t.Execute(w, nil)
}

//restful显示页面
func showPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t, _ := template.ParseFiles("view/" + vars["page"] + ".html")
	t.Execute(w, nil)
}
func main() {
	commons.Router.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	commons.Router.HandleFunc("/", welcome)
	//满足/page/{page}格式的处理
	commons.Router.HandleFunc("/page/{page}", showPage)
	//用户
	user.UserHandler()
	//商品
	item.ItemHandler()
	//商品类目
	cat.ItemCatHandler()
	//规格参数
	param.ParamHandler()
	//商品规格参数
	paramitem.ParamItemHandler()
	http.ListenAndServe(":8090", commons.Router)

}
