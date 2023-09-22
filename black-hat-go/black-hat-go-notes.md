# Black Hat Go

### Go Tool Commands

Run a go main package:

`go run main.go`

Build a go binary:

`go build main.go`

Stripped:

`go build -ldflags "-w -s" main.go`

Cross-compiling (https://golang.org/doc/install/source#environment/):

`GOOS="linux" GOARCH="amd64" go build main.go`
