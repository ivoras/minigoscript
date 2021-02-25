package minigoscript

import (
	"strings"
)

type ScriptAction struct {
	Action string
	Args   []ScriptToken
}

func (a ScriptAction) String() string {
	b := strings.Builder{}
	b.WriteString(a.Action)
	b.WriteString(" ")
	for _, arg := range a.Args {
		b.WriteString(arg.String())
		b.WriteString(" ")
	}
	return b.String()
}
