package timer

import "github.com/gonethopper/libs/concurrent"

//为每个定时器生成id
var timerIds *concurrent.AtomicInt64

func init() {
	//给定初始数值
	timerIds = concurrent.NewAtomicInt64(1)
}
