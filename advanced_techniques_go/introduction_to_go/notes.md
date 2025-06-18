Although Go is a general-purpose programming language, it is primarily used for writing system tools,
commandline utilities, web services and software that work over networks.

Go ca help develop the following kind of applications:

- Professional web services
- Networking tools and servers such as Kubernetes and Istio
- Backend systems
- System utilities
- Powerfule command-line utilities,such as docker and hugo
- Applications that exchange data in JSON format
- Applications that process data from relational databases, NoSQL databases or other popular data storage systems
- Compilers and interpreters for programming laguages we design
- Database systems such as CockroachDB and key/value stored such as etcd

There are many things that Go does better than other programming languages, including the following:

The default behavior of the Go compiler can catch a large set of silly errors that might result in bugs.

Go uses fewer parentheses than C, C++, or Java and no semicolons, which makes the look of Go source code more human-readable and less error-prone.

Go comes with a rich and reliable standard library.

Go has support for concurrency out of the box through goroutines and channels.

Goroutines are really lightweight. We can easily run thousands of goroutines on any modern machine without any performance issues.

Unlike C, Go supports functional programming.

Go code is backward compatible, which means that newer versions of the Go compiler accept programs that were created using a previous version of the language without any modifications. This compatibility guarantee is limited to major versions of Go. For example, there is no guarantee that a Go 1.x program will compile with Go 2.x.

Go comes with concurrency capabilities using a simple concurrency model that is implemented using goroutines
and channels.

Go manages OS threads for us and has a powerful runtime that allows us to spawn lightweight units of work (goroutines) that communicate with each other using channels. Although Go comes with a rich standard library, there are really handy Go packages such as cobra and viper that allow Go to develop complex command-line utilities, such as docker and hugo. This is greatly supported by the fact that Go’s executable binaries are statically linked, which means that once they are generated, they do not depend on any shared libraries and include all required information.

## Goroutines and channels

Go comes with concurrency capabilities using a simple concurrency models that is implemented using goroutines and channels.
Go manages OS threads for us and has a powerful runtime that allows us to spawn lightweight units of work (goroutines)
that communicate with each other using channels. Although Go comes with a rich standard library, there are really handy Go 
packages such as `cobra` and `viper` that allow Go to develop complex command line utilities such as `docker` and `hugo`.
This is greatly supported by the fact that Go's executable binaries are statically linked, which means that once they are generated,
they do not depend on any shared libraries and include all required information.


## Pointers and interfaces

Due to its simplicity, Go code is predictable and does not have strange side effects,
and although Go supports pointers, it does not support pointer arithmetic like C unless we use 
the unsafe package, which is the root of many bugs and security holes. 
Although Go is not an object-oriented programming language, Go interfaces are 
very versatile and allow us to mimic some of the capabilities of object-oriented 
languages, such as polymorphism, encapsulation, and composition.

## Go generics and garbage collection

The latest Go versions offer support for generics, which simplifies our code when working with multiple data types. Go comes
with support for garbage collection, which means that no manual memory management is needed.

Although Go is a very practical and competent programming language, it is not perfect:

Although this is a personal preference rather than an actual technical shortcoming, Go has no direct support for object-oriented programming, which is a popular programming paradigm.

Although goroutines are lightweight, they are not as powerful as OS threads. Depending on the application we are trying to implement, there might exist some rare cases where goroutines will not be appropriate for the job. However, in most cases, designing our application with goroutines and channels in mind will solve our problems.

Although garbage collection is fast enough most of the time and for almost all kinds of applications, there are times when we need to handle memory allocation manually—Go can’t do that. In practice, this means that Go will not allow us to perform any memory management manually.


## Use Cases

 Creating complex command-line utilities with multiple commands, subcommands and command-line parameters
Building highly concurrent applications
Developing servers that work with APIs and clients that interact by exchanginf data in myriad formats, including JSON, XML and CSV
Developing WebSocket servers and clients
Developing gRPC servers and clients
Developing robus UNIX and Windows system tools



## The `go doc` and `godoc` utilities

The Go distribution comes with a plethora of tools that aim to make our lives easier Two of these tools are the go doc subcommand and the godoc utility, which allow us to see the documentation of existing Go functions and packages without needing an internet connection. We can also view the Go documentation online. Because godoc is not installed by default, we might need to install it by running go get golang.org/x/tools/cmd/godoc.

```go
package main
import (
    "fmt"
)
func main() {
    fmt.Println("Hello World!")
}
```

Each Go source code beginds with a `package` declaration. In this case, the name of the package is `main`, which has a special meaning in Go.
The `import` keyword allows us to include functionality from existing packages. In our case, we only need some functionality of the `fmt` 
package that belongs to the standard Go library. Packages that are not part of the standard Go library are imported using their full
internet path. The next important thing when we'ere creating an executable application is the `main()` function. Go considers this the entry point
to the application and beginds the execution of the application with the code found in the `main()` function of the `main` package.

Each Go function definition beginds with the `func` keyword followed by its name, signature and implementation. With the `main` package, we can name
our functions anything we want - there is a global Go rule that also applies to function and variable names and is valid for all
packages except main

Everything that beginds with a lowercase letter is considered private and is only accessible in the current package.

Go programs are organized in packages -even the smallest Go program should be develivered as a package. The `package` keyword helps us define the name of a new package,
which an be anything we want with just one exception: if we're creating an executable application and not just a package, that will be shared
by other applications or packages, we should name our package `main`.

Packages can be used by other packages. The `import` keyword is used for importing other Go packages in our Go programs in order to use some
or all of their functionality. A Go package can either be a part of the rich Standard Go library or come from an external source. Packages of the
standard Go library are imported by name  without the need for a hostname and a path, whereas external packages are imported using their
full internet paths, like github.com/spf13/cobra.

## Running Go Code

We need to know how to execute `hw.go` or any other Go application. There are two ways to execute Go code:

As a compiled language using `go build`.
As a scripting language using `go run`.


### Compiling Go code

In order to compile Go code and create a binary executable file, we need to use the `go build` command. What `go build` does it create an executable
file for us to distribute and execute manually. This means that the `go build` command requires an additional step for running our code.

The generated executable is automatically named after the source code file name without the `.go` file extension. Therefore, because of the hw.go source file, the executable will be called hw. In case this is not what we want, go build supports the -o option, which allows us to change the file name and the path of the generated executable file. As an example, if we want to name the executable file helloWorld, we should execute go build -o helloWorld hw.go instead. If no source files are provided, go build looks for a main package in the current directory.

After that, we need to execute the generated executable binary file on our own. In our case, this means executing either hw or helloWorld. This is shown in the next output:

```txt
# go build hw.go 
# ./hw
Hello World!
```

### Using Go like a scripting language

The `go run` command builds the named Go packages, which in this case is the `main` package implemented in a single file, creates a temporary executable
file, executes that file and deletes it once it is done - to our eyes, this looks like a scripting language.

If we want to test our code, then using `go run` is a better choice. However, if we want to create and distribute an executable binary,
then `go build` is the way to go.

### Important formatting and coding rules

We should know that Go comes with some strict formatting and coding rules that help the developer avoid mistakes and bugs:

- Go code is delivered in packages, and we are free to use the functionality found in existing packages. However, if we are going to import
a package, we should use some of this functionality - there are some exceptions to this rule that mainly have to do with initializing connections.
- We either use a variable or we do not declare it at all
- There is only one way to format curl braces in Go
- Coding blocks in Go are embedded in curly braces, even if they contain just a single or no statements at all.
- Go functions can return multiple values
- We cannot automatically convert between different data types, even if they are of the same kind.

## Variables

Go provides multiple ways to declare new variables in order to make the variable declaration process more natural and convinient. We can delcare
a new variable using the `var` keyword followed by the variable name, followed by the desired data type. If we want, we can follow the 
declarion with `=` and an initial value for our variable. If there is an initial value given, we can omit the data type and the compiler
will guess it for us.

If no initial value is givn to a variable, the Go compiler will automatically initialize that variable to the zero value of its data type.

There is also the `:=` notation, which can be used instead of a `var` declaration. `:=` defines a new variable by inferrinf the data of the value
that follows it. The official name for `:=` is a **short assignment statement** and it is very frequenty used in Go, especially for getting 
the return values from functions and `for` loops with the `range` keyword.

The short assignment statement can be used in place of a var declaration with an implicit type. We rarely see the use of var in Go; the var keyword is mostly used for declaring global or local variables without an initial value. The reason for the former is that every statement that exists outside of the code of a function must begin with a keyword such as func or var.

This means that the short assignment statement cannot be used outside of a function environment because it is not available there. Last, we might need to use var when we want to be explicit about the data type. For example, when we want int8 or int32 instead of int.

Therefore, although we can declare local variables using either var or :=, only const (when the value of a variable is not going to change) and var work for global variables, which are variables that are defined outside of a function and are not embedded in curly braces. Global variables can be accessed from anywhere in a package without the need to explicitly pass them to a function and can be changed unless they were defined as constants using the const keyword.

Go does not allow implicit data conversions like C.

## Go: Control

Go supports the `if/else`and `switch` control structures


curl -sfL https://get.k3s.io | K3S_URL=https://35.184.181.111:6443 K3S_TOKEN=K10e6fb21a3765e322ab61e07a766c1021c46f39f9570432dc97e1b20da005ba5cb::server:3ee5fec654b4e9a7c08f70106304b86f sh -


## Go concurrency model

The Go concurrency model is implemented using goroutines and channels. A goroutine is the smallest executable Go entity. In order to create a new goroutine, we have to use the `go` keyword followed by a predefined function or an anonymous function - both methods are equivalent as far as Go is concerned.

A channel in Go is a mechanism that, among other things, allows goroutines to communicate and exchange data. Although it is easy to create goroutines, there are other difficulties when dealing with concurrent programming, including goroutines synchronization and sharing data between goroutines - this is a Go mechanism for avoiding side effects when running goroutines. As `main()` runs as a goroutine as well, we do not want `main()` to finish before the other gorotuines of the program because when `main()` exits, the entire program, along with any goroutines that have not finished yet will terminate. Although goroutines do no share any variables, they can share memory.

Goroutines are initialized in random order and start running in random order. The Go scheduler is responsible for the execution of goroutines, just like the OS scheduler is responsible for the execution of the OS threads.

## Log files

All UNIX systems have their own log files for writing logging information that comes from running servers and programs. Usually, most system log files of a Unix
system can be found under the `/var/log` directory. However, the log files of many popular services such as Nginx can be found
elsewhere, depending on their configuration.

Generally speaking, using a log file to write some information used to be considered a better practice than writing the same output on the screen for two reasons: firstly, because the output does not get lost as it is stored on a file, and secondly, because we can search and process log files using UNIX tools, such as grep(1), awk(1), and sed(1), which cannot be done when messages are printed on a terminal window. However, this is not true anymore.

Because we usually run our services via systemd, programs should log to stdout so systemd can put logging data in the journal. Additionally, in cloud-native applications, we are encouraged to simply log to stderr and let the container system redirect the stderr stream to the desired destination.


The UNIX logging service has support for two properties named `logging level` and `logging facility`. The logging level is a 
value that specifies the serverity of the log entry.

There are various logging levels, including debug, info, notice, warning, err, crit, alert, and emerg, in reverse order of severity. The log package of the standard Go library does not support working with logging levels. The logging facility is like a category used for logging information. The value of the logging facility part can be one of auth, authpriv, cron, daemon, kern, lpr, mail, mark, news, syslog, user, UUCP, local0, local1, local2, local3, local4, local5, local6, or local7 and is defined inside /etc/syslog.conf, /etc/rsyslog.conf, or another appropriate file depending on the server process used for system logging on our UNIX machine. This means that if a logging facility is not defined correctly, it will not be handled; therefore, the log messages we send to it might get ignored and, therefore, lost.

The `log` package sends log messages to standard error. Part of the `log` package is the `log/syslog` package, which allows us
to send log messages to the `syslog` server of our machine. Although by default, `log` writes to standard error, and the use 
of `log.SetOutput()` modifies that behavior. The list of functions for sending logging data includes log.Printf(), log.Print(), log.Println(), log.Fatalf(), log.Fatalln(), log.Panic(), log.Panicln() and log.Panicf().

In order to write to system logs, we need to call the `syslog.New()` function with the appropriate parameters. Writing to the main system
log file is as easy as calling `syslog.New()` with the syslog.LOG_SYSLOG option. After that, we need to tell our Go program that all logging information goes to the new logger—this is implemented with a call to the log.SetOutput() function


### The log.Fatal() function

The `log.Fatal()` function is used when something erroneous has happened, and we just want to exit our program as soon as 
possible after reporting that bad situation. The call to `log.Fatal()` terminates a Go program at the point where `log.Fatal()`
was called after printing an error message.In most cases, this custom error message can be Not enough arguments, Cannot access file, or something similar. Additionally, it returns back a nonzero exit code, which in UNIX indicates an error.

### The log.Panic() function

There are situations where a program is about to fail for good, and we want to have as much information about the failure as possible—log.Panic() implies that something really unexpected and unknown, such as not being able to find a file that was previously accessed or not having enough disk space, has happened. Analogous to the log.Fatal() function, log.Panic() prints a custom message and immediately terminates the Go program.

Keep in mind that log.Panic() is equivalent to a call to log.Print() followed by a call to panic(). panic() is a built-in function that stops the execution of the current function and begins panicking. After that, it returns to the caller function. On the other hand, log.Fatal() calls log.Print() and then os.Exit(1), which is an immediate way of terminating the current program.


Most of the time, and especially on applications and services that are deployed to production, we just need to write our logging data in a log file
of our choice. This can be for many reasons, including writing debugging data without messing with the system log files or 
keeping our own logging data separate from system logs in order to transfer it or store it in a database of software like
Elasticsearch.

The path of the log file that is used is hardcoded into the code using a global variable named `LOGFILE`. In this example,
the log file is placed inside the /tmp directory. This location is chosen to prevent our file system from running out of space
in case something goes wrong. It's worth noting that the `/tmp` directory is not typically used for storing data because
its contents are usually cleared after each system reboot. 

Additionally, at this point, this will save us from having to execute customLog.go with root privileges and from putting unnecessary files into our precious system directories.

