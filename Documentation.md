# Golang

## Variables

A variable stores a value that can be used and changed while a Go program is running. Go supports type inference, explicit types, and multiple variable declarations.

### Example

```go
package main

import "fmt"

var num = 100
var name = "Anto"
var age int = 20
var a, b, c int = 1, 2, 3

func main() {
    subject := "Maths"

    fmt.Println(num, name, age, subject)
    fmt.Println(a, b, c)
}
```

---

## Constants

A constant is a value that cannot be changed after it is declared. Constants are useful for fixed values such as mathematical constants, application limits, or configuration values.

### Example

```go
package main

import "fmt"

const PI = 3.14

func main() {
    isActive := true

    fmt.Println("PI:", PI)
    fmt.Println("Active:", isActive)
}
```

---

## Arithmetic Operators

Arithmetic operators are used to perform mathematical calculations such as addition, subtraction, multiplication, division, and remainder.

| Operator | Meaning |
|---|---|
| `+` | Addition |
| `-` | Subtraction |
| `*` | Multiplication |
| `/` | Division |
| `%` | Remainder |

### Example

```go
package main

import "fmt"

func main() {
    fmt.Println("Add:", 2+2)
    fmt.Println("Sub:", 5-2)
    fmt.Println("Mul:", 3*4)
    fmt.Println("Div:", 10/2)
    fmt.Println("Mod:", 10%3)
}
```

---

## Comparison Operators

Comparison operators compare two values and return a boolean value: `true` or `false`.

| Operator | Meaning |
|---|---|
| `==` | Equal |
| `!=` | Not equal |
| `>` | Greater than |
| `<` | Less than |
| `>=` | Greater than or equal |
| `<=` | Less than or equal |

### Example

```go
package main

import "fmt"

func main() {
    a := 10
    b := 20

    fmt.Println("a == b:", a == b)
    fmt.Println("a != b:", a != b)
    fmt.Println("a > b:", a > b)
    fmt.Println("a < b:", a < b)
    fmt.Println("a >= b:", a >= b)
    fmt.Println("a <= b:", a <= b)
}
```

---

## Strings

A string represents text in Go. Double quotes are commonly used for normal strings. Backticks can be used for raw multi-line strings.

### Example

```go
package main

import "fmt"

func main() {
    name := "Hello\nworld"

    fmt.Println(name)
}
```

### Raw multi-line string

```go
package main

import "fmt"

func main() {
    message := `Hello
World
Welcome to Go`

    fmt.Println(message)
}
```

---

## String Comparison

The `strings.Compare` function compares two strings. It returns `0` when they are equal, a negative value when the first string sorts before the second, and a positive value when it sorts after the second.

### Example

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    firstName := "Anu"
    lastName := "Anu"

    fmt.Println(strings.Compare(firstName, lastName))
}
```

If both strings are equal, the result is:

```text
0
```

---

## String Contains

`strings.Contains` checks whether one string exists inside another string. It returns `true` or `false`.

### Example

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    firstName := "Golang"
    lastName := "Go"

    fmt.Println(strings.Contains(firstName, lastName))
}
```

Output:

```text
true
```

---

## For Loop

The `for` loop repeats a block of code. Go uses `for` as its main looping keyword.

### Example

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

Output:

```text
0
1
2
3
4
```

---

## continue

`continue` skips the remaining statements in the current loop iteration and moves to the next iteration.

### Example

```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        if i == 5 {
            continue
        }

        fmt.Println(i)
    }
}
```

Here, `5` is skipped.

---

## break

`break` immediately stops the current loop.

### Example

```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        if i == 5 {
            break
        }

        fmt.Println(i)
    }
}
```

The loop stops when `i` becomes `5`.

---

## Arrays

An array is a fixed-size collection of values of the same type. The array length is part of its type.

### Example

```go
package main

import "fmt"

func main() {
    var arr [5]int

    arr[0] = 1
    arr[1] = 2
    arr[2] = 3
    arr[3] = 4
    arr[4] = 5

    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }
}
```

Output:

```text
1
2
3
4
5
```

---

## Slice

A slice in Go is a flexible, dynamic collection of elements of the same type. Unlike an array, a slice does not have a fixed length, so you can add or remove elements as needed.

### Example

```go
package main

import (
    "fmt"
    "golang.org/x/exp/slices"
)

func main() {
    fmt.Println("slices")

    s := []string{"a", "b", "c"}

    fmt.Println(s)

    s = append(s, "d", "e")

    fmt.Println(s)

    s = slices.Delete(s, 1, 2)

    fmt.Println(s)
}
```

The slice starts with:

```text
[a b c]
```

After `append`:

```text
[a b c d e]
```

After deleting the element from index `1` up to, but not including, index `2`:

```text
[a c d e]
```

---

## Maps

A map stores key-value pairs. It is useful when you want to look up a value using a key.

### Example

```go
package main

import "fmt"

func main() {
    personMap := make(map[string]int)

    personMap["Anooj"] = 20
    personMap["Kiran"] = 25

    for key, value := range personMap {
        fmt.Println("Name:", key, "Age:", value)
    }
}
```

Here:

```text
Key   → Name
Value → Age
```

For example:

```text
"Anooj" → 20
"Kiran" → 25
```

---

## Structs

A struct groups related fields of different types into one custom data type. Structs are widely used to represent objects such as users, products, or orders.

### Example

```go
package main

import "fmt"

type Person struct {
    Name  string
    Age   int
    Score int
}

func main() {
    p1 := Person{
        Name:  "Anu",
        Age:   20,
        Score: 100,
    }

    fmt.Println(p1)
    fmt.Println(p1.Name)
}
```

A struct allows related information to be grouped together.

---

## Anonymous Structs

An anonymous struct is a struct type declared without giving the type a reusable name. It is useful for small, one-time data structures.

### Example

```go
package main

import "fmt"

func main() {
    p1 := struct {
        Name  string
        Age   int
        Score int
    }{
        Name:  "Anu",
        Age:   20,
        Score: 100,
    }

    fmt.Println(p1)
}
```

Unlike a named struct such as `Person`, this struct does not create a reusable type name.

---

## Nested Structs

A nested struct is a struct that contains another struct as one of its fields. It helps model related data hierarchically.

### Example

```go
package main

import "fmt"

type Address struct {
    City  string
    State string
}

type Person struct {
    Name    string
    Age     int
    ADR     Address
}

func main() {
    p1 := Person{
        Name: "Anto",
        Age:  20,
        ADR: Address{
            City:  "Idukki",
            State: "Kerala",
        },
    }

    fmt.Println(p1)
    fmt.Println(p1.ADR.City)
}
```

The `Person` struct contains an `Address` struct.

You can access the nested field like this:

```go
p1.ADR.City
```

---

## Functions

A function is a reusable block of code that performs a specific task. Functions can receive parameters and may or may not return a value.

### Example

```go
package main

import "fmt"

func greet(name string, age int) {
    fmt.Println("Hello", name, "you are", age, "years old")
}

func main() {
    greet("Anooj", 20)
}
```

Here:

- `name string` is a parameter.
- `age int` is a parameter.
- `greet("Anooj", 20)` calls the function.

---

## Variadic Functions

A variadic function can accept zero or more arguments of the same type. The `...` syntax is used before the parameter type.

### Example

```go
package main

import "fmt"

func greet(numbers ...int) {
    for i, num := range numbers {
        fmt.Println(i, num)
    }
}

func main() {
    greet(1, 2, 3, 4, 5)
}
```

The function can receive multiple integer arguments:

```go
greet(1, 2, 3, 4, 5)
```

---

## Functions with Return Values

A function can return a value to the code that called it. The return type is written after the parameter list.

### Example

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(1, 2)

    fmt.Println(result)
}
```

Output:

```text
3
```

The `int` after the parameter list means that the function returns an integer.

---

## Anonymous Functions

An anonymous function is a function without a name. It can be assigned to a variable and called through that variable.

### Example

```go
package main

import "fmt"

func main() {
    sum := func(a int, b int) int {
        return a + b
    }

    fmt.Println(sum(2, 4))
}
```

Output:

```text
6
```

The function is stored in the variable `sum`.

---

## Callback Functions

A callback is a function passed as an argument to another function. The receiving function can then call the callback.

### Example

```go
package main

import "fmt"

func processNumber(a, b int, callback func(int, int) int) {
    result := callback(a, b)

    fmt.Println(result)
}

func add(a, b int) int {
    return a + b
}

func multiply(a, b int) int {
    return a * b
}

func main() {
    processNumber(10, 5, add)
    processNumber(10, 5, multiply)
}
```

Here:

```go
processNumber(10, 5, add)
```

passes the `add` function as a callback.

And:

```go
processNumber(10, 5, multiply)
```

passes the `multiply` function as a callback.

---

## defer

`defer` schedules a function call to run when the surrounding function returns. Multiple deferred calls execute in last-in-first-out (LIFO) order.

### Example

```go
package main

import "fmt"

func main() {
    fmt.Println("Start")

    defer fmt.Println("Call defer function 1")
    defer fmt.Println("Call defer function 2")

    fmt.Println("End")
}
```

Output:

```text
Start
End
Call defer function 2
Call defer function 1
```

The deferred calls execute in **LIFO order**:

```text
Last defer added
        ↓
Executes first
```

---

## Pointers

A pointer stores the memory address of another variable. The `&` operator gets an address, while `*` can be used to access the value stored at that address.

### Example

```go
package main

import "fmt"

func main() {
    x := 10

    fmt.Println("Memory address:", &x)

    var ptr *int = &x

    fmt.Println("Pointer:", ptr)
    fmt.Println("Value:", *ptr)
}
```

Here:

```go
&x
```

gets the memory address of `x`.

```go
ptr
```

stores that address.

```go
*ptr
```

gets the value stored at that address.

---

# Quick Revision

| Topic | Main Idea |
|---|---|
| Variables | Store changeable values |
| Constants | Store values that cannot be changed |
| Arithmetic Operators | Perform calculations |
| Comparison Operators | Compare values |
| Strings | Work with text |
| String Comparison | Compare two strings |
| String Contains | Check whether text exists inside another string |
| For Loop | Repeat code |
| `continue` | Skip the current iteration |
| `break` | Stop the loop |
| Arrays | Fixed-size collection |
| Slices | Dynamic collection |
| Maps | Key-value storage |
| Structs | Group related fields |
| Anonymous Structs | One-time struct type |
| Nested Structs | Struct inside another struct |
| Functions | Reusable blocks of code |
| Variadic Functions | Accept multiple arguments |
| Return Values | Return data from a function |
| Anonymous Functions | Functions without names |
| Callback Functions | Pass functions as arguments |
| `defer` | Run a call when the function returns |
| Pointers | Store memory addresses |
