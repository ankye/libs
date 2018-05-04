package common

//MESSAGE ID define
//SSMessage ID 划分区域 0x5001 -- 0xFFFF
const (
	//保留系统ID 0x5001 - 050FF
	//服务器心跳
	SSMessageIDHeartbeat = 0x5001

	//业务ID编号 0x5100 - 0xFFFF
	SSMessageIDGetUserinfo = 0x5100
)
