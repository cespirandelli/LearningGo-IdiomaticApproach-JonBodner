# Learning Go

## Idiomatic Approach to Real-World Programming

## by Jon Bodner

Supplemental material (code examples, exercises, etc.) is available for download at [Learning Go · GitHub](https://github.com/learning-go-book)



# Chapter 1 : Setting Up Your Go Environment

## `go` Commands

`go run` vs. `go build`


### go run:

- Does not create a binary after running the command, but it compile the code, but the binary file is created in a temporary directory.
- After it runs, it's deleted after the program finishes.
- This is useful for testing out small programs, or to use Go as a scripting language.


### go build:

- It compiles Go code into a binary executable, usually for later use.

- By default, it creates an executable called hello.exe in the current directory (if the file was named hello.go).

- Creates an executable called "hello.exe" in the current directory.
  
  ```
  go build hello.go
  ```

- This creates `hello.exe`.


To change the output name, use the -o flag:

```
go build -o hello_world hello.go
```

- This creates `hello_world.exe`.

In summary, `go build` generates a distributable binary, which is often what you want for sharing with others.



## Getting Third-Party Go Tools

```
$ hey https://www.golang.org
$ bash: hey: command not found

$ export GOPATH=$HOME/go
$ export PATH=$PATH:$GOPATH/bin

$ go install github.com/rakyll/hey@latest
```

### **1. Verify Environment Variables**

Check if the environment variables are now properly set:

```
echo $GOPATH
echo $PATH
```

- `$GOPATH` should output `/c/Users/[User]/go`
- `$PATH` should include `/c/Users/[User]/go/bin`

### **2. Verify `hey` Installation**

Check if `hey` is installed and executable:

```
which hey
```

- This should return a path like `/c/Users/[User]/go/bin/hey`

### **3. Test `hey`**

Run a simple test with `hey` to confirm it’s working:

```
$ hey https://go.dev/

Summary:
  Total:        1.6971 secs
  Slowest:      0.8051 secs
  Fastest:      0.1960 secs
  Average:      0.3465 secs
  Requests/sec: 117.8484


Response time histogram:
  0.196 [1]     |■
  0.257 [53]    |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.318 [44]    |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.379 [26]    |■■■■■■■■■■■■■■■■■■■■
  0.440 [38]    |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.501 [25]    |■■■■■■■■■■■■■■■■■■■
  0.561 [7]     |■■■■■
  0.622 [3]     |■■
  0.683 [0]     |
  0.744 [1]     |■
  0.805 [2]     |■■


Latency distribution:
  10% in 0.2147 secs
  25% in 0.2535 secs
  50% in 0.3244 secs
  75% in 0.4134 secs
  90% in 0.4940 secs
  95% in 0.5164 secs
  99% in 0.8007 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0437 secs, 0.1960 secs, 0.8051 secs
  DNS-lookup:   0.0013 secs, 0.0000 secs, 0.0054 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0001 secs
  resp wait:    0.2648 secs, 0.1673 secs, 0.7028 secs
  resp read:    0.0360 secs, 0.0188 secs, 0.0414 secs

Status code distribution:
  [200] 200 responses
```


## Formatting Your Code

> "Most languages allow a great deal of flexibility
> on how code is laid out. Go does not. Enforcing a standard format makes it a
> great deal easier to write tools that manipulate source code."


> "Go programs use tabs to indent, and it is a syntax error if the opening
> brace is not on the same line as the declaration or command that begins the block."



`go fmt` vs.`goimports`

### **`go fmt`**

- Reformats your Go code to adhere to standard styling guidelines.
- Adjusts indentation, aligns struct fields, and ensures proper spacing around operators.
- Helps maintain consistency and readability across projects.

### **`goimports`**

- A more advanced tool that builds on `go fmt` by managing import statements.
- Organizes imports alphabetically, removes unused ones, and attempts to identify and include missing imports.
- The tool's guesses may not always be correct, so reviewing and manually adjusting imports is recommended when necessary.

You can download **`goimports`** with the following command:

```
go install golang.org/x/tools/cmd/goimports@latest
```

Run it using:

```
goimports -l -w .
```

> "The -l flag tells goimports to print the files with incorrect formatting to the console. The -w flag tells goimports to modify the files in-place. The . specifies the files to be scanned: everything in the current directory and all of its subdirectories."

- Note: You can find more details at the [goimports - Go Packages](https://pkg.go.dev/golang.org/x/tools@v0.28.0/cmd/goimports)



### The Semicolon Insertion Rule

`go fmt` won't fix braces on the wrong line, the reason for this is the semicolon (**;**) insertion rule.

Go requires a semicolon at the end of every statement. But Go devs never put the semicolons in themselves, the Go compiler does it for them. Following a very simple rule: 

- The rule is this. If the last token before a newline is an identifier (which includes words like `int` and `float64`), a basic literal such as a number or string constant, or one of the tokens
  
  ```
  break continue fallthrough return ++ -- ) }
  ```
  
  the lexer always inserts a semicolon after the token.

> "Summarized as, if the newline comes after a token that could end a statement, insert a semicolon." - [Effective Go](https://go.dev/doc/effective_go#semicolons)

- Note: Always try to remember using `go fmt` before compiling your code.



### Linting and Vetting

#### Linting:

> "Some of the changes it suggests include properly naming variables, formatting  error messages, and placing comments on public methods and types. These aren’t errors; they don’t keep your programs from compiling or make your program run incorrectly. 
> 
> Also, you cannot automatically assume that golint is 100% accurate: because the kinds of issues that `golint`finds are more fuzzy, it sometimes has false positives and false negatives. 
> 
> This means that you don’t have to make the changes that `golint`suggests. But you should take the suggestions from `golint` seriously."

- Install `golint` with the following command: 
  
  ```
  go install golang.org/x/lint/golint@latest
  ```

- Run it with:
  
  ```
  golint ./...
  ```
  
  That runs golint over your entire project.



#### Vetting:

> "There is another class of errors that developers run into. The code is syntactically valid, but there are mistakes that are not what you meant to do. 
> 
> This includes things like passing the wrong number of parameters to formatting methods or assigning values to variables that are never used. 
> 
> The go tool includes a command called `go vet` to detect these kinds of errors."

Run it with:

```
go vet ./...
```

| Aspect          | Linting                         | Vetting                                   |
| --------------- | ------------------------------- | ----------------------------------------- |
| **Focus**       | Style, syntax, and minor issues | Logical errors, deeper correctness checks |
| **Scope**       | Surface-level analysis          | Deeper static analysis                    |
| **Output**      | Warnings about code quality     | Warnings about correctness and safety     |
| **Tools in Go** | `golint` ./...                  | `go vet` ./...                            |



#### Additional content:

Start off using `go vet` as a required part of your automated build process and `golint` as part of your code review process (since golint might have false positives and false negatives, you can’t fix every issue it reports).

Once you're used to their recommendations, try out **`golangci-lint`** and
tweak its settings until it works for you. [GitHub - golangci-lint](https://github.com/golangci/golangci-lint)


- **Tip:** Integrate tools like `golint`, `go vet`, or `golangci-lint` into your development workflow to catch common bugs and enforce idiomatic Go code.

- Once your tools are properly configured, you can start working on your project. Simply open your favorite text editor and begin writing Go code.