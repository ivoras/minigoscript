package minigoscript

import (
	"errors"
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
