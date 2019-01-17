package value

type Object interface {
	// how to print in repl
	Repr() string
}

type Getter interface {
	Get(name string) Object
}

type Setter interface {
	Set(name string, value Object)
}
