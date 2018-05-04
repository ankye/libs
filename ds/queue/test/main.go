package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gonethopper/libs/ds/queue"
	"github.com/gonethopper/libs/signal"
)

const (
	MaxRunTestNum = 10000000
)

//TConsumer 消费者定义
type TConsumer struct {
	Queue *queue.Queue
}

//Consume 消费
func (t *TConsumer) Consume(queueID, lower, upper int64) {
	if t.Queue == nil {
		return
	}
	for sequence := lower; sequence <= upper; sequence++ {
		message := t.Queue.Pop(sequence) // see performance note on producer sample above
		//fmt.Println("get message:", message)
		if message != nil {

		}
		//	time.Sleep(1 * time.Second)
		// handle the incoming message with your application code
	}
}
func main() {
	runtime.GOMAXPROCS(4)

	disruptorBlock()
	es()
	// gsp()
	disruptor()
	signal.InitSignal()
}
func disruptor() {

	consumer := &TConsumer{}
	q := queue.NewQueue(1, 1024*64, false, consumer)
	consumer.Queue = q

	for index := 0; index < 3; index++ {

		go func(q *queue.Queue, index int) {
			i := 0
			t := time.Now().UnixNano()
			//for t := range time.Tick(1 * time.Second) {
			for a := 0; a <= MaxRunTestNum; a++ {
				i++
				if err := q.Push(fmt.Sprintf("go%d %d-[ %d ]", index, a, 1000*index+i)); err != nil {
					//	fmt.Println("send failed", a, i)
					continue
				}
				//fmt.Println("send success", t.String(), 1000*index+i)

			}
			t = (time.Now().UnixNano() - t) / 1000000 //ms
			fmt.Printf("disruptor done opsPerSecond: %d index:%d\n", t, index)

		}(q, index)

	}

	fmt.Println("hello world")

}

func disruptorBlock() {

	consumer := &TConsumer{}
	q := queue.NewQueue(1, 1024*64, false, consumer)
	consumer.Queue = q

	for index := 0; index < 3; index++ {

		go func(q *queue.Queue, index int) {
			i := 0
			t := time.Now().UnixNano()
			//for t := range time.Tick(1 * time.Second) {
			for a := 0; a <= MaxRunTestNum; a++ {
				i++
				if err := q.BlockPush(fmt.Sprintf("go%d %d-[ %d ]", index, a, 1000*index+i)); err != nil {
					//	fmt.Println("send failed", a, i)
					continue
				}
				//fmt.Println("send success", t.String(), 1000*index+i)

			}
			t = (time.Now().UnixNano() - t) / 1000000 //ms
			fmt.Printf("disruptor block done opsPerSecond: %d index:%d\n", t, index)
		}(q, index)

	}

}

func es() {

	q := queue.NewESQueue(1024 * 64)

	for index := 0; index < 3; index++ {

		go func(q *queue.EsQueue, index int) {
			i := 0
			t := time.Now().UnixNano()
			//for t := range time.Tick(1 * time.Second) {
			for a := 0; a <= MaxRunTestNum; a++ {
				i++
				if ok, _ := q.Put(fmt.Sprintf("go%d %d-[ %d ]", index, a, 1000*index+i)); !ok {
					//	fmt.Println("send failed", a, i)
					continue
				}
				//fmt.Println("send success", t.String(), 1000*index+i)

			}
			t = (time.Now().UnixNano() - t) / 1000000 //ms
			fmt.Printf("es done opsPerSecond: %d index:%d\n", t, index)
		}(q, index)

	}
	go func(q *queue.EsQueue) {
		for {
			q.Get()
		}
	}(q)

}
