//go:build ppc64 || s390x || mips || mips64
// +build ppc64 s390x mips mips64

package bithack

var BytesToUint32 = LittleEndian.BytesToUint32

var Uint32ToBytes = LittleEndian.Uint32ToBytes

var BytesToUint64 = LittleEndian.BytesToUint64

var Uint64ToBytes = LittleEndian.Uint64ToBytes
