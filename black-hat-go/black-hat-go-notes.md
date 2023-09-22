# Black Hat Go

### Go Tool Commands

**Run a go main package:**

`go run main.go`

**Build a go binary:**

`go build main.go`

**Build a stripped go binary:**

`go build -ldflags "-w -s" main.go`

**Cross-compiling (supported arch https://golang.org/doc/install/source#environment/):**

`GOOS="linux" GOARCH="amd64" go build main.go`

**The go doc command:**

`go doc library.function`

Displays documentation for a given function

**The go get command:**

Requires the execution of `go mod init` in the project dir as of version 1.16

`go get github.com/user/package-repo-name`

**The go fmt command:**

`go fmt main.go`

Formats your source code, most IDEs do this on save

**The golint and go vet commands:**

golint is a standalone tool, you’ll need to install it separately by using 
`go get -u golang.org/x/lint/golint.`

`go vet` inspects your code and uses heuristics to identify issues

**Go Playground:**

Execution environment - https://play.golang.org/

**Other commands and tools:**

- `go test` to test projects
- `go tool cover` to check test coverage
- `goimports` to fix imports

### Understanding Go Syntax
