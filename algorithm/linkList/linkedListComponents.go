package linklist


func numComponents(head *ListNode, nums []int) int {

	numsSet := make(map[int]bool)
	for _, num := range nums {
		numsSet[num] = true
	}

	var componentConnected int
	var inComponent bool

	for current := head; current != nil; current = current.Next {
		if numsSet[current.Val] {
			if !inComponent {
				componentConnected++
				inComponent = true
			}
		} else {
			inComponent = false
		}
	}

	return componentConnected
}
