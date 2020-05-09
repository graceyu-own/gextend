package container

/*type hashMap struct {
	root   *hashMapNode
	number int
	colors struct {
		red   bool
		black bool
	}
}

type hashMapNode struct {
	hashCode int
	key      lang.Object_
	value    interface{}
	parent   *hashMapNode
	left     *hashMapNode
	right    *hashMapNode
	color    bool
}

func NewHashMap() *hashMap {
	return &hashMap{nil, 0, struct {
		red   bool
		black bool
	}{red: true, black: false}}
}

func (_this *hashMap) Size() int {
	return _this.number
}

func (_this *hashMap) IsEmpty() bool {
	return 0 == _this.Size()
}

func (_this *hashMap) rotateLeft(node *hashMapNode) *hashMapNode {

	if nil != node && nil != node.right {

		// 节点换位
		oldRight := node.right
		node.right = oldRight.left
		oldRight.left = node

		// 节点重设parent
		oldRight.parent = node.parent
		node.parent = oldRight
		if nil != node.right {
			node.right.parent = node
		}

		// 新节点对接到树中
		// 当前节点是根节点
		if nil == oldRight.parent {
			_this.root = oldRight
		} else {
			if oldRight.parent.left == node {
				oldRight.parent.left = oldRight
			} else {
				oldRight.parent.right = oldRight
			}
		}

		return oldRight
	}

	return node
}

func (_this *hashMap) rotateRight(node *hashMapNode) *hashMapNode {

	if nil != node && nil != node.left {

		oldLeft := node.left
		node.left = oldLeft.right
		oldLeft.right = node

		oldLeft.parent = node.parent
		node.parent = oldLeft
		if nil != node.left {
			node.left.parent = node
		}

		if nil == oldLeft.parent {
			_this.root = oldLeft
		} else {
			if oldLeft.parent.left == node {
				oldLeft.parent.left = oldLeft
			} else {
				oldLeft.parent.right = oldLeft
			}
		}

		return oldLeft
	}

	return node
}

func (_this *hashMap) flipColor(node *hashMapNode) {

	if nil != node {

		node.color = !node.color

		if nil != node.left {
			node.left.color = !node.left.color
		}

		if nil != node.right {
			node.right.color = !node.right.color
		}
	}
}

func (_this *hashMap) Put(key lang.Object_, value interface{}) {

	h := 2 * math.Log2(float64(_this.number+1))

	if h < 4 {
		_this.root = _this.putByRecursion(_this.root, key, value)
	} else {
		_this.putByBasic(key, value)
	}

	_this.root.color = _this.colors.black
}

func (_this *hashMap) putByBasic(key lang.Object_, value interface{}) {

	if nil == _this.root {
		_this.root = &hashMapNode{key.HashCode(), key, value, nil, nil, nil, _this.colors.red}
		_this.number++
		return
	}

	cur := _this.root
	for nil != cur {
		comp := cur.key.HashCode() - key.HashCode()
		if comp < 0 {
			if nil == cur.right {
				cur.right = &hashMapNode{key.HashCode(), key, value, cur, nil, nil, _this.colors.red}
				_this.number++
				_this.putFix(cur)
				return
			}

			cur = cur.right
			continue
		}

		if comp > 0 {
			if nil == cur.left {
				cur.left = &hashMapNode{key.HashCode(), key, value, cur, nil, nil, _this.colors.red}
				_this.number++
				_this.putFix(cur)
				return
			}

			cur = cur.left
			continue
		}

		cur.value = value
		return
	}

}

func (_this *hashMap) putByRecursion(node *hashMapNode, key lang.Object_, value interface{}) *hashMapNode {

	if nil == node {
		_this.number++
		return &hashMapNode{key.HashCode(), key, value, nil, nil, nil, _this.colors.red}
	}

	comp := node.key.HashCode() - key.HashCode()
	if comp < 0 {
		node.right = _this.putByRecursion(node.right, key, value)
		node.right.parent = node
	} else if comp > 0 {
		node.left = _this.putByRecursion(node.left, key, value)
		node.left.parent = node
	} else {
		node.value = value
	}

	return _this.putFix(node)
}

func (_this *hashMap) putFix(node *hashMapNode) *hashMapNode {

	if nil != node {

		if nil != node.left && _this.colors.red == node.left.color && nil != node.right && _this.colors.red == node.right.color {
			_this.flipColor(node)
			return _this.putFix(node)
		}

		if nil != node.right && _this.colors.red == node.right.color {
			nr := _this.rotateLeft(node)
			tempColor := nr.color
			nr.color = nr.left.color
			nr.left.color = tempColor
			return _this.putFix(nr)
		}

		if nil != node.left && _this.colors.red == node.left.color && nil != node.left.left && _this.colors.red == node.left.left.color {
			nr := _this.rotateRight(node)
			tempColor := nr.color
			nr.color = nr.right.color
			nr.right.color = tempColor
			return _this.putFix(nr)
		}
	}

	return node
}

func (_this *hashMap) Get(key lang.Object_) interface{} {

	h := 2 * math.Log2(float64(_this.number+1))
	var result *hashMapNode

	if h < 4 {
		result = _this.getByRecursion(_this.root, key)
	} else {
		result = _this.getByBasic(key)
	}

	if nil != result {
		return result.value
	}

	return nil

}

func (_this *hashMap) getByBasic(key lang.Object_) *hashMapNode {

	cur := _this.root
	for nil != cur {
		comp := cur.key.HashCode() - key.HashCode()
		if comp < 0 {
			cur = cur.right
			continue
		}

		if comp > 0 {
			cur = cur.left
			continue
		}

		return cur
	}

	return nil
}

func (_this *hashMap) getByRecursion(node *hashMapNode, key lang.Object_) *hashMapNode {

	if nil != node {

		comp := node.key.HashCode() - key.HashCode()

		if comp < 0 {
			return _this.getByRecursion(node.right, key)
		}

		if comp > 0 {
			return _this.getByRecursion(node.left, key)
		}
	}

	return node
}

func (_this *hashMap) Remove(key lang.Object_) {

	cur := _this.root
	for nil != cur {
		comp := cur.key.HashCode() - key.HashCode()

		if comp < 0 {
			cur = cur.right
			continue
		}

		if comp > 0 {
			cur = cur.left
			continue
		}

		_this.removeByNode(cur)
		return
	}
}

func (_this *hashMap) removeByNode(node *hashMapNode) {

	// deleting node have two child nodes
	// we need find successor node of deleting node
	// replace deleting node property to be successor node property
	// delete successor node
	if nil != node.left && nil != node.right {

		var cur *hashMapNode
		for cur = node.right; nil != cur.left; cur = cur.left {
		}
		node.key = cur.key
		node.value = cur.value
		_this.removeByNode(cur)
		return
	}

	// only one node, the node color is black and child node color is red
	// we need take deleting node child node to be child node of deleting node parent node
	// warning:: deleting node maybe root node
	if nil != node.left {

		node.left.color = _this.colors.black

		// root node
		if nil == node.parent {
			_this.root = node.left
			_this.root.parent = nil
			return
		}

		if node.parent.left == node {
			node.parent.left = node.left
			node.left.parent = node.parent
		} else {
			node.parent.right = node.left
			node.left.parent = node.parent
		}

		return
	}

	// only one node, the node color is black and child node color is red
	// we need take deleting node child node to be child node of deleting node parent node
	// warning:: deleting node maybe root node
	if nil != node.right {

		node.right.color = _this.colors.black

		// root node
		if nil == node.parent {
			_this.root = node.right
			_this.root.parent = nil
			return
		}

		if node.parent.left == node {
			node.parent.left = node.right
			node.right.parent = node.parent
		} else {
			node.parent.right = node.right
			node.right.parent = node.parent
		}

		return
	}

	// deleting node no any child node and color is red
	if _this.colors.red == node.color {

		if node.parent.left == node {
			node.parent.left = nil
		} else {
			node.parent.right = nil
		}

		return
	}

	// deleting node no any child node and color is black
	// warning:: deleting node maybe root node
	if nil == node.parent {
		_this.root = nil
		return
	}

	_this.removeFix(node)
	if node.parent.left == node {
		node.parent.left = nil
	} else {
		node.parent.right = nil
	}

}

func (_this *hashMap) removeFix(node *hashMapNode) {

	// need using deleting node sibling node
	sibling := _this.sibling(node)

	if _this.colors.red == sibling.color {

		tempColor := sibling.color
		sibling.color = node.parent.color
		node.parent.color = tempColor

		if node.parent.left == node {
			_this.rotateLeft(node.parent)
		} else {
			_this.rotateRight(node.parent)
		}

		sibling = _this.sibling(node)
	}

	if _this.colors.black == sibling.color && nil == sibling.left && nil == sibling.right && _this.colors.red == node.parent.color {
		node.parent.color = _this.colors.black
		sibling.color = _this.colors.red
		return
	}

	if _this.colors.black == sibling.color && nil == sibling.left && nil == sibling.right && _this.colors.black == node.parent.color {
		sibling.color = _this.colors.red
		_this.removeFix(node.parent)
		return
	}

	if nil != sibling.left && _this.colors.red == sibling.left.color {

		// deleting node in left
		if node.parent.left == node {
			sibling.left.color = node.parent.color
			node.parent.color = _this.colors.black
			_this.rotateRight(sibling)
			_this.rotateLeft(node.parent)
			return
		}

		// deleting node in right
		sibling.color = node.parent.color
		node.parent.color = _this.colors.black
		_this.rotateRight(node.parent)
		return
	}

	if nil != sibling.right && _this.colors.red == sibling.right.color {

		// deleting node in left
		if node.parent.left == node {
			sibling.color = node.parent.color
			node.parent.color = _this.colors.black
			_this.rotateLeft(node.parent)
			return
		}

		sibling.right.color = node.parent.color
		node.parent.color = _this.colors.black
		_this.rotateLeft(sibling)
		_this.rotateRight(node.parent)
		return
	}
}

func (_this *hashMap) successor(node *hashMapNode) *hashMapNode {

	if nil == node {
		return nil
	}

	var cur *hashMapNode = nil
	for cur = node.right; nil != cur; cur = cur.left {
	}

	return cur
}

func (_this *hashMap) sibling(node *hashMapNode) *hashMapNode {

	if nil == node || nil == node.parent {
		return nil
	}

	if node.parent.left == node {
		return node.parent.right
	} else {
		return node.parent.left
	}

}

func (_this *hashMap) Iterator() *hashMapIterator {
	panic("implement me")
}

func (_this *hashMap) ForEach(fn func(key lang.Object_, value interface{})) {
	_this.forEachByNode(_this.root, fn)
}

func (_this *hashMap) forEachByNode(node *hashMapNode, fn func(key lang.Object_, value interface{})) {

	if nil != node {
		fn(node.key, node.value)
		_this.forEachByNode(node.left, fn)
		_this.forEachByNode(node.right, fn)
	}
}

type hashMapIterator struct {

	cursor *hashMapNode
}

func (_this *hashMapIterator) hasNext() bool {
	panic("implement me")
}

func (_this *hashMapIterator) next() interface{} {
	panic("implement me")
}

func (_this *hashMapIterator) remove() {
	panic("implement me")
}
*/
