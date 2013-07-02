package pretty

import "testing"

type Bag map[string]interface{}

func TestPrettyPrint(test *testing.T) {

	var ch chan string

	bag := Bag{
		"a": 1,
		"b": false,
		"c": "some stuff",
		"d": []float64{0.0, 0.1, 1.2, 1.23, 1.23456, 999999999999},
		"e": Bag{
			"e1": "here",
			"e2": []int{1, 2, 3, 4},
			"e3": nil,
		},
		"bad": ch,
	}

	arry := []Bag{bag, bag, bag}

	PrettyPrint(arry)
}
