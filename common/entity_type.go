package common

//EntityType 服务器类型定义
const (
	//EntityTypeAgent value -> 1
	EntityTypeAgent = 0x01
	//EntityTypeRouter  value ->3
	EntityTypeRouter = 0x02
	//EntityTypeLogic value -> 2
	EntityTypeLogic = 0x03
	//EntityTypeGamedb value ->4
	EntityTypeGamedb = 0x04
	//EntityTypeUsercenter = 0x05
	EntityTypeUsercenter = 0x05
	//EntityTypeUserdb = 0x06
	EntityTypeUserdb = 0x06

	//MaxEntityTypeNumber 服务器类型最大数量定义
	MaxEntityTypeNumber = 0x7f
)
