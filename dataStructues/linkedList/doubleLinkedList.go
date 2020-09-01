package main

import (
	"fmt"
)

type node struct {
	val  int
	prev *node
	next *node
}

type doubleLinkedList struct {
	head *node
}

func newNode(val int) *node {
	n := &node{}
	n.val = val
	n.prev = nil
	n.next = nil
	return n
}

func (ll *doubleLinkedList) addAtBegain(val int) {
	n := newNode(val)
	n.next = ll.head()
	ll.head = n
}

func (ll *doubleLinkedList) addAtEnd(val int) {
	n := newNode(val)

	if ll.head == nil {
		ll.head = n
		return
	}

	cur := ll.head
	for ; cur.next != nil; cur = cur.next {
	}

	// 尾部双向链接上
	cur.next = n
	n.prev = cur
}

/*
 *
 *
 */
