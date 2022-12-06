package state

func (self *luaState) PC() int {
	return self.pc
}

// 修改指令指针
func (self *luaState) AddPC(n int) {
	self.pc += n
}

// 取出当前指令，将PC指向下一条指令
func (self *luaState) Fetch() uint32 {
	i := self.proto.Code[self.pc]
	self.pc++
	return i
}

// 从常量表取出一个常量值，然后推入栈顶
func (self *luaState) GetConst(idx int) {
	c := self.proto.Constants[idx]
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
