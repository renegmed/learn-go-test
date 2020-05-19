package reflection

import (
	"reflect"
)

/*

If the value is a reflect.String then we just call fn like normal.

Otherwise, our switch will extract out two things depending on the type
	How many fields there are
	How to extract the Value (Field or Index)

Once we've determined those things we can iterate through numberOfValues
calling walk with the result of the getField function.

*/
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	// numberOfValues := 0
	// var getField func(int) reflect.Value

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {

	case reflect.Chan:
		//  iterate through all values sent through channel until it was closed with Recv()
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)
		}

	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}

	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.String:
		fn(val.String())

	}

	// for i := 0; i < numberOfValues; i++ {
	// 	walk(getField(i).Interface(), fn)
	// }

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		return val.Elem()
	}
	return val
}
