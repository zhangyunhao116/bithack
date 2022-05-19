package bithack

const (
	// 32 on 32-bit systems, 64 on 64-bit
	usize = 32 << (^uintptr(0) >> 63)

	// ptrSize is the size of a pointer in bytes - unsafe.Sizeof(uintptr(0)) but as an ideal constant.
	// It is also the size of the machine's native word size (that is, 4 on 32-bit systems, 8 on 64-bit).
	ptrSize = 4 << (^uintptr(0) >> 63)
)

func HasZero32(v uint32) uint32 {
	return ((v) - uint32(0x0101_0101)) & ^(v) & uint32(0x8080_8080)
}

func HasZero64(v uint64) uint64 {
	return (((v) - uint64(0x0101_0101_0101_0101)) & ^(v) & uint64(0x8080_8080_8080_8080))
}

func HasValue32(v uint32, x uint8) uint32 {
	return HasZero32(v ^ (uint32(0x0101_0101) * uint32(x)))
}

// HasValue64 determine if a uint64 has a byte(uint8) equal to x.
// It returns a uint64 indicating all bytes in v *may* have the given value.
//
// This function may return a false positive, the probability is about 1 in 10,000.
func HasValue64(v uint64, x uint8) uint64 {
	return HasZero64(v ^ (uint64(0x0101_0101_0101_0101) * uint64(x)))
}
