Systems programming involves working with files and directories, process control, signal handling, network programming, system files, configuration files and file input and output (I/O).

## stdin, stdout and stderr

- `stdin` (standard input): The default source of input for a program, typically the keyboard.
- `stdout` (standard output): The default destination for output from a program, typically the terminal or console.
- `stderr` (standard error): The default destination for error messages from a program, typically the terminal or console.
- `stdin`, `stdout`, and `stderr` are file descriptors that represent the standard input, output, and error streams of a process.
- In Unix-like operating systems, these file descriptors are represented by the integers 0, 1, and 2, respectively.
- `stdin` is used for reading input from the user or from a file, `stdout` is used for writing output to the console or to a file, and `stderr` is used for writing error messages to the console or to a file.
- `stdin`, `stdout`, and `stderr` can be redirected to files or other processes using shell commands or programming constructs.
- For example, you can redirect `stdout` to a file using the `>` operator in the shell, or you can redirect `stderr` to a file using the `2>` operator.
- You can also use pipes (`|`) to connect the output of one process to the input of another process, allowing for inter-process communication.
- In C, you can use the `printf` function to write to `stdout`, the `fprintf` function to write to `stderr`, and the `scanf` function to read from `stdin`.


Go uses `os.Stdin` for accessing standard input, `os.Stdout` for accessing standard output, and `os.Stderr` for accessing standard error. These are the default file descriptors for reading and writing data in Go programs. You can use these file descriptors to read from and write to the standard input, output, and error streams of a process.

A process is an instance of a program that is being executed. Each process has its own memory space and resources, and it can communicate with other processes through inter-process communication (IPC) mechanisms such as pipes, sockets, and shared memory. Processes can be created and managed by the operating system, and they can run concurrently or sequentially depending on the scheduling policies of the operating system.

A program is a set of instructions that can be executed by a computer. It is typically written in a high-level programming language such as C, Go, or Python, and it can be compiled or interpreted to produce an executable file. A program can consist of one or more processes, and it can perform various tasks such as reading and writing files, processing data, and interacting with the user.

Each running UNIX process has a unique process ID (PID) that is used to identify it. The PID is assigned by the operating system when the process is created, and it can be used to manage the process, such as sending signals or terminating it. The PID is typically a positive integer, and it can be used to refer to the process in various system calls and commands.

### Process Categories

There are 3 process categories: user processes, daemon processes, and kernel processes.

- User processes: These are processes that are created and managed by users. They can be interactive or non-interactive, and they can run in the foreground or background. User processes typically have a user ID (UID) that identifies the user who owns the process.
- Daemon processes: These are background processes that run without user interaction. They are typically started at system boot time and run continuously in the background, waiting for requests or events to occur. Daemon processes usually have a special user ID (UID) that identifies them as system processes. Daemon processes are often used to provide services such as web servers, database servers, and print servers. They can be found in the user space or kernel space, depending on their function and implementation.
- Kernel processes: These are processes that run in the kernel space of the operating system. They have direct access to the hardware and system resources, and they can perform low-level operations such as memory management, process scheduling, and device I/O. Kernel processes are typically created and managed by the operating system itself, and they do not have a user ID (UID) like user processes or daemon processes. Kernel processes are often used to implement system calls, device drivers, and other low-level functions of the operating system.

In Go, although we can fork a new process in Go using the `exec` package, Go does not allow us to control threads - Go offers goroutines, which the user can create on top of threads that are created and handled by the Go runtime.

### Handling UNIX Signals

Signals are a way for the operating system to notify a process that an event has occurred. Signals can be sent to a process by the operating system or by other processes, and they can be used to control the behavior of a process. For example, a signal can be used to terminate a process, suspend it, or continue its execution.
- Signals are typically represented by a number or a name, and they can be sent to a process using system calls or command-line utilities.
- In UNIX-like operating systems, signals are used for various purposes, such as handling interrupts, managing child processes, and implementing inter-process communication (IPC).

### Goroutine and channel

A goroutine is the smallest executable unit in Go. It is a lightweight thread managed by the Go runtime. Goroutines are used to perform concurrent tasks in Go programs, allowing multiple functions to run simultaneously without blocking each other. Goroutines are created using the `go` keyword followed by a function call, and they can communicate with each other using channels.

A channel is a built-in data structure in Go that allows goroutines to communicate with each other by sending and receiving values. Channels provide a way to synchronize the execution of goroutines and share data between them. Channels can be created using the `make` function, and they can be used to send and receive values of a specific type. Channels can be buffered or unbuffered, depending on whether they have a fixed size or not.
- Buffered channels allow sending and receiving values without blocking, while unbuffered channels block the sender until the receiver is ready to receive the value.
- Channels can be used to implement various synchronization patterns, such as fan-out, fan-in, and worker pools.

## Buffered and Unbuffered File I/O

Buffered file I/O happens when there is a buffer for temporarily storing data before reading data or writing data. Thus instead of reading a file byte by byte, we read many bytes at one, We put the data in a buffer and wait for someone to read it in the desired way.

Unbuffered file I/O happens when there is no buffer to temporarily store data before actually reading or writing it - this can affect the performance of our programs

When dealing with critical data, it is often better to use unbuffered I/O to ensure that the data is written immediately and not delayed by buffering. However, buffered I/O can improve performance for large files or when reading/writing small amounts of data frequently.

### The bufio package

`bufio` is a package in Go that provides buffered I/O operations. It allows us to read and write data to files or other I/O streams using buffers, which can improve performance by reducing the number of system calls and minimizing the overhead of reading and writing small amounts of data.
- The `bufio` package provides various types for buffered I/O, such as `Reader`, `Writer`, and `Scanner`.
- The `Reader` type provides methods for reading data from a buffered I/O stream, such as `Read`, `ReadByte`, and `ReadString`.
- The `Writer` type provides methods for writing data to a buffered I/O stream, such as `Write`, `WriteByte`, and `WriteString`.
- The `Scanner` type provides a convenient way to read data from a buffered I/O stream line by line or token by token.