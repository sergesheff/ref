# Ref package

Ref package provides a simple way to assign a reference to a value of the primitive type without creating of temporary variable.

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
func MakeReferenceToInt() {
    var someValue *int
    MakeRef(&someValue, 10)
}
```

Package supports custom types and basic conversions:

```golang
func MakeReferenceToCustomType() {
    type someType string
    const someTypeAny = "any"
    
    var someValue *someType
    MakeRef(&someValue, someTypeAny)
}
```

```golang
func MakeReferenceToInt() {
    var someValue *int
    MakeRef(&someValue, 10.00)
}
```