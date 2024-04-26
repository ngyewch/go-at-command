package go_at_command

import (
	internalParser "github.com/ngyewch/go-at-command/internal/parser"
)

type Parser struct {
	parser *internalParser.Parser
}

func NewParser() (*Parser, error) {
	parser, err := internalParser.NewParser()
	if err != nil {
		return nil, err
	}
	return &Parser{
		parser: parser,
	}, nil
}

func (p *Parser) Parse(text string) (*ATCommands, error) {
	internalAtCommands, err := p.parser.Parse(text)
	if err != nil {
		return nil, err
	}
	var atCommands ATCommands
	for _, internalAtCommand := range internalAtCommands.Commands {
		switch cmd := internalAtCommand.(type) {
		case internalParser.ATTestCommand:
			atCommands.Commands = append(atCommands.Commands, ATTestCommand{
				CommandName: cmd.CommandName,
			})
		case internalParser.ATReadCommand:
			atCommands.Commands = append(atCommands.Commands, ATReadCommand{
				CommandName: cmd.CommandName,
			})
		case internalParser.ATExecuteCommand:
			atCommands.Commands = append(atCommands.Commands, ATExecuteCommand{
				CommandName: cmd.CommandName,
			})
		case internalParser.ATSetCommand:
			atSetCommand := ATSetCommand{
				CommandName: cmd.CommandName,
			}
			for _, value := range cmd.Values {
				atSetCommand.Values = append(atSetCommand.Values, value.String())
			}
			atCommands.Commands = append(atCommands.Commands, atSetCommand)
		}
	}
	return &atCommands, nil
}
