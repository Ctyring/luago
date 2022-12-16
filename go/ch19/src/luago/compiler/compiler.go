package compiler

import (
	"go/ch19/src/luago/binchunk"
	"go/ch19/src/luago/compiler/codegen"
	"go/ch19/src/luago/compiler/parser"
)

func Compile(chunk, chunkname string) *binchunk.Prototype {
	ast := parser.Parse(chunk, chunkname)
	return codegen.GenProto(ast)
}
