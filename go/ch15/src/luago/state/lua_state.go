package state

import "go/ch15/src/luago/api"

type luaState struct {
	registry *luaTable // 注册表
	stack    *luaStack
}

// 创建LuaState实例
func New() *luaState {
	registry := newLuaTable(0, 0)
	registry.put(api.LUA_RIDX_GLOBALS, newLuaTable(0, 0)) // 全局环境

	ls := &luaState{registry: registry}
	ls.pushLuaStack(newLuaStack(api.LUA_MINSTACK, ls)) // 创建Lua栈
	return ls
}

// 向头部添加一个调用帧
func (self *luaState) pushLuaStack(stack *luaStack) {
	stack.prev = self.stack
	self.stack = stack
}

// 从头部移除一个调用帧
func (self *luaState) popLuaStack() {
	stack := self.stack
	self.stack = stack.prev
	stack.prev = nil
}
