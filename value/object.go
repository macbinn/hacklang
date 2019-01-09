package value

type Object interface {
	// getter method
	Get(name string) Object

	// how to print in repl
	Repr() string
}
