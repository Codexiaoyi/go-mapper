package main

import "github.com/Codexiaoyi/go-mapper"

type src struct {
	A string
	B int
	C bool
}

type dest struct {
	A string
	B int
	C bool
}

func main() {
	var src src
	var dest dest
	src.A = "aaa"
	src.B = 1
	src.C = true

	err := mapper.StructMapByFieldName(&src, &dest)
	if err != nil {
		println(err)
	}
	println(dest.A, dest.B, dest.C)
}
