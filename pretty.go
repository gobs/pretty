/*
 Pretty-print Go data structures
*/
package pretty

import (
	"bytes"
	"io"
	"os"
	r "reflect"
	"strconv"
	"strings"
)

const (
	DEFAULT_INDENT = "  "
	DEFAULT_NIL    = "nil"
)

// The context for printing
type Pretty struct {
	// indent string
	Indent string
	// output recipient
	Out io.Writer
	// string for nil
	NilString string
}

// pretty print the input value (to stdout)
func PrettyPrint(i interface{}) {
	PrettyPrintTo(os.Stdout, i)
}

// pretty print the input value (to a string)
func PrettyFormat(i interface{}) string {
	var out bytes.Buffer
	PrettyPrintTo(&out, i)
	return out.String()
}

// pretty print the input value (to specified writer)
func PrettyPrintTo(out io.Writer, i interface{}) {
	p := &Pretty{DEFAULT_INDENT, out, DEFAULT_NIL}
	p.Print(i)
}

// pretty print the input value
func (p *Pretty) Print(i interface{}) {
	p.PrintValue(r.ValueOf(i), 0)
	io.WriteString(p.Out, "\n")
}

// recursively print the input value
func (p *Pretty) PrintValue(val r.Value, level int) {
	if !val.IsValid() {
		io.WriteString(p.Out, p.NilString)
		return
	}

	cur := strings.Repeat(p.Indent, level)
	next := strings.Repeat(p.Indent, level+1)

	switch val.Kind() {
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		io.WriteString(p.Out, strconv.FormatInt(val.Int(), 10))

	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64:
		io.WriteString(p.Out, strconv.FormatUint(val.Uint(), 10))

	case r.Float32, r.Float64:
		io.WriteString(p.Out, strconv.FormatFloat(val.Float(), 'f', -1, 64))

	case r.String:
		io.WriteString(p.Out, strconv.Quote(val.String()))

	case r.Bool:
		io.WriteString(p.Out, strconv.FormatBool(val.Bool()))

	case r.Map:
		l := val.Len()

		io.WriteString(p.Out, "{\n")
		for i, k := range val.MapKeys() {
			io.WriteString(p.Out, next)
			io.WriteString(p.Out, strconv.Quote(k.String()))
			io.WriteString(p.Out, ": ")
			p.PrintValue(val.MapIndex(k), level+1)
			if i < l-1 {
				io.WriteString(p.Out, ",\n")
			} else {
				io.WriteString(p.Out, "\n")
			}
		}
		io.WriteString(p.Out, cur)
		io.WriteString(p.Out, "}")

	case r.Array, r.Slice:
		l := val.Len()

		io.WriteString(p.Out, "[\n")
		for i := 0; i < l; i++ {
			io.WriteString(p.Out, next)
			p.PrintValue(val.Index(i), level+1)
			if i < l-1 {
				io.WriteString(p.Out, ",\n")
			} else {
				io.WriteString(p.Out, "\n")
			}
		}
		io.WriteString(p.Out, cur)
		io.WriteString(p.Out, "]")

	case r.Interface, r.Ptr:
		p.PrintValue(val.Elem(), level)

	case r.Struct:
		l := val.NumField()

		io.WriteString(p.Out, "struct {\n")
		for i := 0; i < l; i++ {
			io.WriteString(p.Out, next)
			io.WriteString(p.Out, val.Type().Field(i).Name)
			io.WriteString(p.Out, ": ")
			p.PrintValue(val.Field(i), level+1)
			if i < l-1 {
				io.WriteString(p.Out, ",\n")
			} else {
				io.WriteString(p.Out, "\n")
			}
		}
		io.WriteString(p.Out, cur)
		io.WriteString(p.Out, "}")

	default:
		io.WriteString(p.Out, "unsupported:")
		io.WriteString(p.Out, val.String())
	}
}
