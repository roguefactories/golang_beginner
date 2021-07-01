# Go Tutorial for Beginner
In this tutorial, you will learn basic go programming.

## [1. Geting Start](https://github.com/roguefactories/golang_beginner/tree/main/01_Getting_Start)

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

## [2. The init function](https://github.com/roguefactories/golang_beginner/tree/main/02_The_Init_Function)

Each source file can define its own ***init*** function to initialize its state. The ***init*** funcitons are executed one after the other in the order of their imports.

```go
func init() {
	fmt.Println("This is init in main.")
}
```
