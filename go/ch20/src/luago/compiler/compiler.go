package compiler

import (
	"go/ch20/src/luago/binchunk"
	"go/ch20/src/luago/compiler/codegen"
	"go/ch20/src/luago/compiler/parser"
)

func Compile(chunk, chunkname string) *binchunk.Prototype {
	ast := parser.Parse(chunk, chunkname)
	return codegen.GenProto(ast)
}
