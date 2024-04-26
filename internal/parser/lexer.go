package parser

import (
	"github.com/alecthomas/participle/v2/lexer"
)

var atLexer = lexer.MustStateful(lexer.Rules{
	"Root": {
		{"AT", "AT", lexer.Push("Commands")},
	},
	"Commands": {
		{"CommandName", "(\\+|&)?[A-Z][A-Z0-9]*", lexer.Push("Command")},
	},
	"Command": {
		{"TestModifier", "=\\?", nil},
		{"ReadModifier", "\\?", nil},
		{"SetModifier", "=", lexer.Push("SetCommand")},
		{"CommandSeparator", ";", lexer.Pop()},
	},
	"SetCommand": {
		{"quotedStringStart", `"`, lexer.Push("QuotedString")},
		{"ValueSeparator", ",", nil},
		lexer.Include("String"),
	},
	"String": {
		{"Escaped", `\\.`, nil},
		{"Chars", `[^",\\]+`, nil},
	},
	"QuotedString": {
		{"QuotedEscaped", `\\.`, nil},
		{"quotedStringEnd", `"`, lexer.Pop()},
		{"QuotedChars", `[^"\\]+`, nil},
	},
})
