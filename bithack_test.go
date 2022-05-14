package bithack

import (
	"math/rand"
	"testing"
)

// Test for UintConv.

func TestUintConv(t *testing.T) {
	// Uint32 <=> Bytes.
	var u32x uint32 = 0x1122_3344
	u32y := BytesToUint32(Uint32ToBytes(u32x))
	if u32x != u32y {
		t.Fatal("invalid conversion")
	}
	u32y = LittleEndian.BytesToUint32(LittleEndian.Uint32ToBytes(u32x))
	if u32x != u32y {
		t.Fatal("invalid conversion")
	}
	u32y = BigEndian.BytesToUint32(BigEndian.Uint32ToBytes(u32x))
	if u32x != u32y {
		t.Fatal("invalid conversion")
	}
	// Uint64 <=> Bytes.
	var u64x uint64 = 0x1122_3344_5566_7788
	u64y := BytesToUint64(Uint64ToBytes(u64x))
	if u64x != u64y {
		t.Fatal("invalid conversion")
	}
	u64y = LittleEndian.BytesToUint64(LittleEndian.Uint64ToBytes(u64x))
	if u64x != u64y {
		t.Fatal("invalid conversion")
	}
	u64y = BigEndian.BytesToUint64(BigEndian.Uint64ToBytes(u64x))
	if u64x != u64y {
		t.Fatal("invalid conversion")
	}
}

// Test for HasZero.

func TestHasZero32(t *testing.T) {
	testHasZero32(t, [4]uint8{0, 0, 0, 0})
	testHasZero32(t, [4]uint8{255, 0, 255, 0})
	testHasZero32(t, [4]uint8{255, 255, 255, 255})
	for i := 0; i < 1000; i++ {
		var x [4]uint8
		for i := range x {
			if rand.Intn(3) == 0 {
				x[i] = 0
			} else {
				x[i] = uint8(rand.Int())
			}
		}
		testHasZero32(t, x)
	}
}

func testHasZero32(t *testing.T, check [4]uint8) {
	var result [4]uint8
	for i := range check {
		if check[i] == 0 {
			result[i] = 1
		}
	}

	got := Uint32ToBytes(HasZero32(BytesToUint32(check)))
	for i, v := range got {
		if result[i] == 1 && v == 0 {
			t.Fatalf("check: %v want: %v got: %v", check, result, got)
		}
		// False positive: result[i] is 0 but v is 1.
	}
}

func TestHasZero64(t *testing.T) {
	testHasZero64(t, [8]uint8{0, 0, 0, 0, 0, 0, 0, 0})
	testHasZero64(t, [8]uint8{255, 255, 255, 255, 255, 255, 255, 255})
	for i := 0; i < 1000; i++ {
		var x [8]uint8
		for i := range x {
			if rand.Intn(3) == 0 {
				x[i] = 0
			} else {
				x[i] = uint8(rand.Int())
			}
		}
		testHasZero64(t, x)
	}
}

func testHasZero64(t *testing.T, check [8]uint8) {
	var result [8]uint8
	for i := range check {
		if check[i] == 0 {
			result[i] = 1
		}
	}

	got := Uint64ToBytes(HasZero64(BytesToUint64(check)))
	for i, v := range got {
		if result[i] == 1 && v == 0 {
			t.Fatalf("check: %v want: %v got: %v", check, result, got)
		}
		// False positive: result[i] is 0 but v is 1.
	}
}

// Test for HasValue.

func TestHasValue32(t *testing.T) {
	var check [4]uint8
	for i := 0; i < 1000; i++ {
		for j := range check {
			check[j] = uint8(rand.Int())
		}
		testHasValue32(t, BytesToUint32(check), uint8(rand.Int()))
	}
}

func testHasValue32(t *testing.T, v uint32, x uint8) {
	var result [4]uint8
	for i, iv := range Uint32ToBytes(v) {
		if iv == x {
			result[i] = 1
		}
	}

	got := Uint32ToBytes(HasValue32(v, x))
	for i, iv := range got {
		if result[i] == 1 && iv == 0 {
			t.Fatalf("(%v,%v) want: %v got: %v", Uint32ToBytes(v), x, result, got)
		}
		// False positive: result[i] is 0 but v is 1.
	}
}

func TestHasValue64(t *testing.T) {
	var check [8]uint8
	for i := 0; i < 1000; i++ {
		for j := range check {
			check[j] = uint8(rand.Int())
		}
		testHasValue64(t, BytesToUint64(check), uint8(rand.Int()))
	}
}

func testHasValue64(t *testing.T, v uint64, x uint8) {
	var result [8]uint8
	for i, iv := range Uint64ToBytes(v) {
		if iv == x {
			result[i] = 1
		}
	}

	got := Uint64ToBytes(HasValue64(v, x))
	for i, iv := range got {
		if result[i] == 1 && iv == 0 {
			t.Fatalf("(%v,%v) want: %v got: %v", Uint64ToBytes(v), x, result, got)
		}
		// False positive: result[i] is 0 but v is 1.
	}
}
