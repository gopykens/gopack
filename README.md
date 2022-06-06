## The Go+ language for engineering, STEM education, and data science

[![Build Status](https://github.com/goplus/gop/actions/workflows/go.yml/badge.svg)](https://github.com/goplus/gop/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/goplus/gop)](https://goreportcard.com/report/github.com/goplus/gop)
[![Coverage Status](https://codecov.io/gh/goplus/gop/branch/main/graph/badge.svg)](https://codecov.io/gh/goplus/gop)
[![GitHub release](https://img.shields.io/github/v/tag/goplus/gop.svg?label=release)](https://github.com/goplus/gop/releases)
[![Tutorials](https://img.shields.io/badge/tutorial-Go+-blue.svg)](https://tutorial.goplus.org/)
[![Playground](https://img.shields.io/badge/play-Go+-blue.svg)](https://play.goplus.org/)
[![VSCode](https://img.shields.io/badge/vscode-Go+-teal.svg)](https://github.com/gopcode/vscode-goplus)
[![Interpreter](https://img.shields.io/badge/interpreter-iGo+-seagreen.svg)](https://github.com/goplus/igop)

* **For engineering**: working in the simplest language that can be mastered by children.
* **For STEM education**: studying an engineering language that can be used for work in the future.
* **For data science**: communicating with engineers in the same language.

## Go vs. Go+

What's the difference between Go and Go+?

The goal of Go is to provide a way to easily build simple, reliable, and efficient software. It wants to be a better C.

The goal of Go+ is to achieve a natural fusion of engineering and low-code. It wants to be a better Python.

The essence of low-code is non-professionals oriented. Today's mainstream programming language designers still take non-professionals as a niche group. But Go+ take this group as vital, and they are a vital breakthrough in today's programming language revolution.

The Python language appeared very early (1991), but its design idea is very far-sighted, and it is a rare non-professionals oriented programming language so far. There are quite a few languages that look relatively concise, such as Ruby and CoffeeScript, but they have many language magics so that people feel they are flexible and powerful, but not easy to master. They cannot be regarded as non-professionals oriented.

Python is very popular today, but it's hard to call it an engineering language. Most people just take it as a tool for rapid prototyping.

Is it possible to minimize the complexity of the project and present it in a low-code form? Go+ hopes to explore a new path at this point.

## How to install

For now, we suggest you install Go+ from source code.

Note: Requires go1.16 or later

```bash
git clone https://github.com/goplus/gop.git
cd gop

# On mac/linux run:
./all.bash
# On Windows run:
all.bat
```

Actually, `all.bash` and `all.bat` will use `go run cmd/make.go` underneath.


## Running in Go+ playground

If you don't want install Go+, you can write your Go+ programs in Go+ playground. This is the fastest way to experience Go+.

* Go+ playground based on Docker: https://play.goplus.org/
* Go+ playground based on GopherJS: https://jsplay.goplus.org/

And you can share your Go+ code with your friends.
Here is my `Hello world` program:
* https://play.goplus.org/p/AAh_gQAKAZR.


## Table of Contents

<table>
    <tr><td width=33% valign=top>

* [Hello world](#hello-world)
* [Running a project folder](#running-a-project-folder-with-several-files)
* [Comments](#comments)
* [Variables](#variables)
* [Go+ types](#go-types)
    * [Strings](#strings)
    * [Numbers](#numbers)
    * [Slices](#slices)
    * [Maps](#maps)
* [Module imports](#module-imports)

</td><td width=33% valign=top>

* [Statements & expressions](#statements--expressions)
    * [If](#if)
    * [For loop](#for-loop)
    * [List comprehension](#list-comprehension)
    * [Select data from a collection](#select-data-from-a-collection)
    * [Check if data exists in a collection](#check-if-data-exists-in-a-collection)

</td><td valign=top>

* [Functions](#functions)
    * [Returning multiple values](#returning-multiple-values)
    * [Variable number of arguments](#variable-number-of-arguments)
    * [Higher order functions](#higher-order-functions)
    * [Lambda expressions](#lambda-expressions)
* [Structs](#structs)
* [Error handling](#error-handling)
* [Unix shebang](#unix-shebang)

</td></tr>
</table>


## Hello World

```go
println "Hello world"
```

Save this snippet into a file named `hello.gop`. Now do: `gop run hello.gop`.

Congratulations - you just wrote and executed your first Go+ program!

You can compile a program without execution with `gop build hello.gop`.
See `gop help` for all supported commands.

[`println`](#println) is one of the few [built-in functions](#builtin-functions).
It prints the value passed to it to standard output.

See https://tutorial.goplus.org/hello-world for more details.


## Running a project folder with several files

Suppose you have a folder with several .gop files in it, and you want 
to compile them all into one program. Just do: `gop run .`.

Passing parameters also works, so you can do:
`gop run . --yourparams some_other_stuff`.

Your program can then use the CLI parameters like this:

```go
import "os"

println os.Args
```

## Comments

```go
# This is a single line comment.

// This is a single line comment.

/*
This is a multiline comment.
*/
```

## Variables

```go
name := "Bob"
age := 20
largeNumber := int128(1 << 65)
println name, age
println largeNumber
```

Variables are declared and initialized with `:=`.

The variable's type is inferred from the value on the right hand side.
To choose a different type, use type conversion:
the expression `T(v)` converts the value `v` to the
type `T`.

### Initialization vs. assignment

Note the (important) difference between `:=` and `=`.
`:=` is used for declaring and initializing, `=` is used for assigning.

```go failcompile
age = 21
```

This code will not compile, because the variable `age` is not declared.
All variables need to be declared in Go+.

```go
age := 21
```

The values of multiple variables can be changed in one line.
In this way, their values can be swapped without an intermediary variable.

```go
a, b := 0, 1
a, b = b, a
println a, b // 1, 0
```

## Go+ Types

### Primitive types

```go ignore
bool

int8    int16   int32   int    int64    int128
uint8   uint16  uint32  uint   uint64   uint128

uintptr // similar to C's size_t

byte // alias for uint8
rune // alias for int32, represents a Unicode code point

string

float32 float64

complex64 complex128

bigint bigrat

unsafe.Pointer // similar to C's void*

any // alias for Go's interface{}
```

### Strings

```go
name := "Bob"
println name.len  // 3
println name[0]   // 66
println name[1:3] // ob
println name[:2]  // Bo
println name[2:]  // b

// or using octal escape `\###` notation where `#` is an octal digit
println "\141a"   // aa

// Unicode can be specified directly as `\u####` where # is a hex digit
// and will be converted internally to its UTF-8 representation
println "\u2605"  // ★
```

String values are immutable. You cannot mutate elements:

```go failcompile
s := "hello 🌎"
s[0] = `H` // not allowed
```

Note that indexing a string will produce a `byte`, not a `rune` nor another `string`.

Strings can be easily converted to integers:

```go
s := "12"
a, err := s.int
b := s.int! // will panic if s isn't a valid integer
```

#### String operators

```go
name := "Bob"
bobby := name + "by" // + is used to concatenate strings
println bobby // Bobby

s := "Hello "
s += "world"
println s // Hello world
```

All operators in Go+ must have values of the same type on both sides. You cannot concatenate an
integer to a string:

```go failcompile
age := 10
println "age = " + age // not allowed
```

We have to either convert `age` to a `string`:

```go
age := 10
println "age = " + age.string
```

### Runes

A `rune` represents a single Unicode character and is an alias for `int32`.

```go
rocket := '🚀'
println rocket         // 128640
println string(rocket) // 🚀
```

### Numbers

```go
a := 123
```

This will assign the value of 123 to `a`. By default `a` will have the
type `int`.

You can also use hexadecimal, binary or octal notation for integer literals:

```go
a := 0x7B
b := 0b01111011
c := 0o173
```

All of these will be assigned the same value, 123. They will all have type
`int`, no matter what notation you used.

Go+ also supports writing numbers with `_` as separator:

```go
num := 1_000_000 // same as 1000000
three := 0b0_11 // same as 0b11
floatNum := 3_122.55 // same as 3122.55
hexa := 0xF_F // same as 255
oct := 0o17_3 // same as 0o173
```

If you want a different type of integer, you can use casting:

```go
a := int64(123)
b := uint8(12)
c := int128(12345)
```

Assigning floating point numbers works the same way:

```go
f1 := 1.0
f2 := float32(3.14)
```

If you do not specify the type explicitly, by default float literals will have the type of `float64`.

Float literals can also be declared as a power of ten:

```go
f0 := 42e1   // 420
f1 := 123e-2 // 1.23
f2 := 456e+2 // 45600
```

We introduce rational numbers as native Go+ types. We use suffix `r` to denote rational literals. For example, `1r << 200` means a big int whose value is equal to 2<sup>200</sup>.

```go
a := 1r << 200
b := bigint(1 << 200)
```

By default, `1r` will have the type of `bigint`.

And `4/5r` means the rational constant `4/5`.
It will have the type of `bigrat`.

```go
a := 4/5r
b := a - 1/3r + 3 * 1/2r
println a, b // 4/5 59/30
```

Casting rational numbers works the same way:

```go
a := 1r
b := bigrat(1r)
c := bigrat(1)
println a/3 // 0
println b/3 // 1/3
println c/3 // 1/3
```

#### Convert bool to number types

```go
println int(true)       // 1
println float64(true)   // 1
println complex64(true) // (1+0i)
```

### Slices

A slice is a collection of data elements of the same type. A slice literal is a
list of expressions surrounded by square brackets. An individual element can be
accessed using an *index* expression. Indexes start from `0`:

```go
nums := [1, 2, 3]
println nums      // [1 2 3]
println nums.len  // 3
println nums[0]   // 1
println nums[1:3] // [2 3]
println nums[:2]  // [1 2]
println nums[2:]  // [3]

nums[1] = 5
println nums // [1 5 3]
```

Type of a slice literal is infered automatically.

```go
a := [1, 2, 3]   // []int
b := [1, 2, 3.4] // []float64
c := ["Hi"]      // []string
d := ["Hi", 10]  // []any
d := []          // []any
```

And casting slices also works.

```go
a := []float64([1, 2, 3]) // []float64
```

### Maps

```go
a := {"Hello": 1, "xsw": 3}     // map[string]int
b := {"Hello": 1, "xsw": 3.4}   // map[string]float64
c := {"Hello": 1, "xsw": "Go+"} // map[string]any
d := {}                         // map[string]any
```

If a key is not found, a zero value is returned by default:

```go
a := {"Hello": 1, "xsw": 3}
c := {"Hello": 1, "xsw": "Go+"}
println a["bad_key"] // 0
println c["bad_key"] // <nil>
```

You can also check, if a key is present, and get its value.

```go
a := {"Hello": 1, "xsw": 3}
if v, ok := a["xsw"]; ok {
    println "its value is", v
}
```

## Module imports

For information about creating a module, see [Modules](#modules).

Modules can be imported using the `import` keyword:

```go
import "strings"

x := strings.NewReplacer("?", "!").Replace("Hello, world???")
println x // Hello, world!!!
```

### Module import aliasing

Any imported module name can be aliased:

```go
import strop "strings"

x := strop.NewReplacer("?", "!").Replace("Hello, world???")
println x // Hello, world!!!
```


## Statements & expressions


### If

`if` statements are pretty straightforward and similar to most other languages.
Unlike other C-like languages,
there are no parentheses surrounding the condition and the braces are always required.

```go
a := 10
b := 20
if a < b {
	println "a < b"
} else if a > b {
	println "a > b"
} else {
	println "a == b"
}
```

### For loop

Go+ has only one looping keyword: `for`, with several forms.

#### `for`/`<-`

This is the most common form. You can use it with an slice, map, numeric range or custom iterators.

For information about creating a custom iterators, see [Custom iterators](#custom-iterators).

##### Slice `for`

The `for value <- arr` form is used for going through elements of an slice.

```go
numbers := [1, 3, 5, 7, 11, 13, 17]
sum := 0
for x <- numbers {
    sum += x
}
println sum // 57
```

If an index is required, an alternative form `for index, value <- arr` can be used.

```go
names := ["Sam", "Peter"]
for i, name <- names {
    println i, name
    // 0 Sam
    // 1 Peter
}
```

##### Map `for`

```go
m := {"one": 1, "two": 2}
for key, val <- m {
	println key, val
	// one 1
	// two 2
}
for key, _ <- m {
	println key
	// one
	// two
}
for val <- m {
	println val
	// 1
	// 2
}
```

##### Range `for`

You can use `range expression` (`start:end:step`) in for loop.

```go
for i <- :5 {
    println i
    // 0
    // 1
    // 2
    // 3
    // 4
}
for i <- 1:5 {
    println i
    // 1
    // 2
    // 3
    // 4
}
for i <- 1:5:2 {
    println i
    // 1
    // 3
}
```

##### `for`/`<-`/`if`

All loops of `for`/`<-` form can have an optional `if` condition.

```go
numbers := [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
for num <- numbers if num%3 == 0 {
    println num
    // 0
    // 3
    // 6
    // 9
}

for num <- :10 if num%3 == 0 {
    println num
    // 0
    // 3
    // 6
    // 9
}
```

#### Condition `for`

```go
sum := 0
i := 1
for i <= 100 {
	sum += i
	i++
}
println sum // 5050
```

This form of the loop is similar to `while` loops in other languages.
The loop will stop iterating once the boolean condition evaluates to false.
Again, there are no parentheses surrounding the condition, and the braces are always required.

#### C `for`

```go
for i := 0; i < 10; i += 2 {
    // Don't print 6
    if i == 6 {
        continue
    }
    println i
    // 0
    // 2
    // 4
    // 8
}
```

Finally, there's the traditional C style `for` loop. It's safer than the `while` form
because with the latter it's easy to forget to update the counter and get
stuck in an infinite loop.


#### Bare `for`

```go
for {
    // ...
}
```

The condition can be omitted, resulting in an infinite loop. You can use `break` or `return` to end the loop.

### List comprehension

```go
a := [x*x for x <- [1, 3, 5, 7, 11]]
b := [x*x for x <- [1, 3, 5, 7, 11] if x > 3]
c := [i+v for i, v <- [1, 3, 5, 7, 11] if i%2 == 1]

arr := [1, 2, 3, 4, 5, 6]
d := [[a, b] for a <- arr if a < b for b <- arr if b > 2]

x := {x: i for i, x <- [1, 3, 5, 7, 11]}
y := {x: i for i, x <- [1, 3, 5, 7, 11] if i%2 == 1}
z := {v: k for k, v <- {1: "Hello", 3: "Hi", 5: "xsw", 7: "Go+"} if k > 3}
```

### Select data from a collection

```go
type student struct {
    name  string
    score int
}

students := [student{"Ken", 90}, student{"Jason", 80}, student{"Lily", 85}]

unknownScore, ok := {x.score for x <- students if x.name == "Unknown"}
jasonScore := {x.score for x <- students if x.name == "Jason"}

println unknownScore, ok // 0 false
println jasonScore // 80
```

### Check if data exists in a collection

```go
type student struct {
    name  string
    score int
}

students := [student{"Ken", 90}, student{"Jason", 80}, student{"Lily", 85}]

hasJason := {for x <- students if x.name == "Jason"} // is any student named Jason?
hasFailed := {for x <- students if x.score < 60}     // is any student failed?
```


## Functions

```go
func add(x int, y int) int {
	return x + y
}

println add(2, 3) // 5
```

### Returning multiple values

```go
func foo() (int, int) {
	return 2, 3
}

a, b := foo()
println a // 2
println b // 3
c, _ := foo() // ignore values using `_`
```

### Variable number of arguments

```go
func sum(a ...int) int {
    total := 0
	for x <- a {
		total += x
	}
	return total
}

println sum(2, 3, 5) // 10
```

Output parameters can have names.

```go
func sum(a ...int) (total int) {
	for x <- a {
		total += x
	}
	return // don't need return values if they are assigned
}

println sum(2, 3, 5) // 10
```

### Higher order functions

Functions can also be parameters.

```go
func square(x float64) float64 {
    return x*x
}

func transform(a []float64, f func(float64) float64) []float64 {
    return [f(x) for x <- a]
}

y := transform([1, 2, 3], square)
println y // [1 4 9]
```

### Lambda expressions

You also can use `lambda expression` to define a anonymous function.

```go
func transform(a []float64, f func(float64) float64) []float64 {
    return [f(x) for x <- a]
}

y := transform([1, 2, 3], x => x*x)
println y // [1 4 9]

z := transform([-3, 1, -5], x => {
    if x < 0 {
        return -x
    }
    return x
})
println z // [3 1 5]
```



## Structs

### For range of UDT

```go
type Foo struct {
}

// Gop_Enum(proc func(val ValType)) or:
// Gop_Enum(proc func(key KeyType, val ValType))
func (p *Foo) Gop_Enum(proc func(key int, val string)) {
    // ...
}

foo := &Foo{}
for k, v := range foo {
    println k, v
}

for k, v <- foo {
    println k, v
}

println {v: k for k, v <- foo}
```

**Note: you can't use break/continue or return statements in for range of udt.Gop_Enum(callback).**


### For range of UDT2

```go
type FooIter struct {
}

// (Iterator) Next() (val ValType, ok bool) or:
// (Iterator) Next() (key KeyType, val ValType, ok bool)
func (p *FooIter) Next() (key int, val string, ok bool) {
    // ...
}

type Foo struct {
}

// Gop_Enum() Iterator
func (p *Foo) Gop_Enum() *FooIter {
    // ...
}

foo := &Foo{}
for k, v := range foo {
    println k, v
}

for k, v <- foo {
    println k, v
}

println {v: k for k, v <- foo}
```

### Deduce struct type

```go
type Config struct {
    Dir   string
    Level int
}

func foo(conf *Config) {
    // ...
}

foo {Dir: "/foo/bar", Level: 1}
```

Here `foo {Dir: "/foo/bar", Level: 1}` is equivalent to `foo(&Config{Dir: "/foo/bar", Level: 1})`. However, you can't replace `foo(&Config{"/foo/bar", 1})` with `foo {"/foo/bar", 1}`, because it is confusing to consider `{"/foo/bar", 1}` as a struct literal.

You also can omit struct types in a return statement. For example:

```go
type Result struct {
    Text string
}

func foo() *Result {
    return {Text: "Hi, Go+"} // return &Result{Text: "Hi, Go+"}
}
```

### Overload operators

```go
import "math/big"

type MyBigInt struct {
    *big.Int
}

func Int(v *big.Int) MyBigInt {
    return MyBigInt{v}
}

func (a MyBigInt) + (b MyBigInt) MyBigInt { // binary operator
    return MyBigInt{new(big.Int).Add(a.Int, b.Int)}
}

func (a MyBigInt) += (b MyBigInt) {
    a.Int.Add(a.Int, b.Int)
}

func -(a MyBigInt) MyBigInt { // unary operator
    return MyBigInt{new(big.Int).Neg(a.Int)}
}

a := Int(1r)
a += Int(2r)
println a + Int(3r)
println -a
```

### Auto property

Let's see an example written in Go+:

```go
import "gop/ast/goptest"

doc := goptest.New(`... Go+ code ...`)!

println doc.Any().FuncDecl().Name()
```

In many languages, there is a concept named `property` who has `get` and `set` methods.

Suppose we have `get property`, the above example will be:

```go
import "gop/ast/goptest"

doc := goptest.New(`... Go+ code ...`)!

println doc.any.funcDecl.name
```

In Go+, we introduce a concept named `auto property`. It is a `get property`, but is implemented automatically. If we have a method named `Bar()`, then we will have a `get property` named `bar` at the same time.


## Error handling

We reinvent the error handling specification in Go+. We call them `ErrWrap expressions`:

```go
expr! // panic if err
expr? // return if err
expr?:defval // use defval if err
```

How to use them? Here is an example:

```go
import (
    "strconv"
)

func add(x, y string) (int, error) {
    return strconv.Atoi(x)? + strconv.Atoi(y)?, nil
}

func addSafe(x, y string) int {
    return strconv.Atoi(x)?:0 + strconv.Atoi(y)?:0
}

println `add("100", "23"):`, add("100", "23")!

sum, err := add("10", "abc")
println `add("10", "abc"):`, sum, err

println `addSafe("10", "abc"):`, addSafe("10", "abc")
```

The output of this example is:

```
add("100", "23"): 123
add("10", "abc"): 0 strconv.Atoi: parsing "abc": invalid syntax

===> errors stack:
main.add("10", "abc")
    /Users/xsw/tutorial/15-ErrWrap/err_wrap.gop:6 strconv.Atoi(y)?

addSafe("10", "abc"): 10
```

Compared to corresponding Go code, It is clear and more readable.

And the most interesting thing is, the return error contains the full error stack. When we got an error, it is very easy to position what the root cause is.

How these `ErrWrap expressions` work? See [Error Handling](https://github.com/goplus/gop/wiki/Error-Handling) for more information.


## Unix shebang

You can use Go+ programs as shell scripts now. For example:

```go
#!/usr/bin/env -S gop run

println "Hello, Go+"

println 1r << 129
println 1/3r + 2/7r*2

arr := [1, 3, 5, 7, 11, 13, 17, 19]
println arr
println [x*x for x <- arr, x > 3]

m := {"Hi": 1, "Go+": 2}
println m
println {v: k for k, v <- m}
println [k for k, _ <- m]
println [v for v <- m]
```

Go [20-Unix-Shebang/shebang](https://github.com/goplus/tutorial/blob/main/20-Unix-Shebang/shebang) to get the source code.


## Compatibility with Go

All Go features will be supported (including partially support `cgo`, see [below](#bytecode-vs-go-code)).

**All Go packages (even these packages use `cgo`) can be imported by Go+.**

```coffee
import (
    "fmt"
    "strings"
)

x := strings.NewReplacer("?", "!").Replace("hello, world???")
fmt.Println "x:", x
```

**And all Go+ packages can also be imported in Go programs. What you need to do is just using `gop` command instead of `go`.**

First, let's make a directory named `14-Using-goplus-in-Go`.

Then write a Go+ package named [foo](https://github.com/goplus/tutorial/tree/main/14-Using-goplus-in-Go/foo) in it:

```go
package foo

func ReverseMap(m map[string]int) map[int]string {
    return {v: k for k, v <- m}
}
```

Then use it in a Go package [14-Using-goplus-in-Go/gomain](https://github.com/goplus/tutorial/tree/main/14-Using-goplus-in-Go/gomain):

```go
package main

import (
    "fmt"

    "github.com/goplus/tutorial/14-Using-goplus-in-Go/foo"
)

func main() {
    rmap := foo.ReverseMap(map[string]int{"Hi": 1, "Hello": 2})
    fmt.Println(rmap)
}
```

How to build this example? You can use:

```bash
gop install -v ./...
```

Go [github.com/goplus/tutorial/14-Using-goplus-in-Go](https://github.com/goplus/tutorial/tree/main/14-Using-goplus-in-Go) to get the source code.


## Bytecode vs. Go code

Go+ supports bytecode backend and Go code generation.

When we use `gop` command, it generates Go code to covert Go+ package into Go packages.

```bash
gop run     # Run a Go+ program
gop install # Build Go+ files and install target to GOBIN
gop build   # Build Go+ files
gop test    # Test Go+ packages
gop fmt     # Format Go+ packages
gop clean   # Clean all Go+ auto generated files
gop go      # Convert Go+ packages into Go packages
```

When we use [`igop`](https://github.com/goplus/igop) command, it generates bytecode to execute.

```bash
igop  # Run a Go+ program
```

In bytecode mode, Go+ doesn't support `cgo`. However, in Go-code-generation mode, Go+ fully supports `cgo`.


## IDE Plugins

* vscode: https://github.com/gopcode/vscode-goplus


## Tutorials

* https://tutorial.goplus.org/


## Contributing

The Go+ project welcomes all contributors. We appreciate your help!

Here are [list of Go+ Contributors](https://github.com/goplus/gop/wiki/Goplus-Contributors). We award an email account (XXX@goplus.org) for every contributor. And we suggest you commit code by using this email account:

```bash
git config --global user.email XXX@goplus.org
```

If you did this, remember to add your `XXX@goplus.org` email to https://github.com/settings/emails.

What does `a contributor to Go+` mean? You must meet one of the following conditions:

* At least one pull request of a full-featured implemention.
* At least three pull requests of feature enhancements.
* At least ten pull requests of any kind issues.

Where can you start?

* [![Issues](https://img.shields.io/badge/ISSUEs-Go+-blue.svg)](https://github.com/goplus/gop/issues)
* [![Issues](https://img.shields.io/badge/ISSUEs-NumGo+-blue.svg)](https://github.com/numgoplus/ng/issues)
* [![Issues](https://img.shields.io/badge/ISSUEs-PandasGo+-blue.svg)](https://github.com/goplus/pandas/issues)
* [![Issues](https://img.shields.io/badge/ISSUEs-vscode%20Go+-blue.svg)](https://github.com/gopcode/vscode-goplus/issues)
* [![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/goplus/gop)](https://www.tickgit.com/browse?repo=github.com/goplus/gop)
