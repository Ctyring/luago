package compiler

import (
	"go/ch21/src/luago/binchunk"
	"go/ch21/src/luago/compiler/codegen"
	"go/ch21/src/luago/compiler/parser"
)

func Compile(chunk, chunkname string) *binchunk.Prototype {
	ast := parser.Parse(chunk, chunkname)
	return codegen.GenProto(ast)
}
