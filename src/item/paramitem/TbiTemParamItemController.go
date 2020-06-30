package paramitem

import (
	"commons"
	"net/http"
	"strconv"
	"encoding/json"
)

func ParamItemHandler(){
	commons.Router.HandleFunc("/item/paramitem/selid",selidController)
}
//显示商品的规格参数
func selidController(w http.ResponseWriter,r *http.Request){
	id,_:=strconv.Atoi(r.FormValue("id"))
	er:=showItemByIdService(id)
	b,_:=json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE,commons.JSON_HEADER)
	w.Write(b)
}
