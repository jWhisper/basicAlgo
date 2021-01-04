package link

type linkNode struct {
	val  int
	next *linkNode
}

func NewLink(vals []int) *linkNode {
	var head, cur *linkNode
	for i := 0; i < len(vals); i++ {
		if cur == nil {
			cur = &linkNode{val: vals[i]}
			if i == 0 {
				head = cur
			}
		}
		cur = cur.next
	}
	return head
}

func ReverseLink(root *linkNode) *linkNode {
	if root.next == nil {
		return root
	}
	last := ReverseLink(root.next)
	root.next.next = root
	root.next = nil
	return last
}
