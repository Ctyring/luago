package main

import (
	"fmt"
	"go/ch11/src/luago/api"
	"go/ch11/src/luago/state"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		data, err := os.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		ls := state.New()
		ls.Register("print", print)
		ls.Register("getmetatable", getMetable)
		ls.Register("setmetatable", setMetable)
		ls.Load(data, os.Args[1], "b")
		ls.Call(0, 0)
	}
}

func print(ls api.LuaState) int {
	nArgs := ls.GetTop()
	for i := 1; i <= nArgs; i++ {
		if ls.IsBoolean(i) {
			fmt.Printf("%t", ls.ToBoolean(i))
		} else if ls.IsString(i) {
			fmt.Printf("%q", ls.ToString(i))
		} else {
			fmt.Printf("%v", ls.TypeName(ls.Type(i)))
		}
		if i < nArgs {
			fmt.Printf("\t")
		}
	}
	fmt.Printf("\n")
	return 0
}

func getMetable(ls api.LuaState) int {
	if !ls.GetMetatable(1) {
		ls.PushNil()
	}
	return 1
}

func setMetable(ls api.LuaState) int {
	ls.SetMetatable(1)
	return 1
}
