package cloak

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Secret[T any, M MaskObject] struct {
	data   T
	masker M
}

type MaskObject interface {
	Mask(any) []byte
}

func NewSecret[T any, M MaskObject](data T, masker M) Secret[T, M] {
	return Secret[T, M]{data, masker}
}

func (t *Secret[T, M]) UnmarshalJSON(data []byte) error {
	var d T

	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}

	*t = Secret[T, M]{data: d}
	return nil
}

func (t Secret[T, M]) MarshalJSON() ([]byte, error) {
	if val, err := json.Marshal(t.Sneak()); err != nil {
		return nil, err
	} else {
		return val, nil
	}
}

func RegisterCustomType[T any, M MaskObject](field reflect.Value) interface{} {
	if val, ok := field.Interface().(Secret[T, M]); ok {
		return val.Sneak()
	}
	return nil
}

func (s Secret[T, M]) Format(f fmt.State, c rune) {
	masked := s.masker.Mask(s.data)
	f.Write(masked)
}

func (s Secret[T, M]) Sneak() T {
	return s.data
}
