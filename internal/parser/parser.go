package parser

import (
	"github.com/alecthomas/participle/v2"
	"strconv"
)

type ATCommands struct {
	AT       string      `@AT`
	Commands []ATCommand `( @@ ( CommandSeparator @@ )* )?`
}

type ATCommand interface {
	value()
}

type ATTestCommand struct {
	CommandName  string `@CommandName`
	TestModifier string `@TestModifier`
}

func (v ATTestCommand) value() {}

type ATReadCommand struct {
	CommandName  string `@CommandName`
	ReadModifier string `@ReadModifier`
}

func (v ATReadCommand) value() {}

type ATExecuteCommand struct {
	CommandName string `@CommandName`
}

func (v ATExecuteCommand) value() {}

type ATSetCommand struct {
	CommandName string   `@CommandName`
	SetModifier string   `@SetModifier`
	Values      []String `(@@ (ValueSeparator @@)*)?`
}

func (v ATSetCommand) value() {}

type String struct {
	Fragments []Fragment `@@*`
}

func (s String) String() string {
	var v string
	for _, fragment := range s.Fragments {
		v += fragment.String()
	}
	return v
}

type Fragment struct {
	Escaped string `(  (@Escaped | @QuotedEscaped )`
	Text    string ` | (@Chars | @QuotedChars) )`
}

func (f Fragment) String() string {
	if f.Escaped != "" {
		s, err := strconv.Unquote(`"` + f.Escaped + `"`)
		if err != nil {
			panic(err)
		}
		return s
	} else {
		return f.Text
	}
}

type Parser struct {
	parser *participle.Parser[ATCommands]
}

func NewParser() (*Parser, error) {
	parser, err := participle.Build[ATCommands](
		participle.Lexer(atLexer),
		participle.Union[ATCommand](
			ATTestCommand{},
			ATReadCommand{},
			ATSetCommand{},
			ATExecuteCommand{},
		),
	)
	if err != nil {
		return nil, err
	}
	return &Parser{
		parser: parser,
	}, nil
}

func (p *Parser) Parse(text string) (*ATCommands, error) {
	return p.parser.ParseString("", text)
}
