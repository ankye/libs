package db

//RetType Define query return result type
type RetType int

const (
	//Result 单值
	Result RetType = iota
	//Row 单条记录
	Row
	//Rows 记录集
	Rows
)

//Processer db Processer
type Processer struct {
	Command string  //命令字，处理进程自己设置
	SQL     string  //执行语句
	RetType RetType //return type
	Data    interface{}
}
