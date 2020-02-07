Go is an open source programming language that makes it easy to build simple, reliable and efficient software. The goals of the language and its accompanying tools were to be expressive, efficient in both compilation and execution and effective in writing reliable and robust programs.

Go's facilities for concurrency are new and efficient, and its approach to data abstraction and object-oriented programming is unusually flexible. It has automatic memory management or garbage collection.

Go is especially well suited for building infrastructure like networked servers, and tools and systems for programmers, but it is truly a general-purpose language and finds use in domains as diverse as graphics, mobile applications and machine learning. It has become popular as a replacement for untyped scripting languages because it balances expressiveness with safety: Go programs typically run faster than programs written in dynamic languages and suffer far fewer crashes due to unexpected type errors.

## Go Application Domain

- web apps
- network servers
- mobile applications
- machine learning
- image processing
- load balancers
- system admin
- hardware
- scripts
- crypto

## GOPATH and Code Organization

- Go requires you to organize your code in a specific way.
- Go programmers typically keep all their Go code is a single workspace
- A workspace is nothing more than a directory in your file system whose path is stored in the environment variable called GOPATH

The workspace directory (GOPATH) contains the following 3 subdirectories at its root

- **src** - Contains the source files for your own or other downloaded packages.
- **pkg** - contains go package objects aka package archives. These are non-executable files or shared libraries ending in .a
- **bin** - contains compiled and executable Go programs. When you run go install on a program directory, Go will put the executable file under this folder.

## Go Application Structure

In the most basic terms, A package is nothing but a directory inside your Go workspace containing one or more Go source files or other Go packages. Every Go source file belongs to a package. To declare a source file to be part of a package,we use the following syntax 

```C#
package <packagename>

```
The above package declaration must be the first line of code in your Go source file. All the functions, types and variables defined in your Go source file become part of the declared package.

You can choose to export a member defined in your package to outside packages, or keep them private to the same package. Other packages can import and reuse the functions or types that are exported from your package.

Packaging functionalities in this way has the following benefits:-
- It reduces naming conflicts. You can have the same function names in different packages. This keeps our function names short and concise
- It organizes related code together so that it is easier to find the code you want to reuse
- It speeds up the compilation process by only requiring recompilation of smaller parts of the program that has actually changed.

Go programs start running in the `main` package. It is a special package that is used with programs that are meant to be executable. By convention, Executable programs (the ones with the `main` package) are called Commands. Others are simply called PAckages. The `main()` function is a special function that is the entry point of an executable program.

## Compiling and Running Go Applications
1. **go run**: it compiles and runs the application. It doesn't produce an executable. It doesn't produce an executable
    - `go run file.go` compiles and immediately runs Go programs
2. **go build**: it just compiles the application. It produces an executable
    - `go build file.go` compiles a bunch of Go source files. It compiles packages and dependencies
    - If you run `go build`, it will compile the files in the current directory and will produce an executable file with the name of the current working directory

- In Go, all values are initialized
