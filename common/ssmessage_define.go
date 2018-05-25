package common

//MESSAGE ID define
//SSMessage ID 划分区域 0x5001 -- 0xFFFF
const (
	//保留系统ID 0x5001 - 050FF
	//服务器心跳
	SSMessageIDHeartbeat = 0x5001
	SSMessageIDTimeout   = 0x5002

	//业务ID编号 0x5100 - 0xFFFF
	//usercenter 0x5100 - 0x51FF
	SSMessageIDUserLogin    = 0x5100
	SSMessageIDUserRegister = 0x5101
	SSMessageIDUserGetInfo  = 0x5102
)
