package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	size := 0
	current := l.Head

	for current != nil {
		size++
		current = current.Next
	}

	return size
}

// func main() {
// 	// Example usage
// 	node1 := &NodeL{Data: 10, Next: nil}
// 	node2 := &NodeL{Data: 20, Next: nil}
// 	node3 := &NodeL{Data: 30, Next: nil}

// 	list := &List{Head: node1, Tail: node3}
// 	node1.Next = node2
// 	node2.Next = node3

// 	size := ListSize(list)
// }
