package bithack

const (
	// 32 on 32-bit systems, 64 on 64-bit
	usize = 32 << (^uintptr(0) >> 63)

	// PtrSize is the size of a pointer in bytes - unsafe.Sizeof(uintptr(0)) but as an ideal constant.
	// It is also the size of the machine's native word size (that is, 4 on 32-bit systems, 8 on 64-bit).
	PtrSize = 4 << (^uintptr(0) >> 63)
)

func HasZero(v uint) uint {
	if usize == 32 {
		return uint(HasZero32(uint32(v)))
	}
	return uint(HasZero64(uint64(v)))
}

func HasZero32(v uint32) uint32 {
	return ((v) - uint32(0x0101_0101)) & ^(v) & uint32(0x8080_8080)
}

func HasZero64(v uint64) uint64 {
	return (((v) - uint64(0x0101_0101_0101_0101)) & ^(v) & uint64(0x8080_8080_8080_8080))
}

func HasValue(v uint, x uint8) uint {
	if usize == 32 {
		return uint(HasValue32(uint32(v), x))
	}
	return uint(HasValue64(uint64(v), x))
}

func HasValue32(v uint32, x uint8) uint32 {
	return HasZero32(v ^ (uint32(0x0101_0101) * uint32(x)))
}

func HasValue64(v uint64, x uint8) uint64 {
	return HasZero64(v ^ (uint64(0x0101_0101_0101_0101) * uint64(x)))
}
