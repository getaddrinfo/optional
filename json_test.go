package optional_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/getaddrinfo/optional"
)

type JsonTestRepresentation struct {
	Field optional.Optional[int]
}

var jsonTestSome = []byte("{\"Field\":1}")
var jsonTestNone = []byte("{\"Field\":null}")
var jsonTestFieldMissing = []byte("{}")

func TestJsonUnmarshalWithNonNullValueCorrectlyUnmarshalledAsSome(t *testing.T) {
	var out JsonTestRepresentation

	if err := json.Unmarshal(jsonTestSome, &out); err != nil {
		t.Errorf("opt.UnmarshalJSON: error unmarshaling: %v", err.Error())
	}

	if !out.Field.Has() {
		t.Errorf("opt.Has(): unmarshalled from JSON with a non-null value (Some), but got None")
	}

	if out.Field.Value() != 1 {
		t.Errorf("opt.Value(): unmarshalled from JSON but got wrong value (expected: 1, got: %d)", out.Field.Value())
	}
}

func TestJsonMarshalWithSomeValueCorrectlyMarshalled(t *testing.T) {
	in := JsonTestRepresentation{
		Field: optional.Some[int](1),
	}

	data, err := json.Marshal(in)

	if err != nil {
		t.Errorf("json.Marshal: failed to marshal JsonTestRepresentation: %v", err.Error())
	}

	if !bytes.Equal(data, jsonTestSome) {
		t.Errorf("opt.MarshalJSON(): optional.Some incorrectly encoded, expected: %v, got: %v", jsonTestSome, data)
	}
}

func TestJsonUnmarshalWithNullValueCorrectlyUnmarshalledAsNone(t *testing.T) {
	var out JsonTestRepresentation

	if err := json.Unmarshal(jsonTestNone, &out); err != nil {
		t.Errorf("opt.UnmarshalJSON: error unmarshaling: %v", err.Error())
	}

	if out.Field.Has() {
		t.Errorf("opt.Has(): unmarshalled from JSON with a null value (None), but got Some")
	}
}

func TestJsonMarshalWithNoneValueCorrectlyMarshalled(t *testing.T) {
	in := JsonTestRepresentation{
		Field: optional.None[int](),
	}

	data, err := json.Marshal(in)

	if err != nil {
		t.Errorf("json.Marshal: failed to marshal JsonTestRepresentation: %v", err.Error())
	}

	if !bytes.Equal(data, jsonTestNone) {
		t.Errorf("opt.MarshalJSON(): optional.Some incorrectly encoded, expected: %v, got: %v", jsonTestSome, data)
	}
}

func TestJsonMarshalWithUnspecifiedFieldCorrectlyUnmarshalledAsNone(t *testing.T) {
	var out JsonTestRepresentation

	if err := json.Unmarshal(jsonTestFieldMissing, &out); err != nil {
		t.Errorf("json.Unmarshal: failed to unmarshal JsonTestRepresentation: %v", err.Error())
	}

	if out.Field.Has() {
		t.Error("opt.Has(): missing field incorrectly unmarshalled as Some")
	}
}
