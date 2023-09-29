package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	value := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = value
		l.Tail = value
	} else {
		value.Next = l.Head
		l.Head = value
	}
}
