package ref

import (
	"reflect"
	"testing"
)

var tests = []*Tests{
	&Tests{
		fn: func() interface{} {
			var v *bool
			v = Ref(true)
			return v
		},
		expectedValue: true,
	}, // 1
	&Tests{
		fn: func() interface{} {
			var v bool
			v = *Ref(true)
			return v
		},
		expectedValue: true,
	}, // 2
	&Tests{
		fn: func() interface{} {
			var v *int
			v = Ref(10)
			return v
		},
		expectedValue: 10,
	}, // 3
	&Tests{
		fn: func() interface{} {
			var v int
			v = *Ref(10)
			return v
		},
		expectedValue: 10,
	}, // 4
	&Tests{
		fn: func() interface{} {
			var v *float64
			v = Ref(10.00)
			return v
		},
		expectedValue: float64(10),
	}, // 5
	&Tests{
		fn: func() interface{} {
			var v float64
			v = *Ref(10.00)
			return v
		},
		expectedValue: float64(10),
	}, // 6
	&Tests{
		fn: func() interface{} {
			var v *string
			v = Ref("A")
			return v
		},
		expectedValue: "A",
	}, // 7
	&Tests{
		fn: func() interface{} {
			var v string
			v = *Ref("A")
			return v
		},
		expectedValue: "A",
	}, // 8
	&Tests{
		fn: func() interface{} {
			var v *CustomType
			v = Ref(CustomTypeTest)
			return v
		},
		expectedValue: CustomTypeTest,
	}, // 9
	&Tests{
		fn: func() interface{} {
			var v CustomType
			v = *Ref(CustomTypeTest)
			return v
		},
		expectedValue: CustomTypeTest,
	}, //10
}

func TestInt(t *testing.T) {

	for i, test := range tests {
		res := test.fn()

		t.Logf("test #%d \n", i+1)

		var tmpVal interface{} = nil
		v := reflect.ValueOf(res)
		if !v.IsZero() {

			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}

			tmpVal = v.Interface()
		}

		t.Logf("Expected value: %v; Received value: %v \n", test.expectedValue, tmpVal)

		if tmpVal != test.expectedValue {
			t.Error(`FAIL`)
			continue
		}

		t.Log("OK")
	}
}

type Tests struct {
	fn            func() (val interface{})
	expectedValue interface{}
}

type CustomType string

const CustomTypeTest CustomType = "test"
