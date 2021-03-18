package minigoscript

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// ScriptParserError is the error which is ultimately wrapped in the returned errors.
var ScriptParserError = errors.New("Error parsing script")

const alphaString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphanumericString = alphaString + "0123456789"
const identifierString = alphanumericString + "_"
const numberString = "0123456789.-"
const quoteString = "\"'"
const operatorString = "=+-*/"

type ScriptParser struct {
}

// DefaultParser is the default mini script parser. Use it for parsing mini script code.
var DefaultParser ScriptParser

// Parse accepts a string containing the newline-delimited mini script code, and parses
// this code into a slice of `ScriptAction`s.
func (p *ScriptParser) Parse(inputStr string) (actions []ScriptAction, err error) {
	for lineCount, rawLine := range strings.Split(inputStr, "\n") {
		line := p.gobbleWhiteSpace(rawLine)
		if len(line) == 0 {
			continue
		}
		l, s := p.gobbleIdentifier(line)
		if l == 0 {
			err = fmt.Errorf("%w: Cannot find action in line %d: '%s'", ScriptParserError, lineCount+1, rawLine)
			return
		}
		action := ScriptAction{
			Action: s,
			Args:   []ScriptToken{},
		}
		line = line[l:]

		for len(line) > 0 {
			line = p.gobbleWhiteSpace(line)
			if len(line) == 0 {
				break
			}

			l, b := p.gobbleBool(line)
			if l > 0 {
				action.Args = append(action.Args, newScriptTokenBool(b))
				line = line[l:]
				continue
			}

			l, n := p.gobbleNumber(line)
			if l > 0 {
				action.Args = append(action.Args, newScriptTokenNumber(n))
				line = line[l:]
				continue
			} else if l == -1 {
				err = fmt.Errorf("%w: error parsing number in line %d: '%s'", ScriptParserError, lineCount+1, line)
				return
			}

			l, s = p.gobbleIdentifier(line)
			if l > 0 {
				action.Args = append(action.Args, newScriptTokenIdentifier(s))
				line = line[l:]
				continue
			}

			l, s = p.gobbleString(line)
			if l > 0 {
				action.Args = append(action.Args, newScriptTokenString(s))
				line = line[l:]
				continue
			} else if l == -1 {
				err = fmt.Errorf("%w: error parsing string in line %d: '%s'", ScriptParserError, lineCount+1, line)
				return
			}

			l, s = p.gobbleOperator(line)
			if len(s) > 0 {
				action.Args = append(action.Args, newScriptTokenOperator(s))
				line = line[l:]
				continue
			}

			err = fmt.Errorf("%w: Unknown token at line %d: '%s'", ScriptParserError, lineCount+1, line)
			return
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
	for i < len(s) && strings.IndexByte(identifierString, s[i]) != -1 {
		i++
	}
	return i, s[0:i]
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

func (p *ScriptParser) gobbleNumber(s string) (i int, f float32) {
	if len(s) == 0 {
		return 0, 0
	}
	for i < len(s) && strings.IndexByte(numberString, s[i]) != -1 {
		i++
	}
	if i == 0 {
		return 0, 0
	}
	f64, err := strconv.ParseFloat(s[0:i], 32)
	if err != nil {
		log.Println("Error parsing number:", s[0:i], err)
		return -1, 0
	}
	f = float32(f64)
	return
}

func (p *ScriptParser) gobbleString(s string) (i int, outString string) {
	if len(s) == 0 {
		return 0, ""
	}
	q := s[0]
	if strings.IndexByte(quoteString, q) == -1 {
		return 0, ""
	}
	s = s[1:]

	inEscape := false
	sb := strings.Builder{}

	for i < len(s) {
		ch := s[i]
		if ch == '\\' {
			inEscape = true
			i++
			continue
		}
		if inEscape {
			if ch == q {
				sb.WriteByte(ch)
				inEscape = false
			} else if ch == 'n' {
				sb.WriteByte('\n')
				inEscape = false
			} else if ch == 't' {
				sb.WriteByte('\t')
				inEscape = false
			} else {
				return -1, ""
			}
			i++
			continue
		}
		if ch == q {
			i += 2
			break
		}
		sb.WriteByte(ch)
		i++
	}

	outString = sb.String()
	return
}

func (p *ScriptParser) gobbleOperator(s string) (int, string) {
	i := 0
	for i < len(s) && strings.IndexByte(operatorString, s[i]) != -1 {
		i++
	}
	return i, s[0:i]
}
