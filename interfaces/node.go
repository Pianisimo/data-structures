package interfaces

type Node interface {
	String() string
	GetNode() Node
	GetLeft() Node
	GetRight() Node
	GetValue() int
}
