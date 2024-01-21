// Package optional provides a generic Optional type to use within your code.
//
// Optional values can either be Some or None, which makes it harder to incorrectly use
// a null-pointer as you must explicitly call Value() on an Optional.
//
// The packed Optional struct has the same memory size as a pointer to the underlying type.
package optional

// An Optional represents a value that can be either 'Some' (has a value),
// or 'None' (no value).
//
// Under the hood this is just `*T`
type Optional[T any] struct {
	val *T
}

// Gets an optional's Value if it is Some
//
// Will panic if the optional is None.
func (o Optional[T]) Value() T {
	if o.val == nil {
		panic("Value() called on an optional with a None value")
	}

	return *o.val
}

// Gets the value in the Optional if it is 'Some',
// or the provided fallback if the Optional is 'None'.
func (o Optional[T]) Or(other T) T {
	if !o.Has() {
		return other
	}

	return o.Value()
}

// Gets the value in the Optional if it is 'Some', or
// calls the provided function if 'None'.
//
// Useful when you don't want to pre-compute the 'Or' case.
func (o Optional[T]) OrLazy(other func() T) T {
	if !o.Has() {
		return other()
	}

	return o.Value()
}

// Returns true if the optional is 'Some'
func (o Optional[T]) Has() bool {
	return o.val != nil
}

// Creates an optional that represents the 'Some' variant.
func Some[T any](val T) Optional[T] {
	return Optional[T]{&val}
}

// Creates an optional that represents the 'None' variant.
func None[T any]() Optional[T] {
	return Optional[T]{nil}
}
