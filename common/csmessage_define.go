package common

//MESSAGE ID define
//CSMessage ID 划分区域 0x0001 -- 0x4FFF
const (
	//保留系统ID 0x0001 - 000FF
	//客户端心跳
	CSMessageIDHeartbeat = 0x0001

	//业务ID编号 0x0100 - 0x4FFF
	//获取用户信息
	CSMessageIDGetUserinfo = 0x0100
)
