package main

import (
	"fmt"

	"github.com/mediocregopher/radix.v2/cluster"
)

func main() {
	cl, err := cluster.New("10.0.5.134:7000")
	if err != nil {
		panic(err)
	}

	s, err := cl.Cmd("SET", "foo", "bar").Str()
	if err != nil {
		panic(err)
	}

	fmt.Println(s)

	s, err = cl.Cmd("GET", "foo").Str()
	if err != nil {
		panic(err)
	}

	fmt.Println(s)

}
