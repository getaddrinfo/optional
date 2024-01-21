package optional_test

import (
	"testing"
	"unsafe"

	"github.com/getaddrinfo/optional"
)

func TestOptionalWithValue(t *testing.T) {
	const expected int = 42

	opt := optional.Some[int](expected)

	if opt.Value() != expected {
		t.Errorf("opt.Value(): expected %d, got: %d", expected, opt.Value())
	}

	if !opt.Has() {
		t.Error("opt.Has(): expected true, got false")
	}
}

func TestEmptyOptional(t *testing.T) {
	opt := optional.None[int]()

	if opt.Has() {
		t.Error("opt.Has(): expected false, got true")
	}
}

func TestCallingValueOnEmptyOptionalShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("opt.Value() called on empty optional, but did not panic")
		}
	}()

	opt := optional.None[int]()
	opt.Value()
}

func TestMemoryUsageSameAsPackedPointer(t *testing.T) {
	structSize := unsafe.Sizeof(optional.None[int]())
	ptrSize := unsafe.Sizeof((*int)(nil))

	if structSize != ptrSize {
		t.Fatalf("sizeof(optional.Optional[T]) != sizeof(*T): %d != %d", structSize, ptrSize)
	}
}
