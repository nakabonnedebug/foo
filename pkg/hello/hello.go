package hello

import "github.com/k0kubun/pp"

type Hello struct {
}

func (h *Hello) Greet(s string) {
	pp.Printf("Hello, %s", s)
}
