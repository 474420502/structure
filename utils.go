package utils

func TryPanic(do func()) (err error) {
	defer func() {
		err = recover().(error)
	}()
	do()
	return nil
}
