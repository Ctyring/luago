package state

import "go/ch07/src/luago/binchunk"

type luaState struct {
	stack *luaStack
	proto *binchunk.Prototype
	pc    int
}

// 创建LuaState实例
func New(stackSize int, proto *binchunk.Prototype) *luaState {
	return &luaState{
		stack: newLuaStack(20),
		proto: proto,
		pc:    0,
	}
}
