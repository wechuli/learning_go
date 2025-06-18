The main focus of this chapter is Go packages, which are Go's way of organizing, delivering and using code. The most common component of packages is `functions` which are pretty flexible and powerfule and are used for data processing and manipulation. Go also supports modules, which are packages with version numbers.


Regarding the visibility of package elements, Go follows a simple rule that states that functions, variables, data types, structure fields, and so forth that begin with an uppercase letter are public, whereas functions, variables, types, and so on that begin with a lowercase letter are private. This is the reason why fmt.Println() is named Println() instead of just println(). The same rule applies not only to the name of a struct variable but to the fields of a struct variable—in practice, this means that we can have a struct variable with both private and public fields. However, this rule does not affect package names, which are allowed to begin with either uppercase or lowercase letters.

- Go packages
- Functions
- Developing our own packages
- Using GitHub to store Go packages
- A package for working with a database
- Modules
- Creating better packages
- Generating documentation
- GitHub Actions and Go
- Versioning utilities

## Go Packages

Everything in Go is delivered in the form of a package. A Go package is a Go source file that begins with the `package` keyword, followed by the name of the package:

> Note: Note that packages can have structure. For example, the net package has several subdirectories named http, mail, rpc, smtp, textproto, and url, which should be imported as net/http, net/mail, net/rpc, net/smtp, net/textproto, and net/url, respectively.

Apart from the packages of the Go standard library, there are external packages that cab be imported using their full address, and that should be downloaded on the local machine before their first use

Packages are mainly used for grouping related functions, variables, and constants so that we can transfer them easily and use them in our own programs. Note that apart from the `main` package, Go packages are not autonomous programs and can't be compiled into executable files on their own. Packages need to be called directly or indirectly from a `main` package in order to be used.

The single most important package in Go is the `main` package, which is the entry point of a Go program. The `main` package must contain a function named `main()`, which is the first function to be executed when the program starts. The `main` package can import other packages, but it cannot be imported by other packages. This means that the `main` package is not a library and cannot be used as a library by other packages.
The `main` package is the only package that can be compiled into an executable file. All other packages are libraries that can be used by the `main` package or by other packages. The `main` package is the only package that can be run as a standalone program. All other packages are libraries that can be used by the `main` package or by other packages.

### Anonymous functions

Anonymous functions can be defined inline without the need for a name, and they are usually used for implementing things that require a small amount of code. In Go, a function can return an anonymous function or take an anonymous function as one of its arguments. Additionally, anonymous functions can be attached to Go variables. Note that anonymous functions are called `lambdas` in functional programming terminology. An anonymous function in Go has access to the variables defined in the context where it is itself defined. This property of anonmous functions is called closure.

It is considered a good practice for anonymous functions to have a small implementation and a local focus. If an anonymous function does not have a local focus, then we might need to consider making it a regular function. When an anonymous function is suitable for a job, it is extremely convenient and makes our life easier; just do not use too many anonymous functions in your programs without having a good reason to. We will look at anonymous functions in action in a while.

Unlike C, Go allows us to name the return values of a Go function. Additionally, when such a function has a `return` statement without any arguments, the function automatically returns the current value of each named return value in the order in which they are declared in the function signature.

### Functions That Accept Functions as Arguments

Functions can accept other functions are parameters. This is a powerful feature of Go that allows us to create higher-order functions. A higher-order function is a function that takes one or more functions as arguments or returns a function as its result. Higher-order functions are used in functional programming to create more abstract and reusable code.

Example of a higher-order function:
```go
package main
import (
    "fmt"
    "math"
)
func apply(f func(float64) float64, x float64) float64 {
    return f(x)
}
func main() {
    fmt.Println(apply(math.Sin, math.Pi/2))
    fmt.Println(apply(math.Cos, math.Pi))
    fmt.Println(apply(math.Tan, math.Pi/4))
}
```


Apart from accepting functions as arguments, Go functions can also return functions. This is a powerful feature of Go that allows us to create closures. A closure is a function that captures the environment in which it was created. This means that a closure can access variables from its surrounding scope even after that scope has exited.
```go
package main
import (
    "fmt"
)
func makeIncrementer(increment int) func(int) int {
    return func(x int) int {
        return x + increment
    }
}
func main() {
    incrementBy2 := makeIncrementer(2)
    incrementBy3 := makeIncrementer(3)
    fmt.Println(incrementBy2(5)) // 7
    fmt.Println(incrementBy3(5)) // 8
}
```

### Variadic Functions

Variadic functions are functions that can accept a variable number of parameters. In Go, we can define a variadic function by using the `...` syntax in the function signature. The `...` syntax indicates that the function can accept zero or more arguments of the specified type. The arguments are passed to the function as a slice of the specified type.

#### General ideas and rules

The general ideas and rules behind variadic functions are as follows:

- Variadic functions use the pack operator, which consists of three dots (`...`) followed by the type of the arguments. The pack operator indicates that the function can accept a variable number of arguments of the specified type.
- The pack operator can onle be used one in any given funcrion
- The variable that holds the pack operation is a slice, and is accessed as a slice inside the variadic function
- The variable name that is related to the pack operator is always last in the list of function parameters
- When calling a varidic function, we should put a list of separated values spearated by `,` in the place of the variable with the pack operator os a slice with the upack operator

Note that the pack operator can also be used with an empty interface. In fact, most functions in the `fmt` package are variadic functions that accept an empty interface as their last argument. This means that we can pass any type of value to these functions, and they will be converted to a string representation.


#### PAssing os.Args as ...interface{}

However, there is a situation that needs special carehere. If we try to pass `os.Args`, which is a slice of strings (`[]string`) as `...interface{}` to a variadic function, our code will not compile and will generate an error message similar to cannot use os.Args (type []string) as type []interface {} in argument to <function_name>.

This happens because the two data type ([]string and []interface{}) do not have the same representations in memory - this applies to all data types. In practice, this means that we can't use `os.Args...` to pass each individual value of the `os.Args` slice to a variadic function.

On the other hand, if we just use `os.Args`, it will work, but this passes the entire sliece as a single entity instead of passing its individual values! 
The silution to this problem is converting the slice of strings - or any other slice into a slice of interface{}. One way to do this is by using the code that follows:

```go
empty := make([]interface{}, len(os.Args))
for i, arg := range os.Args {
    empty[i] = arg
}
```


### The `defer` Keyword


The `defer` keyword postpones the execution of a function until the surrounding function returns. Usually, `defer` is used in file I/O operations to keep the function call that closes an opened file close to call thay opened it. This helps to avoid the common programmer error where we forget to close a file - that we had opened earlier -just before the function exits.

It is very important to remember that `deferred function` are executed in `last in, first out (LIFO)` order after the surrounding function has been returned. Putting it simply, this means that if we `defer`function `f1()` first, function `f2()` second, and function `f3()` third, then the order of execution will be `f3()`, `f2()`, and `f1()`. This is a very important point to remember when using the `defer` keyword.
```go
```

## Developing our own packages

At some point, we are going to need to develop our own packages to organize our code and distribute it if needed. In Go, everything that begins with an uppercase latter is considered public, and can be accessed from outside its package, whears all other elements are considered private. The only exception to this Go rule is package names - it is a best practice to use lowercase package names, even though uppercase package names are allowed.

Compiling a Go package can be done manually if it exists on the local machine. When a Go package is downloaded from the internet, it is compiled automatically. Additionally, if the downloaded package contains any errors, we'll learn about them at the time of the download.

The main reason for compiling Go packages on our own is to check for syntax errors or any other errors that might be present in the package. This is done by using the `go build` command, which will compile the package and create an executable file. If there are any errors in the package, they will be displayed in the terminal.

## The init() function

Each Go package can have a function named `init()`, which is executed automatically when the package is imported. The `init()` function is used to initialize the package and perform any setup that is needed before the package is used. The `init()` function is executed before the `main()` function, and it is executed only once per package, even if the package is imported multiple times.

The `init()` function has the following characteristics:
- `init()` takes no arguments
- `init()` has no return values
- The `init()` function is optional
- The `init()` function is called implicitly by Go
- We can have an `init()` function in the `main` package. In that case, `init()` is executed before the `main()` function. In fact, all `init()` functions are always executed prior to the `main()` function.
- A source file can contain multiple `init()` functions - these are executed in the order of declaration.
- The `init()` function or functions of a package are executed only one, even if the package is imported multiple times.
- Go packages can contain multiple files. Each source file can contain one or more `init()` functions.

The fact that the `init()` function is a private function by design means that it can't be called from outside the package in which it is contained. Additionally, because the user of a package has no control over the `init()` function, we should think carefully before using an `init()` function or changing any global state in `init()`.

There are some exceptions where the use of `init()` makes sense:

- For initializing network connections that might take time prior to the execution of package functions or methods
- For initializing connections to one or more servers prior to the execution of package functions or methods
- For creating required files and directories
- For checking whether required resources are available or not. As an example, if a `main` package imports package `A`, and package `A` depends on package `B`, then the following will take place:

- The process starts with the `main` package
- The `main` package imports package `A`
- Package `A` imports package `B`
- The global variables if any in package `B` are initialized
- The `init()` function or functions of package `B`, if they exist, run. This is the first `init()` function that gets executed
- The global variables, if any, in package `A` are initalized
- The `init()` function or functions of package `A`, if ther are any run.
- The global variables in the `main` package are initialized
- The `init()` function or functions of the `main` package, if they exist, run
- The `main()` function of the `main` package begins execution

## Modules

In Go Modules are used to manage dependencies and versioning of packages. A Go module is a collection of Go packages that are versioned together. A Go module is defined by a `go.mod` file, which contains information about the module, including its name, version, and dependencies.

Go uses semantic versioning for versioning modules. This means that versions begin with the letter v, followed by the major.minor.patch version numbers. Therefore, we can have versions such as v1.0.0, v1.0.5, and v2.0.2. The v1, v2, and v3 parts signify the major version of a Go package that is usually not backward compatible. This means that if our Go program works with v1, it will not necessarily work with v2 or v3—it might work, but we can’t count on it. The second number in a version is about features. Usually, v1.1.0 has more features than v1.0.2 or v1.0.0 while being compatible with all older versions. Lastly, the third number is just about bug fixes without having any new features. Note that semantic versioning is also used for Go versions.

### Creating better packages

- The first unofficial rule of a successful package is thaht its elements should be connected in some way. It's better to split the functionality of a package into multiple packages, even if it may seem unncecessary, than to add too much functionslity to a single Go package.
- - Ensure our packaage has a clear and useful API so that any cunsumer can quickly become productive with it.
- We should try and limit the public API of our packages to only what is absolutely necessary. Additionally, we should give our functions descriptive but concise names.
- Interfaces, and in future Go versions, generics, can improve the usefulness of our functions, so when it’s appropriate, we can use an interface instead of a single type as a function parameter or return type.
- When updating one of our packages, we should make an effort to avoid breaking things and creating incompatibilities with older versions unless it is absolutely necessary.
- When we develop a new Go package, we should try to use multiple files in order to group similar tasks or concepts.
- We should avoid re-creating a package from scratch that already exists. Instead, we can make changes to the existing package and maybe create our own version of it.
- Nobody wants a Go package that prints logging information on the screen. It’s more professional to have a flag for turning on logging when needed. The Go code of our package should align with the Go code of our programs. This means that if someone looks at a program that uses our packages and our function names stand out in the code in a bad way, it would be better to change the names of our functions. Since the name of a package is used almost everywhere, we should try to use concise and expressive package names.
- It is more convenient if we put new Go-type definitions near where they are used the first time because we wouldn’t want to search source files for definitions of new data types.
- We should try to create test files for our packages because packages with test files are considered more professional than ones without them; small details make all the difference and give people confidence in our abilities as developers!


## Generating Documentation

Go follows a simple rule regarding documentation: in order to document a function, a method, a variable or even the package itself, we can write comments. As usual, that should be located directly before the elemtn we want to document, without any empty lines in between. We can use one or more single-line comments, which are lines beginning with //, or block comments, which begin with /* and end with */—everything in between is considered a comment.