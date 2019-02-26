package util

import (
	"encoding/binary"
	"math"
)

// BytesToUint16 converts a little endian array of bytes to an array of unit16s
func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}
