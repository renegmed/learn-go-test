package reflection

import "reflect"

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

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {

	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field // method reflect.Value.Field(int) reflect.Value satisfies getField requirement

	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index // method reflect.Value.Index(int) reflect.Value satisfies getField requirement

	case reflect.String:
		fn(val.String())

	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		return val.Elem()
	}
	return val
}
