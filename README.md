# Function Option Pattern in Golang

The function option pattern in Golang helps to create a struct constructor with optional configuration parameters. This pattern is useful when you have a struct with many fields and you want to provide a way to set only the fields that you need.

## What is the problem that this pattern solves?

When you have a struct with many fields and it can grow over time, this can lead to a constructor with many parameters. This can make the constructor hard to read and maintain.

See the example below:

```go
type Server struct {
  maxConn int
  id      string
  tls     bool
  name    string
  port    int
  ...
}

func newServer(maxConn int, id string, tls bool, name string, port int) *Server {
  return &Server{
    maxConn: maxConn,
    id:      id,
    tls:     tls,
    name:    name,
    port:    port,
    ...
    }
}

func main() {
  newServer(10, "server-id", false, "server-name", 8080)
}
```

## How does the function option pattern solve this problem?

The function option pattern solves this problem by providing a way to set only the fields that you need. This pattern uses a function that takes a pointer to the struct and sets a field of the struct.

See the example below and see the full code in the `main.go` file.:

```go
type Opts struct {
  maxConn int
  id      string
  tls     bool
  name    string
  port    int
}

type ServerOpt struct {
  Opts
}

func newServerOpt(opts ...OptFunc) *ServerOpt {
  o := defaultOpts()
  for _, fn := range opts {
    fn(&o)
  }

  return &ServerOpt{
    Opts: o,
    }
}

func main() {
  // now there is a default configuration and you can set only the fields that you need
  newServerOpt(withTls, withMaxConn(20))
}
```

## How to run the code

To run the code, you need to have Go installed on your machine. Then, you can run the following command:

```bash
go run main.go
```
