package timer

import (
	"sync"
	"time"
)

var (
	onceTM     sync.Once
	instanceTM *TimerManager
)

//GetTimerManagerInstance timer manager singleton
func GetTimerManagerInstance() *TimerManager {
	onceTM.Do(func() {
		instanceTM = NewTimerManager()
	})
	return instanceTM
}

//NewTimerManager create TimerManager
func NewTimerManager() *TimerManager {
	s := new(TimerManager)
	s.timerwheel = NewTimerWheel(time.Millisecond * 10) //定时器时钟设置
	return s
}

//TimerManager 定时器管理类
type TimerManager struct {
	timerwheel *TimerWheel //时间轮
}

//AddTimer 添加定时器
func (tm *TimerManager) AddTimer(d time.Duration, f func()) *Timer {
	return tm.timerwheel.AddTimer(d, f)
}
