package html

import "io"

type NodeType int32

type Attribute struct {
	Key, Val string
}

type Node struct {
	Type NodeType
	Data string
	Attr []Attribute
	FirstChild, NextSibling *Node
}

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

func Parse(r io.Reader) (*Node, error)

