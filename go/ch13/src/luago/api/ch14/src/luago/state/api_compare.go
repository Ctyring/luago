package state

import "go/ch13/src/luago/api/ch14/src/luago/api"

func (self *luaState) RawEqual(idx1, idx2 int) bool {
	if self.stack.isValid(idx1) || self.stack.isValid(idx2) {
		return false
	}

	a := self.stack.get(idx1)
	b := self.stack.get(idx2)
	return _eq(a, b, nil)
}

// 根据操作码调用相应的函数进行比较
func (self *luaState) Compare(idx1, idx2 int, op api.CompareOp) bool {
	a := self.stack.get(idx1)
	b := self.stack.get(idx2)
	switch op {
	case api.LUA_OPEQ:
		return _eq(a, b, self)
	case api.LUA_OPLT:
		return _lt(a, b, self)
	case api.LUA_OPLE:
		return _le(a, b, self)
	default:
		panic("invalid compare op!")
	}
}

func _eq(a, b luaValue, ls *luaState) bool {
	switch x := a.(type) {
	case nil:
		return b == nil
	case bool:
		y, ok := b.(bool)
		return ok && x == y
	case string:
		y, ok := b.(string)
		return ok && x == y
	case int64:
		switch y := b.(type) {
		case int64:
			return x == y
		case float64:
			return float64(x) == y
		default:
			return false
		}
	case float64:
		switch y := b.(type) {
		case float64:
			return x == y
		case int64:
			return x == float64(y)
		default:
			return false
		}
	case *luaTable:
		// 两个不同表直接比较：调用元方法
		if y, ok := b.(*luaTable); ok && x != y && ls != nil {
			if result, ok := callMetamethod(x, y, "__eq", ls); ok {
				return convertToBoolean(result)
			}
		}
		return a == b
	default:
		return a == b
	}
}

func _lt(a, b luaValue, ls *luaState) bool {
	switch x := a.(type) {
	case string:
		if y, ok := b.(string); ok {
			return x < y
		}
	case int64:
		switch y := b.(type) {
		case int64:
			return x < y
		case float64:
			return float64(x) < y
		}
	case float64:
		switch y := b.(type) {
		case float64:
			return x < y
		case int64:
			return x < float64(y)
		}
	}
	if result, ok := callMetamethod(a, b, "__lt", ls); ok {
		return convertToBoolean(result)
	} else {
		panic("comparison error!")
	}
}

func _le(a, b luaValue, ls *luaState) bool {
	switch x := a.(type) {
	case string:
		if y, ok := b.(string); ok {
			return x <= y
		}
	case int64:
		switch y := b.(type) {
		case int64:
			return x <= y
		case float64:
			return float64(x) <= y
		}
	case float64:
		switch y := b.(type) {
		case float64:
			return x <= y
		case int64:
			return x <= float64(y)
		}
	}
	if result, ok := callMetamethod(a, b, "__le", ls); ok {
		return convertToBoolean(result)
	} else if result, ok := callMetamethod(b, a, "__lt", ls); ok {
		return !convertToBoolean(result)
	} else {
		panic("comparison error!")
	}
}
