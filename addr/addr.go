package addr

// Of returns a reference to the given argument.
// It can be useful to obtain a reference to a literal value easily.
// It also creates a copy of the value before referencing it, thus
// preventing mutations to the original value.
func Of[T any](t T) *T {
	return &t
}
