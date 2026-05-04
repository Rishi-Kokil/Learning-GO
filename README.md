# Introduction

I am planning to learn GO by reading the official Documentation. This section will server as the notes that I will take whil learning GO

# Installing GO

- There is a way to manage GO Version and there is a straigh forward installation
- I am going with downloading the 1.26.2 version of GO

# Getting Started

- When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a `go.mod` file.
- That go.mod file stays with your code, including in your source code repository.
- To enable dependency tracking for your code by creating a go.mod file, run the `go mod init` command, giving it the name of the module your code will be in. The name is the module's module path.
- In actual development, the module path will typically be the **repository location where your source code will be kept**. For example, the module path might be **github.com/mymodule**. If you plan to publish your module for others to use, the module path must be a location from which Go tools can download your module.

```shell
mkdir hello
cd hello
go mod init eample/hello
```

- After running this command I can see a go.mod file created and also it has the following structure

```markdown
module example/hello

go 1.26.2
```

## Writing your first code

- In the `hello` directory, I initialized the `go.mod` file and created a `hello.go` file.
- I observed that importing packages such as `fmt` requires double quotes (`"fmt"`). Using single quotes throws an error because single quotes are used for `rune literals (Single Chatacters)` in Go, not strings.
- `Println()` is a function in the `fmt` package used to print output to the console.
- Another useful observation is naming and visibility in Go: identifiers starting with an uppercase letter are exported (accessible from other packages), while identifiers starting with a lowercase letter are unexported (package-private behavior). Go does not use explicit `public` or `private` keywords.
- To run the hello world program, I use the command:

```shell
go run hello.go
```

- In the official documentation, when I am inside the `hello` directory and run the command:

```shell
go run .
```

Go builds and runs the package in the current directory (it does not select only `hello.go`). All non-test `.go` files in the same directory should belong to the same package.

```shell
PS C:\hello> go run .
# example/hello
.\hello2.go:7:6: main redeclared in this block
        .\hello.go:7:6: other declaration of main
```

- This error occurs because both files define `func main()` in the same package.
- If I provide a specific file name (instead of `.`), only that file is run, so this conflict may not appear.
- Another obervation if you run a file using the go run command and if the func main is not defined then it throws error when you run that file. (# Read about this)

## Call code in an external package

- When you need your code to do something that might have been implemented by someone else, you can look for a package that has functions you can use in your code.

```link
https://pkg.go.dev/search?q=quote
```

- This is the repository for searching packages created by others, similar to `npm`
- if you look carefully you can see a lot of versions of the package listed over there.
- In the top you can see the pacakge it is rsc.io/quote.
  ![alt text](image.png)
  You can use the pkg.go.dev site to find published modules whose packages have functions you can use in your own code. Packages are published in modules -- like rsc.io/quote -- where others can use them. Modules are improved with new versions over time, and you can upgrade your code to use the improved versions.

## `go tidy`

- When we imported the rsc.io/quote in our code for the first time it could not resolve the module as it was un aware where it is coming from
- when we run the `go tidy` command it automatically searched the package repository and adds it to the go.mod file. Basically it adds it as dependency
- When you ran go mod tidy, it located and downloaded the rsc.io/quote module that contains the package you imported. By default, it downloaded the latest version -- v1.5.2.



# GO Lang book

## Go Basics: Compilation & Structure

Go is a **compiled language**. The Go toolchain converts source code and its dependencies into native machine language instructions. All tools are accessed via the `go` command.

### Core Toolchain Commands

- **`go run`**: Compiles source code from `.go` files, links libraries, and runs the resulting executable immediately.
```bash
  $ go run helloworld.go
```
- **`go build`**: Compiles the code and saves it as a **reusable executable binary** file.
```bash
  $ go build helloworld.go
  $ ./helloworld
```
- **`go get`**: Fetches source code from a repository and places it in the corresponding directory.

### Code Organization & Packages

- **Definition**: A **package** consists of one or more `.go` source files in a single directory that define what the package does. It is similar to a library or module in other languages.
- **Standard Library**: Go provides over 100 packages for common tasks (e.g., `fmt` for formatted I/O, sorting, text manipulation).
- **File Structure**: Every source file follows this order:
  1.  **Package Declaration**: (e.g., `package main`)
  2.  **Import List**: Other packages needed for the code.
  3.  **Program Declarations**: Functions, variables, etc.


# Naming Conventions in Go

## 1. Module Naming

### Recommended Pattern

```
github.com/<org-or-username>/<repo-name>
```

### Rules

* Use a **domain-based path** for global uniqueness
* Prefer **kebab-case** (hyphen-separated) for repo/module names
* Keep it **short and descriptive**

### Examples

```
github.com/rahul/go-basics
github.com/acme/payment-service
go.opentelemetry.io/otel
```

### Avoid

```
github.com/rahul/myProject123   // ❌ mixed case
github.com/rahul/gobasicsstuff  // ❌ unclear
github.com/rahul/Go_Basics      // ❌ underscores + caps
```

---

## 2. Package Naming

### Core Rules

* **Lowercase only**
* **No underscores or hyphens**
* **Short, meaningful names**
* Usually matches the **directory name**

### Good Examples

```go
package user
package auth
package config
package logger
```

### Avoid

```go
package user_service   // ❌ underscore
package User           // ❌ uppercase
package utils          // ❌ too generic
package common         // ❌ vague
```

### Two-word package names

👉 Use **single lowercase word if possible**

If needed:

* Prefer combining words naturally

Examples:

```go
package userprofile   // acceptable
package httpserver    // acceptable
```

Better approach:

```
/user/profile   → package profile
/http/server    → package server
```

👉 Go prefers **structure over long names**

---

## 3. Directory Naming

### Rules

* Lowercase
* No spaces
* Avoid underscores
* Keep names short

### Examples

```
/user
/auth
/payment
/config
```

### For multiple words

👉 Prefer splitting into directories

Instead of:

```
/user_profile
```

Use:

```
/user/profile
```

---

## 4. File Naming

### Rules

* Lowercase
* Use underscores if needed (allowed here)

### Examples

```
user.go
user_service.go
http_handler.go
```

---

## 5. Function & Variable Naming

### Exported (public)

* Start with **Capital letter**

```go
func GetUser()
var AppConfig
```

### Unexported (private)

* Start with **lowercase**

```go
func getUser()
var appConfig
```

### Two-word names

👉 Use **camelCase / PascalCase**

```go
getUserName()     // private
GetUserName()     // public
userID            // common style (ID not Id)
```

---

## 6. Constants Naming

### Rules

* Use camelCase or PascalCase
* Avoid screaming snake case

```go
const maxRetries = 3
const DefaultTimeout = 30
```

Avoid:

```go
const MAX_RETRIES = 3  // ❌ not Go style
```

---

## 7. Interface Naming

### Convention

* Often end with `-er`

```go
type Reader interface {}
type Writer interface {}
type Stringer interface {}
```

---

## 8. Avoid Stutter in Names

Bad:

```go
package user

func user.GetUser() // ❌ stutter
```

Good:

```go
package user

func Get() // ✅ cleaner
```

---

## 9. Summary Rules

* Module → kebab-case (with domain)
* Package → lowercase, single word preferred
* Directory → lowercase, simple
* File → lowercase, underscores allowed
* Functions → camelCase / PascalCase
* Exported → Capitalized
* Prefer **clarity over cleverness**

---

## Golden Principle

> Use simple, predictable, and readable names. Let folder structure do the heavy lifting instead of long names.


--

### The Entry Point

`Read the program hello/hello.go`

- **`package main`**: This is a special package that defines a **standalone executable** program rather than a library.
- **`func main()`**: A special function within `package main` where program execution begins.
- **`fmt.Println`**: A basic function from the `fmt` package that prints values separated by spaces with a newline at the end.

## Imports and Program Structure

### The `import` Declaration

- **Requirement**: You must import **exactly** the packages you need.
- **Strictness**: A program will not compile if there are missing imports or if there are unnecessary (unused) ones. This prevents the accumulation of unused dependencies.
- **Placement**: The `import` declarations must follow the `package` declaration at the top of the file.

### Program Declarations

A Go program primarily consists of four types of declarations:

1. **`func`**: Functions
2. **`var`**: Variables
3. **`const`**: Constants
4. **`type`**: Types

> **Note**: For the most part, the order of these declarations within a file does not matter.

### Function Syntax

A function declaration consists of:

- The keyword `func`.
- The **name** of the function.
- A **parameter list** (enclosed in parentheses).
- A **result list** (the return types).
- The **body**: The statements defining what the function does, enclosed in braces `{ }`.

### Semicolons and Newlines

- **Semicolons**: Go does not require semicolons at the ends of statements or declarations, except when placing multiple statements on a single line.
- **Automatic Insertion**: The Go compiler automatically converts certain newlines into semicolons.
- **Brace Rule**: Because of how newlines are parsed, the opening brace `{` of a function **must** be on the same line as the `func` declaration, not on a line by itself.
- **Operators**: In expressions (like `x + y`), a newline is permitted **after** the `+` operator, but not before it.

### Code Formatting Tooling

- **`gofmt`**: A tool that rewrites Go source code into a standard format. It eliminates debates over coding style and enables automated code transformations.
- **`go fmt`**: The command used to apply `gofmt` to all files in a package or directory.
- **`goimports`**: A related tool (not part of the standard distribution) that automatically manages the insertion and removal of `import` declarations as you code. Note: add this in your project when you create one, this will add this tool in your go.mod file
```bash
$ go get golang.org/x/tools/cmd/goimports
```

## Command-Line Arguments

Most programs require external input to operate. One of the simplest ways to provide this is via command-line arguments.

### The `os.Args` Variable
* **Definition**: `os.Args` is a **slice of strings** provided by the `os` package that contains the command-line arguments used to run the program.
* **Indexing**:
    * `os.Args[0]` is the name of the command itself.
    * `os.Args[1:]` contains the actual arguments passed by the user.
* **Slices**: These are dynamically sized sequences. Go uses **half-open intervals** for slicing (`s[m:n]`), which include the first index but exclude the last (i.e., elements $m$ through $n-1$).



---

### The `for` Loop
Go only has one looping construct: the `for` loop. It has three primary forms:

1.  **The Traditional Form**: Includes initialization, condition, and post-iteration statement.
    ```go
    for i := 1; i < len(os.Args); i++ {
        // logic
    }
    ```
2.  **The "While" Form**: Only the condition is present.
    ```go
    for condition {
        // logic
    }
    ```
3.  **The Infinite Loop**: No condition; it runs until a `break` or `return` is reached.
    ```go
    for {
        // logic
    }
    ```



---

### Iterating with `range`
The `range` keyword is used to iterate over a slice or map. In each iteration, it produces a pair of values: the **index** and the **element value**.

* **The Blank Identifier (`_`)**: Go does not allow unused local variables. If you only need the element value and not the index, use an underscore to discard the index.
    ```go
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    ```

---

### Variable Declarations
There are several ways to declare and initialize variables in Go:

* **`s := ""`**: Short variable declaration (most compact; only available inside functions).
* **`var s string`**: Relies on default **zero-value** initialization (empty string `""`).
* **`var s = ""`**: Used rarely, typically when declaring multiple variables.
* **`var s string = ""`**: Explicitly defines the type; redundant if the initializer is the same type.

---

### String Concatenation & Efficiency
* **Operators**: The `+` operator concatenates strings, and `+=` is the assignment operator for it.
* **Performance**: Repeatedly using `s += string` is inefficient for large datasets because it creates a brand-new string every time (quadratic complexity).
* **`strings.Join`**: The most efficient way to combine many strings into one.
    ```go
    import "strings"
    fmt.Println(strings.Join(os.Args[1:], " "))
    ```

---

### Important Syntax Details
* **Increments**: `i++` and `i--` are **statements**, not expressions. Therefore, `j = i++` is illegal. `They are also postfix only (`++i` is not allowed).`
* **Braces**: Braces `{ }` are mandatory for loops, and the opening brace `{` must be on the same line as the `for` or `post` statement.
* **Comments**: Begin with `//` and continue to the end of the line.

<hr />

# Packages and Modules

## 1. What is a Go Module?

A **module** is a collection of Go packages with a unique identity.

Initialize a module using:

```bash
go mod init <module-name>
```

Example:

```bash
go mod init github.com/your-username/go-basics
```

### Key Idea:

* Module name is a **global identifier**
* Used in **import paths**
* Not the same as package name

---

## 2. Why Module Names Look Like URLs

Example:

```
github.com/user/project
example.com/hello
```

### Reasons:

* Ensures **global uniqueness**
* Allows Go to **download code from the internet**
* Maps directly to repository locations

### Symbols Explained:

* `/` → separates module and sub-packages
* `.` → part of domain naming (no special Go meaning)

---

## 3. Module vs Package (Important Difference)

| Concept | Example                 | Purpose           |
| ------- | ----------------------- | ----------------- |
| Module  | github.com/user/project | Global identity   |
| Package | main, auth, user        | Code organization |

---

## 4. What is a Package?

Defined inside `.go` files:

```go
package main
```

or

```go
package user
```

### Rules:

* Must be **lowercase**
* Usually matches **folder name**
* Should be **short and meaningful**

---

## 5. How Modules and Packages Work Together

### Example Structure:

```
go-basics/
  go.mod
  loops/
  arrays/
```

`go.mod`:

```
module github.com/you/go-basics
```

Importing a package:

```go
import "github.com/you/go-basics/loops"
```

---

## 6. Package Naming Conventions (Production)

### Best Practices:

* Use **lowercase only**
* Keep names **short**
* Avoid underscores
* Match folder name

### Good Examples:

```go
package user
package auth
package config
```

### Avoid:

```go
package user_service   // ❌
package utils          // ❌ too generic
package mypackage123   // ❌
```

---

## 7. Common Project Structure

```
project/
  go.mod
  cmd/
    server/
      main.go
  internal/
    user/
    auth/
  pkg/
    logger/
```

### Special Folders:

* `cmd/` → entry points
* `internal/` → private code
* `pkg/` → reusable code (optional)

---

## 8. Single Module vs Multiple Modules

### Recommended (Single Module):

```
repo/
  go.mod
  loops/
  arrays/
```

### Multiple Modules (Advanced):

```
repo/
  loops/go.mod
  arrays/go.mod
```

### Problems with Multiple Modules:

* Harder versioning
* Complicated dependencies
* More tooling complexity

---

## 9. How Go Resolves Imports

When you write:

```go
import "github.com/you/project/loops"
```

Go:

1. Reads module path from `go.mod`
2. Downloads module (if needed)
3. Locates package inside module

---

## 10. Naming Your Learning Project

### Recommended:

```bash
go mod init github.com/your-username/go-basics
```

### Alternatives:

```bash
go mod init go-basics
```

Avoid:

```bash
go mod init myproject123
```

---

## 11. Key Mental Models

* **Module = Project**
* **Package = Folder inside project**
* **Import path = Module + Folder**

---

## 12. Golden Rule

> Clarity > Cleverness

Keep names simple, predictable, and easy to read.

---

## 13. Note 

* Use **one module per repo** (especially while learning)
* Follow **real-world naming conventions**
* Think about how your code will be **imported and read by others**

---
## Names in Go

Names for functions, variables, constants, types, labels, and packages follow specific rules that determine their behavior and visibility.

### Naming Rules
* **Format**: A name must begin with a **letter** (Unicode) or an **underscore**. It can be followed by any number of letters, digits, and underscores.
* **Case Sensitivity**: Case matters. `heapSort` and `Heapsort` are distinct names.
* **Keywords**: There are **25 reserved keywords** that cannot be used as names:
  `break`, `default`, `func`, `interface`, `select`, `case`, `defer`, `go`, `map`, `struct`, `chan`, `else`, `goto`, `package`, `switch`, `const`, `fallthrough`, `if`, `range`, `type`, `continue`, `for`, `import`, `return`, `var`.

### Predeclared Names
Go has several dozen **predeclared names** for built-in constants, types, and functions. These are **not reserved** (you can redeclare them), but doing so is generally discouraged to avoid confusion.
* **Constants**: `true`, `false`, `iota`, `nil`.
* **Types**: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `float32`, `float64`, `bool`, `byte`, `rune`, `string`, `error`, etc.
* **Functions**: `make`, `len`, `cap`, `new`, `append`, `copy`, `close`, `delete`, `panic`, `recover`, etc.

---

### Scope and Visibility
The visibility of a name is determined by where it is declared and the **case of its first letter**:

* **Local**: Declared within a function; only visible inside that function.
* **Package-level**: Declared outside a function; visible in all files within that package.
* **Exported**: If a name begins with an **upper-case letter**, it is exported. This means it is accessible to other packages (e.g., `fmt.Printf`).
* **Unexported**: If a name begins with a **lower-case letter**, it is not visible outside its own package.
* **Package Names**: These are always written in **lower case**.

---

## Program Structure: Declarations, Variables, and Scope

### 2.2 Declarations
A declaration names a program entity and specifies its properties. There are four major kinds:
* **`var`**: Variables
* **`const`**: Constants
* **`type`**: Types
* **`func`**: Functions

**Structure of a `.go` file:**
1.  **Package declaration** (identifies the package).
2.  **Import declarations** (dependencies).
3.  **Package-level declarations** (visible across all files in the package).

---

### 2.3 Variables
The general form of a variable declaration is:  
`var name type = expression`

* **Type or Expression omission**: You can omit the `type` (inferred from the initializer) or the `expression` (initializes to the **Zero Value**).
* **Zero Value**: Go ensures every variable has a well-defined value. 
    * `0` for numbers.
    * `false` for booleans.
    * `""` for strings.
    * `nil` for interfaces, slices, pointers, maps, channels, and functions.
* **Multiple declarations**: `var i, j, k int` or `var b, f, s = true, 2.3, "four"`.

#### Short Variable Declarations (`:=`)
Used for the majority of local variables within functions.
* **Syntax**: `name := expression`
* **Shadowing**: A short declaration can act as an assignment to existing variables in the same block, provided at least one new variable is being declared.
* **Note**: `:=` is a declaration; `=` is an assignment.

---

### 2.3.2 Pointers
A **pointer** value is the **address** of a variable.
* **`&x`**: The "address-of" operator. Yields a pointer to the variable `x`.
* **`*p`**: The "dereference" operator. Yields the value of the variable the pointer `p` points to.
* **Addressability**: Every variable has an address; not every value (like a literal constant) does.
* **Safety**: It is safe for a function to return the address of a local variable. Go's compiler manages this via **Escape Analysis**.



---

### 2.3.3 The `new` Function
`new(T)` creates an **unnamed variable** of type `T`, initializes it to the zero value of `T`, and returns its address (`*T`).
* It is a syntactic convenience: `new(int)` is equivalent to declaring a dummy `int` variable and taking its address.

---

### 2.3.4 Lifetime of Variables
* **Package-level variables**: Exist for the entire execution of the program.
* **Local variables**: Have dynamic lifetimes. They are created when the declaration is executed and live until they become **unreachable** (garbage collected).
* **Escape Analysis**: If a local variable is still reachable after its function returns (e.g., its address was assigned to a global pointer), it is said to "escape" and is allocated on the **heap** instead of the **stack**.

---

### 2.4 Assignments
* **Tuple Assignment**: Allows assigning multiple values at once (e.g., `x, y = y, x` to swap).
* **Assignability**: A value can be assigned to a variable only if its type matches exactly (with some flexibility for constants and interfaces).

---

### 2.5 Type Declarations
A type declaration defines a new **named type** with the same underlying type as an existing one.  
`type name underlying-type`

* **Purpose**: To separate different/incompatible uses of the same underlying type (e.g., `Celsius` vs `Fahrenheit` both being `float64`).
* **Conversions**: `Celsius(f)` changes the type, not the value. It makes the change of meaning explicit.
* **Methods**: You can associate functions (methods) with named types to define specific behaviors.

---

### 2.6 Packages and Initialization
* **Exported Names**: Identifiers starting with an **upper-case letter** are visible outside their package.
* **`init` Functions**: Any file can contain `init()` functions. They are executed automatically when the program starts, after package-level variables are initialized.
* **Initialization Order**: Dependencies are resolved first. If package `A` imports `B`, `B` is fully initialized before `A`.

---

### 2.7 Scope
Scope is a **compile-time** property defining where a name is visible.
* **Lexical Blocks**: Scopes are defined by blocks (universe, package, file, function, etc.).
* **Shadowing**: An inner declaration can "hide" an outer one if they share the same name.
* **Implicit Blocks**: Constructs like `if`, `for`, and `switch` create implicit blocks for variables declared in their initialization headers.

> **Common Pitfall**: Using `:=` in an inner block (like an `init` function) can accidentally shadow a package-level variable instead of updating it.