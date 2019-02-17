package utils

type Node struct {
	userID        int
	d             float64
	relationships map[int]Node //0=parent, 1=left child, 2=right child
	h             int
}

func NewNode(userID int, distance float64, parent Node) *Node {

	n := Node{
		userID: userID,
		d:      distance,
	}

	n.relationships = make(map[int]Node, 3)
	n.relationships[0] = parent

	n.h = n.relationships[0].h + 1

	return &n
}

func (m *MinHeap) GetBottleIDs() []int {
	return m.root.traverse()
}

func (m *MinHeap) Traverse() []int {
	return m.root.traverse()
}

func (n *Node) traverse() []int {
	//assume start at root
	bottleIDs := make([]int, 1)
	bottleIDs = append(bottleIDs, n.userID)

	if n.Left().userID != 0 {
		bottleIDs = append(bottleIDs, (n.Left()).traverse()...)
	}
	if n.Right().userID != 0 {
		bottleIDs = append(bottleIDs, (n.Right()).traverse()...)
	}

	return bottleIDs
}

func (n *Node) Swap() {
	for {
		node := *n
		node2 := n.relationships[0]

		if node.d < node.relationships[0].d {
			node.userID, node2.userID = node2.userID, node.userID
			node.d, node2.d = node2.d, node.d

		} else {
			break
		}
	}
}

func (n *Node) GetRoot() *Node {
	var m Node

	for !m.IsRoot() {
		m = m.relationships[0]
	}

	return &m
}

func (n Node) IsRoot() bool {
	h := n.relationships[0].h

	return h == 0
}

//Minheap is used for tracking the 15 closest bottles

type MinHeap struct {
	root      Node
	deepest   Node
	len       int
	cap       int
	maxH      int
	Initiated int
}

func NewHeap(size int) *MinHeap {
	m := MinHeap{
		root: Node{
			h: 0,
		},
		len:       0,
		cap:       size,
		maxH:      4,
		Initiated: 1,
	}

	return &m
}

/*0. If heap is empty place element at root.
Compare the added element with its parent; if they are in the correct order, stop.
If not, swap the element with its parent and return to the previous step.
*/
func (m *MinHeap) Insert(n *Node) {

	if m.root.h == 0 {
		m.len++
		m.root = *n
		return
	}

	y := m.FindYoungestParent()
	y.AddChild(n)
	n.Swap()
	m.root.Trim(m.maxH)

}

func (n *Node) KillChild(userID int) {
	if n.Left().userID == userID {
		n.relationships[1] = Node{}
	}
	if n.Right().userID == userID {
		n.relationships[2] = Node{}
	}
}

func (n *Node) Trim(allowedHeight int) {
	if n.h > allowedHeight {
		n.Parent().KillChild(n.userID)
	} else if n.Left().userID != 0 {
		n.Left().Trim(allowedHeight)
	} else if n.Right().userID != 0 {
		n.Right().Trim(allowedHeight)
	}
}

func (n *Node) AddChild(o *Node) {
	if n.Right().userID != 0 {
		return
	}

	if n.Left().userID != 0 {
		n.AddLeft(o)
	}

	if n.Right().userID == 0 {
		n.AddRight(o)
	}
}

func (m *MinHeap) FindYoungestParent() *Node {
	if m.cap == m.len {
		return &m.deepest
	}

	n := m.root

	for {
		if n.Right().userID != 0 {
			n = *n.Right()
			continue
		}
		if n.Left().userID != 0 {
			n = *n.Left()
			continue
		}
		return &n
	}
}

func (n *Node) Parent() *Node {
	k := n.relationships[0]
	return &k
}

func (n *Node) Left() *Node {
	k := n.relationships[1]
	return &k
}

func (n *Node) Right() *Node {
	k := n.relationships[2]
	return &k
}

func (n *Node) AddRight(o *Node) {
	n.relationships[2] = *o

}

func (n *Node) AddLeft(o *Node) {
	n.relationships[1] = *o
}
