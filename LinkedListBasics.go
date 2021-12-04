package Leetcode

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next *node
}

type linkedList struct {
	head *node
	len int
}

func (l *linkedList) add(value int) {
	newNode := new(node)
	newNode.value = value

	if l.head == nil {
		l.head = newNode
		l.len++
	} else {
		iterator := l.head
		for ; iterator.next != nil; iterator = iterator.next {
		}
		iterator.next = newNode
		l.len++
	}
}

func (l *linkedList) remove(value int) {
	nodeToRemove := l.get(value)

	if nodeToRemove == l.head {
		l.head = l.head.next
	} else {
		for iterator := l.head; iterator.next != nil; iterator = iterator.next{
			if iterator.next == nodeToRemove {
				iterator.next = iterator.next.next
				l.len--
				break
			}
		}
	}
}

func (l linkedList) get(value int) *node {
	iterator := l.head
	for ; iterator.next != nil; iterator = iterator.next{
		if iterator.value == value {
			return iterator
		}
	}
	return nil
}

func (n node) String() string {
	return fmt.Sprintf("%d", n.value)
}

func (l linkedList) String() string{
	sb := strings.Builder{}

	for iterator := l.head; iterator != nil; iterator = iterator.next{
		sb.WriteString(fmt.Sprintf("%d ", iterator.value))
	}

	return sb.String()
}