package state

import "go/ch07/src/luago/api"

// 创建一个空lua表，将其推入栈顶，两个参数指定数组部分和哈希表部分的初始大小
func (self *luaState) CreateTable(nArr, nRec int) {
	t := newLuaTable(nArr, nRec)
	self.stack.push(t)
}

// 属于CreateTable的特殊情况，无法预估大小，所以直接创建一个空表
func (self *luaState) NewTable() {
	self.CreateTable(0, 0)
}

// 根据键(从栈顶弹出)从表中取值，将值推入栈顶
func (self *luaState) GetTable(idx int) api.LuaType {
	t := self.stack.get(idx)
	k := self.stack.pop()
	return self.getTable(t, k)
}

func (self *luaState) getTable(t, k luaValue) api.LuaType {
	if tbl, ok := t.(*luaTable); ok {
		v := tbl.get(k)
		self.stack.push(v)
		return typeOf(v)
	}
	panic("not a table!")
}

// 根据参数传入的字符串键从表中取值，将值推入栈顶
func (self *luaState) GetField(idx int, k string) api.LuaType {
	t := self.stack.get(idx)
	return self.getTable(t, k)
}

// 传入数字键从表中取值，将值推入栈顶
func (self *luaState) GetI(idx int, i int64) api.LuaType {
	t := self.stack.get(idx)
	return self.getTable(t, i)
}
