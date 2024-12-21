# Learning Go
## Idiomatic Approach to Real-World Programming
## by Jon Bodner

`go run` vs. `go build`

### go run:
- Does not create a binary after running the command, but it compile the code, but the binary file is created in a temporary directory.
- After it runs, it's deleted after the program finishes.
- This is useful for testing out small programs, or to use Go as a scripting language.

### go build
- It compiles Go code into a binary executable, usually for later use.
- By default, it creates an executable called hello.exe in the current directory.

To change the output name, use the -o flag:
```bash
go build -o hello_world hello.go
```
This creates `hello_world.exe`.

In summary, `go build` generates a distributable binary, which is often what you want for sharing with others.