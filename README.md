# Meth

Meth is a lightweight integer math lexer, parser, and evaluator, designed to provide simple mathematical operations. It is equipped with a Read-Eval-Print Loop (REPL) interface for convenient usage.

## Installation and Usage

To compile and run the project, follow these steps:

Obs: you need to have Go installed!

1. Clone the repository:

```bash
git clone https://github.com/AndreyArthur/meth.git
cd meth
go build ./main.go
./main
```

Alternatively, you can directly run the project using:

```bash
go run ./main.go
```

## Modes

Meth supports three modes: lex, parse, and eval. The default mode is eval.

## REPL Interface

Upon running the executable, you will enter the Meth language REPL.

```bash
$ ./main
Welcome to the Meth language REPL!
Select a mode by typing "mode lex | parse | eval" (default: eval).
>> 9 + 9
18
>> 80 = 8 ^ 2 + 16
true
>> 5 * 3 > 15
false
>> mode lex
Switched to the lex mode.
>> 4 + 4 / 5
Integer "4"
Plus "+"
Integer "4"
Slash "/"
Integer "5"
Eof ""
>> mode parse
Switched to the parse mode.
>> 5 + 5 * 6
(5 + (5 * 6))
>> mode eval
Switched to the eval mode.
>> 5 >= 2^2  
true
```
## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
