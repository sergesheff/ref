# Ref package

Ref package provides a simple way to assign a pointer to a value of the primitive type without creating of temporary variable.

The golang code for assigning of the pointer to a bool value is looking as:

```golang
func MakeReferenceToInt() {
/*
    someValue := &10 // doesn't work
*/
  var someValue *int
  tmpValue := 10 // creating a temporary variable
  someValue = &tmpValue // assigning pointer to it to a new variable
  _ = someValue
}
```

Using of Ref package:

```golang
import . "github.com/sergesheff/ref"

func MakeReferenceToInt() {
    var someValue *int = Ref(10)
    _ = someValue
}
```

Or

```golang
import . "github.com/sergesheff/ref"

func MakeReferenceToInt() {
    someValue := Ref(10) // someValue is using a *int type and pointing to 10
    _ = someValue
}
```

This function is very useful when need to assign many values to pointer fields of any small or huge struct (for example when struct is defining a DB table with nullable columns).

This is how it's looking right now:  

```golang
import . "github.com/sergesheff/ref"

type customType string

// defining the struct type
type customStruct struct {
    BoolValue   *bool
    IntValue    *int
    FloatValue  *float64
    StringValue *string
    CustomType  *customType
}

func createStruct() {
    // defining temporary variables
    var boolValue bool = true
    var intValue int = 1
    var floatValue float64 = 2
    var stringValue = "string"
    var cType = customType("customType")
    
    // assigning pointers to temporary variables to a struct instance
    structValue := customStruct{
        BoolValue:   &boolValue,
        IntValue:    &intValue,
        FloatValue:  &floatValue,
        StringValue: &stringValue,
        CustomType:  &cType,
    }
    
    _ = structValue
}
```

This is how it's looking with using of the Ref package:

```golang
import . "github.com/sergesheff/ref"

type customType string

// defining the struct type
type customStruct struct {
    BoolValue   *bool
    IntValue    *int
    FloatValue  *float64
    StringValue *string
    CustomType  *customType
}

func createStruct() {
    // assigning pointers to values directly without using of temporary variables
    structValue := customStruct{
        BoolValue:   Ref(true),
        IntValue:    Ref(1),
        FloatValue:  Ref(2.00),
        StringValue: Ref("string"),
        CustomType:  Ref(customType("customType")),
    }

    _ = structValue
}
```
