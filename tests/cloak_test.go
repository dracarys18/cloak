package cloak_test

import (
	"encoding/json"
	"fmt"
	"github.com/dracarys18/cloak"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMasking(t *testing.T) {
	assert := assert.New(t)
	val := "Wow this is a secret"
	s := cloak.NewSecret(val, cloak.StringMasking{})

	type_name := fmt.Sprintf("%s", reflect.TypeOf(s.Sneak()))

	secret_type := fmt.Sprintf("***%s(%d)***", type_name, len(val))

	test := fmt.Sprintf("%s", s)

	assert.Equal(test, secret_type, "Expected %s got %s", secret_type, test)
}

func TestSerializeDeserialize(t *testing.T) {
	assert := assert.New(t)
	type Test struct {
		W cloak.Secret[string, cloak.StringMasking] `json:"w"`
		S cloak.Secret[string, cloak.StringMasking] `json:"s"`
	}

	v := Test{
		W: cloak.NewSecret("Such secret much wow", cloak.StringMasking{}),
		S: cloak.NewSecret("Another secret much wow", cloak.StringMasking{}),
	}

	actual, ser_err := json.Marshal(v)

	if ser_err != nil {
		t.Errorf("Unable to deserialize the secret value")
	}

	expected := "{\"w\":\"Such secret much wow\",\"s\":\"Another secret much wow\"}"
	assert.Equal(string(actual), expected, "Expected %s got %s", expected, actual)

	var deser Test
	err := json.Unmarshal(actual, &deser)

	if err != nil {
		t.Errorf("Unable to deserialize the secret value")
	}

	assert.Equal(deser, v, "Expected %s got %s", deser, v)
}
