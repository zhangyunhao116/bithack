//go:build 386 || amd64 || arm || arm64 || ppc64le || mips64le || mipsle || riscv64 || wasm
// +build 386 amd64 arm arm64 ppc64le mips64le mipsle riscv64 wasm

package bithack

var BytesToUint32 = BigEndian.BytesToUint32

var Uint32ToBytes = BigEndian.Uint32ToBytes

var BytesToUint64 = BigEndian.BytesToUint64

var Uint64ToBytes = BigEndian.Uint64ToBytes
