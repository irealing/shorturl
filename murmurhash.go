package shorturl

import "encoding/binary"

const magicNumber uint32 = 0x5bd1e995

func MurmurHash2(data []byte, seed uint32) uint32 {
	var h, k uint32
	h = seed ^ uint32(len(data))
	buf := data
	for len(buf) >= 4 {
		k = binary.LittleEndian.Uint32(buf[:4])
		buf = buf[4:]
		k *= magicNumber
		k ^= k >> 24
		k *= magicNumber
		h *= magicNumber
		h ^= k
	}
	switch len(buf) {
	case 3:
		h ^= uint32(buf[2]) << 16
	case 2:
		h ^= uint32(buf[1]) << 8
	case 1:
		h ^= uint32(buf[0])
		h *= magicNumber
	}
	h ^= h >> 13
	h *= magicNumber
	h ^= h >> 15
	return h
}
