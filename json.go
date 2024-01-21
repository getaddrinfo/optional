package optional

import "encoding/json"

// Marshals an Optional:
// - If Some, then it marshals the underlying value.
// - If None, it marshals to null.
func (o Optional[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.val)
}

// Unmarshals an Optional:
// - If null, then it unmarshals to a None
// - If not null, then it unmarshals to a Some
func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &o.val); err != nil {
		return err
	}

	return nil
}
