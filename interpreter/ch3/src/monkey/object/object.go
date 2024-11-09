package object

import "fmt"

type ObjectType string

const (
	INTEG_OBJ   = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

// Object represents interpreted values
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer Object represents an Integer
type Integer struct {
	Value int64
}

func (i *Integer) Type() string    { return INTEG_OBJ }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Boolean Object represents a Boolean
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() string    { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%v", b.Value) }

// Null object
type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }
