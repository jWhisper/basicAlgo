package tree

import "testing"

func toll(t *tree) {
	if t == nil {
		return
	}

	toll(t.left)
	toll(t.right)

	if t.right != nil {
		t.right.right = t.left
		t.left = nil
	}
	if t.right == nil {
		t.right = t.left
		t.left = nil
	}
}
func TestTree2Linked(t *testing.T) {
	li := []int{2, 4, 5, 1, 2, 6, 3}
	tr := NewBTree(li)
	inlist := tr.Inorder()
	t.Log(inlist)
	toll(tr)

	for tr != nil {
		t.Log(tr)
		tr = tr.right
	}
}
