package state

import (
	"go/ch10/src/luago/api"
	"go/ch10/src/luago/number"
	"math"
)

// Arith方法
// 使用两个参数返回一个参数的形式表达，如果是一元运算符，忽略第二个参数即可
var (
	iadd  = func(a, b int64) int64 { return a + b }
	fadd  = func(a, b float64) float64 { return a + b }
	isub  = func(a, b int64) int64 { return a - b }
	fsub  = func(a, b float64) float64 { return a - b }
	imul  = func(a, b int64) int64 { return a * b }
	fmul  = func(a, b float64) float64 { return a * b }
	imod  = number.IMod
	fmod  = number.FMod
	pow   = math.Pow
	div   = func(a, b float64) float64 { return a / b }
	iidiv = number.IFloorDiv
	fidiv = number.FFloorDiv
	band  = func(a, b int64) int64 { return a & b }
	bor   = func(a, b int64) int64 { return a | b }
	bxor  = func(a, b int64) int64 { return a ^ b }
	shl   = number.ShiftLeft
	shr   = number.ShiftRight
	iunm  = func(a, _ int64) int64 { return -a }
	funm  = func(a, _ float64) float64 { return -a }
	bnot  = func(a, _ int64) int64 { return ^a }
)

// 容纳整数运算和浮点数运算
type operator struct {
	integerFunc func(int64, int64) int64
	floatFunc   func(float64, float64) float64
}

var operators = []operator{
	{iadd, fadd},   // OP_ADD
	{isub, fsub},   // OP_SUB
	{imul, fmul},   // OP_MUL
	{imod, fmod},   // OP_MOD
	{nil, pow},     // OP_POW
	{nil, div},     // OP_DIV
	{iidiv, fidiv}, // OP_IDIV
	{band, nil},    // OP_BAND
	{bor, nil},     // OP_BOR
	{bxor, nil},    // OP_BXOR
	{shl, nil},     // OP_SHL
	{shr, nil},     // OP_SHR
	{iunm, funm},   // OP_UNM
	{bnot, nil},    // OP_BNOT
}

func (self *luaState) Arith(op api.ArithOp) {
	var a, b luaValue
	b = self.stack.pop()
	if op != api.LUA_OPUNM && op != api.LUA_OPBNOT {
		a = self.stack.pop()
	} else {
		a = b
	}

	operator := operators[op]
	if result := _arith(a, b, operator); result != nil {
		self.stack.push(result)
	} else {
		panic("Arith error!")
	}
}

// 执行计算
func _arith(a, b luaValue, op operator) luaValue {
	if op.integerFunc != nil { // 整数运算
		if x, ok := convertToInteger(a); ok {
			if y, ok := convertToInteger(b); ok {
				return op.integerFunc(x, y)
			}
		}
	}
	if op.floatFunc != nil {
		if x, ok := convertToFloat(a); ok {
			if y, ok := convertToFloat(b); ok {
				return op.floatFunc(x, y)
			}
		}
	}
	return nil
}
