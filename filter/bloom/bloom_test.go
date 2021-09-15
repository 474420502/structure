package bloom

import (
	"log"
	"math/rand"
	"testing"

	"github.com/474420502/random"
)

var basechars []byte = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789")

func TestBloom(t *testing.T) {
	var keynum = uint64(100000)

	bloom := New(keynum)
	r := random.New()
	var collect [][]byte

	var truecount int
	for i := uint64(0); i < keynum; i++ {
		var chars []byte
		for n := 0; n < rand.Intn(32)+5; n++ {
			s := r.Intn(len(basechars))
			chars = append(chars, basechars[s])
		}
		collect = append(collect, chars)
		if bloom.AddBytes(chars) {
			truecount++
		}
	}

	for _, chars := range collect {
		if !bloom.ContainsBytes(chars) {
			t.Error("")
		}
	}

	hitsize := bloom.HitSize()

	buf := bloom.Encode()
	bloom.Reset()
	bloom.Decode(buf)

	ratio := float64(keynum-bloom.HitSize()) / float64(keynum)
	if ratio == 0 || bloom.HitSize() == 0 || bloom.hitSize != hitsize {
		t.Error("Encode and Decode error")
	}

	log.Println(ratio, bloom.HitRatio())

}

func TestBloomPut(t *testing.T) {
	var keynum = uint64(100000)

	bloom := New(keynum * 5)
	r := random.New()
	var collect []uint64

	var truecount int
	for i := uint64(0); i < keynum; i++ {

		var v = r.Uint64()
		collect = append(collect, v)
		if bloom.Add(v) {
			truecount++
		}
	}

	for _, chars := range collect {
		if !bloom.Contains(chars) {
			t.Error("")
		}
	}

	if truecount != int(keynum-bloom.HitSize()) {
		t.Error("hit size error")
	}

	hitsize := bloom.HitSize()

	buf := bloom.Encode()
	bloom.Reset()
	bloom = NewByDecode(buf)

	ratio := float64(keynum-bloom.HitSize()) / float64(keynum)
	if ratio == 0 || bloom.HitSize() == 0 || bloom.hitSize != hitsize {
		t.Error("Encode and Decode error")
	}

	log.Println(ratio)
	bloom.Reset()
}

func TestBloomPutForce(t *testing.T) {
	r := random.New()

	for n := 0; n < 200; n++ {

		var keynum = uint64(r.Int63n(10000) + 1000)
		bloom := New(keynum * uint64(r.Intn(7)+3))
		var collect []uint64

		var truecount int
		for i := uint64(0); i < keynum; i++ {

			var v = r.Uint64()
			collect = append(collect, v)
			if bloom.Add(v) {
				truecount++
			}
		}

		for _, chars := range collect {
			if !bloom.Contains(chars) {
				t.Error("")
			}
		}

		if truecount != int(keynum-bloom.HitSize()) {
			t.Error("hit size error")
		}

		hitsErrorRatio := float64(keynum-bloom.HitSize()) / float64(keynum)
		if hitsErrorRatio >= 0.2 {
			log.Println(hitsErrorRatio)
		}

		bloom.Reset()
	}
}

func BenchmarkPut(b *testing.B) {
	r := random.New()
	bloom := New(10000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bloom.Add(r.Uint64())
	}
}
