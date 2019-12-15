package main

import (
	"errors"
	"fmt"

	"github.com/nakabonnedebug/foo/pkg/hello"
)

func main() {
	h := &hello.Hello{}
	h.Greet("foo")
}

func Bar() error {
	msg := "bar"
	return errors.New(fmt.Sprintf("error: %s", msg))
}
