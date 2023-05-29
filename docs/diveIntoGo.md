# Dive into Go

https://go.dev/doc/effective_go

📖 Instruction:

- ✅: Completed goal
- 👉: Short-term goal
- 📝: Long-term goal
- 📖: Some relative note (📌: highlight)
- 🤔: Not sure the credit
- 🔜: Haven't completed but browsed
- ❌: Abandoned or failed

🫡 Reading list:

- https://go.dev/tour/concurrency/2
- https://go.dev/tour/concurrency/4
- https://go.dev/tour/concurrency/5
- https://go.dev/tour/concurrency/7

## 🤖 Robot summary to my random questions

### 📖 Can you give me an example of struct in Go with using BACKTICK ANNOTATION?

A struct in Go can use **BACKTICK ANNOTATION** to specify tags for the fields. Tags are optional string literals that can 📌**provide extra information or metadata for the fields, such as how to encode or validate them**. Tags are made visible through a reflection interface and take part in type identity for structs but are otherwise ignored².

For example, you can use backtick annotation to specify **how a struct should be encoded as JSON**:

```Go
type Employee struct {
    name string `json:"name"`
    age int `json:"age"`
    job string `json:"job"`
    salary int `json:"salary"`
}
```

**The json tag tells the json package to use the specified name as the key for each field in the JSON object.** For example, if you have an Employee value like this:

```Go
emp1 := Employee{"Alice", 25, "Developer", 5000}
```

You can encode it as JSON **using the `json.Marshal` function**:

```Go
data, err := json.Marshal(emp1)
```

The data variable will contain a JSON object like this:

```json
{ "name": "Alice", "age": 25, "job": "Developer", "salary": 5000 }
```

If you want to use a backtick inside a tag value, **you can use regular quotes instead of backticks and escape the characters as needed³**. For example:

```Go
type User struct {
    email string "validate:\"regexp=`\""
    password string `validate:"min=8"`
}
```

Source: Conversation with Bing, 5/6/2023(1) How to escape a backtick inside a Go struct tag - Stack Overflow. https://stackoverflow.com/questions/61682237/how-to-escape-a-backtick-inside-a-go-struct-tag Accessed 5/6/2023.
(2) Structs in Go (Golang) | Detailed Tutorial with Examples | golangbot.com. https://golangbot.com/structs/ Accessed 5/6/2023.
(3) go - What is the usage of backtick in golang structs definition? - Stack Overflow. https://stackoverflow.com/questions/30681054/what-is-the-usage-of-backtick-in-golang-structs-definition Accessed 5/6/2023.
(4) Proposal: Allow backticks in raw strings (by doubling them) · Issue #18221 · golang/go · GitHub. https://github.com/golang/go/issues/18221 Accessed 5/6/2023.

# Topics

## 🌶️ interface

🫡 Reading list:

- https://go.dev/tour/methods/1
  - https://go.dev/tour/methods/4
  - https://go.dev/tour/methods/7
  - In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.
- 📌 https://go.dev/tour/methods/9
  - [Interfaces are implemented implicitly](https://go.dev/tour/methods/10)
  - [Interface values with nil underlying values](https://go.dev/tour/methods/12)

### 📖 How to ASSERT A VARIABLE IS A CERTAIN TYPE in Go? Can you give more examples such as the provided code `publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)`? and explain more about it.

A type assertion is a way of checking if an 📌**interface value** holds 🤔**a specific type and accessing the underlying value of that type**. The syntax for a type assertion is 📌**`interfaceVariable.(type)`**, where interfaceVariable is an expression of an interface type and **type is the asserted type**. A type assertion can return either one or two values, depending on how it is used.

If a type assertion is used as a single-value expression, such as **`v := i.(string)`**, it will return the underlying value of the interface variable if it holds the asserted type, or 📌**panic if it does not**. 🤔**This can be useful when you are certain that the interface value has a specific type and you want to access its value directly.**

If a type assertion is used as a two-value assignment, such as **`v, ok := i.(string)`**, **it will return the underlying value of the interface variable and a boolean value that indicates whether the assertion succeeded or not**. If the interface value does not hold the asserted type, 📌**it will not panic, but instead return the zero value of the asserted type and false**. This can be useful when you want to check if the interface value has a specific type and handle the error gracefully.

Here are some more examples of type assertions in Go:

- **i.(int)**: asserts that the interface value `i` holds an `int` and returns its value
- **i.(fmt.Stringer)**: asserts that 🤔**the interface value `i` satisfies the `fmt.Stringer` interface** and returns its value
- **i.(interface{})**: asserts that 🤔**the interface value `i` holds 🤔any type and returns its value**
- **i.(nil)**: asserts that the interface value `i` is `nil` and returns `nil`

Source: Conversation with Bing, 5/8/2023(1) A Tour of Go. https://go.dev/tour/methods/15 Accessed 5/8/2023.
(2) Type assertions in Golang - Golang Docs. https://golangdocs.com/type-assertions-in-golang Accessed 5/8/2023.
(3) Golang Type Assertion. Using interfaces to determine behavior | by Yair Fernando | Better Programming - Medium. https://betterprogramming.pub/golang-type-assertion-d5517d81c366 Accessed 5/8/2023.
(4) Golang Type Assertion Explained with Examples | GoLinuxCloud. https://www.golinuxcloud.com/golang-type-assertion/ Accessed 5/8/2023.

### Can I assert multiple types at once in Go?

A type assertion in Go is a way of checking and extracting the dynamic type of an interface value¹. **For example, if you have a variable x of type `interface{}` and you want to check if it holds a string value, you can write:**

```go
s, ok := x.(string)
```

This will assign the string value of x to s and a boolean value to ok, indicating whether the assertion was successful or not.

If you want to assert multiple types at once, you can use a switch statement with type cases¹. For example:

```go
switch v := x.(type) {
    case string:
        // v is a string
    case int:
        // v is an int
    case error:
        // v is an error
    default:
        // v is some other type
}
```

This will assign the value of x to v with the appropriate type in each case.

You cannot use a type assertion on a function that returns multiple values, such as foobar() in your example. You need to assign the return values to variables first, and then use type assertions on them².

Source: Conversation with Bing, 5/8/2023
(1) Explain Type Assertions in Go - Stack Overflow. https://stackoverflow.com/questions/38816843/explain-type-assertions-in-go.
(2) Multiple return types with interface {} and type assertions (in Go). https://stackoverflow.com/questions/7045848/multiple-return-types-with-interface-and-type-assertions-in-go.
(3) How to assert on error types in Go tests? - Stack Overflow. https://stackoverflow.com/questions/57718024/how-to-assert-on-error-types-in-go-tests.
(4) assert package - github.com/stretchr/testify/assert - Go Packages. https://pkg.go.dev/github.com/stretchr/testify/assert.

### 🌶️ Show me more examples of using interfaces in Go.

🙋‍♂️: IMO, after browsing documents about Go interfaces, I regard Go interfaces as a blueprint (or a DEFINATION of a group of types) to let future coder have ability to reuse them. However, it is also easy to compose, since the composing way to meet the blueprint is easy to achieve. In the past, coder have to inherit the parent class then have abiliy to reuse the definition, but the implemented class can't conveniently be the implementation of the new definition. In this case, any old struct in Go can easily meet the new blueprint with a requiring function. (Furtermore, interfaces is the best way to abstract a struct.)

Sure, I can show you some more examples of using interfaces in Go. Here are some common scenarios where interfaces are useful:

- **Writing generic and reusable code that works with different types that share a common behavior.** For example, you can use the `io.Reader` and `io.Writer` interfaces to **read and write data from various sources and destinations, such as files, buffers, network connections, etc**. 📌**You don't need to know the specific type of the data source or destination, as long as it implements the `Read()` or `Write()` methods of the interface³.**
- **Defining function parameters or return values that can accept any type that implements an interface.** For example, you can use the `sort.Interface` interface to **sort any collection of elements that implements the `Len()`, `Less()` and `Swap()` methods of the interface.** You can also use the error interface to return any type that implements the Error() method of the interface⁴.
- 📌**Implementing multiple interfaces by a single type to achieve composition and polymorphism.** For example, you can define a type that implements both the `fmt.Stringer` and the `json.Marshaler` interfaces, which means it can provide a string representation of itself and also a JSON encoding of itself³.
- **Embedding interfaces within other interfaces to create larger interfaces that inherit the methods of the embedded interfaces.** For example, you can define an interface that embeds the `io.Reader` and `io.Writer` interfaces, which means it inherits both the `Read()` and `Write()` methods of those interfaces².

Here are some links to more examples of using interfaces in Go:

- https://www.programiz.com/golang/interface
- https://golangbot.com/interfaces-part-1/
- https://www.alexedwards.net/blog/interfaces-explained
- https://stackoverflow.com/questions/1734673/what-are-some-examples-of-go-interfaces

Source: Conversation with Bing, 5/8/2023
(1) Golang Interfaces Explained – Alex Edwards. https://www.alexedwards.net/blog/interfaces-explained.
(2) What are some examples of Go interfaces? - Stack Overflow. https://stackoverflow.com/questions/1734673/what-are-some-examples-of-go-interfaces.
(3) A Practical Guide to Interfaces in Go (Golang) - golangbot.com. https://golangbot.com/interfaces-part-1/.
(4) Go Interface (With Examples) - Programiz. https://www.programiz.com/golang/interface.

### 📖 What is the DIFFERENCE between TYPE ASSERTION and TYPE CONVERSION?

Type assertion and type conversion are two different operations in Go that can change the type of a variable. The main difference is that type assertion is used to **check and access THE UNDERLYING TYPE OF AN INTERFACE VALUE**, while type conversion is used to **change one NON-INTERFACE TYPE to another**.

Type assertion has the syntax **interfaceVariable.(type)**, where interfaceVariable is an expression of an interface type and type is the asserted type. A type assertion can return either one or two values, depending on how it is used. If a type assertion fails, it will panic unless it is used as a two-value assignment.

Type conversion has the syntax **type(variable)**, where type is the desired type and variable is the expression to be converted. A type conversion can only return one value, which is the converted value of the variable. **If a type conversion is not possible, it will cause a compile-time error.**

Here are some examples of type assertion and type conversion in Go:

- **i.(int)**: asserts that the interface value `i` holds an int and returns its value
- **i.(fmt.Stringer)**: asserts that the interface value `i` satisfies the `fmt.Stringer` interface and returns its value
- **i.(interface{})**: asserts that the interface value `i` holds any type and returns its value
- **i.(nil)**: asserts that the interface value `i` is `nil` and returns `nil`
- **int(x)**: converts the value `x` to an `int` and returns its value
- **string(b)**: converts the byte slice `b` to a `string` and returns its value
- **float64(n)**: converts the numeric value `n` to a `float64` and returns its value

Source: Conversation with Bing, 5/8/2023(1) go - What is the difference between type conversion and type assertion? - Stack Overflow. https://stackoverflow.com/questions/20494229/what-is-the-difference-between-type-conversion-and-type-assertion Accessed 5/8/2023.
(2) Type assertions vs. type conversions in Go - LogRocket Blog. https://blog.logrocket.com/type-assertions-vs-type-conversions-go/ Accessed 5/8/2023.
(3) Type Assertions vs Type Conversions in Golang - Soham Kamani. https://www.sohamkamani.com/golang/type-assertions-vs-type-conversions/ Accessed 5/8/2023.

### 🔜 Can I convert any type to any other type?

No, you cannot convert any type to any other type in Go. There are some rules and restrictions for type conversions, depending on the types involved. Here are some of them:

- You can only convert **a non-constant value to another non-constant type**. For example, you cannot convert a string literal to an int, but you can convert a string variable to an int using the strconv.Atoi function.
- You can only convert a value to a type that can represent its value. For example, you cannot convert a negative int to an unsigned int, or a large int to a byte, as they would cause overflow or truncation errors.
- You can only convert a value to a type that has the same underlying type as its original type, or to an interface type that it implements. For example, you can convert a named int type to an int, or an int to an interface{}, but you cannot convert an int to a string directly.
- You can only convert a pointer value to another pointer type if they have the same underlying base types, or to an unsafe.Pointer type. For example, you can convert a *int to a *uint if they have the same size, or to an unsafe.Pointer, but you cannot convert a *int to a *string.
- You can only convert a slice value to another slice type if they have the same element types. For example, you can convert a []byte to a []uint8, but you cannot convert a []byte to a []int.
- You can only convert a map value to another map type if they have the same key and element types. For example, you can convert a map[string]int to a map[string]uint, but you cannot convert a map[string]int to a map[int]string.
- You can only convert a function value to another function type if they have the same parameter and result types. For example, you can convert a func(int)int to a func(x int)int, but you cannot convert a func(int)int to a func(int)string.
- You can only convert a channel value to another channel type if they have the same element types and direction. For example, you can convert a chan int to a chan<- int, but you cannot convert a chan int to a chan string or a <-chan int.

Source: Conversation with Bing, 5/8/2023(1) The Go Programming Language Specification. https://go.dev/ref/spec Accessed 5/8/2023.
(2) Idiomatic Type Conversion in Go - Stack Overflow. https://stackoverflow.com/questions/13367979/idiomatic-type-conversion-in-go Accessed 5/8/2023.
(3) Value Conversion, Assignment and Comparison Rules in Go. https://go101.org/article/value-conversions-assignments-and-comparisons.html Accessed 5/8/2023.

## 🌶️ Composition and polymorphism

### So there is no class in Go?

That’s right. Go does not have classes in the traditional sense of object-oriented programming. Instead, Go uses structs to define custom types that can have fields and methods. Structs can also embed other structs or interfaces to achieve composition and polymorphism.

### My futher question is that would there be a downside in this case?

That depends on what you mean by "this case". If you are referring to using structs instead of classes, then there are some trade-offs to consider. Some possible downsides are:

- You cannot have private fields or methods in a struct. 🤔**All fields and methods are either exported (public) or unexported (package-level).** This means you have less control over the encapsulation and abstraction of your types.
- **You cannot have constructors or destructors in a struct.** You have to use functions or methods to initialize or finalize your structs. **This means you have to be more careful about memory management and resource allocation.** (use `defer` instead)
- You cannot have inheritance or subtyping in a struct. **You can only use embedding or interfaces to achieve composition and polymorphism. This means you have to be more explicit about the relationships and behaviors of your types.**

Of course, these downsides also come with some benefits, such as simplicity, performance, and flexibility. Go's approach to structs and interfaces is designed to encourage composition over inheritance, which is often considered a better design principle in software engineering. Ultimately, it depends on your preferences and needs as a programmer.

### Show me how to use functions or methods to initialize or finalize structs in Go.

To initialize or finalize structs in Go, you can use functions or methods that take a struct type or a pointer to a struct type as a parameter or a receiver, and perform some operations on the struct fields. For example:

- To initialize a struct, you can use a function that returns a pointer to a struct with some default or computed values for the fields. This is similar to a constructor in other languages. For example¹:

```go
package people

type Person struct {
    name string
    age  int
}

// NewPerson is a constructor function for Person
func NewPerson(name string, age int) *Person {
    if age < 0 {
        panic("NewPerson: age cannot be a negative number")
    }
    return &Person{name: name, age: age}
}
```

- To finalize a struct, you can use a method that performs some cleanup or release actions on the struct fields. This is similar to a destructor in other languages. For example²:

```go
package files

type File struct {
    name string
    data []byte
}

// Close is a finalizer method for File
func (f *File) Close() error {
    // release the memory used by the data field
    f.data = nil
    // perform other cleanup actions
    // ...
    return nil
}
```

You can also use the defer keyword to ensure that the finalizer method is called when the function scope exits. For example:

```go
package main

import "files"

func main() {
    // create a new file
    f := files.NewFile("test.txt")
    // defer the close method
    defer f.Close()
    // do something with the file
    // ...
}
```

This way, you don't have to worry about forgetting to call the close method manually. The defer statement will ensure that it is called automatically when the main function returns.

Source: Conversation with Bing, 5/8/2023
(1) Different Ways to Initialize Go structs | Anton Sankov's Blog. https://asankov.dev/blog/2022/01/29/different-ways-to-initialize-go-structs/.
(2) Struct initialization and method declaration in Go. https://stackoverflow.com/questions/32007158/struct-initialization-and-method-declaration-in-go.
(3) go - How to initialize struct fields - Stack Overflow. https://stackoverflow.com/questions/30428571/how-to-initialize-struct-fields.

### Can I forbid user from using a variable in a struct directly?

Yes, you can. In Go, the visibility of a variable or a field is determined by its first letter. If the first letter is uppercase, then the variable or field is exported (public), which means it can be accessed from outside the package where it is defined. If the first letter is lowercase, then the variable or field is unexported (private), which means it can only be accessed from within the same package where it is defined.

For example, if you have a struct type called Person with a field called name, and you want to forbid users from using the name field directly, you can make the name field unexported by using a lowercase n:

```go
package people

type Person struct {
    name string // unexported field
    Age  int    // exported field
}
```

Now, users can only access the name field from within the people package. If they try to access it from another package, they will get a compilation error. For example:

```go
package main

import (
    "fmt"
    "people"
)

func main() {
    p := people.NewPerson("Alice", 25) // assume NewPerson is a constructor function
    fmt.Println(p.name) // compilation error: p.name undefined (cannot refer to unexported field or method name)
    fmt.Println(p.Age)  // OK
}
```

If you want to allow users to read or write the name field indirectly, you can provide getter and setter methods for the Person type. For example:

```go
package people

type Person struct {
    name string // unexported field
    Age  int    // exported field
}

// NewPerson is a constructor function for Person
func NewPerson(name string, age int) *Person {
    if age < 0 {
        panic("NewPerson: age cannot be a negative number")
    }
    return &Person{name: name, age: age}
}

// GetName is a getter method for the name field
func (p *Person) GetName() string {
    return p.name
}

// SetName is a setter method for the name field
func (p *Person) SetName(name string) {
    if name == "" {
        panic("SetName: name cannot be empty")
    }
    p.name = name
}
```

Now, users can use the GetName and SetName methods to access the name field indirectly. For example:

```go
package main

import (
    "fmt"
    "people"
)

func main() {
    p := people.NewPerson("Alice", 25)
    fmt.Println(p.GetName()) // OK: prints Alice
    p.SetName("Bob")        // OK: changes the name to Bob
    fmt.Println(p.GetName()) // OK: prints Bob
}
```

**This way, you can control how users can access the name field and enforce some validation or logic on it. You can also change the implementation of the name field without affecting the users of your package.**

### Can I create a private function for a struct?

Yes, you can. Similar to fields, the visibility of a function or a method is determined by its first letter. If the first letter is uppercase, then the function or method is exported (public), which means it can be accessed from outside the package where it is defined. If the first letter is lowercase, then the function or method is unexported (private), which means it can only be accessed from within the same package where it is defined.

For example, if you have a struct type called Person with a method called greet, and you want to make the greet method private, you can use a lowercase g:

```go
package people

type Person struct {
    name string // unexported field
    Age  int    // exported field
}

// greet is an unexported method for Person
func (p *Person) greet() string {
    return "Hello, I am " + p.name
}
```

Now, the greet method can only be called from within the people package. If you try to call it from another package, you will get a compilation error. For example:

```go
package main

import (
    "fmt"
    "people"
)

func main() {
    p := people.NewPerson("Alice", 25) // assume NewPerson is a constructor function
    fmt.Println(p.greet()) // compilation error: p.greet undefined (cannot refer to unexported field or method greet)
}
```

If you want to allow users to call the greet method indirectly, you can provide another exported method that wraps the greet method. For example:

```go
package people

type Person struct {
    name string // unexported field
    Age  int    // exported field
}

// NewPerson is a constructor function for Person
func NewPerson(name string, age int) *Person {
    if age < 0 {
        panic("NewPerson: age cannot be a negative number")
    }
    return &Person{name: name, age: age}
}

// greet is an unexported method for Person
func (p *Person) greet() string {
    return "Hello, I am " + p.name
}

// Greet is an exported method for Person that wraps the greet method
func (p *Person) Greet() string {
    return p.greet()
}
```

Now, users can use the Greet method to call the greet method indirectly. For example:

```go
package main

import (
    "fmt"
    "people"
)

func main() {
    p := people.NewPerson("Alice", 25)
    fmt.Println(p.Greet()) // OK: prints Hello, I am Alice
}
```

This way, you can hide the implementation details of the greet method and change it without affecting the users of your package. You can also use unexported functions or methods for internal logic or helper functions that are not meant to be exposed to the users of your package.

### Show me how to use unexported functions or methods for internal logic or helper functions.

Unexported functions or methods are useful for internal logic or helper functions that are not meant to be exposed to the users of your package. They can perform some common tasks or calculations that are used by multiple exported functions or methods, or they can encapsulate some implementation details that are irrelevant or confusing for the users.

To use unexported functions or methods for internal logic or helper functions, you can follow these steps:

- Define your unexported function or method with a lowercase name in the same package where you want to use it. For example:

```go
package math

// square is an unexported function that returns the square of a number
func square(x int) int {
    return x * x
}

// SumOfSquares is an exported function that returns the sum of squares of two numbers
func SumOfSquares(a, b int) int {
    return square(a) + square(b)
}
```

- Call your unexported function or method from within the same package where you defined it. You don't need to import anything or use any qualifiers. For example:

```go
package math

import "fmt"

func main() {
    // call the unexported function square
    fmt.Println(square(5)) // prints 25
    // call the exported function SumOfSquares
    fmt.Println(SumOfSquares(3, 4)) // prints 25
}
```

- Do not try to call your unexported function or method from outside the package where you defined it. You will get a compilation error. For example:

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    // try to call the unexported function square from math package
    fmt.Println(math.square(5)) // compilation error: math.square undefined (cannot refer to unexported name math.square)
    // call the exported function SumOfSquares from math package
    fmt.Println(math.SumOfSquares(3, 4)) // prints 25
}
```

- If you want to test your unexported function or method, you can use a test file with the same package name as your function or method. For example¹:

```go
package math

import "testing"

// TestSquare is a test function for the unexported function square
func TestSquare(t *testing.T) {
    tests := []struct {
        input  int
        output int
    }{
        {0, 0},
        {1, 1},
        {2, 4},
        {-1, 1},
        {-2, 4},
    }
    for _, test := range tests {
        result := square(test.input)
        if result != test.output {
            t.Errorf("square(%d) = %d; want %d", test.input, result, test.output)
        }
    }
}
```

This way, you can use unexported functions or methods for internal logic or helper functions in your package.

Source: Conversation with Bing, 5/8/2023
(1) go - Getting "implicit assignment of unexported field" - Stack Overflow. https://stackoverflow.com/questions/61412331/getting-implicit-assignment-of-unexported-field.
(2) How to test a unexported (private) function in go (golang)?. https://stackoverflow.com/questions/24622388/how-to-test-a-unexported-private-function-in-go-golang.
(3) View documentation for internal / unexported function?. https://stackoverflow.com/questions/59623734/view-documentation-for-internal-unexported-function.
(4) Internal functions in R packages | R-bloggers. https://www.r-bloggers.com/2019/12/internal-functions-in-r-packages/.

## 🌶️ Concurrency

🫡 Reading list:

- https://go.dev/tour/concurrency/8
  - https://stackoverflow.com/questions/12224042/go-tour-exercise-7-binary-trees-equivalence
  - https://go.dev/tour/concurrency/4
  - https://pkg.go.dev/golang.org/x/tour/tree#pkg-index
- 🌶️ [sync.Mutex](https://go.dev/tour/concurrency/9)

```go
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			ch <- 0
			return
		}
		ch <- t.Value
		walk(t.Left)
		walk(t.Right)
	}
	walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if ok1 == false {
			break
		}
	}
	return true
}

func main() {
	t1, t2, t3 := tree.New(10), tree.New(10), tree.New(20)
	c := make(chan int)
	go Walk(t1, c)
	for tVal := range c {
		fmt.Printf("%v, ", tVal)
	}
	fmt.Printf("\n %v", Same(t1, t1))
	fmt.Printf("\n %v", Same(t1, t2))
	fmt.Printf("\n %v", Same(t1, t3))
}
```

- https://go.dev/tour/concurrency/10
  - https://go.dev/tour/concurrency/3
  - 📌 https://stackoverflow.com/questions/40326723/go-vet-range-variable-captured-by-func-literal-when-using-go-routine-inside-of-f
  - 📌 https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver
  - 📌 https://stackoverflow.com/questions/44370277/type-is-pointer-to-interface-not-interface-confusion
  - https://yourbasic.org/golang/for-loop/

```go
package main

import (
	"fmt"
	"sync"
)

type Cache interface {
	Set(key string, data string)
	Get(key string) (string, bool)
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses the given fetcher (sth. like dependency injection)
// to recursively crawl pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache Cache) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	if _, isExist := cache.Get(url); isExist {
		fmt.Println(url, "has been traced. SKIP")
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	cache.Set(url, body)
	fmt.Printf("found: %s %q\n", url, body)
	done := make(chan bool, len(urls))
	for _, u := range urls {
		go func(u string) {
			Crawl(u, depth-1, fetcher, cache)
			done <- true
		}(u)
	}

	for cunt := 0; cunt < len(urls); cunt++ {
		<-done
	}
	return
}

func main() {
	var cache Cache
	cache = &SafeCache{data: make(map[string]string)}
	Crawl("https://golang.org/", 4, fetcher, cache)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

type SafeCache struct {
	mu   sync.Mutex
	data map[string]string
}

func (c *SafeCache) Set(key string, data string) {
	c.mu.Lock()
	c.data[key] = data
	c.mu.Unlock()
}

func (c *SafeCache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	res, isExist := c.data[key]
	return res, isExist
}
```

## 🔜 Design parrtern

### 🔜 How to achieve "dependency injection" in Go?

Dependency injection is a design pattern that allows you to write code 📌**that depends on abstractions rather than concrete implementations**. It helps you to **decouple your components and make them easier to test and reuse**. In Go, dependency injection is often achieved using 📌**interfaces**, which define a set of methods that a concrete type must implement¹³. By injecting dependencies at runtime, you can easily swap out implementations for testing or for different environments³. For example, you can inject a mock database or a logger into your component instead of creating them inside the component.

There are different ways to achieve dependency injection in Go, such as 🤔**using constructor functions**, 🤔**setter methods**, or 🤔**struct fields**³. You can also use code generation tools like 🤔**Wire**¹, which automates connecting components using dependency injection. Wire uses **function parameters** to represent dependencies and generates code that initializes your components without reflection or runtime state¹.

Source: Conversation with Bing, 5/10/2023
(1) GitHub - google/wire: Compile-time Dependency Injection for Go. https://github.com/google/wire.
(2) Dependency Injection in Go: The better way | by Cloud_Freak | FAUN Publication. https://faun.pub/dependency-injection-in-go-the-better-way-833e2cc65fab.
(3) What is Dependency Injection in Go? | Developer.com. https://www.developer.com/languages/golang-dependency-injection/.
(4) Dependency Injection in Go - drewolson.org. https://blog.drewolson.org/dependency-injection-in-go.
