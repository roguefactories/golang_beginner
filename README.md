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



