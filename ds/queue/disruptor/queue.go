package disruptor

import (
	"errors"
)

//NewQueue 创建一个队列，单消费者和(1-n)生产者模式队列,通过isSingleProducer参数来判断是否使用单生产者模式
func NewQueue(queueID int64, capacity uint32, isSingleProducer bool, consumer Consumer) *Queue {
	q := new(Queue)
	q.Buffer = NewBuffer(capacity)
	q.IsSingleProducer = isSingleProducer
	//// Build() = single producer vs BuildShared() = multiple producers
	if isSingleProducer {
		q.controller = Configure(queueID, int64(q.Buffer.RingBufferCapacity)).WithConsumerGroup(consumer).Build()
	} else {
		q.sharedController = Configure(queueID, int64(q.Buffer.RingBufferCapacity)).WithConsumerGroup(consumer).BuildShared()
	}
	q.Start()
	return q
}

//Queue 队列定义，对go-disruptor的封装
type Queue struct {
	Buffer           *Buffer
	IsSingleProducer bool //是否使用单生产者模式
	controller       Disruptor
	sharedController SharedDisruptor
}

//Start start queue
func (q *Queue) Start() {
	if q.IsSingleProducer {
		q.controller.Start()

	} else {
		q.sharedController.Start()

	}
}

//Push 非阻塞入队列生产者产生数据入队列，成功返回nil，失败返回错误描述
func (q *Queue) Push(v interface{}) error {
	if q.IsSingleProducer {
		writer := q.controller.Writer()
		// for each item received from a network socket, e.g. UDP packets, HTTP request, etc. etc.
		sequence := writer.TryReserve(1) // reserve 1 slot on the ring buffer and give me the upper-most sequence of the reservation
		if sequence < 0 {
			return errors.New("queue is full")
		}
		// this could be written like this: ringBuffer[sequence%RingBufferCapacity] but the Mask and & operator is faster.
		q.Buffer.RingBuffer[sequence&int64(q.Buffer.RingBufferMask)] = v // data from network stream

		writer.Commit(sequence, sequence) // the item is ready to be consumed

	} else {
		writer := q.sharedController.Writer()
		// for each item received from a network socket, e.g. UDP packets, HTTP request, etc. etc.
		sequence := writer.TryReserve(1) // reserve 1 slot on the ring buffer and give me the upper-most sequence of the reservation
		if sequence < 0 {
			return errors.New("queue is full")
		}
		// this could be written like this: ringBuffer[sequence%RingBufferCapacity] but the Mask and & operator is faster.
		q.Buffer.RingBuffer[sequence&int64(q.Buffer.RingBufferMask)] = v // data from network stream

		writer.Commit(sequence, sequence) // the item is ready to be consumed
	}

	return nil
}

//BlockPush 阻塞入队列生产者产生数据入队列，成功返回nil，失败返回错误描述
func (q *Queue) BlockPush(v interface{}) error {

	if q.IsSingleProducer {
		writer := q.controller.Writer()
		// for each item received from a network socket, e.g. UDP packets, HTTP request, etc. etc.
		sequence := writer.Reserve(1) // reserve 1 slot on the ring buffer and give me the upper-most sequence of the reservation
		// this could be written like this: ringBuffer[sequence%RingBufferCapacity] but the Mask and & operator is faster.

		q.Buffer.RingBuffer[sequence&int64(q.Buffer.RingBufferMask)] = v // data from network stream

		writer.Commit(sequence, sequence) // the item is ready to be consumed

	} else {
		writer := q.sharedController.Writer()
		// for each item received from a network socket, e.g. UDP packets, HTTP request, etc. etc.
		sequence := writer.Reserve(1) // reserve 1 slot on the ring buffer and give me the upper-most sequence of the reservation
		// this could be written like this: ringBuffer[sequence%RingBufferCapacity] but the Mask and & operator is faster.
		q.Buffer.RingBuffer[sequence&int64(q.Buffer.RingBufferMask)] = v // data from network stream

		writer.Commit(sequence, sequence) // the item is ready to be consumed
	}

	return nil
}

//Pop 弹出数据
func (q *Queue) Pop(sequence int64) interface{} {
	return q.Buffer.RingBuffer[sequence&int64(q.Buffer.RingBufferMask)]
}

//Close clean shutdown which stops all idling consumers after all published items have been consumed
func (q *Queue) Close() {
	if q.IsSingleProducer {
		q.controller.Stop()
	} else {
		q.sharedController.Stop()
	}
}
