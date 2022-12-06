package api

// 扩展LuaState接口
type LuaVM interface {
	LuaState
	PC() int          // 获取当前指令指针
	AddPC(n int)      // 修改指令指针
	Fetch() uint32    // 取出当前指令，将PC指向下一条指令
	GetConst(idx int) // 将指定常量推入栈顶
	GetRK(rk int)     // 将指定常量或栈值推入栈顶
}
