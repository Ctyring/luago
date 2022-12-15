package parser

import "go/ch18/src/luago/compiler/lexer"
import "go/ch18/src/luago/compiler/ast"

// 创建Block结构体实例
func parseBlock(l *lexer.Lexer) *ast.Block {
	return &ast.Block{
		Stats:    parseStats(l),
		RetExps:  parseRetExps(l),
		LastLine: l.Line(),
	}
}

// 解析语句序列
func parseStats(l *lexer.Lexer) []ast.Stat {
	stats := make([]ast.Stat, 0, 8)
	for !_isReturnOrBlockEnd(l.LookAhead()) {
		stat := parseStat(l)
		if _, ok := stat.(*ast.EmptyStat); !ok {
			stats = append(stats, stat)
		}
	}
	return stats
}

// 判断块是否结束
func _isReturnOrBlockEnd(tokenKind int) bool {
	switch tokenKind {
	case lexer.TOKEN_KW_RETURN, lexer.TOKEN_KW_END, lexer.TOKEN_EOF, lexer.TOKEN_KW_ELSE, lexer.TOKEN_KW_ELSEIF, lexer.TOKEN_KW_UNTIL:
		return true
	}
	return false
}

// 解析返回值表达式
func parseRetExps(l *lexer.Lexer) []ast.Exp {
	// 如果不是return说明没有返回值
	if l.LookAhead() != lexer.TOKEN_KW_RETURN {
		return nil
	}

	l.NextToken() // skip `return`
	switch l.LookAhead() {
	// 如果发现是分号或者块结束符号，说明没有返回值
	case lexer.TOKEN_EOF, lexer.TOKEN_KW_END, lexer.TOKEN_KW_ELSE, lexer.TOKEN_KW_ELSEIF, lexer.TOKEN_KW_UNTIL:
		return []ast.Exp{}
	case lexer.TOKEN_SEP_SEMI:
		l.NextToken() // 跳过分号
		return []ast.Exp{}
	default:
		// 解析返回值序列
		exps := parseExpList(l)
		if l.LookAhead() == lexer.TOKEN_SEP_SEMI {
			l.NextToken()
		}
		return exps
	}
}
