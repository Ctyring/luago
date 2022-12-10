package state

import "go/ch10/src/luago/api"

// 把键值写入表，键和值都从栈顶弹出
func (self *luaState) SetTable(idx int) {
	t := self.stack.get(idx)
	v := self.stack.pop()
	k := self.stack.pop()
	self.setTable(t, k, v)
}

func (self *luaState) setTable(t, k, v luaValue) {
	if tbl, ok := t.(*luaTable); ok {
		tbl.put(k, v)
		return
	}
	panic("not a table!")
}

// 把值写入表，键从参数传入(字符串)，值从栈顶弹出
func (self *luaState) SetField(idx int, k string) {
	t := self.stack.get(idx)
	v := self.stack.pop()
	self.setTable(t, k, v)
}

// 把值写入表，键从参数传入(数字)，值从栈顶弹出
func (self *luaState) SetI(idx int, i int64) {
	t := self.stack.get(idx)
	v := self.stack.pop()
	self.setTable(t, i, v)
}

// 向全局变量写入一个值
func (self *luaState) SetGlobal(name string) {
	t := self.registry.get(api.LUA_RIDX_GLOBALS)
	v := self.stack.pop()
	self.setTable(t, name, v)
}

// 给全局环境注册Go函数值
func (self *luaState) Register(name string, f api.GoFunction) {
	// 把go函数压入栈
	self.PushGoFunction(f)
	// 放入全局环境
	self.SetGlobal(name)
}
