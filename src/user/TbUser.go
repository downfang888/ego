package user

//对应数据库中用户表
type TbUser struct {
	//属性首字母大写:1. 要转换为json   2. 可能出现跨包访问
	Id int64
	Username string
	Password string
	Phone string
	Email string
	Created string
	Updated string
}
