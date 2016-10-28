package shared

// Arithmetic interface
type Arithmetic interface {
	Multiply(args *Args, reply *int) error
	Divide(args *Args, quo *Quotient) error
}
