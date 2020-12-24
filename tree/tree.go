package tree

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

func (t *tree) Inorder() []int {
	var ret []int
	if t == nil {
		return ret
	}
	ret = append(ret, t.val)

	l := t.left.Inorder()
	ret = append(l, ret...)

	r := t.right.Inorder()
	ret = append(ret, r...)

	return ret
}

func (t *tree) Preorder() []int {
	var ret []int
	if t == nil {
		return ret
	}
	ret = append(ret, t.val)

	l := t.left.Preorder()
	ret = append(ret, l...)

	r := t.right.Preorder()
	ret = append(ret, r...)

	return ret
}

func (t *tree) Postorder() []int {
	var ret []int
	if t == nil {
		return ret
	}

	l := t.left.Postorder()
	r := t.right.Postorder()
	ret = append(l, r...)
	ret = append(ret, t.val)

	return ret
}
