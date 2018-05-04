package common

//EntryType 服务器类型定义
const (
	//EntryTypeAgent value -> 1
	EntryTypeAgent = 0x01

	//EntryTypeRouter  value ->3
	EntryTypeRouter = 0x02
	//EntryTypeLogic value -> 2
	EntryTypeLogic = 0x03
	//EntryTypeGamedb value ->4
	EntryTypeGamedb = 0x04
	//MaxEntryTypeNumber 服务器类型最大数量定义
	MaxEntryTypeNumber = 0x7f
)
