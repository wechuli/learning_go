# Goroutines and Channels

Concurrent programming, the expression of a program as a composition of sevreal autonomous activities has never been more important than it is today. Web servers handle requests for thousands of clients at once. Tablet and phone apps render animations in the user interface while simultaneously performing computation and network requests in the background.

Go enables two styles of concurrent programming. Goroutines and channels support communicating sequential processes or CSP, a model of concurrency in which values are passed between independent activitities (goroutines) but variables are for the most part confined to a single activity.

## Goroutines

In Go, each concurrently executing activity is called a goroutine. The differences between threads and goroutines are essentially quantitative, not qualitative.

When a program starts, its only goroutine is the one that calles the main function, so we call it the main goroutine. New goroutines are created by the go statement. Syntatically, a go statement is an ordinary function or method call prefixed by the keyword `go`. A go statement causes the function to be called in a newly created goroutine.

```GO

f()  // call f(); wait for it to return
go f() // create a new goroutine that calls f(); dont wait


package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d \n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r %c", r)
			time.Sleep(delay)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}


```

After several seconds of animations, the fib(45) call returns and the `main` funnction prints its result.

The `main` function then returns. When this happens, all goroutines are abruptyl terminated and the program exits. Other than by returning from main or exiting the program, there is no programmatic way for one goroutine to stop another, but there are ways to communicate with a goroutine to request that it stop itself.

Notice how the program is expressed as the composition of two autonomous activities, spinning and Fibonacci computation. Each is written as a separate function but both make progress concurrently.

Networking is a natural domain in which to use concurrency since servers typically handle many connections from their clients at once, each client being essentially independent of the others.

## Channels

If goroutines are the activities of a concurrent Go program, channels are the connections between them. A channel is a communication mechanism that lets one goroutine send values to another. Each channel is a conduit for values of a particular type, called the channel's element type.

As with maps, a channel is a reference to the data structure created by `make`. When we copy a channel or pass one as an argument to a function, we are copying a reference, so caller and callee refer to the same data structure. As with other reference types, the zero value of a channel is nil.

A channel has two principal operations, send and receive, collectively known as communications. A send statement transmits a value from one goroutine, through the channel to another goroutine executing a corresponding receive expression. Both operations are written using the <- operator. In a send statement, the <-  separates the channel and value operands. In a receive expression, <- precedes the channel operand. A receive expression whose result is not used is a valid statement.

Channels support a third operation, close, which sets a flag indicating that no more values will ever be sent on this channel, subsequent attempts to send will panic. Receive operations on a closed channel yield the values that have been sent until no more values are left; any receive operations thereafter complete immedia