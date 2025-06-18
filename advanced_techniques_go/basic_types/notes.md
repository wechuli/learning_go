Data is stored and used in variables, and all Go variables should have a data type that is determined either implicitly or explicitly. Knowing the built in data types in Go allows us to understand how to manipulate somple data values and construct more complex data structures when simple daa types are not enough or not efficient for a given job


## The error Data Type

Go provides a special data type for representing error conditions and error messages named `error` - in practice, this means that Go treats errors as values. In order to program successfully in Go, we should be aware of the error conditions that might occur with the functions and methods we are using and handle them accordingly.

## Global error handling practoces

We should have a global error-handling tactic in each application that should not change. In practice, this means the following

- All error messages should be handled at the same level, which means that all errors should either be returned to the calling function or handled at the place they occurred.
- It should be clearly documented how to handle critical errors. This means that there will be situations where a critical error should terminate the program and other times where  critical error might just create a warning message onscreen.
- All error messages should be logged in a consistent way. This means that all error messages should be logged in the same way, and the log should be easily accessible for the developer to review.
- It is considered a good practice to send all error messages to the log service of our machine because this way, the error messages can be examined at a later time. However, this is not always true, so exercise caution when setting this up—for example, cloud-native apps do not work that way.

## Numeric Data types

Go supports integer, floating-point and complex nmber values in various versions depending on the memory space they coonsume:

| Data Type   | Description                        |
|-------------|------------------------------------|
| int8        | 8-bit signed integer              |
| int16       | 16-bit signed integer             |
| int32       | 32-bit signed integer             |
| int64       | 64-bit signed integer             |
| int         | 32- or 64-bit signed integer      |
| uint8       | 8-bit unsigned integer            |
| uint16      | 16-bit unsigned integer           |
| uint32      | 32-bit unsigned integer           |
| uint64      | 64-bit unsigned integer           |
| uint        | 32- or 64-bit unsigned integer    |
| float32     | 32-bit floating-point number      |
| float64     | 64-bit floating-point number      |
| complex64   | Complex number with float32 parts |
| complex128  | Complex number with float64 parts |


## Non-numeric Data Types

Go has support for strings, characters, runes and dates and times. However, Go does not have a dedicated `char` data type.

Go supports the `string` data type for representing strings. A Go string is just a collection of bytes and can be accessed as a whole or as an array. A single byte can store any ASCII character - however, multiple bytes are usually needed to store a single Unicode character. Go uses the UTF-8 encoding for Unicode characters, which means that a single Unicode character can be stored in one to four bytes. Four bytes is equivalent to 32 bits, which is the size of a `rune` data type in Go. The `rune` data type is used to represent a single Unicode character in Go.

We can create a new byte slice from a given string by using the `[]byte("A String")` statement. Given a byte slice variable `b`, we can convert it into a strin using the `string(b)` statement.When working with byte slices that contain Unicode characters, the number of bytes in a byte slice is not always connected to the number of characters in the byte slice because most Unicode characters require more than one byte for their representation. As a result, when we try to print each single byte of a byte slice using fmt.Println() or fmt.Print(), the output is not text presented as characters but integer values. If we want to print the contents of a byte slice as text, we should either print it using string(byteSliceVar) or using fmt.Printf() with %s to tell fmt.Printf() that we want to print a string. We can initialize a new byte slice with the desired string b

We can define a rune using single quotes `r := '$'` and we can print the integer value of the byte that compose it using `fmt.Println(r)`. We can also print the rune as a character using `fmt.Printf("%c", r)`. We can also print the hexadecimal value of the byte that compose the rune using `fmt.Printf("%x", r)`. We can also print the binary value of the byte that compose the rune using `fmt.Printf("%b", r)`. We can also print the type of the rune using `fmt.Printf("%T", r)`.


We can convert an integer value into a string in two main ways: using `string()` and using a function from the `strconv` package. The two methods are fundamentally different. The `string()` function converts an integer value into a Unicode code point, which is a single character, whereas functions such as `strconv.FormatInt()` and `strconv.Itoa()` converts an integer value into a string value with the same represention and the same number of characters

## The unicode package

The `unicode` standard Go package contains various handy functions for working with Unicode code points. One of them is called `unicode.ISPrint()` and it can help us identify the parts of a string that are printable using runes.

### The strings package

The `strings` standard Go package allows us to manipulate UTF-8 strings in Go and includes many powerful functions for working with strings. For example, we can use the `strings.ToUpper()` function to convert a string to uppercase and the `strings.ToLower()` function to convert a string to lowercase. We can also use the `strings.Contains()` function to check if a string contains a substring and the `strings.Replace()` function to replace a substring with another substring. We can also use the `strings.Split()` function to split a string into a slice of substrings and the `strings.Join()` function to join a slice of substrings into a single string.

### Times and Dates

Often, we need to work with date and time information to store the time and entry was last used in a database or the time an entry was inserted into a database.

You can use the `time.Time` data type, which represents an instant in time with nanosecond precision. Each `time.Time` value is associated with a location (time zone)

## Introducting Go Constants

Go supports constants, which are variables that cannot change their values. Constants in Go are defined with the help of the `const`
keyword. Generally speaking, constants can be either global or local variables

## Arrays and Slices

Arrays in Go have the following characteristics and limitations

- When defining an array variable, we must define its size. Otherwise, we should put [...] in the array declaration and let the Go compiler find out the length for us.
- We can't change the size of an array after it has been defined.
- When we pass an array to a function, we are passing a copy of the array, not a reference to the array. This means that any changes made to the array inside the function will not be reflected outside the function.


## Slices

Slices in Go are more powerful than arrays mainly because they are dynamic, which means they can grow or shrink after creation if needed. Additionally, any changes we make to a slice inside a function also affect the original slice because they are pointing toward the same memory. In reality, a slice value is a header that contains the following:

- A pointer to the underlying array where the elements are actually stored
- The length of the array
- The array's capacity

Note that the slice value does not include its elements, just a pointer to the underlying array. So, when we pass a slice to a function, Go makes a copy of that header and passes it to the function. This copy of the slice header includes the pointer to the underlying array. That slice header is defined in the reflect package as follows:

```go
type sliceHeader struct {
    Data uintptr
    Len  int
    Cap  int
}
```

A side effect of passing the slice header is that it is faster to pass a slice to a function because Go does not need to make a copy of the slice and its elements, just the slice header.

We can create a slice using `make()` or like an array without specifying its size or using `[...]`. If we do not want to initializa a slice, then using `make()` is better and faster. 

However, if we want to initialize it at the time of creation, then make() can't help us. As a result, we can create a slice with three float64 elements as aSlice := []float64{1.2, 3.2, -4.5}.

Creating a slice with space for three float64 elements with make() is as simple as executing make([]float64, 3). Each element of that slice has a value of 0, which is the zero value of the float64 data type.

Both slices and arrays can have many dimensions—creating a slice with two dimensions with make() is as simple as writing make([][]int, 2). This returns a slice with two dimensions where the first dimension is 2 (rows) and the second dimension (columns) is unspecified and should be explicitly specified when adding data to it.

If we want to define and initialize a slice with two dimensions at the same time, we should execute something similar to twoD := [][]int{{1, 2, 3}, {4, 5, 6}}.

We can find the length of an array or a slice using len(). We can add new elements to a full slice using the append() function. The append() function automatically allocates the required memory space.

Both arrays and slices support the `len()` function for finding out their length. However, slices also have an additional property called `capacity` that can be found using the `cap()` function.

The capacity of a slice is really important when we want to select a part of a slice or when we want to reference an array using a slice. The capacity shows how much a slice can be expanded without the need to allocate more memory and change the underlying array. Although the capacity of a slice is handled by Go after slice creation, a developer can define the capacity of a slice at creation time using the make() function—after that, the capacity of the slice doubles each time the length of the slice is about to become bigger than its current capacity. The first argument of make() is the type of the slice and its dimensions, the second is its initial length, and the third, which is optional, is the capacity of the slice. Although the data type of a slice can’t change after creation, the other two properties can change.

Go allows us to select parts of a slice, provide all desired elements are next to each other. This can be pretty handt when we select a range of elements and we do not wanto go give their indexes one by one. We select a part of a slice by defining two indexes: the first one is the beginning of the selection, whereas the second one is the end of the selection, without including the element at that index, separated by `:`.

A byte slice is a slice of the `byte` data type ([]byte). Go knows that most `byte` slices are used to store strings and so it makes it easy to swutch between this type and the `string` type. There is nothing special in the way we can access a byte slice compared to the other types of slices. What is special is that Go uses byte slices for performing file I/O operations because they allow us to determine with precision the amount of data we want to read or write to a file. This happens because bytes are a universal unit among computer systems.

Because Go does not have a `char` data type, it uses `byte` and `rune` for storing character values. A single `byte` can only store a single ASCII character, whereas a `rune` can store Unicode characters. However, a rune can occupy multiple bytes.

As mentioned before, behind the scenes, each slice is implemented using an underlying array. The length of the underlying array is the same as the capacity of the slice, and there exists pointers that connect the slice elements to the appropriate array elements.

We can understand that by connecting an existing array with a slice, Go allows us to reference an array or part of an array using a slice. This has some strange capabilities, including the fact that the changes to the slice affect the referenced array. However, when the capacity of the slice changes, the connection to the array ceases to exist! This happens because when the capacity of a slice changes, so does the underlying array, and the connection between the slice and the original array does not exist anymore.

## The copy() function

Go offers the `copy()` function for copying an existing array to a slice or an existing slice to another slice. However, the use of `copy()` can be tricky because the destination slice is not automatically expanded if the source slice is bigger than the destination slice. Additionally, if the destination slice is bigger than the source slice, then `copy()` does not empty the elements from the destination slice that did not get copied.

## Sorting Slices

There are times when we want to present our information sorted, and we want Go to do the job for us.

The `sort` package can sort slices of built-in data types without the need to write any extra code. Additionally, Go provides the `sort.Reverse()` function for sorting in reverse rather than the default order. However, what is really interesting is that sort allows us to write our own sorting function for custom data types by implementing the `sort.Interface` interface. This is a powerful feature because it allows us to sort slices of custom data types without the need to write any extra code. 

## Pointers in Go

Go has support for pointers but not pointer arithmetic, which is the case of many bugs and errors in programming languages like C. A pointer is the memory address of a variable. We need to dereference a pointer in order to get its value - derefencing is performed using the * character in from of the pointer variable. Additionally, we can get the memory address of a normal variable using an `&` in front of it

If a pointer variable points to an existing variable, then any changes we make to the stored value using the pointer variable will modify the regula variable.

The main benefit we get from pointers is that passing a variable to a function as a pointer (we can call that by reference) does not discard any changes we make to the value of that variable inside the function when  the function returns. There exists times when we want tthat functionality because it simplifies our code, but the price we pay for that simplicity is being extra careful with what we do with a pointer variable. Remember that slices are passed to functions without the need to use a pointer - it is Go that passes the pointer to the underlying array of a slice, and there is no way to change that behavior.

Apart from reasons of simplicity, there exists three more reasons for using pointers:

- Pointers allow us to share data between functions. However, when sharing data between functions and goroutines, we should be extra careful with race condition issues.
- Pointers are also very handy when we want to tell the difference between the zero value of a variable and a value that is not set (`nil`). This is particularly useful with structures because pointers to structures can have the `nil` value, which means that we can compare a pointer to a structure with the `nil` value, which is not allowed for normal structure variables.
- Having support for pointers and, more specifically, pointers to structures allows Go to support data structures such as linked lists and binary trees, which are widely used in computer science. Therefore, we are allowed to define a structure field of a Node structure as Next *Node, which is a pointer to another `Node` structure. Without pointers, this would have been difficult to implement and may be too slow.

## Generating Random Numbers


Each random number generator needs a seed to start producing numbers. The seed is used for initizlizing the entire process and is extremely important because we always start with the same seed we will always get the same sequence of pseudorandom numbers. This means that everybody can regenerate that sequence and that particular squence will not be random after all.

If we intend to use these pseudorandom numbers for security-related work, it is important that we use the `crypto/rand` package, which is a part of the Go standard library. The `crypto/rand` package is a cryptographically secure pseudorandom number generator (CSPRNG) that is suitable for generating random numbers for security-related work. The `crypto/rand` package is not as fast as the `math/rand` package, but it is more secure. The `crypto/rand` package is used in the `math/big` package, which is used for working with big numbers.