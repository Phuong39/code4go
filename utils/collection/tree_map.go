package collection

const (
	red   = false
	black = true
)

var null = &treeNode{color: black}

type TreeMap struct {
	len  int
	root *treeNode
}

type treeNode struct {
	key    Comparable
	value  interface{}
	color  bool
	left   *treeNode
	right  *treeNode
	parent *treeNode
}

func (t *TreeMap) Put(key Comparable, value interface{}) bool {
	t.checkRoot()
	temp := &treeNode{key: key, value: value}
	y := null
	x := t.root
	for x != null {
		y = x
		if temp.key.Compare(x.key) < 0 {
			x = x.left
		} else if temp.key.Compare(x.key) > 0 {
			x = x.right
		} else {
			x.value = temp.value
			return false
		}
	}
	temp.parent = y
	if y == null {
		t.root = temp
	} else if temp.key.Compare(x.key) < 0 {
		y.left = temp
	} else {
		y.right = temp
	}
	temp.left = null
	temp.right = null
	temp.color = red
	t.insertFixup(temp)
	t.len++
	return true
}

func (t *TreeMap) Get(key Comparable) interface{} {
	t.checkRoot()
	temp := t.findNode(key)
	if temp != nil {
		return temp.value
	}
	return nil
}

func (t *TreeMap) Size() int {
	return t.len
}

func (t *TreeMap) IsEmpty() bool {
	return t.len == 0
}

func (t *TreeMap) ContainsKey(key Comparable) bool {
	t.checkRoot()
	if t.root != null {
		temp := t.root
		for temp != null {
			if key.Compare(temp.key) > 0 {
				temp = temp.right
			} else if key.Compare(temp.key) < 0 {
				temp = temp.left
			} else {
				return true
			}
		}
	}
	return false
}

func (t *TreeMap) ContainsValue(value interface{}) bool {
	t.checkRoot()
	var isFind bool
	t.ForEach(func(key Comparable, value interface{}) {
		if value == value {
			isFind = true
			return
		}
	})
	return isFind
}

func (t *TreeMap) minNode(node *treeNode) *treeNode {
	for node.left != null {
		node = node.left
	}
	return node
}

func (t *TreeMap) maxNode(node *treeNode) *treeNode {
	for node.right != null {
		node = node.right
	}
	return node
}

func (t *TreeMap) FirstKey() Comparable {
	t.checkRoot()
	min := t.minNode(t.root)
	if min != null {
		return nil
	} else {
		return min.key
	}
}

func (t *TreeMap) LastKey() Comparable {
	t.checkRoot()
	max := t.maxNode(t.root)
	if max != nil {
		return nil
	} else {
		return max.key
	}
}

func (t *TreeMap) FirstValue() interface{} {
	t.checkRoot()
	min := t.minNode(t.root)
	if min != null {
		return nil
	} else {
		return min.value
	}
}

func (t *TreeMap) LastValue() interface{} {
	t.checkRoot()
	max := t.maxNode(t.root)
	if max != null {
		return nil
	} else {
		return max.value
	}
}

func (t *TreeMap) Remove(key Comparable) {
	if t.root == nil || t.root == null {
		return
	}
	temp := t.root
	for temp != null {
		if key.Compare(temp.key) < 0 {
			temp = temp.left
		} else if key.Compare(temp.key) > 0 {
			temp = temp.right
		} else {
			t.delete(temp)
			t.len--
			break
		}
	}
}

func (t *TreeMap) delete(node *treeNode) {
	y := node
	yOldColor := y.color
	var x *treeNode
	if node.left == null {
		x = node.right
		t.transplant(node, node.right)
	} else if node.right == null {
		x = node.left
		t.transplant(node, node.left)
	} else {
		y = t.minNode(node.right)
		yOldColor = y.color
		x = y.right
		if x.parent == node {
			x.parent = y
		} else {
			t.transplant(y, y.right)
			y.right = node.right
			y.right.parent = y
		}
		t.transplant(node, y)
		y.left = node.left
		y.left.parent = y
		y.color = node.color
	}
	if yOldColor == black {
		t.deleteFixup(x)
	}
}

func (t *TreeMap) findNode(key Comparable) *treeNode {
	if t.root != null {
		temp := t.root
		for temp != null {
			if key.Compare(temp.key) > 0 {
				temp = temp.right
			} else if key.Compare(temp.key) < 0 {
				temp = temp.left
			} else {
				return temp
			}
		}
	}
	return null
}

func (t *TreeMap) ForEach(each func(key Comparable, value interface{})) {
	t.checkRoot()
	t.forEach(t.root, each)
}

func (t *TreeMap) forEach(node *treeNode, each func(key Comparable, value interface{})) {
	if node != null {
		if node.left != null {
			t.forEach(node.left, each)
		}
		each(node.key, node.value)
		if node.right != null {
			t.forEach(node.right, each)
		}
	}
}

func (t *TreeMap) insertFixup(node *treeNode) {
	for node.parent.color == red {
		if node.parent == node.parent.parent.left {
			y := node.parent.parent.right
			if y.color == red {
				node.parent.color = black
				y.color = black
				node.parent.parent.color = red
				node = node.parent.parent
			} else if node == node.parent.right {
				node = node.parent
				t.leftRotate(node)
			} else {
				node.parent.color = black
				node.parent.parent.color = red
				t.rightRotate(node.parent.parent)
			}
		} else {
			x := node.parent.parent.left
			if x.color == red {
				node.parent.color = black
				x.color = black
				node.parent.parent.color = red
				node = node.parent.parent
			} else if node == node.parent.left {
				node = node.parent
				t.rightRotate(node)
			} else {
				node.parent.color = black
				node.parent.parent.color = red
				t.leftRotate(node.parent.parent)
			}
		}
	}
	t.root.color = black
}

func (t *TreeMap) deleteFixup(node *treeNode) {
	for node != t.root && node.color == black {
		if node == node.parent.left {
			w := node.parent.right
			if w.color == red {
				w.color = black
				node.parent.color = red
				t.leftRotate(node.parent)
				w = node.parent.right
			}
			if w.left.color == black && w.right.color == black {
				w.color = red
				node = node.parent
			} else if w.right.color == black {
				w.left.color = black
				w.color = red
				t.rightRotate(w)
				w = node.parent.right
			} else {
				w.color = node.parent.color
				node.parent.color = black
				w.right.color = black
				t.leftRotate(node.parent)
				node = t.root
			}
		} else {
			w := node.parent.left
			if w.color == red {
				w.color = black
				node.parent.color = red
				t.leftRotate(node.parent)
				w = node.parent.left
			}
			if w.right.color == black && w.left.color == black {
				w.color = red
				node = node.parent
			} else if w.left.color == black {
				w.right.color = black
				w.color = red
				t.leftRotate(w)
				w = node.parent.left
			} else {
				w.color = node.parent.color
				node.parent.color = black
				w.left.color = black
				t.rightRotate(node.parent)
				node = t.root
			}
		}
	}
	node.color = black
}

func (t *TreeMap) transplant(old, new *treeNode) {
	if old.parent == null {
		t.root = new
	} else if old == old.parent.left {
		old.parent.left = new
	} else {
		old.parent.right = new
	}
	new.parent = old.parent
}

func (t *TreeMap) leftRotate(x *treeNode) {
	y := x.right
	x.right = y.left
	if y.left != null {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == null {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *TreeMap) rightRotate(y *treeNode) {
	x := y.left
	y.left = x.right
	if x.right != null {
		x.right.parent = y
	}
	x.parent = y.parent
	if x.parent == null {
		t.root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}
	x.right = y
	y.parent = x
}

func (t *TreeMap) Clear() {
	t.len = 0
	t.root = null
}

func (t *TreeMap) checkRoot() {
	if t.root == nil {
		t.root = null
	}
}

//func (t *TreeMap) Remove(key Comparable) {
//	temp := t.findNode(key)
//	if temp != nil {
//		if temp.left != nil && temp.right != nil {
//			replace := t.maxNode(temp.right)
//			t.Remove(replace.key)
//			temp.key = replace.key
//			temp.value = replace.value
//		} else {
//			if temp.left != nil {
//				temp.left.parent = temp.parent
//				if temp.parent != nil {
//					temp.parent.left = temp.left
//				} else {
//					t.root = temp.left
//				}
//				temp.left = nil
//			} else if temp.right != nil {
//				temp.right.parent = temp.parent
//				if temp.parent != nil {
//					temp.parent.right = temp.right
//				} else {
//					t.root = temp.right
//				}
//				temp.right = nil
//			} else {
//				if temp.parent != nil {
//					if temp.parent.left == temp {
//						temp.parent.left = nil
//					} else {
//						temp.parent.right = nil
//					}
//				}
//			}
//			temp = nil
//		}
//		t.len--
//	}
//}

//func (t *TreeMap) Put(key Comparable, value interface{}) bool {
//	if t.root == nil {
//		t.root = &treeNode{key: key, value: value}
//	} else {
//		temp := t.root
//		for {
//			if key.Compare(temp.key) > 0 {
//				if temp.right == nil {
//					temp.right = &treeNode{key: key, value: value, parent: temp}
//					break
//				}
//				temp = temp.right
//			} else if key.Compare(temp.key) < 0 {
//				if temp.left == nil {
//					temp.left = &treeNode{key: key, value: value, parent: temp}
//					break
//				}
//				temp = temp.left
//			} else {
//				temp.value = value
//				return false
//			}
//		}
//	}
//	t.len++
//	return true
//}
