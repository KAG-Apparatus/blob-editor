package cfg

import (
	"bytes"
	"fmt"
	"reflect"
)

//go:generate stringer -type=CFGType
type CFGType int

const (
	INT_TYPE CFGType = iota
	INT_ARRAY_TYPE

	STRING_TYPE
	STRING_ARRAY_TYPE

	FLOAT_TYPE
	FLOAT_ARRAY_TYPE

	BOOL_TYPE
	BOOL_ARRAY_TYPE
)

type CFGFile struct {
	Comments    []Comment
	Assignments []Assignment
}

func (p *CFGFile) String() string {
	var out bytes.Buffer
	for _, s := range p.Assignments {
		out.WriteString(s.String())
	}
	return out.String()
}

type Assignment struct {
	Line       int
	Type       CFGType
	Identifier string
	Value      interface{}
}

func (a *Assignment) String() string {
	return fmt.Sprintf("%s %s = %v", a.Type, a.Identifier, a.Value)
}

func (a *Assignment) Equal(b *Assignment) bool {
	if a.Line != b.Line || a.Type != b.Type || a.Identifier != b.Identifier {
		return false
	}

	if !reflect.DeepEqual(a.Value, b.Value) {
		return false
	}

	return true
}

type Comment struct {
	Line  int
	Value string
}

func (c *Comment) String() string {
	return c.Value
}
