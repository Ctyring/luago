package state

func (self *luaState) PC() int {
	return self.stack.pc
}

// 修改指令指针
func (self *luaState) AddPC(n int) {
	self.stack.pc += n
}

// 取出当前指令，将PC指向下一条指令
func (self *luaState) Fetch() uint32 {
	i := self.stack.closure.proto.Code[self.stack.pc]
	self.stack.pc++
	return i
}

// 从常量表取出一个常量值，然后推入栈顶
func (self *luaState) GetConst(idx int) {
	c := self.stack.closure.proto.Constants[idx]
	self.stack.push(c)
}

func (self *luaState) GetRK(rk int) {
	if rk > 0xFF { // constant
		// 把常量推入栈顶
		self.GetConst(rk & 0xFF)
	} else { // register
		// 把某个索引处的栈值推入栈顶
		self.PushValue(rk + 1)
	}
}

// 获取寄存器数量
func (self *luaState) RegisterCount() int {
	return int(self.stack.closure.proto.MaxStackSize)
}

// 将可变参数推入栈顶
func (self *luaState) LoadVararg(n int) {
	if n < 0 {
		n = len(self.stack.varargs)
	}
	self.stack.check(n)
	self.stack.pushN(self.stack.varargs, n)
}

// 将指定子函数原型推入栈顶
func (self *luaState) LoadProto(idx int) {
	proto := self.stack.closure.proto.Protos[idx]
	closure := newLuaClosure(proto)
	self.stack.push(closure)
}
