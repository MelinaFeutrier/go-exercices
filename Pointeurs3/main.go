package main

import "fmt"

// Définition de la structure ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}

	return dummy.Next
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, "->")
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {

	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	fmt.Println("Liste 1 :")
	printList(list1)

	fmt.Println("Liste 2 :")
	printList(list2)

	mergedList := mergeTwoLists(list1, list2)

	fmt.Println("Liste fusionnée :")
	printList(mergedList)
}
