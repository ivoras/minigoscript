# A very simple scripting language parser in Go

This is an ad-hoc parser for a small scripting language. It's not even a complete language - it only deals with syntax, you need to provide the semantics yourself.

The language looks like this:

```
let a = 42
let b = "hello, world"
let c = true
perform_action a b "some argument"
do_something_else
```

The language is newline-delimited, each action is a separate line. Each action must start with an identifier (you can think of it as a command) which can be followed by arbitrary arguments.

At the very simplest case, it's not a Turing complete language. The output of the parser is a list of parsed actions (aka commands, statements),
and the caller will walk through the list and perform whatever action is needed. That includes almost everything relating to the syntax, maintaining the symbol table,
validating the arguments are correct, etc.

The parser only knows about 3 data types:

* String
* Number
* Bool

and there are two special token types:

* Operator
* Identifier

## Examples

A real-world-ish example might be:


## Notes

* The design of this parser will not win any awards in efficiency or performance, you really shouldn't use it for anything complex.
* This same approach is implemented for .Net in C# at https://github.com/ivoras/minidotnetscript
