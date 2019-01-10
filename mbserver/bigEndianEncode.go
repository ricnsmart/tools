package mbserver

import (
	"encoding/binary"
	"errors"
	"math"
)

// BytesToUint16 converts a big endian array of bytes to an array of unit16s
func BytesToUint16(bytes []byte) []uint16 {
	values := make([]uint16, len(bytes)/2)

	for i := range values {
		values[i] = binary.BigEndian.Uint16(bytes[i*2 : (i+1)*2])
	}
	return values
}

// Uint16ToBytes converts an array of uint16s to a big endian array of bytes
func Uint16ToBytes(values []uint16) []byte {
	bytes := make([]byte, len(values)*2)

	for i, value := range values {
		binary.BigEndian.PutUint16(bytes[i*2:(i+1)*2], value)
	}
	return bytes
}

// BytesToUint32 converts a big endian array of bytes to an array of unit32s
func BytesToUint32(bytes []byte) []uint32 {
	values := make([]uint32, len(bytes)/4)

	for i := range values {
		values[i] = binary.BigEndian.Uint32(bytes[i*4 : (i+1)*4])
	}
	return values
}

// Uint32ToBytes converts an array of uint32s to a big endian array of bytes
func Uint32ToBytes(values []uint32) []byte {
	bytes := make([]byte, len(values)*4)

	for i, value := range values {
		binary.BigEndian.PutUint32(bytes[i*4:(i+1)*4], value)
	}
	return bytes
}

// BytesToFloat32 converts a big endian array of bytes to an array of float32
func BytesToFloat32(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}

// Float32ToBytes converts an array of float32 to a big endian array of bytes
func Float32ToBytes(value float32) []byte {
	bits := math.Float32bits(value)

	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, bits)
	return bytes
}

func EncodeUint16(bytes *[]byte, value uint16) {
	bArr := make([]byte, 2)
	binary.BigEndian.PutUint16(bArr[0:2], value)
	*bytes = append(*bytes, bArr...)
}

func EncodeUint32(bytes *[]byte, value uint32) {
	bArr := make([]byte, 4)
	binary.BigEndian.PutUint32(bArr[0:4], value)
	*bytes = append(*bytes, bArr...)
}

func EncodeFloat32(bytes *[]byte, value float32) {
	bArr := Float32ToBytes(value)
	*bytes = append(*bytes, bArr...)
}

func DecodeUint16s(bytes *[]byte, num uint) (vals []uint16, err error) {
	needLen := (int)(2 * num)
	if len(*bytes) < needLen {
		err = errors.New("Bytes is not Enough!")
		return
	}

	vals = BytesToUint16((*bytes)[:needLen])
	*bytes = (*bytes)[needLen:]

	return
}

func DecodeUint32s(bytes *[]byte, num uint) (vals []uint32, err error) {
	needLen := (int)(4 * num)
	if len(*bytes) < needLen {
		err = errors.New("Bytes is not Enough!")
		return
	}

	vals = BytesToUint32((*bytes)[0:needLen])
	*bytes = (*bytes)[needLen:]

	return
}

func DecodeFloat32s(bytes *[]byte, num uint) (vals []float32, err error) {
	needLen := (int)(4 * num)
	if len(*bytes) < needLen {
		err = errors.New("Bytes is not Enough!")
		return
	}

	fp32vals := make([]float32, num)

	for i := (uint)(0); i < num; i++ {
		fp32vals[i] = BytesToFloat32((*bytes)[i*4 : (i+1)*4])
	}

	*bytes = (*bytes)[needLen:]

	return fp32vals, nil
}
