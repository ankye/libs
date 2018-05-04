package queue

//NewBuffer new queue buffer
func NewBuffer(capacity uint32) *Buffer {
	buf := new(Buffer)
	buf.RingBufferCapacity = MinQuantity(capacity)
	buf.RingBufferMask = buf.RingBufferCapacity - 1
	buf.RingBuffer = make([]interface{}, buf.RingBufferCapacity)

	return buf
}

//Buffer 队列buffer
type Buffer struct {
	RingBufferCapacity uint32 // must be a power of 2
	RingBufferMask     uint32 // ringBufferCapacity - 1
	RingBuffer         []interface{}
}

//MinQuantity round 到最近的2的倍数
func MinQuantity(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}
