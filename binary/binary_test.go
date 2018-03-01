package binary

import (
	"encoding/binary"
	"testing"
)

func BenchmarkPutVarint(b *testing.B) {
	for i := int64(0); i < int64(b.N); i++ {
		key := make([]byte, 8)
		binary.PutVarint(key, i)
	}
}

func BenchmarkPutUint64BigEndian(b *testing.B) {
	for i := int64(0); i < int64(b.N); i++ {
		key := make([]byte, 8)
		binary.BigEndian.PutUint64(key, uint64(i))
	}
}

func BenchmarkPutUint64LittleEndian(b *testing.B) {
	for i := int64(0); i < int64(b.N); i++ {
		key := make([]byte, 8)
		binary.LittleEndian.PutUint64(key, uint64(i))
	}
}
