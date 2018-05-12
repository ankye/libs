package timer

//timer wheel https://github.com/lileeei/timer
import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

const (
	TIME_NEAR_SHIFT  = 8
	TIME_NEAR        = 1 << TIME_NEAR_SHIFT
	TIME_LEVEL_SHIFT = 6
	TIME_LEVEL       = 1 << TIME_LEVEL_SHIFT
	TIME_NEAR_MASK   = TIME_NEAR - 1
	TIME_LEVEL_MASK  = TIME_LEVEL - 1
)

type TimerWheel struct {
	near [TIME_NEAR]*list.List
	t    [4][TIME_LEVEL]*list.List
	sync.Mutex
	time uint32        //时间轮time
	tick time.Duration //时间轮的tick
	quit chan struct{} //时间轮退出信号
}

type Timer struct {
	//TimerId uint32  //定时器id，便于以后查找
	expire       uint32 //任务过期时间
	callBackFunc func() //回调函数
}

func (n *Timer) String() string {
	return fmt.Sprintf("Timer:expire,%d", n.expire)
}

//创建时间轮
func NewTimerWheel(d time.Duration) *TimerWheel {
	t := new(TimerWheel)
	t.time = 0
	t.tick = d
	t.quit = make(chan struct{})

	var i, j int
	for i = 0; i < TIME_NEAR; i++ {
		t.near[i] = list.New()
	}

	for i = 0; i < 4; i++ {
		for j = 0; j < TIME_LEVEL; j++ {
			t.t[i][j] = list.New()
		}
	}

	return t
}

func (tw *TimerWheel) AddTimer(d time.Duration, f func()) *Timer {
	nd := new(Timer)
	nd.callBackFunc = f

	tw.Lock()
	nd.expire = uint32(d/tw.tick) + tw.time
	tw.addTimer(nd)
	tw.Unlock()

	return nd
}

//向时间轮中添加任务结点
func (tw *TimerWheel) addTimer(nd *Timer) {
	expire := nd.expire
	current := tw.time

	//将定时节点添加到near中
	if (expire | TIME_NEAR_MASK) == (current | TIME_NEAR_MASK) {
		tw.near[expire&TIME_NEAR_MASK].PushBack(nd)
	} else { //将定时节点加入到另外的4个轮中（根据duration的大小加入到不同的层级中）
		var i uint32
		var mask uint32 = TIME_NEAR << TIME_LEVEL_SHIFT
		for i = 0; i < 3; i++ {
			if (expire | (mask - 1)) == (current | (mask - 1)) {
				break
			}
			mask <<= TIME_LEVEL_SHIFT
		}

		tw.t[i][(expire>>(TIME_NEAR_SHIFT+i*TIME_LEVEL_SHIFT))&TIME_LEVEL_MASK].PushBack(nd)
	}

}

//开启时间轮
func (tw *TimerWheel) Start() {
	tick := time.NewTicker(tw.tick)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			tw.update()
		case <-tw.quit:
			return
		}
	}
}

//更新时间轮
func (tw *TimerWheel) update() {
	tw.execute()
	tw.shift()
	tw.execute()
}

//调整时间轮上的定时节点
func (tw *TimerWheel) shift() {
	tw.Lock()
	var mask uint32 = TIME_NEAR
	tw.time++
	ct := tw.time

	time := ct >> TIME_NEAR_SHIFT
	var i int = 0

	//每执行TIME_NEAR次循环就重建一次层级， 每循环TIME_LEVEL次重建一次下一个level
	for (ct & (mask - 1)) == 0 {
		idx := int(time & TIME_LEVEL_MASK)

		if idx != 0 {
			tw.moveList(i, idx)

			break
		}

		mask <<= TIME_LEVEL_SHIFT
		time >>= TIME_LEVEL_SHIFT

		i++
	}

	tw.Unlock()
}

//将level层上hash值为idx的任务列表进行调整
func (tw *TimerWheel) moveList(level, idx int) {
	l := tw.t[level][idx]
	front := l.Front()
	l.Init() //将该list清空

	for e := front; e != nil; e = e.Next() {
		nd := e.Value.(*Timer)
		//将定时节点重新加入到时间轮中
		tw.addTimer(nd)
	}
}

//执行timeout的任务，并把它们从时间轮中删除，每次都从near中删除
func (tw *TimerWheel) execute() {
	tw.Lock()
	idx := tw.time & TIME_NEAR_MASK
	l := tw.near[idx]
	if l.Len() > 0 {
		front := l.Front()
		l.Init()
		tw.Unlock()

		dispatchList(front)
		return
	}

	tw.Unlock()
}

//将timeout的任务从时间轮中删除
func dispatchList(front *list.Element) {
	for e := front; e != nil; e = e.Next() {
		timer := e.Value.(*Timer)
		go timer.callBackFunc()
	}
}

//Stop 关闭时间轮
func (tw *TimerWheel) Stop() {
	close(tw.quit)
}

//便于将时间轮的信息以字符串的形式输出
func (tw *TimerWheel) String() string {
	return fmt.Sprintf("Timer:time:%d, tick:%s", tw.time, tw.tick)
}
