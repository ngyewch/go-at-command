package go_at_command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	parser, err := NewParser()
	if err != nil {
		t.Fatal(err)
		return
	}

	_, err = parser.Parse("boo")
	if err == nil {
		t.Fatal("expected error")
		return
	}

	{
		atCmds, err := parser.Parse(`AT`)
		if err != nil {
			t.Fatal(err)
			return
		}
		assert.Equal(t, 0, len(atCmds.Commands))
	}

	{
		atCmds, err := parser.Parse(`AT+CGMI?;+CGMI=?;+CGMI;+CGMI=abc,123,"boo,hoo",a\"b,"c\"d"`)
		if err != nil {
			t.Fatal(err)
			return
		}
		assert.Equal(t, 4, len(atCmds.Commands))

		assertReadCommand(t, atCmds.Commands[0], `+CGMI`)
		assertTestCommand(t, atCmds.Commands[1], `+CGMI`)
		assertExecuteCommand(t, atCmds.Commands[2], `+CGMI`)
		assertSetCommand(t, atCmds.Commands[3], `+CGMI`, `abc`, `123`, `boo,hoo`, `a"b`, `c"d`)
	}

	{
		atCmds, err := parser.Parse(`AT+CMEE=2;+CMGF=1;+CMGD=,4;+CNMP=38;+CMNB=1;+CGDCONT=1,"IP","super"`)
		if err != nil {
			t.Fatal(err)
			return
		}
		assert.Equal(t, 6, len(atCmds.Commands))

		assertSetCommand(t, atCmds.Commands[0], `+CMEE`, `2`)
		assertSetCommand(t, atCmds.Commands[1], `+CMGF`, `1`)
		assertSetCommand(t, atCmds.Commands[2], `+CMGD`, ``, `4`)
		assertSetCommand(t, atCmds.Commands[3], `+CNMP`, `38`)
		assertSetCommand(t, atCmds.Commands[4], `+CMNB`, `1`)
		assertSetCommand(t, atCmds.Commands[5], `+CGDCONT`, `1`, `IP`, `super`)
	}

	{
		atCmds, err := parser.Parse(`AT#SELINT=2`)
		if err != nil {
			t.Fatal(err)
			return
		}
		assert.Equal(t, 1, len(atCmds.Commands))

		assertSetCommand(t, atCmds.Commands[0], `#SELINT`, `2`)
	}

	{
		atCmds, err := parser.Parse(`AT&K0`)
		if err != nil {
			t.Fatal(err)
			return
		}
		assert.Equal(t, 1, len(atCmds.Commands))

		assertExecuteCommand(t, atCmds.Commands[0], `&K0`)
	}

	{
		atCmds, err := parser.Parse(`AT%Q`)
		if err != nil {
			t.Fatal(err)
			return
		}
		assert.Equal(t, 1, len(atCmds.Commands))

		assertExecuteCommand(t, atCmds.Commands[0], `%Q`)
	}

	{
		atCmds, err := parser.Parse(`AT\R1`)
		if err != nil {
			t.Fatal(err)
			return
		}
		assert.Equal(t, 1, len(atCmds.Commands))

		assertExecuteCommand(t, atCmds.Commands[0], `\R1`)
	}
}

func assertTestCommand(t *testing.T, cmd ATCommand, commandName string) {
	atCmd, ok := cmd.(ATTestCommand)
	assert.Truef(t, ok, "expected ATTestCommand")
	if !ok {
		return
	}
	assert.Equalf(t, commandName, atCmd.CommandName, "command name mismatch")
	assert.Equal(t, commandName+"=?", atCmd.String())
}

func assertReadCommand(t *testing.T, cmd ATCommand, commandName string) {
	atCmd, ok := cmd.(ATReadCommand)
	assert.Truef(t, ok, "expected ATReadCommand")
	if !ok {
		return
	}
	assert.Equalf(t, commandName, atCmd.CommandName, "command name mismatch")
	assert.Equal(t, commandName+"?", atCmd.String())
}

func assertExecuteCommand(t *testing.T, cmd ATCommand, commandName string) {
	atCmd, ok := cmd.(ATExecuteCommand)
	assert.Truef(t, ok, "expected ATExecuteCommand")
	if !ok {
		return
	}
	assert.Equalf(t, commandName, atCmd.CommandName, "command name mismatch")
	assert.Equal(t, commandName, atCmd.String())
}

func assertSetCommand(t *testing.T, cmd ATCommand, commandName string, values ...string) {
	atCmd, ok := cmd.(ATSetCommand)
	assert.Truef(t, ok, "expected ATSetCommand")
	if !ok {
		return
	}
	assert.Equalf(t, commandName, atCmd.CommandName, "command name mismatch")
	assert.Equal(t, len(values), len(atCmd.Values), "argument count mismatch")
	maxIndex := min(len(values), len(atCmd.Values))
	for i := range maxIndex {
		assert.Equalf(t, values[i], atCmd.Values[i], "mismatch at argument #%d", i)
	}
}
