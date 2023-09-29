package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

func ListAt(head *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}

	current := head
	for i := 0; current != nil && i < pos; i++ {
		current = current.Next
	}

	return current
}
