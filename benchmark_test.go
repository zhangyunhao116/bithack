package bithack

import (
	"math/rand"
	"testing"
)

func BenchmarkHasValue(b *testing.B) {
	var x [8]byte
	rand.Read(x[:])
	y := uint8(rand.Int())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HasValue64(BytesToUint64(x), y)
		b.StopTimer()
		rand.Read(x[:])
		y = uint8(rand.Int())
		b.StartTimer()
	}
}

func BenchmarkHasValueLoop(b *testing.B) {
	var x [8]byte
	rand.Read(x[:])
	y := uint8(rand.Int())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopMatch(x, y)
		b.StopTimer()
		rand.Read(x[:])
		y = uint8(rand.Int())
		b.StartTimer()
	}
}

func LoopMatch(x [8]byte, value uint8) int {
	for _, v := range x {
		if value == v {
			return 1
		}
	}
	return 0
}
