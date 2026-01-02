package trees

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{
		Data:  value,
		Left:  nil,
		Right: nil,
	}
}

func (bt *Node) Insert(value int) {
	if bt != nil {
		if bt.Data < value {
			//move right
			if bt.Right == nil {
				bt.Right = &Node{Data: value}
			} else {
				bt.Right.Insert(value)
			}
		} else if bt.Data > value {
			//move left
			if bt.Left == nil {
				bt.Left = &Node{Data: value}
			} else {
				bt.Left.Insert(value)
			}
		}
	}
	return
}

// NonRecursiveInsert has no callstack so it has better space complexity of O(1)
func (bt *Node) NonRecursiveInsert(value int) {
	current := bt

	for current != nil {
		if current.Data > value {
			if current.Left == nil {
				current.Left = &Node{Data: value}
				return
			}
			current = current.Left
		} else if current.Data < value {
			if current.Right == nil {
				current.Right = &Node{Data: value}
				return
			}
			current = current.Right
		}
	}
}

func (bt *Node) isLeaf() bool {
	if bt.Left == nil && bt.Right == nil {
		return true
	}
	return false
}

func (bt *Node) PrintSelf() {
	if bt == nil {
		fmt.Printf("(empty tree)\n")
		return
	}
	bt.printSelfHelper("", "", true)
}

func (bt *Node) printSelfHelper(prefix string, childPrefix string, isRoot bool) {
	if bt == nil {
		return
	}

	// Print current node
	if isRoot {
		fmt.Printf("%d\n", bt.Data)
	} else {
		fmt.Printf("%s%d\n", prefix, bt.Data)
	}

	// Print children
	if bt.Left != nil {
		if bt.Right != nil {
			// Has both children
			bt.Left.printSelfHelper(childPrefix+"├── ", childPrefix+"│   ", false)
		} else {
			// Only left child
			bt.Left.printSelfHelper(childPrefix+"└── ", childPrefix+"    ", false)
		}
	}

	if bt.Right != nil {
		// Right child is always last
		bt.Right.printSelfHelper(childPrefix+"└── ", childPrefix+"    ", false)
	}
}
