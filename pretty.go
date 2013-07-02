package pretty

import (
	"fmt"
	r "reflect"
	"strings"
)

func PrettyPrint(i interface{}) {
	PrettyPrintValue(r.ValueOf(i), 0)
}

func PrettyPrintValue(val r.Value, indent int) {
	spaces := strings.Repeat("  ", indent)

	if !val.IsValid() {
		fmt.Printf("null")
		return
	}

	switch val.Kind() {
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		fmt.Printf("%v", val.Int())

	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64:
		fmt.Printf("%v", val.Uint())

	case r.Float32, r.Float64:
		fmt.Printf("%v", val.Float())

	case r.String:
		fmt.Printf("%q", val)

	case r.Bool:
		fmt.Printf("%v", val.Bool())

	case r.Map:
		l := val.Len()

		fmt.Printf("{\n")
		for i, k := range val.MapKeys() {
			fmt.Printf("%s%q: ", spaces, k)
			PrettyPrintValue(val.MapIndex(k), indent+1)
			if i < l {
				fmt.Printf(",\n")
			}
		}
		fmt.Printf("%s}", spaces)

	case r.Array, r.Slice:
		l := val.Len()

		fmt.Printf("[\n")
		for i := 0; i < l; i++ {
			fmt.Printf(spaces)
			PrettyPrintValue(val.Index(i), indent+1)
			if i < l {
				fmt.Printf(",\n")
			}
		}
		fmt.Printf("%s]", spaces)

	case r.Interface, r.Ptr:
		PrettyPrintValue(val.Elem(), indent)

	default:
		fmt.Printf("%v %v", val.Kind(), val)
	}
}
