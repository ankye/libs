package disruptor

import "testing"

const QueueID = 1111

func BenchmarkWriterReserve(b *testing.B) {
	read, written := NewCursor(), NewCursor()
	writer := NewWriter(QueueID, written, read, 1024)
	iterations := int64(b.N)
	b.ReportAllocs()
	b.ResetTimer()

	for i := int64(0); i < iterations; i++ {
		sequence := writer.Reserve(1)
		read.Store(sequence)
	}
}

func BenchmarkWriterNextWrapPoint(b *testing.B) {
	read, written := NewCursor(), NewCursor()
	writer := NewWriter(QueueID, written, read, 1024)
	iterations := int64(b.N)
	b.ReportAllocs()
	b.ResetTimer()

	read.Store(MaxSequenceValue)
	for i := int64(0); i < iterations; i++ {
		writer.Reserve(1)
	}
}

func BenchmarkWriterAwait(b *testing.B) {
	written, read := NewCursor(), NewCursor()
	writer := NewWriter(QueueID, written, read, 1024*64)
	iterations := int64(b.N)
	b.ReportAllocs()
	b.ResetTimer()

	for i := int64(0); i < iterations; i++ {
		writer.Await(i)
		read.Store(i)
	}
}

func BenchmarkWriterCommit(b *testing.B) {
	writer := NewWriter(QueueID, NewCursor(), nil, 1024)
	iterations := int64(b.N)
	b.ReportAllocs()
	b.ResetTimer()

	for i := int64(0); i < iterations; i++ {
		writer.Commit(i, i)
	}
}
