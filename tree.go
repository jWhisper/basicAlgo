package main

type tree struct {
	val         int
	left, right *tree
}

//完全二叉树
func NewBTree(l []int) *tree {
	if len(l) < 1 {
		return nil
	}
	// branch := 2
	head := new(tree)
	ch := make(chan *tree, (len(l)+1)/2)

	defer func() {
		close(ch)
	}()

	head.val = l[0]
	ch <- head
	//保证l[i]不会outindex
	i := 1
	for {
		head := <-ch

		if i >= len(l) {
			break
		}
		left := &tree{val: l[i]}
		//有些节点是没有子节点
		if i <= len(l)/2 {
			ch <- left
		}
		head.left = left
		i++

		if i >= len(l) {
			break
		}
		right := &tree{val: l[i]}
		head.right = right
		if i <= len(l)/2 {
			ch <- right
		}
		i++

	}

	return head
}

func (t *tree) inorder() []int {
	var ret []int
	if t == nil {
		return ret
	}
	ret = append(ret, t.val)

	l := t.left.inorder()
	ret = append(l, ret...)

	r := t.right.inorder()
	ret = append(ret, r...)

	return ret
}

func (t *tree) preorder() []int {
	var ret []int
	if t == nil {
		return ret
	}
	ret = append(ret, t.val)

	l := t.left.preorder()
	ret = append(ret, l...)

	r := t.right.preorder()
	ret = append(ret, r...)

	return ret
}

func (t *tree) postorder() []int {
	var ret []int
	if t == nil {
		return ret
	}

	l := t.left.postorder()
	r := t.right.postorder()
	ret = append(l, r...)
	ret = append(ret, t.val)

	return ret
}
