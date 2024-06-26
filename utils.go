package utils

import (
	"fmt"
	"log"

	"github.com/474420502/random"
)

func TryPanic(do func()) (err error) {
	defer func() {
		err = recover().(error)
	}()
	do()
	return nil
}

var basechars []byte = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789")

func Rangdom(s, e int, seeds ...interface{}) []byte {
	var r *random.Random
	if len(seeds) > 0 {
		r = seeds[0].(*random.Random)
	} else {
		r = random.NewNoLog()
	}

	size := len(basechars)
	msize := r.Intn(e+s) + s
	var result []byte = make([]byte, msize)
	for i := 0; i < msize; i++ {
		result[i] = basechars[r.Intn(size)]
	}

	return result
}

func Expect(exceptValid string, inputs ...any) {
	result := fmt.Sprint(inputs...)
	if result != exceptValid {
		log.Panicf("inputs != except: %s != %s", result, exceptValid)
	}
}
