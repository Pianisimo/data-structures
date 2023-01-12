package AVLTree

import (
	"data-structures/interfaces"
)

type AVLTree struct {
	Root *Node
}

func (t *AVLTree) String() string {
	return "implement me"
}

func (t *AVLTree) GetRoot() interfaces.Node {
	return t.Root
}

func (t *AVLTree) Remove(value int) {
	t.removeByNode(t.Root, value)
}

func (t *AVLTree) removeByNode(currentNode *Node, value int) *Node {
	if currentNode == nil {
		return currentNode
	}

	if value > currentNode.value {
		currentNode.right = t.removeByNode(currentNode.right, value)
	} else if value < currentNode.value {
		currentNode.left = t.removeByNode(currentNode.left, value)
	} else {
		if currentNode.left == nil || currentNode.right == nil {
			var temp *Node
			if temp == currentNode.left {
				temp = currentNode.right
			} else {
				temp = currentNode.left
			}

			if temp == nil {
				temp = currentNode
				currentNode = nil
			} else {
				currentNode = temp
			}
		} else {
			temp := currentNode.right.getMin()
			currentNode.value = temp.value
			currentNode.right = t.removeByNode(currentNode.right, temp.value)
		}
	}

	if currentNode == nil {
		return currentNode
	}

	currentNode.height = 1 + maxInt(t.getHeight(currentNode.left), t.getHeight(currentNode.right))
	b := t.getNodeBalance(currentNode)

	if b > 1 && t.getNodeBalance(currentNode.left) >= 0 {
		return t.rotateRight(currentNode)
	}
	if b > 1 && t.getNodeBalance(currentNode.left) < 0 {
		currentNode.left = t.rotateLeft(currentNode.left)
		return t.rotateRight(currentNode)
	}
	if b < -1 && t.getNodeBalance(currentNode.right) <= 0 {
		return t.rotateLeft(currentNode)
	}
	if b < -1 && t.getNodeBalance(currentNode.right) > 0 {
		currentNode.right = t.rotateRight(currentNode.right)
		return t.rotateLeft(currentNode)
	}

	return currentNode
}

func (t *AVLTree) Add(value int) {
	t.Root = t.addByNode(value, t.Root)
}

func (t *AVLTree) addByNode(value int, currentNode *Node) *Node {
	if currentNode == nil {
		return &Node{
			value:  value,
			height: 1,
		}
	}
	if value < currentNode.value {
		currentNode.left = t.addByNode(value, currentNode.left)
	} else {
		currentNode.right = t.addByNode(value, currentNode.right)
	}

	currentNode.height = 1 + maxInt(t.getHeight(currentNode.left), t.getHeight(currentNode.right))
	b := t.getNodeBalance(currentNode)

	if b > 1 && value < currentNode.left.value {
		return t.rotateRight(currentNode)
	}
	if b < -1 && value > currentNode.right.value {
		return t.rotateLeft(currentNode)
	}
	if b > 1 && value > currentNode.left.value {
		currentNode.left = t.rotateLeft(currentNode.left)
		return t.rotateRight(currentNode)
	}
	if b < -1 && value < currentNode.right.value {
		currentNode.right = t.rotateRight(currentNode.right)
		return t.rotateLeft(currentNode)
	}

	return currentNode
}

func (t *AVLTree) rotateRight(currentNode *Node) *Node {
	n := currentNode.left
	n2 := n.right
	n.right = currentNode
	currentNode.left = n2

	currentNode.height = 1 + maxInt(t.getHeight(currentNode.left), t.getHeight(currentNode.right))
	n.height = 1 + maxInt(t.getHeight(n.left), t.getHeight(n.right))
	return n
}

func (t *AVLTree) rotateLeft(currentNode *Node) *Node {
	n := currentNode.right
	n2 := n.left
	n.left = currentNode
	currentNode.right = n2

	currentNode.height = 1 + maxInt(t.getHeight(currentNode.left), t.getHeight(currentNode.right))
	n.height = 1 + maxInt(t.getHeight(n.left), t.getHeight(n.right))
	return n
}

func (t *AVLTree) getNodeBalance(currentNode *Node) int {
	if currentNode == nil {
		return 0
	}
	return t.getHeight(currentNode.left) - t.getHeight(currentNode.right)
}

func (t *AVLTree) getHeight(currentNode *Node) int {
	if currentNode == nil {
		return 0
	} else {
		return currentNode.height
	}
}

func maxInt(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
