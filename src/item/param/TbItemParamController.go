package param

import (
	"net/http"
	"strconv"
	"encoding/json"
	"commons"
)

func ParamHandler(){
	commons.Router.HandleFunc("/item/param/show",showParamController)
	commons.Router.HandleFunc("/item/param/delete",delByIdsController)//删除规格参数
	commons.Router.HandleFunc("/item/param/iscat",isCatController)//删除规格参数
	commons.Router.HandleFunc("/item/param/add",insertParamController)//规格参数新增
	commons.Router.HandleFunc("/item/param/edit",updateParamController)//规格参数修改

}
//编辑规格参数
func updateParamController(w http.ResponseWriter,r *http.Request){
	id,_:=strconv.Atoi(r.FormValue("id"))
	itemCatId,_:=strconv.Atoi(r.FormValue("itemCatId"))
	paramData:=r.FormValue("paramData")
	er:=updateParamService(id,itemCatId,paramData)
	b,_:=json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE,commons.JSON_HEADER)
	w.Write(b)
}
//规格参数新增
func insertParamController(w http.ResponseWriter,r *http.Request){

	catid,_:=strconv.Atoi(r.FormValue("itemCatId"))
	paramData:=r.FormValue("paramData")
	er:=insertParamService(catid,paramData)
	b,_:=json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE,commons.JSON_HEADER)
	w.Write(b)
}
//判断类目是否已经添加了规格参数
func isCatController(w http.ResponseWriter,r *http.Request){
	catid,_:=strconv.Atoi(r.FormValue("catid"))
	er:=catidService(catid)
	b,_:=json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE,commons.JSON_HEADER)
	w.Write(b)
}

//删除规格参数
func delByIdsController(w http.ResponseWriter,r *http.Request){
	er:=delByIdsService(r.FormValue("ids"))
	b,_:=json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE,commons.JSON_HEADER)
	w.Write(b)
}

//显示规格参数
func showParamController(w http.ResponseWriter,r *http.Request){
	page,_:=strconv.Atoi(r.FormValue("page"))
	rows,_:=strconv.Atoi(r.FormValue("rows"))
	datagrid:=showParamService(page,rows)
	b,_:=json.Marshal(datagrid)
	w.Header().Set(commons.HEADER_CONTENT_TYPE,commons.JSON_HEADER)
	w.Write(b)
}
