package options

// Option is a type that can be used to configure a value.
type Option[T any] interface {
	Apply(*T)
}

// OptionFunc is a function type that implements the Option interface.
type OptionFunc[T any] func(*T)

// Apply calls the OptionFunc.
func (f OptionFunc[T]) Apply(opt *T) {
	f(opt)
}
