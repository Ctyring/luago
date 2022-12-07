package state

type luaState struct {
	stack *luaStack
}

// 创建LuaState实例
func New() *luaState {
	return &luaState{
		stack: newLuaStack(20),
	}
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
