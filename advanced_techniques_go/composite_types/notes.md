Go offers support for maps and structures, which are composite data types.

- maps
- structures
- pointers and structures
- regular expressions and pattern matching

Maps are powerful data structures because they allow us to use indexes of various data types as keys to look up our data as long as these keys are comparable.

### Advantages of maps

- Maps are versatile
- working with maps in Go is fast because we can access all elements of a map in linear time.

## Storing to a Nil map

We are allowed to assign a map variable to `nil`. In that case, we will not be able to use that variable until we assign it to a new map variable. If we try to store data on a `nil` map, our program will crash.

In real-world applications, if a function accepts a map argument, then it should check that the map is not nil before working with it.


## Iterating over a map

We can iterate over a map using a `for` loop. The order of the elements in a map is not guaranteed, so we cannot rely on the order of the elements in a map.

```go
for key, value := range mapVariable {
    // do something with key and value
}
```

# Go Strictures

Structures in Go are both very powerful and very popular and are used for organizing and grouping various types of data under the same name. Structures are the more versatile data types in Go, and they can even be associated with functions, which are called methods.


Note: Structures, as well as other user-defined data types, are usually defined outside the main() function or any other package function so that they have a global scope and are available to the entire Go package. Therefore, unless we want to make clear that a type is only useful within the current local scope and is not expected to be used elsewhere, we should write the definitions of new data types outside functions.


When we define a new structure, we group a set of values into a single data type, which allows us to pass and receive this set of values as a single entity. A structure has fields, and each field has its own data type, which can even be another structure or slice of structures. Additionally, because a structure is a new data type, it is defined using the `type` keyword followed by the name of the structure and ending with the `struct` keyword.

```go
type structName struct {
    field1 dataType1
    field2 dataType2
    ...
    fieldN dataTypeN
}
```

## How to use structures

There exist two ways to work with structure variables. The first one is as regular variables, and the second one is as pointer variables that point to the memory address of a structure. Both ways are equally good and are usually embedded into separate functions because they allow us to initilize some or all of the fields od structure variables properly and/or do any other tasks we want before using the structure variable.As a result, there exist two main ways to create a new structure variable using a function. The first one returns a regular structure variable, whereas the second one returns a pointer to a structure. Each one of these two ways has two variations. The first variation returns a structure instance that is initialized by the Go compiler, whereas the second variation returns a structure instance that is initialized by the user.

## Type identity

The order in which we put the fields in the definition of a structure type is significant for the `type identity` of the defined structure. Put simply, two structures with the same fields will not be considered identical in Go if their fields are not in the same order.

## Using the new keyword

Additionally, we can create new structure instances using the `new()` keyword: `ps := new(Entry)`. The `new()` keyword has the following properties

- It allocates the proper memory space, which depends on the data type and then it zeroes it
- It always returns a pointer to the newly allocated memory space
- It works for all data types except channels and maps

If no initial value is given to a variable, the Go compiler automatically initializes that variable to the zero value of its data type.

## Slices of structures

We can create slices of structures in order to group and handle multiple structures under a single variable name. However, accessing a field of a given structure required knowinf the exact place of the structure in the slice. So, each slice element is a structure that is accessed using a slice index. Once we select the slice element we want, we can select its field.


## Regular Expressions and Pattern Matching

Pattern matching is a technique for searching a string for some set of characters based on a specific search pattern that is based on regular expressions and grammr.

Regular expressions are a powerful tool for searching and manipulating strings. They are a domain-specific language that is present in many programming languages, including Go. Regular expressions are used to search for patterns in strings, and they are created using a sequence of characters that define the search pattern. In Go, regular expressions are created using the `regexp` package, which is part of the standard library.

A regular expression is a sequence of characters that defines a search pattern. Every regular expression is compiled into a recognizer by building a generalized transition diagram called a finite automation. 

### The `regexp` package

The Go package responsible for defining regular expressions and performing pattern matching is called `regexp`. We use the `regexp.MustCompile()` function to create the regular expression and the `Match()` function to see whether the given string is a match or not.

The `regexp.MustCompile()` function parses the given regular expression and returns a `regexp.Regexp` variable that can be used for matching - `regexp.Regexp` is the representation of a compiled regular expression. The function panics if the expression cannot be parsed, which is good because we will know that our expression is invalid early in the process. The `re.Match()` method returns `true` if the given `byte slice` matches the `re` regular expression, whhich is a `regexp.Regexp` variable, and `false` otherwise.

```go