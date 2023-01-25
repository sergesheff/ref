package ref

import (
	"errors"
	"reflect"
)

var (
	NotPointerError           = errors.New("variable should be a pointer to pointer (**variable)")
	NotAssignedVariablesError = errors.New("variable and value are not assignable to each other")
)

// MakeRef is saving the pointer to the value to the variable.
//
// The Variable parameter should be a pointer to the pointer of the variable (**variable)
//
// The Value parameter should be a primitive value (true, 1, 10.00, "string" etc)
func MakeRef(variable interface{}, value interface{}) error {

	_variable := reflect.ValueOf(variable)
	_value := reflect.ValueOf(value)

	// _variable should be a reference
	if _variable.Kind() != reflect.Ptr || _variable.Elem().Kind() != reflect.Ptr {
		return NotPointerError
	}

	// getting pointer to an initial value
	_variable = _variable.Elem()

	// checking if _variable and values can be assigned to each other
	canConvert, err := compareTypes(_variable.Type().Elem(), _value.Type())
	if err != nil {
		return err
	}

	// if conversion is needed, then convert the _value
	if canConvert {
		_value = _value.Convert(_variable.Type().Elem())
	}

	// making a pointer to a new _value
	tmpValue := reflect.New(_value.Type())
	tmpValue.Elem().Set(_value)

	// updating pointer of the _variable
	_variable.Set(tmpValue)

	return nil
}

func compareTypes(v reflect.Type, val reflect.Type) (bool, error) {
	// checking if types are assignable to each other
	if !val.AssignableTo(v) {

		// if no, then checking if value can be converted
		if val.ConvertibleTo(v) {
			return true, nil
		}

		return false, NotAssignedVariablesError
	}

	return false, nil
}
