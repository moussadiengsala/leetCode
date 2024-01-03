package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func hasCycle(head *ListNode) bool {
	var visited = make(map[*ListNode]bool)
	var current = head
	for current != nil {
		var _, isVisited = visited[current]
		if isVisited {
			return true
		}
		visited[current] = true
		current = current.Next
	}
	return false
}


/*
// JS

const hasCycle = (head) => {
  let fast = head;
  while (fast && fast.next) {
    head = head.next;
    fast = fast.next.next;
    if (head === fast) return true;
  }
  return false;
};
*/