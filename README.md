pretty
======

A Pretty-printer for Go data structures

## Installation

    $ go get github.com/gobs/pretty

## Documentation
http://godoc.org/github.com/gobs/pretty

## Example

    import "github.com/gobs/pretty"

    stuff := map[string]interface{} {
      "a": 1,
      "b", "due",
      "c", []int { 1, 2, 3 },
      "d": false
    }
    
		pretty.PrettyPrint(stuff)
