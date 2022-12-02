package main

import (
	"fmt"
	"go/ch02/src/luago/binchunk"
	"go/ch03/src/luago/vm"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		data, err := os.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		proto := binchunk.Undump(data)
		list(proto)
	}
}

// 把函数原型的信息打印到控制台
func list(f *binchunk.Prototype) {
	printHeader(f)
	printCode(f)
	printDeatils(f)
	for _, p := range f.Protos {
		list(p)
	}
}

// 打印函数原型的头部信息
func printHeader(f *binchunk.Prototype) {
	funcType := "main"
	if f.LineDefined > 0 {
		funcType = "function"
	}
	varargFlag := ""
	if f.IsVararg > 0 {
		varargFlag = "+"
	}
	fmt.Printf("\n%s <%s:%d,%d> (%d instructions)\n", funcType, f.Source, f.LineDefined, f.LastLineDefined, len(f.Code))
	fmt.Printf("%d%s params, %d slots, %d upvalues, ", f.NumParams, varargFlag, f.MaxStackSize, len(f.Upvalues))
	fmt.Printf("%d locals, %d constants, %d functions\n", len(f.LocVars), len(f.Constants), len(f.Protos))
}

// 打印指令的序号、行号和十六进制表示
func printCode(f *binchunk.Prototype) {
	for pc, c := range f.Code {
		line := "-"
		if len(f.LineInfo) > 0 {
			line = fmt.Sprintf("%d", f.LineInfo[pc])
		}
		i := vm.Instruction(c)
		fmt.Printf("\t%d\t[%s]\t%s\t", pc+1, line, i.OpName())
		printOperands(i)
		fmt.Println("\n")
	}
}

// 打印指令的操作数
func printOperands(i vm.Instruction) {
	switch i.OpMode() {
	// 对应iABC模式的指令，首先打出操作数A
	case vm.IABC:
		a, b, c := i.ABC()
		fmt.Printf("%d", a)
		// 如果使用了操作数B或C，打印出来
		// 如果操作数B或C的最高位是1，则认为它表示常量表索引，按负数输出
		if i.BMode() != vm.OpArgN {
			if b > 0xFF {
				fmt.Printf(" %d", -1-(b&0xFF))
			} else {
				fmt.Printf(" %d", b)
			}
		}
		if i.CMode() != vm.OpArgN {
			if c > 0xFF {
				fmt.Printf(" %d", -1-(c&0xFF))
			} else {
				fmt.Printf(" %d", c)
			}
		}
	case vm.IABx:
		a, bx := i.ABx()
		fmt.Printf("%d", a)
		if i.BMode() == vm.OpArgK {
			fmt.Printf(" %d", -1-bx)
		} else if i.BMode() == vm.OpArgU {
			fmt.Printf(" %d", bx)
		}
	case vm.IAsBx:
		a, sbx := i.AsBx()
		fmt.Printf("%d %d", a, sbx)
	case vm.IAx:
		ax := i.Ax()
		fmt.Printf("%d", -1-ax)
	}

}

// 打印常量表、局部变量表和Upvalue表
func printDeatils(f *binchunk.Prototype) {
	fmt.Printf("constants (%d):\n", len(f.Constants))
	for i, c := range f.Constants {
		fmt.Printf("\t%d\t%s\n", i+1, constantToString(c))
	}

	fmt.Printf("locals (%d):\n", len(f.LocVars))
	for i, v := range f.LocVars {
		fmt.Printf("\t%d\t%s\t%d\t%d\n", i, v.VarName, v.StartPC+1, v.EndPC+1)
	}

	fmt.Printf("upvalues (%d):\n", len(f.Upvalues))
	for i, u := range f.Upvalues {
		fmt.Printf("\t%d\t%s\t%d\t%d\n", i, upvalName(f, i), u.Instack, u.Idx)
	}
}

// 把常量转换成字符串表示
func constantToString(c interface{}) string {
	switch c.(type) {
	case nil:
		return "nil"
	case bool:
		return fmt.Sprint(c)
	case int64:
		return fmt.Sprint(c)
	case float64:
		return fmt.Sprint(c)
	case string:
		return fmt.Sprintf("%q", c)
	default:
		return "?"
	}
}

// 获取Upvalue的名字
func upvalName(f *binchunk.Prototype, idx int) string {
	if len(f.UpvalueNames) > 0 {
		return f.UpvalueNames[idx]
	}
	return "-"
}
