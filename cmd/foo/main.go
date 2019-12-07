package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("foo")
}

func Bar() error {
	msg := "msg"
	return errors.New(fmt.Sprintf("error: %s", msg))
}
