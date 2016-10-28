package main

import (
	"errors"
	"net/rpc"
	"shared" //Path to the package contains shared struct
)

// ArithS implements Arithmetic interface on the server side
type ArithS int

// Multiply ...
func (t *ArithS) Multiply(args *shared.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Divide ...
func (t *ArithS) Divide(args *shared.Args, quo *shared.Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func registerArith(server *rpc.Server, arith shared.Arithmetic) {
	// registers Arith interface by name of `Arithmetic`.
	// If you want this name to be same as the type name, you
	// can use server.Register instead.
	server.RegisterName("Arithmetic", arith)
}

func maina() {

}
