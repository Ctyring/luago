package main

import (
	"fmt"
	"go/ch15/src/luago/api"
	"go/ch15/src/luago/compiler/lexer"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		data, err := os.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		testLexer(string(data), os.Args[1])
	}
}

func testLexer(chunk, chunkName string) {
	lexer := lexer.NewLexer(chunk, chunkName)
	for {
		line, kind, token := lexer.NextToken()
		fmt.Printf("[%2d] %-10s %s\n", line, kindToCategory(kind), token)
		if kind == lexer.TOKEN_EOF {
			break
		}
	}
}

func kindToCategory(kind int) string {
	switch {
	case kind < lexer.TOKEN_SEP_SEMI:
		return "other"
	case kind <= lexer.TOKEN_SEP_RCURLY:
		return "separator"
	case kind <= lexer.TOKEN_OP_NOT:
		return "operator"
	case kind <= lexer.TOKEN_KW_WHILE:
		return "keyword"
	case kind == lexer.TOKEN_IDENTIFIER:
		return "identifier"
	case kind == lexer.TOKEN_NUMBER:
		return "number"
	case kind == lexer.TOKEN_STRING:
		return "string"
	default:
		return "other"
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

func next(ls api.LuaState) int {
	ls.SetTop(2) // next的第二个参数是可选的，所以首先调用settop以便在用户没有提供第二个参数时，将第二个参数设为nil(这种情况就是遍历的是第一个元素)
	if ls.Next(1) {
		return 2
	} else {
		ls.PushNil()
		return 1
	}
}

// paris函数实际上就是返回了next函数的三个值
func pairs(ls api.LuaState) int {
	ls.PushGoFunction(next) // 将next函数压入栈
	ls.PushValue(1)         // 将第一个参数(表)压入栈
	ls.PushNil()            // 将nil压入栈
	return 3
}

func iPairs(ls api.LuaState) int {
	ls.PushGoFunction(_iPairsAux) // 将_iPairsAux函数压入栈
	ls.PushValue(1)
	ls.PushInteger(0)
	return 3
}

func _iPairsAux(ls api.LuaState) int {
	i := ls.ToInteger(2) + 1
	ls.PushInteger(i)
	if ls.GetI(1, i) == api.LUA_TNIL {
		return 1
	} else {
		return 2
	}
}

func error(ls api.LuaState) int {
	return ls.Error()
}

func pCall(ls api.LuaState) int {
	nArgs := ls.GetTop() - 1
	status := ls.PCall(nArgs, -1, 0)
	ls.PushBoolean(status == api.LUA_OK)
	ls.Insert(1)
	return ls.GetTop()
}
