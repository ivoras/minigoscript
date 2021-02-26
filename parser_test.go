package minigoscript_test

import (
	"fmt"
	"testing"

	"github.com/ivoras/minigoscript"
)

func TestParser(t *testing.T) {
	actions, err := minigoscript.DefaultParser.Parse(`
	let a = 1
	let b = true
	let c = 'hello'
	`)

	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range actions {
		fmt.Println(a.Action, a.Args)
	}
}
