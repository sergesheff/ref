package ref

import (
	"errors"
	"reflect"
	"testing"
)

var tests = []*Tests{
	&Tests{
		fn: func() (interface{}, error) {
			var v *bool
			err := MakeRef(v, true)
			return v, err
		},
		expectedValue: nil,
		expectedErr:   NotPointerError,
	}, // 1
	&Tests{
		fn: func() (interface{}, error) {
			var v *bool
			err := MakeRef(&v, true)
			return v, err
		},
		expectedValue: true,
		expectedErr:   nil,
	}, // 2
	&Tests{
		fn: func() (interface{}, error) {
			var v *int
			err := MakeRef(&v, 10)
			return v, err
		},
		expectedValue: 10,
		expectedErr:   nil,
	}, // 3
	&Tests{
		fn: func() (interface{}, error) {
			var v *int
			err := MakeRef(&v, 10.00)
			return v, err
		},
		expectedValue: 10,
		expectedErr:   nil,
	}, // 4
	&Tests{
		fn: func() (interface{}, error) {
			var v *float32
			err := MakeRef(&v, 10.00)
			return v, err
		},
		expectedValue: float32(10),
		expectedErr:   nil,
	}, // 5
	&Tests{
		fn: func() (interface{}, error) {
			var v *float32
			err := MakeRef(&v, 10)
			return v, err
		},
		expectedValue: float32(10),
		expectedErr:   nil,
	}, // 6
	&Tests{
		fn: func() (interface{}, error) {
			var v *float32
			err := MakeRef(&v, "10")
			return v, err
		},
		expectedValue: float32(0),
		expectedErr:   NotAssignedVariablesError,
	}, // 78
	&Tests{
		fn: func() (interface{}, error) {
			var v *string
			err := MakeRef(&v, 65)
			return v, err
		},
		expectedValue: "A",
		expectedErr:   nil,
	}, // 8
	&Tests{
		fn: func() (interface{}, error) {
			var v *string
			err := MakeRef(&v, 10.5)
			return v, err
		},
		expectedValue: "",
		expectedErr:   NotAssignedVariablesError,
	}, // 9
	&Tests{
		fn: func() (interface{}, error) {
			var v *string
			err := MakeRef(&v, "test")
			return v, err
		},
		expectedValue: "test",
		expectedErr:   nil,
	}, // 10
	&Tests{
		fn: func() (interface{}, error) {
			var v *CustomType
			err := MakeRef(&v, CustomTypeTest)
			return v, err
		},
		expectedValue: CustomTypeTest,
		expectedErr:   nil,
	}, // 11
}

func TestInt(t *testing.T) {

	for i, test := range tests {
		res, err := test.fn()

		t.Logf("test #%d \n", i+1)

		var tmpVal interface{} = nil
		v := reflect.ValueOf(res)
		if !v.IsZero() {
			tmpVal = v.Elem().Interface()
		}

		t.Logf("Expected error: %v; Received error: %v \n", test.expectedErr, err)
		t.Logf("Expected value: %v; Received value: %v \n", test.expectedValue, tmpVal)

		if !errors.Is(err, test.expectedErr) {

			if tmpVal != test.expectedValue {
				t.Error(`FAIL`)
				continue
			}

		}

		t.Log("OK")
	}
}

type Tests struct {
	fn            func() (val interface{}, err error)
	expectedValue interface{}
	expectedErr   error
}

type CustomType string

const CustomTypeTest CustomType = "test"
