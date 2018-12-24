package tree

import (
	"errors"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

func Build(records []Record) (*Node, error) {
	treeNodes := map[int]*Node{}
	tempNodes := map[int]*Node{}
	var root *Node
	maxLength := len(records)
	if maxLength == 0 {
		return root, nil
	}
	for _, record := range records {
		if treeNodes[record.ID] != nil {
			return nil, errors.New("duplicate node")
		}
		if record.ID >= maxLength {
			return nil, errors.New("over maxlength")
		}
		if record.ID == record.Parent {
			if root != nil {
				return nil, errors.New("multiple roots")
			}
			root = getNode(treeNodes, tempNodes, record.Parent)
		} else {
			parent := getNode(treeNodes, tempNodes, record.Parent)
			child := getNode(treeNodes, tempNodes, record.ID)
			if parent.ID > child.ID {
				return nil, errors.New("invalid order")
			}
			parent.addChild(child)
			treeNodes[record.ID] = child
		}
	}
	if root == nil {
		return nil, errors.New("no root")
	}
	return root, nil
}

func (parent *Node) addChild(node *Node) {
	if len(parent.Children) == 0 {
		parent.Children = append(parent.Children, node)
	} else {
		index := 0
		for _, child := range parent.Children {
			if child.ID > node.ID {
				break
			}
			index++
		}
		parent.Children = append(parent.Children, nil)
		copy(parent.Children[index+1:], parent.Children[index:])
		parent.Children[index] = node
	}

}

func getNode(treeNodes map[int]*Node, dummyNodes map[int]*Node, id int) *Node {
	if treeNodes[id] != nil {
		return treeNodes[id]
	}
	if dummyNodes[id] == nil {
		dummyNodes[id] = &Node{ID: id}
	}
	return dummyNodes[id]
}
