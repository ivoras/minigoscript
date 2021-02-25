package minigoscript

import "fmt"

type ScriptTokenType int32

const (
	TokenTypeIdentifier ScriptTokenType = iota
	TokenTypeString
	TokenTypeBool
	TokenTypeNumber
	TokenTypeOperator
)

type ScriptToken struct {
	Type ScriptTokenType

	s string
	f float32
	b bool
}

func (v ScriptToken) GetString() (s string, err error) {
	if v.Type == TokenTypeString {
		s = v.s
	} else {
		err = fmt.Errorf("%w: Not a string (%d)", ScriptParserError, v.Type)
	}
	return
}

func (v ScriptToken) GetIdentifier() (s string, err error) {
	if v.Type == TokenTypeIdentifier {
		s = v.s
	} else {
		err = fmt.Errorf("%w: Not an identifier (%d)", ScriptParserError, v.Type)
	}
	return
}

func (v ScriptToken) GetBool() (b bool, err error) {
	if v.Type == TokenTypeBool {
		b = v.b
	} else {
		err = fmt.Errorf("%w: Not a bool (%d)", ScriptParserError, v.Type)
	}
	return
}

func (v ScriptToken) GetNumber() (f float32, err error) {
	if v.Type == TokenTypeNumber {
		f = v.f
	} else {
		err = fmt.Errorf("%w: Not a number (%d)", ScriptParserError, v.Type)
	}
	return
}

func (v ScriptToken) GetOperator() (s string, err error) {
	if v.Type == TokenTypeOperator {
		s = v.s
	} else {
		err = fmt.Errorf("%w: Not an operator (%d)", ScriptParserError, v.Type)
	}
	return
}

func (v ScriptToken) String() string {
	switch v.Type {
	case TokenTypeBool:
		if v.b {
			return "<f true>"
		} else {
			return "<f false>"
		}
	case TokenTypeIdentifier:
		return fmt.Sprintf("<i %s>", v.s)
	case TokenTypeNumber:
		return fmt.Sprintf("<n %f>", v.f)
	case TokenTypeOperator:
		return fmt.Sprintf("<o %s>", v.s)
	case TokenTypeString:
		return fmt.Sprintf("<s %s>", v.s)
	}
	return "<unknown>"
}

func NewScriptTokenBool(b bool) ScriptToken {
	return ScriptToken{
		Type: TokenTypeBool,
		b:    b,
	}
}

func NewScriptTokenIdentifier(s string) ScriptToken {
	return ScriptToken{
		Type: TokenTypeIdentifier,
		s:    s,
	}
}

func NewScriptTokenNumber(f float32) ScriptToken {
	return ScriptToken{
		Type: TokenTypeNumber,
		f:    f,
	}
}

func NewScriptTokenOperator(s string) ScriptToken {
	return ScriptToken{
		Type: TokenTypeOperator,
		s:    s,
	}
}

func NewScriptTokenString(s string) ScriptToken {
	return ScriptToken{
		Type: TokenTypeString,
		s:    s,
	}
}
