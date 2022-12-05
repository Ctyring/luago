package state

import (
	"go/ch04/src/luago/api"
	"go/ch05/src/luago/number"
)

type luaValue interface{}

func typeOf(val luaValue) api.LuaType {
	switch val.(type) {
	case nil:
		return api.LUA_TNIL
	case bool:
		return api.LUA_TBOOLEAN
	case int64:
		return api.LUA_TNUMBER
	case float64:
		return api.LUA_TNUMBER
	case string:
		return api.LUA_TSTRING
	default:
		panic("todo!")
	}
}

func convertToFloat(val luaValue) (float64, bool) {
	switch x := val.(type) {
	case int64:
		return float64(x), true
	case float64:
		return x, true
	case string:
		return number.ParseFloat(x)
	default:
		return 0, false
	}
}

func convertToInteger(val luaValue) (int64, bool) {
	switch x := val.(type) {
	case int64:
		return x, true
	case float64:
		return int64(x), true
	case string:
		return number.ParseInteger(x)
	default:
		return 0, false
	}
}
