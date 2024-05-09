package cloak

import (
	"database/sql/driver"
	"fmt"
)

func (s *Secret[T, M]) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	v, ok := value.(T)

	if !ok {
		return fmt.Errorf("Failed to Unmarshal database value, Invalid type (%T)", value)
	}

	*s = NewSecret(v, s.masker)

	return nil
}

func (s Secret[T, M]) Value() (driver.Value, error) {
	return s.data, nil
}
