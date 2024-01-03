package linklist

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Insert(head *ListNode, data int) *ListNode {
	var newNode = &ListNode{Val: data, Next: nil}

	if head == nil {
		return newNode
	}

	var current = head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	return head
}

func Read(head *ListNode) []int {
	var current = head
	var integers = []int{}
	for current != nil {
		integers = append(integers, current.Val)
		current = current.Next
	}
	return integers
}

func sumTwoArray(arr1, arr2 []int) []int {
	var maxLen = int(math.Max(float64(len(arr1)), float64(len(arr2))))
	var result = []int{}
	var reminder = 0

	for index := 0; index < maxLen; index++ {
		var op1, op2 = 0, 0

		if len(arr1) > index {
			op1 = arr1[index]
		}

		if len(arr2) > index {
			op2 = arr2[index]
		}

		sum := op1 + op2 + reminder
		reminder = sum / 10

		result = append(result, sum%10)
	}

	if reminder > 0 {
		result = append(result, reminder)
	}

	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var list *ListNode
	var integers = sumTwoArray(Read(l1), Read(l2))

	for i := 0; i < len(integers); i++ {
		list = Insert(list, integers[i])
	}
	return list
}
