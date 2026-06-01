# 📜 yasl

> A collection of programs written in YASL, a postfix, multi-stack esoteric language — a 42 rush where the language itself was the puzzle.

![Language](https://img.shields.io/badge/language-YASL-lightgrey)
![School](https://img.shields.io/badge/school-42-black)
![Status](https://img.shields.io/badge/status-finished-brightgreen)

---

## Description

`yasl_linux_x64` is an interpreter provided by the subject; the actual work was writing programs
against it. YASL has no variables in the usual sense — everything happens by pushing, popping and
swapping values across several numbered stacks (`!`, `^`, `#`) with postfix operators and
`?( ) : ( )` / `@( )` for branching and looping.

Programs in this repo, roughly from simplest to most involved:

| Program | What it does |
| --- | --- |
| `yasl_hw` | Hello world |
| `yasl_aff_param` | Prints its arguments |
| `yasl_fact` | Factorial of a positive integer, with input validation |
| `yasl_repeat` | Repeats a value N times |
| `yasl_split` | Splits a string on a one-character separator |
| `yasl_do` | Defines and calls named blocks, evaluates code passed on the stack |
| `yasl_interactive` | A minimal read-eval loop |
| `mandel` | Renders the Mandelbrot set |
| `display_b64` | Decodes base64-encoded RGB pixel data and prints it as ANSI-coloured blocks in the terminal |

## Run

```sh
./yasl_linux_x64 yasl_fact 5
./yasl_linux_x64 mandel
cat img1.rgb.b64 | ./yasl_linux_x64 display_b64
```
