package vm

import "go/ch09/src/luago/api"

// 把某个全局变量的值写入寄存器
func getTabUp(i Instruction, vm api.LuaVM) {
	a, _, c := i.ABC()
	a += 1

	// 把全局环境压入栈中
	vm.PushGlobalTable()
	// 把常量表中的索引为c的值压入栈中
	vm.GetRK(c)
	// 从栈中弹出两个值，然后把第二个值作为key，第一个值作为table，把table[key]的值写入寄存器a
	vm.GetTable(-2)
	vm.Replace(a)
	// 弹出全局环境
	vm.Pop(1)
}
