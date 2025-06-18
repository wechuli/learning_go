An interface is a contract that defines a set of methods and properties that a class must implement. It allows for polymorphism, where different classes can be treated as the same type if they implement the same interface.
An interface is a way to define a common set of behaviors that can be shared across different classes. It allows for loose coupling and promotes code reusability. In Python, interfaces are typically defined using abstract base classes (ABCs) from the `abc` module.

Interfaces are not just about data manipulation ad sorting. Interfaces are about expressing abstractions and identifying and defining behaviors that can be shared among different data types. Once we have implemented an interaface for a data type, a new world of functionality becomes available to the variables and the values of that type, which can save us time and increase our productivity.

Interfaces work with methods on types or type methods, which are like functions attached to given data types, which in Go are usually structures. Remember that once we implement the required type methods of an interface, that interface is satisfied implicitly.

Another handy Go feature is reflection, which allows us to examine the structure of a data type at execution time.

- Reflection
- Type methods
- Interfaces
- Working with two different CSV file formats
- Object-oriented programming in Go
- Updating the phone book application


## Reflection

Reflection is a powerful feature in Go that allows you to inspect the type and value of variables at runtime. It is part of the `reflect` package, which provides functions to work with types and values dynamically.


The most useful parts of the reflect package are two data types named reflect.Value and reflect.Type. Now, reflect.Value is used for storing values of any type, whereas reflect.Type is used for representing Go types. There exist two functions named reflect.TypeOf() and reflect.ValueOf() that return the reflect.Type and reflect.Value values, respectively. Note that reflect.TypeOf() returns the actual type of a variable—if we are examining a structure, it returns the name of the structure.

As structures are really important in Go, the reflect package offers the reflect.NumField() method for listing the number of fields in a structure, as well as the Field() method for getting the reflect.Value value of a specific field of a structure.

The reflect package also defines the reflect.Kind data type, which is used for representing the specific data type of a variable: int, struct, etc. The documentation of the reflect package lists all possible values of the reflect.Kind data type. The Kind() function returns the kind of a variable.

Lastly, the Int() and String() methods return the integer and string value of a reflect.Value, respectively.

Note: Reflection code can look unpleasant and hard to read sometimes. Therefore, according to the Go philosophy, we should rarely use reflection unless it is absolutely necessary because, despite its cleverness, it does not create clean code.

### The three disadvantages of reflection

Without a doubt, reflection is a powerful Go feature. However, as with all tools, reflection should be used sparingly for three main reasons:

- The first reason is that extensive use of reflection will make our programs hard to read and maintain. A potential solution to this problem is good documentation.
- The second reason is that Go code that uses reflection makes out programs slower. Generally speaking, Go code that works with a particular data type is always faster than Go code that uses reflection to dynamically work with any Go data type. Additionally, such dynamic code makes it difficult for tools to refactor or analyze our code.
- The last reason is that reflection errors cannot be caught at build time and are reported at runtime as panics which means that relfection errors can potentially crash our programs. This can happen months or even years after the development of a Go program.

## Type Methods

A type method is a function that is attached to a specific data type. Although type methods (or methods on types) are in reality functions, they are defined and used in a slightly different way. Defining new type methods is as simple as creating new functions, provided that we follow certain rules that associate the function with a data type.

Having a data type called `ar2x2`, we can create a type method named `FunctionName` for it as follows:

```go
func (a ar2x2) FunctionName() <returnType> {
    // Function body
}
```

The `(a ar2X2)` part is what makes the `FunctionName()` function a type method because it associates `FunctionName()` with `ar2x2` data type. No other data type can use that function. However, we are free to implement `FucntionName()` for other data types or as a regular function. If we habe a `ar2X2` variable named `varAr`, we can invoke `FunctionName()` as `varAr.FunctionName()`, which looks like selecting the field of a structure variable.

Keep in mind that under the hood, the Go compiler does turn methods into regular function calls with the `self` value as the first parameter. However, interfaces require the use of type methods to work.

Most of the time, and when there is such a need, the results of a type method are saved in the variable that invoked the type method - in order to implement that for the `ar2x2` data type, we pass a pointer to the array that invoked the type method, so that the function signature looks like: `func (a *ar2x2) FunctionName() <returnType>`. This way, the type method can modify the array that invoked it.

If we are defining type methods for a structure, we should make sure that the names of the type methods do not conflict with any field name of the structure because the Go compiler will reject such ambiguities.

## Interfaces

An interface is a Go mechanism for defining behave that is implemented using a set of methods. Interfaces play a key role in Go and can simplify the code of our programs when they have to deal with multiple data types that perform the same task.

Interfaces work with methods on types (or type methods), which are like functions attached to given data types, which in Go are usually structures(although we can use any data type we want). Once we implement the required type methods of an interface, that interface is satisfied implicitly. The empty interface is defined as just `interface{}`. As the empty interface has no methods, it means that it is already implemented by all data types.

In a more formal way, a Go interface type defines (or describes) the behavior of other types by specifying a set of methods that need to be implemented to support the behavior. For a data type to satisfy an interface, it needs to implement all the type methods required by that interface. Therefore, interfaces are abstract types that specify a set of methods that need to be implemented  so that another type can be considered an instance of the interface. So, an interface is two things, a set of methods and a type.

### Advantages

The biggest advantage we get from interfaces is that, if needed, we can pass a variable of a data type that implements a particular interface to any function that expects a parameter of that specific interface, which saves us from having to write separate functions of each supported data type. However, Go offers an alternative to this with recent addtions of generics to the language.

Interfaces can also be used to provide a kind of polymorphism in Go, which is an object-oriented concept. Polymorphism offers a way of accessing objects of different types in the same uniform way when they share a common behavior.

Lastly, interfaces can be used for composition. In practice, this means that we can combine existing interfaces and create new ones that offer the combined behavior of the interfaces that were brought together.

When we combine existing interfaces , it is better that the interfaces do not contain methods with the same name. What we should keep in mind is that there is no need for an interface to be impressive and require the implementation of a large number of methods. In fact, the fewer methods an interface has, the more generic and widely used it can be, which improves its usefulness and, therefore, its usage.

### The Empty Interface

The empty interface is defined as just `interface{}` and is already implemented by all data types. This, variables of any data type can be put in the place of a parameter of the empty interface data type. Thus, a function with an `interface{}` parameter can accept variables of any data type in this place. However, if we intend to work with `interface{}` function parameters without examining their data type inside the function, we should process them with statements that work on all data types; otherwise, our code may crash or misbehave.

## Type Assertions and Type Switches

A type assertion is a mechanism for working with the underlying concrete value of an interface. This mainly happens because
interfaces are virtual data types without their own values - interfaces just define behavior and do not hold data of
their own. But what happens when we do not know the data type before attempting a type assertion? How can we differentiate
between the supported data types and unsupprted ones?

## Type switches

Type switches use `switch` blocks for data types and allow us to differentiate between type assertion values, which are 
data types and process each data type the way we want. On the other hand, in order to use the empty interface in type switches,
we need to use type assertions.

Type assertions use the `x.(T)` notation, where `x` is an interface type and `T` is a type, and help us extract the value that
is hidden behind the empty interface. For a type assertion to work, `x` should not be `nil` and the dynamic type `x` should
be identical to the `T` type


### Benefits of type assertions

Strictly speaking, type assertions allow us to perform two main tasks:

- Checking whether an interface value keeps a particular type: When used this way, a type assertion returns two values: the underlying value and a bool value. The underlying value is what we might want to use. However, it is the value of the `bool` variable that tells us whether the type assertion was successful or not and, therefore,whether can use the underlying value or not. Checking whether a variable name `aVar` is of the type `int` type requires the use of the `aVar.(int)` notation, which returns two values. If successful, it returns the real `int` value of `aVar` and `true`. Otherwise, it returns `false` as the second value, whicn means that the type assettion was not successulf and that the real value could not be extracted.
- Using the concrete value stored in an interface or assigning it to a new variable. Type assertions in Go enable us to access the actual value stored within an interface or assign it to a new variable.To illustrate this, let’s consider an example where we have an interface variable called myInterface that holds a value of type float64. By performing a type assertion on myInterface, we can extract the underlying float64 value. This allows us to work directly with the numeric value and perform operations on it. Without the type assertion, we would only have access to the interface’s methods and not the specific value it encapsulates. The type assertion allows us to retrieve and utilize the concrete value stored in the interface, in this case, the float64 value.

## Safely Handling Arbitrary Data

The biggest advantage we get from using a `map[string]interface{}` map, or any map that stores an `interface{}` value, in general, is that we still have our data in its original state and data type. If we use `map[string]string` instead, or anything similar, then any data we have is going to be converted into a `string`, which means that we are going to lose information about the original data type and the structure of he data we are storing in the map.


### The `error` interface

The `error` data type, which is an interface defined as follows:

```go
type error interface{
	Error() string
}
```

So, in order to satisfy the `error` interface, we just need to implement the `Error() string` type method. This does not change the way we use errors to find out whether the execution of a function or method was successful or not, but it shows how important interfaces are in Go as they are being used transparrently all the time.

After learning about using existing interfaces, we will write another command-line utility that sorts 3D shapes according to their volumes. This taks requires learning the following tasks:
- Creating new interfaces
- Combining existing interfaces
- Implementing `sort.Interface` for 3D shapes


### Implementing `sort.Interface` for 3D shapes

Create a utility for sorting various 3D shapes based on their volume. We will use a single slice for storing all kinds of data structures that all satisfy a given interface. The fact that Go considers interfaces as data types allows us to create slices with elements that satisfy a given interface without getting any error messages.

This kind of scenario can be useful in various cases because it illustrates how to store elements with different data types that all satisfy a common interface on the same slice and how to sort them using `sort.Interface`. Put simply, the presented utility sorts different structures with different numbers and names of fields that all share a common behavior through an interface implementation. The dimensions of the shapes are created using random numbers, which means that each time we execute the utility, we get a different output.

## Limitations of Object-Oriented Programming in Go

As Go does not support all object-oriented features, it can't replace an object-oriented programming language fully. However, it can mimic some object-oriented concepts.

First of all, a Go structure with its type methods is like an object with its methods. Second, interfaces are like abstract data types that define behaviors and objects of the same class, which is similar to polymorphism. Third, Go supports encapsulation, which means it supports hiding data and functions from the user by making them private to the structure and the current Go package. Lastly, combining interfaces and structures is like composition in object-oriented terminology.

