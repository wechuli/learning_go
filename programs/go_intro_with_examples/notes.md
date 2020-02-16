# GO

Go is an open source programming language that makes it easy to build simple, reliable and efficient software. The goals of the language and its accompanying tools were to be expressive, efficient in both compilation and execution, and effective in writing reliable and robust programs.

Go is especially well suited for building infrastructure like networked servers and tools and systems for programmers, but it is truly a general-purpose language and finds use in domains as diverse as graphics, mobile applications and machine learning. It has become popular as a replacement for untyped scripting languages because it balances expressiveness wit h safety :
Go programs typically run faster than programs written in dynamic languages and suffer far
fewer crashes due to unexpected type errors.

Go is a compiled language. The Go toolchain converts a source program and the things it depends on into instructions in the native machine language of a computer. These tools are accessed through a single command called `go` that has a number of subcommands. The simplest of these subcommands is `run`, which compiles the source code from one or more source files whose end in `.go`, links it with libraries, then runs the resulting executable file.

Go code is organized into packages, which are similar to libraries or modules in other languages. A package consists of one or more `.go` source files in a single directory that define what the package does. Each source file begins with a package declaration.

The go standard library has over 100 packages for common tasks like input and output, sorting and text manipulation.

Package `main` is special. It defines a standalone executable program, not a library. Within package `main` the function `main` is also special - it's where execution of the program begins. Whatever `main` does is what the program does. Of course, `main` will normally call upon functions in other packages to do much of the work.

We must tell the compiler what packages are needed by this source file, that's the role of the `import` declaration that follows the `package` declaration.


A map is a reference to the data structure. When a map is passed to a function, the function receives a copy of the reference, so any changes the called function makes to the underlying data structure will be visible through the caller's map reference too.
