package util

import (
	"fmt"
	"go/ch20/src/luago/api"
)

func printStack(ls api.LuaState) {
	top := ls.GetTop()
	for i := 1; i <= top; i++ {
		t := ls.Type(i)
		switch t {
		case api.LUA_TBOOLEAN:
			fmt.Printf("[%t]", ls.ToBoolean(i))
		case api.LUA_TNUMBER:
			fmt.Printf("[%g]", ls.ToNumber(i))
		case api.LUA_TSTRING:
			fmt.Printf("[%q]", ls.ToString(i))
		default:
			fmt.Printf("[%s]", ls.TypeName(t))
		}
	}
	fmt.Println()
}
