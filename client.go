package main

import (
	"log"
	"net/rpc"

	"shared" //Path to the package contains shared struct
)

// ArithC implements Arithmetic interface on the client side
type ArithC struct {
	client *rpc.Client
}

// Divide ...
func (t *ArithC) Divide(a, b int) shared.Quotient {
	args := &shared.Args{a, b}
	var reply shared.Quotient
	err := t.client.Call("Arithmetic.Divide", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return reply
}

// Multiply ...
func (t *ArithC) Multiply(a, b int) int {
	args := &shared.Args{a, b}
	var reply int
	err := t.client.Call("Arithmetic.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return reply
}
