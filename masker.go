package cloak

import "fmt"

type StringMasking struct{}

func (_ StringMasking) Mask(a any) []byte {
	v, ok := a.(string)

	if !ok {
		return []byte{}
	}
	return []byte(fmt.Sprintf("***string(%d)***", len(v)))
}

type PasswordMasking struct{}

func (_ PasswordMasking) Mask(a any) []byte {
	return []byte(fmt.Sprintf("***HashedSecretPassword**"))
}
