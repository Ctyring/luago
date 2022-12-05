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
