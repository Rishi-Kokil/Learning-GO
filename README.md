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

# Writing your first code
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

# Call code in an external package
- When you need your code to do something that might have been implemented by someone else, you can look for a package that has functions you can use in your code.
```link
https://pkg.go.dev/search?q=quote
```
- This is the repository for searching packages created by others, similar to `npm`
- if you look carefully you can see a lot of versions of the package listed over there. 
- In the top you can see the pacakge it is rsc.io/quote.
![alt text](image.png)
You can use the pkg.go.dev site to find published modules whose packages have functions you can use in your own code. Packages are published in modules -- like rsc.io/quote -- where others can use them. Modules are improved with new versions over time, and you can upgrade your code to use the improved versions.

# `go tidy`
- When we imported the rsc.io/quote in our code for the first time it could not resolve the module as it was un aware where it is coming from
- when we run the `go tidy` command it automatically searched the package repository and adds it to the go.mod file. Basically it adds it as dependency
- When you ran go mod tidy, it located and downloaded the rsc.io/quote module that contains the package you imported. By default, it downloaded the latest version -- v1.5.2.