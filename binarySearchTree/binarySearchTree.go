package binarySearchTree

import (
	"data-structures/interfaces"
	"fmt"
	"strings"
)

type BinarySearchTree struct {
	Root   *Node
	Length int
}

func (b *BinarySearchTree) GetRoot() interfaces.Node {
	return b.Root
}

func (b *BinarySearchTree) String() string {
	return b.inOrderTraversal()
}

// inOrderTraversal also known as BreadthDepthSearch, see inOrderTraversalByNode for it's 3 variants
func (b *BinarySearchTree) inOrderTraversal() string {
	sb := &strings.Builder{}
	b.inOrderTraversalByNode(sb, b.Root)
	return sb.String()
}

// inOrderTraversalByNode move the expression [sb.WriteString(fmt.Sprintf("%s ", currentNode))]
// to get a different variant of this method
func (b *BinarySearchTree) inOrderTraversalByNode(sb *strings.Builder, currentNode *Node) {
	if currentNode == nil {
		return
	}
	// PreOrder: sb.WriteString(fmt.Sprintf("%s ", currentNode))
	b.inOrderTraversalByNode(sb, currentNode.Left)

	// Ordered: sb.WriteString(fmt.Sprintf("%s ", currentNode))
	sb.WriteString(fmt.Sprintf("%s ", currentNode)) // moving this will change variant
	b.inOrderTraversalByNode(sb, currentNode.Right)

	// PostOrder: sb.WriteString(fmt.Sprintf("%s ", currentNode))
}

func (b *BinarySearchTree) Add(value int) {
	b.Root = b.addByNode(b.Root, value)
	b.Length++
}

func (b *BinarySearchTree) addByNode(currentNode *Node, value int) *Node {
	if currentNode == nil {
		return &Node{
			Value: value,
			Left:  nil,
			Right: nil,
		}
	}

	if value < currentNode.Value {
		currentNode.Left = b.addByNode(currentNode.Left, value)
	} else {
		currentNode.Right = b.addByNode(currentNode.Right, value)
	}

	return currentNode
}

func (b *BinarySearchTree) Search(value int) (*Node, bool) {
	return b.searchByNode(b.Root, value)
}

func (b *BinarySearchTree) searchByNode(root *Node, value int) (*Node, bool) {
	if root == nil {
		return nil, false
	}

	if value == root.Value {
		return root, true
	}

	if value < root.Value {
		return b.searchByNode(root.Left, value)
	} else {
		return b.searchByNode(root.Right, value)
	}
}

func (b *BinarySearchTree) Remove(value int) {
	b.removeByNode(b.Root, value)
}

func (b *BinarySearchTree) removeByNode(currentNode *Node, value int) *Node {
	if currentNode == nil {
		return currentNode
	}

	if value > currentNode.Value {
		currentNode.Right = b.removeByNode(currentNode.Right, value)
	} else if value < currentNode.Value {
		currentNode.Left = b.removeByNode(currentNode.Left, value)
	} else {
		if currentNode.Left == nil {
			return currentNode.Right
		} else {
			temp := currentNode.Left
			for temp.Right != nil {
				temp = temp.Right
			}

			currentNode.Value = temp.Value
			currentNode.Left = b.removeByNode(currentNode.Left, value)
		}
	}

	return currentNode
}

func (b *BinarySearchTree) BreadthFirstSearch() []int {
	currentNode := b.Root
	list := make([]int, 0)
	queue := make([]*Node, 0)

	queue = append(queue, currentNode)

	for len(queue) > 0 {
		currentNode = queue[0]
		queue = queue[1:]
		list = append(list, currentNode.Value)

		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
	}

	return list
}

func (b *BinarySearchTree) BreadthFirstSearchRecursive() []int {
	list := make([]int, 0)
	queue := []*Node{b.Root}
	list = b.breadthFirstSearchRecursive(queue, list)
	return list
}

func (b *BinarySearchTree) breadthFirstSearchRecursive(queue []*Node, list []int) []int {
	if len(queue) == 0 {
		return list
	}

	currentNode := queue[0]
	queue = queue[1:]
	list = append(list, currentNode.Value)

	if currentNode.Left != nil {
		queue = append(queue, currentNode.Left)
	}
	if currentNode.Right != nil {
		queue = append(queue, currentNode.Right)
	}

	return b.breadthFirstSearchRecursive(queue, list)
}
