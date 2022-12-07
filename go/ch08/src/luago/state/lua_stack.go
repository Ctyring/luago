package state

// lua栈
type luaStack struct {
	slots   []luaValue // 栈的数据
	top     int        // 栈顶
	prev    *luaStack
	closure *closure   // 当前栈对应的闭包
	varargs []luaValue // 可变参数
	pc      int
}

// 创建指定容量的栈
func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
	}
}

// 检查空闲空间是否还可以容纳至少n个值
func (self *luaStack) check(n int) {
	free := len(self.slots) - self.top
	// 如果空闲空间不足，则扩容
	for i := free; i < n; i++ {
		self.slots = append(self.slots, nil)
	}
}

// 将值压入栈顶
func (self *luaStack) push(val luaValue) {
	// 如果溢出。终止
	if self.top == len(self.slots) {
		panic("stack overflow!")
	}
	self.slots[self.top] = val
	self.top++
}

// 将栈顶的值弹出
func (self *luaStack) pop() luaValue {
	// 如果栈顶是空的，终止
	if self.top < 1 {
		panic("stack underflow!")
	}
	self.top--
	val := self.slots[self.top]
	self.slots[self.top] = nil // GC
	return val
}

// 把索引转换成绝对索引
func (self *luaStack) absIndex(idx int) int {
	if idx >= 0 {
		return idx
	}
	return idx + self.top + 1
}

// 检查索引是否合法
func (self *luaStack) isValid(idx int) bool {
	absIds := self.absIndex(idx)
	return absIds > 0 && absIds <= self.top
}

// 根据索引从栈里取值
func (self *luaStack) get(idx int) luaValue {
	absIds := self.absIndex(idx)
	if absIds > 0 && absIds <= self.top {
		return self.slots[absIds-1]
	}
	return nil
}

// 根据索引设置栈里的值
func (self *luaStack) set(idx int, val luaValue) {
	absIds := self.absIndex(idx)
	if absIds > 0 && absIds <= self.top {
		self.slots[absIds-1] = val
		return
	}
	// 如果索引无效，终止
	panic("invalid index!")
}

// 反转
func (self *luaStack) reverse(from, to int) {
	slots := self.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}
