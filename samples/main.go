package main

import (
	"fmt"

	"github.com/Codexiaoyi/go-mapper"
)

type src struct {
	A string `mapper:"a"`
	B int    `mapper:"B"`
	C bool
}

type dest struct {
	A string `mapper:"a"`
	B int    `mapper:"B"`
	C bool
}

func main() {
	var src src
	var dest1 dest
	src.A = "aaa"
	src.B = 1
	src.C = true

	err := mapper.StructMapByFieldName(&src, &dest1)
	if err != nil {
		println(err)
	}
	println(dest1.A, dest1.B, dest1.C)

	var dest2 dest

	err = mapper.StructMapByTag(&src, &dest2)
	if err != nil {
		fmt.Println(err)
	}
	println(dest2.A, dest2.B, dest2.C)
}
