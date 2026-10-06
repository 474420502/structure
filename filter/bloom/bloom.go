package bloom

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
)

// FNV-1 64-bit constants, matching hash/fnv.New64 so persisted filters keep the
// same bit layout.
const (
	fnvOffset64 = 14695981039346656037
	fnvPrime64  = 1099511628211
)

// fnv164 computes the FNV-1 64-bit hash of data without an interface call, so
// callers can pass a stack buffer without it escaping.
func fnv164(data []byte) uint64 {
	h := uint64(fnvOffset64)
	for _, b := range data {
		h *= fnvPrime64
		h ^= uint64(b)
	}
	return h
}

type Bloom struct {
	bits    []byte
	hitSize uint64
	cap     uint64
}

func NewByDecode(reader io.Reader) *Bloom {

	bl := &Bloom{}
	binary.Read(reader, binary.BigEndian, &bl.hitSize)
	binary.Read(reader, binary.BigEndian, &bl.cap)
	binary.Read(reader, binary.BigEndian, &bl.bits)

	return bl
}

// New  推荐 bitsCap = key预估容量 x 10
func New(bitsCap uint64) *Bloom {
	bsize := (bitsCap + 7) >> 3
	return &Bloom{
		bits: make([]byte, bsize),
		cap:  bsize << 3,
	}
}

// withKey encodes key to big-endian bytes and applies AddBytes or
// ContainsBytes. Fixed-width scalar types use a stack buffer, so the common
// scalar paths do not allocate.
func (bl *Bloom) withKey(key interface{}, add bool) bool {
	apply := func(b []byte) bool { return false } // placeholder
	_ = apply
	switch v := key.(type) {
	case []byte:
		if add {
			return bl.AddBytes(v)
		}
		return bl.ContainsBytes(v)
	case string:
		data := []byte(v)
		if add {
			return bl.AddBytes(data)
		}
		return bl.ContainsBytes(data)
	case int:
		var b [8]byte
		binary.BigEndian.PutUint64(b[:], uint64(v))
		if add {
			return bl.AddBytes(b[:])
		}
		return bl.ContainsBytes(b[:])
	case int64:
		var b [8]byte
		binary.BigEndian.PutUint64(b[:], uint64(v))
		if add {
			return bl.AddBytes(b[:])
		}
		return bl.ContainsBytes(b[:])
	case uint:
		var b [8]byte
		binary.BigEndian.PutUint64(b[:], uint64(v))
		if add {
			return bl.AddBytes(b[:])
		}
		return bl.ContainsBytes(b[:])
	case uint64:
		var b [8]byte
		binary.BigEndian.PutUint64(b[:], v)
		if add {
			return bl.AddBytes(b[:])
		}
		return bl.ContainsBytes(b[:])
	case int32:
		var b [4]byte
		binary.BigEndian.PutUint32(b[:], uint32(v))
		if add {
			return bl.AddBytes(b[:])
		}
		return bl.ContainsBytes(b[:])
	case uint32:
		var b [4]byte
		binary.BigEndian.PutUint32(b[:], v)
		if add {
			return bl.AddBytes(b[:])
		}
		return bl.ContainsBytes(b[:])
	case float64:
		var b [8]byte
		binary.BigEndian.PutUint64(b[:], math.Float64bits(v))
		if add {
			return bl.AddBytes(b[:])
		}
		return bl.ContainsBytes(b[:])
	case float32:
		var b [4]byte
		binary.BigEndian.PutUint32(b[:], math.Float32bits(v))
		if add {
			return bl.AddBytes(b[:])
		}
		return bl.ContainsBytes(b[:])
	default:
		var keybuf bytes.Buffer
		if err := binary.Write(&keybuf, binary.BigEndian, key); err != nil {
			panic(err)
		}
		if add {
			return bl.AddBytes(keybuf.Bytes())
		}
		return bl.ContainsBytes(keybuf.Bytes())
	}
}

// Add add the key(interface{}) to bloom. interface{} will be encoded to binary data
func (bl *Bloom) Add(key interface{}) (isExists bool) {
	return bl.withKey(key, true)
}

// AddBytes add the key([]byte) to bloom.
func (bl *Bloom) AddBytes(key []byte) (isExists bool) {
	h := fnv164(key)
	bitnum := h % bl.cap
	bytenum := bitnum >> 3
	bitnum = bitnum % 8
	bitnum = 1 << bitnum
	bset := byte(bitnum)
	isExists = (bl.bits[bytenum] & bset) != 0
	if !isExists {
		bl.hitSize++
	}
	bl.bits[bytenum] |= bset
	return
}

// Contains  if the key is in bloom. return true.
func (bl *Bloom) Contains(key interface{}) (isExists bool) {
	return bl.withKey(key, false)
}

// ContainsBytes if the key is in bloom. return true.
func (bl *Bloom) ContainsBytes(key []byte) (isExists bool) {
	h := fnv164(key)
	bitnum := h % bl.cap
	bytenum := bitnum >> 3 // byte = 8 == 1 >> 3
	bitnum = bitnum % 8
	bitnum = 1 << bitnum
	bset := byte(bitnum)

	return bl.bits[bytenum]&bset != 0
}

// Cap bits 位的数量
func (bl *Bloom) Cap() uint64 {
	return bl.cap
}

// HitSize 占用bit的size数量
func (bl *Bloom) HitSize() uint64 {
	return bl.hitSize
}

// HitRatio 占用bit的比率 == float64(bl.hitSize) / float64(bl.cap)
//
// Ratio of occupied bits. equal to float64(bl.hitSize) / float64(bl.cap)
func (bl *Bloom) HitRatio() float64 {
	return float64(bl.hitSize) / float64(bl.cap)
}

// Reset 重置
func (bl *Bloom) Reset() {

	if len(bl.bits) == 0 {
		return
	}

	bl.bits[0] = 0
	for bp := 1; bp < len(bl.bits); bp *= 2 {
		copy(bl.bits[bp:], bl.bits[:bp])
	}

	bl.hitSize = 0
}

// Encode 序列化为buf
func (bl *Bloom) Encode() *bytes.Buffer {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, &bl.hitSize)
	binary.Write(&buf, binary.BigEndian, &bl.cap)
	binary.Write(&buf, binary.BigEndian, &bl.bits)
	return &buf
}

// Decode 从buf反序列化
func (bl *Bloom) Decode(reader io.Reader) {
	bl.bits = bl.bits[:0]
	binary.Read(reader, binary.BigEndian, &bl.hitSize)
	binary.Read(reader, binary.BigEndian, &bl.cap)
	binary.Read(reader, binary.BigEndian, &bl.bits)
}
