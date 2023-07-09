package main

import "fmt"

type resp struct {
	code int64
	msg  string
}

func f() (r resp) {

	defer func() {
		r.msg = "111"
		r.code = 200
	}()

	r.msg = "test"
	r.code = 500
	return
}

func main() {
	fmt.Println(f())
}
