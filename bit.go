package bithack

// LittleEndian is the little-endian bithack implementation.
var LittleEndian littleEndian

type littleEndian struct{}

func (c littleEndian) BytesToUint32(v [4]uint8) uint32 {
	return uint32(v[0]) | uint32(v[1])<<8 | uint32(v[2])<<16 | uint32(v[3])<<24
}

func (c littleEndian) BytesToUint64(v [8]uint8) uint64 {
	return uint64(v[0]) | uint64(v[1])<<8 | uint64(v[2])<<16 | uint64(v[3])<<24 | uint64(v[4])<<32 | uint64(v[5])<<40 | uint64(v[6])<<48 | uint64(v[7])<<56
}

func (c littleEndian) Uint32ToBytes(v uint32) [4]uint8 {
	return [4]uint8{uint8(v), uint8(v >> 8), uint8(v >> 16), uint8(v >> 24)}
}

func (c littleEndian) Uint64ToBytes(v uint64) [8]uint8 {
	return [8]uint8{uint8(v), uint8(v >> 8), uint8(v >> 16), uint8(v >> 24), uint8(v >> 32), uint8(v >> 40), uint8(v >> 48), uint8(v >> 56)}
}

// BigEndian is the big-endian bithack implementation.
var BigEndian bigEndian

type bigEndian struct{}

func (c bigEndian) BytesToUint32(v [4]uint8) uint32 {
	return uint32(v[3]) | uint32(v[2])<<8 | uint32(v[1])<<16 | uint32(v[0])<<24
}

func (c bigEndian) BytesToUint64(v [8]uint8) uint64 {
	return uint64(v[7]) | uint64(v[6])<<8 | uint64(v[5])<<16 | uint64(v[4])<<24 | uint64(v[3])<<32 | uint64(v[2])<<40 | uint64(v[1])<<48 | uint64(v[0])<<56
}

func (c bigEndian) Uint32ToBytes(v uint32) [4]uint8 {
	return [4]uint8{uint8(v >> 24), uint8(v >> 16), uint8(v >> 8), uint8(v)}
}

func (c bigEndian) Uint64ToBytes(v uint64) [8]uint8 {
	return [8]uint8{uint8(v >> 56), uint8(v >> 48), uint8(v >> 40), uint8(v >> 32), uint8(v >> 24), uint8(v >> 16), uint8(v >> 8), uint8(v)}
}
