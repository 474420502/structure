package utils

import "github.com/474420502/random"

func TryPanic(do func()) (err error) {
	defer func() {
		err = recover().(error)
	}()
	do()
	return nil
}

var basechars []byte = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789")

func Rangdom(s, e int) []byte {
	r := random.NewNoLog()

	size := len(basechars)
	msize := r.Intn(e+s) + s
	var result []byte = make([]byte, msize)
	for i := 0; i < msize; i++ {
		result[i] = basechars[r.Intn(size)]
	}

	return result
}
