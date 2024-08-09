package main

import "fmt"

// This projects is an example of how to create optional configurations using functions

// The Server represents the standard pattern of options, without using the function options pattern

type Server struct {
	maxConn int
	id      string
	tls     bool
}

func newServer(maxConn int, id string, tls bool) *Server {
	return &Server{
		maxConn: maxConn,
		id:      id,
		tls:     tls,
	}
}

// The ServerOpt represents the function options pattern

type OptFunc func(*Opts)

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "server1",
		tls:     false,
	}
}

func withTls(opts *Opts) {
	opts.tls = true
}

func withMaxConn(n int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

func withId(id string) OptFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}

type Opts struct {
	maxConn int
	id      string
	tls     bool
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
	//s := newServer(10, "server1", false)

	sOpt := newServerOpt(withTls, withMaxConn(20)) // now the configuration is optional

	fmt.Printf("%+v\n", sOpt)
}
