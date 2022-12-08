package state

import "go/ch09/src/luago/binchunk"
import "go/ch09/src/luago/api"

// 闭包
// proto和goFunc必须有一个不为空
type closure struct {
	proto  *binchunk.Prototype // Lua函数原型
	goFunc api.GoFunction      // Go函数原型
}

// 创建lua闭包
func newLuaClosure(proto *binchunk.Prototype) *closure {
	return &closure{proto: proto}
}

// 创建go闭包
func newGoClosure(f api.GoFunction) *closure {
	return &closure{goFunc: f}
}
