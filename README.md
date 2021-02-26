# A very simple scripting language parser in Go

This is an ad-hoc parser for a small scripting language. It's not even a complete language - it only deals with tokenisation and a bit of syntax, you need to provide all the semantics yourself. See the example section below for an idea what this does.

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

```
func TestParserLang(t *testing.T) {
	actions, err := minigoscript.DefaultParser.Parse(`
	let a = 1
	let b = true
	let c = 'Hello'
	print c
	print "World"
	print c "World"
	`)

	if err != nil {
		t.Error(err)
		return
	}
	symbolMap := map[string]interface{}{}

	for i, a := range actions {
		if a.Action == "let" {
			if len(a.Args) != 3 {
				t.Error("Expecting exactly 3 arguments for 'let' in line", i)
				continue
			}
			if !a.Args[0].IsIdentifier() {
				t.Error("Expecting identifier in line", i)
				continue
			}
			if !a.Args[1].IsOperator() || a.Args[1].MustGetOperator() != "=" {
				t.Error("Expecting operator = in line", i)
				continue
			}
			symbolMap[a.Args[0].MustGetIdentifier()] = a.Args[2].Value()
		} else if a.Action == "print" {
			for _, a := range a.Args {
				if a.IsIdentifier() {
					fmt.Print(symbolMap[a.MustGetIdentifier()])
				} else {
					fmt.Print(a.Value())
				}
				fmt.Print(" ")
			}
			fmt.Println()
		}
	}
}
```

## Notes

* The design of this parser will not win any awards in efficiency or performance, you really shouldn't use it for anything complex.
* This same approach is implemented for .Net in C# at https://github.com/ivoras/minidotnetscript
* This module has no external dependencies

A simple benchmark on an input script of 100 bytes:

```
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkParser-4         327212              3295 ns/op
```

I.e. it takes about 3us to parse that string on this hardware.
