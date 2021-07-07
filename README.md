# Go Tutorial for Beginners
In this tutorial, you will learn basic go programming.

## [1. Geting Start](./Tutorial_01)

### Install
Refer to [Download and install Go](https://golang.org/doc/install)

### Mode init
The go.mod file in the root directory contains module infromation for managing packages dependencies in your project. Run the ***go mod init command*** to create a go.mod file.
```bash
go mode init github.com/roguefactories/golang_beginner
```

### Run your code
```bash
go run hello.go
```

## [2. The init function](./Tutorial_02)

Each source file can define its own ***init*** function to initialize its state. The ***init*** funcitons are executed only one after the other in the order of their imports.

```go
func init() {
	fmt.Println("This is init in main.")
}
```

The <span style="color:lightblue">*init.go*</span> imports pkg1 and pkg2 and calls Pkg1() and Pkg1_1() in the pkg1 and Pkg2() in the pkg2. The <span style="color:lightblue">*pkg1.go*</span> imports pkg3 and calls Pkg3() in the pkg3. The <span style="color:lightblue">*pkg2.go*</span> imports pkg3 and calls Pkg3() in the pkg3.

```bash
This is init in pkg3.
This is init in pkg1.
This is init in pkg1_1.
This is init in pkg2.
This is init in main.
This is main.
This is pkg1.
This is pkg3.
This is pkg1_1.
This is pkg2.
This is pkg3.
```

## [3. Scope](./Tutorial_03)

### Local Variables
Local variables are declared inside a function or block and they can be used within the function or block.

### Global Variables
Global variables are defined outside of a function and they can be accessed by any functions. Howevver, if there is a local variable with the same name as a global variable, the local variable is accessed by its name.

```bash
1 + 2 = 3
x = 10 and y = 20 in main()
x = 1 and y = 20 in sum()
1 + 20 = 21
```

## [4. Format](./Tutorial_04)

### Disply an object
|fmt|Description|
|:---:|:---|
|%v|Default format for an object|
|%+v|Show field names|
|%#v|Show a Go-syntax representation of the value|
|%T|Show a Go-syntax representation of the type of the value|

```bash
{RogueFactories 23 123.456789 true}
{Name:RogueFactories Age:23 Height:123.456789 Trainable:true}
main.Dog{Name:"RogueFactories", Age:23, Height:123.456789, Trainable:true}
main.Dog
```

### Boolean
|fmt|Description|
|:---:|:---|
|%t|Show the word true or false|

```bash
true
```

### Integer
|fmt|Description|
|:---:|:---|
|%b|Show binary number|
|%c|Show the character represented by the corresponding Unicode code point|
|%d|Show decimal number|
|%o|Show octal number|
|%O|Show octal number with 0o prefix|
|%q|Show a single-quoted character literal safely escaped with Go syntax.|
|%x|Show Hexadecimal in lower-case letters for a-f|
|%X|Show Hexadecimal in upper-case letters for A-F|
|%U|Show Unicode format: U+1234; same as "U+%04X"|

```bash
10111

23
27
0o27
'\x17'
17
17
U+0017
```

### Floating-point
|fmt|Description|
|:---:|:---|
|%b|Decimalless scientific notation with exponent a power of two, 	in the manner of strconv.FormatFloat with the 'b' format, 	e.g. -123456p-78|
|%e|Scientific notation, e.g. -1.234456e+78|
|%E|Scientific notation, e.g. -1.234456E+78|
|%f|Decimal point but no exponent, e.g. 123.456|
|%F|Synonym for %f|
|%g|%e for large exponents, %f otherwise. Precision is discussed below.|
|%G|%E for large exponents, %F otherwise|

```bash
8687499202136843p-46
1.234568e+02
1.234568E+02
1234567.890000
123456.789000
1.23456789e+06
123456.789
```

### String
|fmt|Description|
|:---:|:---|
|%s|The uninterpreted bytes of the string or slice|
|%q|A double-quoted string safely escaped with Go syntax|

```bash
RogueFactories
"RogueFactories"
```
