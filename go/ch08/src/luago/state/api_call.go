package state

import (
	"fmt"
	"go/ch08/src/luago/binchunk"
	"go/ch08/src/luago/vm"
)

// 加载二进制chunk，第一个参数是二进制chunk，第二个参数是chunk名字，第三个参数指定加载模式("b" 二进制 "t" 文本 "bt" 二进制或文本)
func (self *luaState) Load(chunk []byte, chunkName, mode string) int {
	proto := binchunk.Undump(chunk)
	c := newLuaClosure(proto)
	self.stack.push(c)
	return 0
}

// 调用Lua函数
// 第一个参数是参数个数，第二个参数是返回值个数
func (self *luaState) Call(nArgs, nResults int) {
	// 根据索引取出函数，判断是否真的是Lua函数
	val := self.stack.get(-(nArgs + 1))
	if c, ok := val.(*closure); ok {
		fmt.Printf("call %s<%d,%d>\n", c.proto.Source, c.proto.LineDefined, c.proto.LastLineDefined)
		// 如果是真的，则调用
		self.callLuaClosure(nArgs, nResults, c)
	} else {
		panic("not function!")
	}
}

// 调用Lua函数
func (self *luaState) callLuaClosure(nArgs, nResults int, c *closure) {
	// 拿到编译器为我们事先准备好的信息
	nRegs := int(c.proto.MaxStackSize)
	nParams := int(c.proto.NumParams)
	isVararg := c.proto.IsVararg == 1

	// 创建Lua栈帧
	newStack := newLuaStack(nRegs + 20)
	// 把闭包和调用帧联系起来
	newStack.closure = c

	// 把参数传递给新的Lua栈帧
	funcAndArgs := self.stack.popN(nArgs + 1) // 把函数和参数弹出
	newStack.pushN(funcAndArgs[1:], nParams)  // 把参数传递给新的Lua栈帧
	newStack.top = nRegs                      // 设置栈顶
	if nArgs > nParams && isVararg {          // 如果参数个数大于参数个数，且是可变参数
		newStack.varargs = funcAndArgs[nParams+1:] // 把多余的参数传递给可变参数
	}

	// 把新的Lua栈帧压入Lua虚拟机栈
	self.pushLuaStack(newStack)
	// 执行Lua函数
	self.runLuaClosure()
	// 把返回值传递给调用者
	self.popLuaStack()

	// 根据期望的返回值个数，从新的Lua栈帧中弹出返回值
	if nResults != 0 {
		results := newStack.popN(newStack.top - nRegs) // 弹出返回值
		self.stack.check(len(results))                 // 检查栈空间
		self.stack.pushN(results, nResults)            // 把返回值传递给调用者
	}
}

// 执行被调函数
func (self *luaState) runLuaClosure() {
	for {
		inst := vm.Instruction(self.Fetch())
		inst.Execute(self)
		if inst.Opcode() == vm.OP_RETURN {
			break
		}
	}
}

// 从栈顶弹出n个值
func (self *luaStack) popN(n int) []luaValue {
	vals := make([]luaValue, n)
	for i := n - 1; i >= 0; i-- {
		vals[i] = self.pop()
	}
	return vals
}

// 把n个值压入栈顶
func (self *luaStack) pushN(vals []luaValue, n int) {
	nVals := len(vals)
	if n < 0 {
		n = nVals // n < 0时, 压入全部值
	}
	for i := 0; i < n; i++ {
		if i < nVals {
			self.push(vals[i])
		} else {
			self.push(nil) // 压入nil补齐
		}
	}
}
