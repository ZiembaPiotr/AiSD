package Leetcode

type bst struct {
	root *children
	len  int
}

type children struct {
	value      int
	leftChild  *children
	rightChild *children
}

func (b *bst) addNewElement(value int, root *children) {
	if value > root.value {
		if root.rightChild == nil {
			root.rightChild = &children{
				value:      value,
				leftChild:  nil,
				rightChild: nil,
			}
			b.len++
		} else {
			b.addNewElement(value, root.rightChild)
		}
	} else {
		if root.leftChild == nil {
			root.leftChild = &children{
				value:      value,
				leftChild:  nil,
				rightChild: nil,
			}
			b.len++
		} else {
			b.addNewElement(value, root.leftChild)
		}
	}
}

func (b bst) searchForChildren(value int, root *children) bool {
	if root != nil {
		if root.value == value {
			return true
		}
		if value > root.value {
			return b.searchForChildren(value, root.rightChild)
		} else {
			return b.searchForChildren(value, root.leftChild)
		}
	}
	return false
}
