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
	for _, arg := range a.Args {
		b.WriteString(" ")
		b.WriteString(arg.String())
	}
	return b.String()
}
