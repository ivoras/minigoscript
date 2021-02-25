package minigoscript

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ScriptParserError = errors.New("Error parsing script")

const alphaString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphanumericString = alphaString + "0123456789"
const identifierString = alphanumericString + "_"
const numberString = "0123456789."
const quoteString = "\"'"
const operatorString = "=+-*/"

type ScriptParser struct {
}

func (p *ScriptParser) Parse(inputStr string) (actions []ScriptAction, err error) {
	for lineCount, rawLine := range strings.Split(inputStr, " ") {
		line := p.gobbleWhiteSpace(rawLine)
		l, s := p.gobbleIdentifier(line)
		if l == 0 {
			err = fmt.Errorf("%w: Cannot find action in line %d: '%s'", ScriptParserError, lineCount+1, rawLine)
			return
		}
		action := ScriptAction{
			Action: s,
			Args:   []ScriptToken{},
		}

		for len(line) > 0 {
			line = p.gobbleWhiteSpace(line)
			if len(line) == 0 {
				break
			}

			l, b := p.gobbleBool(line)
			if l > 0 {
				action.Args = append(action.Args, NewScriptTokenBool(b))
				line = line[l:]
				continue
			}

			l, n := p.gobbleNumber(line)
			if l > 0 {
				action.Args = append(action.Args, NewScriptTokenNumber(n))
				line = line[l:]
				continue
			} else if l == -1 {
				err = fmt.Errorf("%w: error parsing number in line %d: '%s'", ScriptParserError, lineCount+1, line)
				return
			}

			l, s = p.gobbleIdentifier(line)
			if l > 0 {
				action.Args = append(action.Args, NewScriptTokenIdentifier(s))
				line = line[l:]
				continue
			}

			l, s = p.gobbleString(line)

		}

		actions = append(actions, action)
	}
	return
}

func (p *ScriptParser) gobbleWhiteSpace(s string) string {
	i := 0
	for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
		i++
	}
	return s[i:]
}

func (p *ScriptParser) gobbleIdentifier(s string) (int, string) {
	i := 0
	for i < len(s) && strings.IndexByte(identifierString, byte(s[i])) != -1 {
		i++
	}
	return i, s[i:]
}

func (p *ScriptParser) gobbleBool(s string) (l int, b bool) {
	if len(s) == 0 {
		return 0, false
	}
	l, tok := p.gobbleIdentifier(s)
	if l != 0 {
		if tok == "true" {
			return l, true
		} else if tok == "false" {
			return l, false
		}
	}
	return 0, false
}

func (p *ScriptParser) gobbleNumber(s string) (l int, f float32) {
	if len(s) == 0 {
		return 0, 0
	}
	i := 0
	for i < len(s) && strings.IndexByte(numberString, byte(s[i])) != -1 {
		i++
	}
	num := s[0:i]
	f64, err := strconv.ParseFloat(num, 32)
	if err != nil {
		return -1, 0
	}
	f = float32(f64)
	return
}

func (p *ScriptParser) gobbleString(s string) (l int, outString string) {
	return
}
