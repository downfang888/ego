package commons

//客户端服务器端数据交互模版
type EgoResult struct {
	Status int//状态,为200表示成功
	Data interface{}//返回的数据
	Msg string//返回的消息
}