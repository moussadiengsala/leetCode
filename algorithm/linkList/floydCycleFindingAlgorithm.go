package linklist

func floydCycleFindingAlgorithm(head *ListNode) bool {
	// Initialize two pointers, slow and fast, to the head of the linked list
	slow := head
	fast := head

	// Traverse the linked list
	for fast != nil && fast.Next != nil {
		// Move slow one step at a time
		slow = slow.Next
		// Move fast two steps at a time
		fast = fast.Next.Next

		// Check if the two pointers meet
		if slow == fast {
			return true // Cycle detected
		}
	}

	// If fast reaches the end of the list, there is no cycle
	return false
}
