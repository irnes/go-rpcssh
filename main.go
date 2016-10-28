package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// Mode ...
type Mode int

// Process running mode options
const (
	NONE Mode = iota
	CLIENT
	SERVER
)

func (m *Mode) String() string {
	if *m == CLIENT {
		return "client"
	} else if *m == SERVER {
		return "server"
	}
	return "unknown"
}

// Set value of flag
func (m *Mode) Set(value string) error {
	if value == "client" {
		*m = CLIENT
	} else if value == "server" {
		*m = SERVER
	}
	return nil
}

// CmdFlags defines a structure holding command line arguments
type CmdFlags struct {
	RunMode Mode
}

var (
	cmdFlags CmdFlags
)

func init() {
	flag.Var(&cmdFlags.RunMode, "mode", "Running mode server/client")
	flag.Parse()
}

func main() {
	fmt.Printf("Hello %s\n", cmdFlags.RunMode.String())

	if cmdFlags.RunMode == CLIENT {
		// Tries to connect to localhost:1234 (The port on which rpc server is listening)
		conn, err := net.Dial("tcp", "localhost:1234")
		if err != nil {
			log.Fatal("Connectiong:", err)
		}
		defer conn.Close()

		// Create a struct, that mimics all methods provided by interface.
		// It is not compulsory, we are doing it here, just to simulate a traditional method call.
		arith := &ArithC{client: rpc.NewClient(conn)}

		fmt.Println(arith.Multiply(5, 6))
		fmt.Println(arith.Divide(500, 11))
	} else if cmdFlags.RunMode == SERVER {
		//Creating an instance of struct which implement Arithmetic interface
		arith := new(ArithS)

		// Register a new rpc server (In most cases, you will use default server only)
		// And register struct we created above by name "Arith"
		// The wrapper method here ensures that only structs which implement Arith interface
		// are allowed to register themselves.
		server := rpc.NewServer()
		registerArith(server, arith)

		// Listen for incoming tcp packets on specified port.
		l, e := net.Listen("tcp", ":1234")
		if e != nil {
			log.Fatal("listen error:", e)
		}

		// This statement links rpc server to the socket, and allows rpc server to accept
		// rpc request coming from that socket.
		server.Accept(l)
	}

}
