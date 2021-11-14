package main

import (
	"log"
	"math/rand"

	"github.com/474420502/random"
	"github.com/474420502/structure/filter/bloom"
)

var basechars []byte = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789")

func main() {
	var keynum = uint64(100000)

	bloom := bloom.New(keynum * 10) // create bloom
	r := random.New()
	var collect [][]byte // collect the bytes

	var truecount int
	for i := uint64(0); i < keynum; i++ {
		var chars []byte
		// random create bytes
		for n := 0; n < rand.Intn(32)+5; n++ {
			s := r.Intn(len(basechars))
			chars = append(chars, basechars[s])
		}
		collect = append(collect, chars)
		if bloom.AddBytes(chars) { // add the bloom
			truecount++
		}
	}

	for _, chars := range collect {
		if !bloom.ContainsBytes(chars) { // is in bloom?
			log.Panic("bloom.ContainsBytes error")
		}
	}

	hitsize := bloom.HitSize()

	buf := bloom.Encode() // encode to buffer
	bloom.Reset()
	bloom.Decode(buf) // decode from buffer

	ratio := float64(keynum-bloom.HitSize()) / float64(keynum)
	if ratio == 0 || bloom.HitSize() == 0 || bloom.HitSize() != hitsize {
		log.Println("Encode and Decode error")
	}

	log.Println(bloom.HitSize(), ratio) // 95192(bit used by bloom) 0.04843(percentage of duplicates)
}
