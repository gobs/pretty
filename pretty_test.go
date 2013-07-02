package pretty

import "testing"

type Bag map[string]interface{}

func TestPrettyPrint(test *testing.T) {

	bag := Bag{
		"a": 1,
		"b": false,
		"c": "some stuff",
		"d": 1.5,
		"e": Bag{
			"e1": "here",
			"e2": []int{1, 2, 3, 4},
			"e3": nil,
		},
	}

	PrettyPrint(bag)
}
