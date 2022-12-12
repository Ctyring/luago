package state

import "go/ch13/src/luago/api/ch14/src/luago/api"

func (self *luaState) PushNil() {
	self.stack.push(nil)
}

func (self *luaState) PushBoolean(b bool) {
	self.stack.push(b)
}

func (self *luaState) PushInteger(n int64) {
	self.stack.push(n)
}

func (self *luaState) PushNumber(n float64) {
	self.stack.push(n)
}

func (self *luaState) PushString(s string) {
	self.stack.push(s)
}

func (self *luaState) PushGoFunction(f api.GoFunction) {
	self.stack.push(newGoClosure(f, 0))
}

// 把全局环境压入栈中
func (self *luaState) PushGlobalTable() {
	global := self.registry.get(api.LUA_RIDX_GLOBALS)
	self.stack.push(global)
}

func (self *luaState) PushGoClosure(f api.GoFunction, n int) {
	// 创建Go闭包
	closure := newGoClosure(f, n)
	for i := n; i > 0; i-- {
		// 从栈中取出n个值，作为upvalue
		val := self.stack.pop()
		// 将值封装成upvalue
		closure.upvals[i-1] = &upvalue{&val}
	}
	// 将闭包压入栈中
	self.stack.push(closure)
}
