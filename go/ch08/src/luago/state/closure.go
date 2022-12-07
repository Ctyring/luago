package state

import "go/ch08/src/luago/binchunk"

// 闭包
type closure struct {
	proto *binchunk.Prototype // Lua函数原型
}

// 创建lua闭包
func newLuaClosure(proto *binchunk.Prototype) *closure {
	return &closure{proto}
}
