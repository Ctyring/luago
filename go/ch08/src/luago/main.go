package main

import (
	"go/ch08/src/luago/state"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		data, err := os.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		ls := state.New()
		ls.Load(data, os.Args[1], "b")
		ls.Call(0, 0)
	}
}
