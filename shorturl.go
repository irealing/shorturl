package shorturl

import (
	"strings"
	"errors"
)

const (
	baseNum     uint32 = 0x24924925
	alphaBat    string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaBatLen uint32 = 62
	hashSeed    uint32 = 0x31c96
)

type URLHash struct {
	Hash     uint32
	Fragment uint32
	Value    uint32
	Index    uint32
}
type ShortedURL struct {
	URL     string
	Hash    *URLHash
	Shorted string
}

func (su *ShortedURL) ReMake() {
	su.Shorted = makeStrHash(su.Hash)
}
func NewURLHash(url string) *URLHash {
	bits := []byte(url)
	hash := MurmurHash2(bits, hashSeed)
	f, v := hash/baseNum, hash%baseNum
	return &URLHash{Hash: hash, Fragment: f, Value: v}
}
func makeStrHash(hash *URLHash) string {
	f := hash.Fragment << 3
	if hash.Index > 0 {
		f += hash.Index & 0x7
	}
	value := hash.Value
	var builder strings.Builder
	builder.WriteString(alphaBat[f : f+1])
	for {
		i := value % alphaBatLen
		builder.WriteString(alphaBat[i : i+1])
		value = value / alphaBatLen
		if value < 1 {
			break
		}
	}
	return builder.String()
}

func NewShortedURL(url string) *ShortedURL {
	hash := NewURLHash(url)
	s := makeStrHash(hash)
	return &ShortedURL{URL: url, Hash: hash, Shorted: s}
}

func Decode(s string) (*URLHash, error) {
	if len(s) < 2 {
		return nil, errors.New("error format")
	}
	tag := s[0:1]
	n := strings.Index(alphaBat, tag)
	if n < 0 {
		return nil, errors.New("unknown option")
	}
	f, index := (n&0x38)>>3, n&0x7
	s = s[1:]
	var value uint32
	var base uint32 = 1
	for i := 0; i < len(s); i++ {
		v := strings.Index(alphaBat, s[i:i+1])
		value += uint32(v) * base
		base *= alphaBatLen
	}
	hash := uint32(f)*baseNum + value
	return &URLHash{Hash: hash, Fragment: uint32(f), Value: value, Index: uint32(index)}, nil
}
