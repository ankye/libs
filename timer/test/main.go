package main

import (
	"fmt"
	"time"

	"github.com/gonethopper/libs/signal"
	"github.com/gonethopper/libs/timer"
)

var (
	wheel *timer.TimerWheel
)

func main() {

	wheel = timer.NewTimerWheel()
	task1 := timer.NewTimerTaskTimeOut("测试任务1", func(val interface{}) {
		fmt.Printf("测试任务1111 callback: %v;timeis:%s \n", val, time.Now().Format("2006-01-02 15:04:05"))
	})

	task2 := timer.NewTimerTaskTimeOut("测试任务2", func(val interface{}) {
		fmt.Printf("测试任务2222 callback: %v;timeis:%s \n", val, time.Now().Format("2006-01-02 15:04:05"))
	})

	timerID1 := wheel.AddTask(time.Duration(1)*time.Second, -1, task1)
	timerID2 := wheel.AddTask(time.Duration(5)*time.Second, -1, task2)

	go delTimer(timerID1)
	go delTimer(timerID2)
	signal.InitSignal()
}

func delTimer(timerID int64) {

	//先过个5秒再删除
	time.Sleep(time.Duration(timerID) * time.Second)
	fmt.Printf("删除任务：%d\n", timerID)
	wheel.CancelTimer(timerID)
	for i := 0; i < 5; i++ {

		go addNewTask()
	}
}

func addNewTask() {
	task3 := timer.NewTimerTaskTimeOut("测试任务不知道多少了", func(val interface{}) {
		fmt.Printf("测试任务tttt callback: %v;timeis:%s \n", val, time.Now().Format("2006-01-02 15:04:05"))
	})
	timerID := wheel.AddTask(time.Duration(3)*time.Second, -1, task3)

	go delTimer(timerID)
}
