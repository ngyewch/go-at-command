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
		atCmds, err := parser.Parse("AT")
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

		atCmd0, ok := atCmds.Commands[0].(ATReadCommand)
		assert.True(t, ok)
		assert.Equal(t, "+CGMI?", atCmd0.String())

		atCmd1, ok := atCmds.Commands[1].(ATTestCommand)
		assert.True(t, ok)
		assert.Equal(t, "+CGMI=?", atCmd1.String())

		atCmd2, ok := atCmds.Commands[2].(ATExecuteCommand)
		assert.True(t, ok)
		assert.Equal(t, "+CGMI", atCmd2.String())

		atCmd3, ok := atCmds.Commands[3].(ATSetCommand)
		assert.True(t, ok)
		assert.Equal(t, 5, len(atCmd3.Values))
		assert.Equal(t, "abc", atCmd3.Values[0])
		assert.Equal(t, "123", atCmd3.Values[1])
		assert.Equal(t, "boo,hoo", atCmd3.Values[2])
		assert.Equal(t, `a"b`, atCmd3.Values[3])
		assert.Equal(t, `c"d`, atCmd3.Values[4])
		assert.Equal(t, `+CGMI=abc,123,"boo,hoo","a\"b","c\"d"`, atCmd3.String()) // HACK

		assert.Equal(t, `AT+CGMI?;+CGMI=?;+CGMI;+CGMI=abc,123,"boo,hoo","a\"b","c\"d"`, atCmds.String()) // HACK
	}
}
