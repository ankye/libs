package timer

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

var sum int32 = 0
var N int32 = 30
var tw *TimerWheel

func callBack() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	atomic.AddInt32(&sum, 1)
	v := atomic.LoadInt32(&sum)
	if v == 2*N {
		tw.Stop()
	}

}

func TestTimer(t *testing.T) {
	timerwheel := NewTimerWheel(time.Millisecond * 10)
	tw = timerwheel
	fmt.Println(timerwheel)

	go func() {
		var i int32
		for i = 0; i < 2*N; i++ {
			timerwheel.AddTimer(time.Millisecond*time.Duration(10*i), callBack)
		}
	}()

	timerwheel.Start()
	if sum != 2*N {
		t.Error("failed")
	}
}
