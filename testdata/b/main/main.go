package main

import (
	"github.com/ahilbig/go-package-plantuml/testdata/b"
	"github.com/ahilbig/go-package-plantuml/testdata/b/sub"
)

func m1(a1 b.IA) {
	a1.Add(sub2.SubSA{})
}

func main() {
	m1(&b.SB{})
}

