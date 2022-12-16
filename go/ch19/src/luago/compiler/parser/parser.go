package parser

import (
	"go/ch19/src/luago/compiler/ast"
	"go/ch19/src/luago/compiler/lexer"
)

func Parse(chunk, chunkName string) *ast.Block {
	l := lexer.NewLexer(chunk, chunkName)
	block := parseBlock(l)
	l.NextTokenOfKind(lexer.TOKEN_EOF)
	return block
}
