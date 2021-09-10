package bloom

import (
	"bytes"
	"encoding/binary"
	"hash"
	"hash/fnv"
	"io"
)

type Bloom struct {
	bits     []byte
	hitSize  uint64
	cap      uint64
	hashFunc hash.Hash64
}

func NewByDecode(reader io.Reader) *Bloom {

	bl := &Bloom{hashFunc: fnv.New64()}
	binary.Read(reader, binary.BigEndian, &bl.hitSize)
	binary.Read(reader, binary.BigEndian, &bl.cap)
	binary.Read(reader, binary.BigEndian, &bl.bits)

	return bl
}

// New  推荐 bitsCap = key预估容量 x 10
func New(bitsCap uint64) *Bloom {
	bsize := (bitsCap + 7) >> 3
	return &Bloom{
		bits:     make([]byte, bsize),
		cap:      bsize << 3,
		hashFunc: fnv.New64(),
	}
}

func (bl *Bloom) Put(key interface{}) (isExists bool) {
	var keybuf bytes.Buffer
	err := binary.Write(&keybuf, binary.BigEndian, key)
	if err != nil {
		panic(err)
	}
	return bl.PutBytes(keybuf.Bytes())
}

func (bl *Bloom) PutBytes(key []byte) (isExists bool) {
	defer bl.hashFunc.Reset()
	_, err := bl.hashFunc.Write(key)
	if err != nil {
		panic(err)
	}
	h := bl.hashFunc.Sum64()
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

func (bl *Bloom) Contains(key interface{}) (isExists bool) {
	var keybuf bytes.Buffer
	err := binary.Write(&keybuf, binary.BigEndian, key)
	if err != nil {
		panic(err)
	}
	return bl.ContainsBytes(keybuf.Bytes())
}

func (bl *Bloom) ContainsBytes(key []byte) (isExists bool) {
	defer bl.hashFunc.Reset()
	_, err := bl.hashFunc.Write(key)
	if err != nil {
		panic(err)
	}
	h := bl.hashFunc.Sum64()
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
