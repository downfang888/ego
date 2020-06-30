package item

import (
	"net/http"
	"strconv"
	"encoding/json"
	"commons"
)

func ItemHandler() {
	commons.Router.HandleFunc("/showItem", showItemController)
	commons.Router.HandleFunc("/item/delete", delByIdsController)
	commons.Router.HandleFunc("/item/instock", instockController)
	commons.Router.HandleFunc("/item/offstock", offstockController)
	commons.Router.HandleFunc("/item/imageupload", imagesUploadController) //图片上传
	commons.Router.HandleFunc("/item/add", insertControllew) //商品新增
	commons.Router.HandleFunc("/item/showItemById", showItemDescCatController) //商品修改页面信息显示
	commons.Router.HandleFunc("/item/update", updateController) //商品修改页面信息显示
}
//修改
func updateController(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	er:=updateService(r.Form)
	b,_:=json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE,commons.JSON_HEADER)
	w.Write(b)
}

//显示修改页面信息
func showItemDescCatController(w http.ResponseWriter, r *http.Request){
	id,_:= strconv.Atoi(r.FormValue("id"))
	c:=showItemDescCatService(id)
	b,_:=json.Marshal(c)
	w.Header().Set(commons.HEADER_CONTENT_TYPE,commons.JSON_HEADER)
	w.Write(b)
}
//商品新增
func insertControllew (w http.ResponseWriter, r *http.Request) {
	//需要先进行解析
	r.ParseForm()
	er:=insetService(r.Form)
	b,_:=json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE,commons.JSON_HEADER)
	w.Write(b)
}

//图片上传
func imagesUploadController(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("imgFile")
	if err != nil {
		m := make(map[string]interface{})
		m["error"] = 1
		m["message"] = "接收图片失败"
		b, _ := json.Marshal(m)
		w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
		w.Write(b)
		return
	}
	m := imageUploadService(file, fileHeader)
	b, _ := json.Marshal(m)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}

//显示商品信息
func showItemController(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	rows, _ := strconv.Atoi(r.FormValue("rows"))
	datagrid := showItemService(page, rows)
	b, _ := json.Marshal(datagrid)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}

//商品删除
func delByIdsController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	er := delByIdsService(ids)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}

//商品上架
func instockController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	er := instockService(ids)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}

//商品下架
func offstockController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	er := offStockService(ids)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}
