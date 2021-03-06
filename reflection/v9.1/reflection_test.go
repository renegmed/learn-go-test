package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	type Profile struct {
		Age  int
		City string
	}

	type Person struct {
		Name    string
		Profile Profile
	}

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{
				Name:    "Chris",
				Profile: Profile{Age: 33, City: "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				Name:    "Chris",
				Profile: Profile{Age: 33, City: "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{Age: 33, City: "London"},
				{Age: 34, City: "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		{
			"Arrays",
			[2]Profile{
				{Age: 33, City: "London"},
				{Age: 34, City: "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

}
