package pretty

import "testing"

type Bag map[string]interface{}

type Struct struct {
	N int
	S string
	B bool
}

var (
	ch chan string

	bag = Bag{
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

	arry = []Bag{bag, bag, bag}

	strutty = Struct{N: 42, S: "Hello", B: true}
)

func TestPrettyPrint(test *testing.T) {
	PrettyPrint(arry)
}

func TestPrettyFormat(test *testing.T) {
	test.Log(PrettyFormat(bag))
}

func TestStruct(test *testing.T) {
	test.Log(PrettyFormat(strutty))
}

func TestTabPrint(test *testing.T) {
	tp := NewTabPrinter(8)

	for i := 0; i < 33; i++ {
		tp.Print(i)
	}

	tp.Println()

	for _, v := range []string{"one", "two", "three", "four", "five", "six"} {
		tp.Print(v)
	}

	tp.Println()
}
