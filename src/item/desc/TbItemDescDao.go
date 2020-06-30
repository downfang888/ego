package desc

import (
	"commons"
	"fmt"
)
//新增描述
func insertDescDao(t TbItemDesc ) int{
	count,err:=commons.Dml("insert into tb_item_desc values(?,?,?,?)",t.ItemId,t.ItemDesc,t.Created,t.Updated)
	if err!=nil{
		fmt.Println(err)
		return -1
	}
	return int(count)
}

//根据主键查询
func selByIdDao(id int) *TbItemDesc{
	r,err:=commons.Dql("select * from tb_item_desc where item_id=?",id)
	if err!=nil{
		fmt.Println(err)
		return nil
	}
	if r.Next(){
		t := new(TbItemDesc)
		r.Scan(&t.ItemId,&t.ItemDesc,&t.Created,&t.Updated)
		return t
	}
	return nil
}

//根据主键修改商品描述,带有事务
func UpdDescByIdWithTxDao(t TbItemDesc)  int{
	return commons.PrepareWithTx("update tb_item_desc set item_desc=?,updated=? where item_id=?",t.ItemDesc,t.Updated,t.ItemId)
}

//根据主键删除
func DelDescByIdDao(id int) int{
	count,err:=commons.Dml("delete from tb_item_desc where item_id = ?",id)
	if err!=nil{
		fmt.Println(err)
		return -1
	}
	return int(count)
}