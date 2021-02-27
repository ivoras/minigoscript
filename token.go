package minigoscript

import "fmt"

// ScriptTokenType is the type identifier for parsed tokens
type ScriptTokenType int32

const (
	TokenTypeIdentifier ScriptTokenType = iota
	TokenTypeString
	TokenTypeBool
	TokenTypeNumber
	TokenTypeOperator
)

// ScriptToken represents a single token from the parsed script
type ScriptToken struct {
	Type ScriptTokenType

	s string
	f float32
	b bool
}

// IsString returns true if the token is a string
func (v ScriptToken) IsString() bool {
	return v.Type == TokenTypeString
}

// GetString returns the string content of the token, if the token's type is `TokenTypeString`
// or returns an error.
func (v ScriptToken) GetString() (s string, err error) {
	if v.Type == TokenTypeString {
		s = v.s
	} else {
		err = fmt.Errorf("%w: Not a string (%d)", ScriptParserError, v.Type)
	}
	return
}

// MustGetString returns the string content of the token, if the token's type is `TokenTypeString`
// or an empty string otherwise.
func (v ScriptToken) MustGetString() (s string) {
	if v.Type == TokenTypeString {
		s = v.s
	}
	return
}

func (v ScriptToken) IsIdentifier() bool {
	return v.Type == TokenTypeIdentifier
}

func (v ScriptToken) GetIdentifier() (s string, err error) {
	if v.Type == TokenTypeIdentifier {
		s = v.s
	} else {
		err = fmt.Errorf("%w: Not an identifier (%d)", ScriptParserError, v.Type)
	}
	return
}

func (v ScriptToken) MustGetIdentifier() (s string) {
	if v.Type == TokenTypeIdentifier {
		s = v.s
	}
	return
}

func (v ScriptToken) IsBool() bool {
	return v.Type == TokenTypeBool
}

func (v ScriptToken) GetBool() (b bool, err error) {
	if v.Type == TokenTypeBool {
		b = v.b
	} else {
		err = fmt.Errorf("%w: Not a bool (%d)", ScriptParserError, v.Type)
	}
	return
}

func (v ScriptToken) MustGetBool() bool {
	if v.Type == TokenTypeBool {
		return v.b
	}
	return false
}

func (v ScriptToken) IsNumber() bool {
	return v.Type == TokenTypeNumber
}

func (v ScriptToken) GetNumber() (f float32, err error) {
	if v.Type == TokenTypeNumber {
		f = v.f
	} else {
		err = fmt.Errorf("%w: Not a number (%d)", ScriptParserError, v.Type)
	}
	return
}

func (v ScriptToken) MustGetNumber() float32 {
	if v.Type == TokenTypeNumber {
		return v.f
	}
	return 0
}

func (v ScriptToken) IsOperator() bool {
	return v.Type == TokenTypeOperator
}

func (v ScriptToken) GetOperator() (s string, err error) {
	if v.Type == TokenTypeOperator {
		s = v.s
	} else {
		err = fmt.Errorf("%w: Not an operator (%d)", ScriptParserError, v.Type)
	}
	return
}

func (v ScriptToken) MustGetOperator() (s string) {
	if v.Type == TokenTypeOperator {
		s = v.s
	}
	return
}

// String Returns a human-readable representation of the token.
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

// Value returns the token's value, either a string, float32 or a bool, as an empty interface.
func (v ScriptToken) Value() interface{} {
	switch v.Type {
	case TokenTypeBool:
		return v.b
	case TokenTypeIdentifier:
		return v.s
	case TokenTypeNumber:
		return v.f
	case TokenTypeOperator:
		return v.s
	case TokenTypeString:
		return v.s
	}
	return nil
}

func newScriptTokenBool(b bool) ScriptToken {
	return ScriptToken{
		Type: TokenTypeBool,
		b:    b,
	}
}

func newScriptTokenIdentifier(s string) ScriptToken {
	return ScriptToken{
		Type: TokenTypeIdentifier,
		s:    s,
	}
}

func newScriptTokenNumber(f float32) ScriptToken {
	return ScriptToken{
		Type: TokenTypeNumber,
		f:    f,
	}
}

func newScriptTokenOperator(s string) ScriptToken {
	return ScriptToken{
		Type: TokenTypeOperator,
		s:    s,
	}
}

func newScriptTokenString(s string) ScriptToken {
	return ScriptToken{
		Type: TokenTypeString,
		s:    s,
	}
}
