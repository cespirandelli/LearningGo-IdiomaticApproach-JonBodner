# Learning Go

## Idiomatic Approach to Real-World Programming

## by Jon Bodner

Supplemental material (code examples, exercises, etc.) is available for download at [Learning Go · GitHub](https://github.com/learning-go-book)





`go run` vs. `go build`

### go run:

- Does not create a binary after running the command, but it compile the code, but the binary file is created in a temporary directory.
- After it runs, it's deleted after the program finishes.
- This is useful for testing out small programs, or to use Go as a scripting language.

### go build

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



## Go Tools Install Procedure

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