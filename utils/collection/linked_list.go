package collection

import "errors"

/**
* 作者：刘时明
* 时间：2019/8/13-17:37
* 作用：LinkedList实现类
 */

// 节点
type node struct {
	data interface{}
	prev *node
	next *node
}

// 链表结构体
type LinkedList struct {
	// 当前链表长度
	len int
	// 头节点
	head *node
	// 尾节点
	tail *node
}

func (l *LinkedList) IsEmpty() bool {
	return l.len == 0
}

func (l *LinkedList) Remove(ele interface{}) bool {
	index := l.IndexOf(ele)
	if index == -1 {
		return false
	}
	_ = l.RemoveByIndex(index)
	return true
}

func (l *LinkedList) IndexOf(ele interface{}) int {
	temp := l.head
	var curr int
	for temp != nil {
		if temp.data == ele {
			return curr
		}
		temp = temp.next
		curr++
	}
	return -1
}

func (l *LinkedList) LastIndexOf(ele interface{}) int {
	temp := l.tail
	var curr int
	for temp != nil {
		if temp.data == ele {
			return l.len - curr - 1
		}
		temp = temp.prev
		curr++
	}
	return -1
}

func (l *LinkedList) Add(ele interface{}) bool {
	l.AddLast(ele)
	return true
}

func (l *LinkedList) Contains(ele interface{}) bool {
	temp := l.head
	for temp != nil {
		if temp.data == ele {
			return true
		}
		temp = temp.next
	}
	return false
}

func (l *LinkedList) AddByIndex(index int, ele interface{}) error {
	if index < 0 || index > l.len {
		return errors.New("索引越界")
	}
	if index == 0 {
		l.AddFirst(ele)
	} else if index == l.len {
		l.AddLast(ele)
	} else {
		var curr int
		if index < l.len/2-1 {
			temp := l.head
			for curr < index-1 {
				temp = temp.next
				curr++
			}
			newNode := &node{data: ele, prev: temp, next: temp.next}
			temp.next = newNode
			temp.next.prev = newNode
		} else {
			temp := l.tail
			for curr < l.len-index-1 {
				temp = temp.prev
				curr++
			}
			newNode := &node{data: ele, prev: temp.prev, next: temp}
			temp.prev.next = newNode
			temp.prev = newNode
		}
		l.len++
	}
	return nil
}

func (l *LinkedList) AddLast(ele interface{}) {
	if l.head == nil {
		l.head = &node{data: ele, prev: nil, next: l.tail}
	} else {
		if l.tail == nil {
			l.tail = &node{data: ele, prev: l.head, next: nil}
			l.head.next = l.tail
		} else {
			temp := &node{data: ele, prev: nil, next: nil}
			l.tail.next = temp
			temp.prev = l.tail
			l.tail = temp
		}
	}
	l.len++
}

func (l *LinkedList) AddFirst(ele interface{}) {
	if l.head == nil {
		l.head = &node{data: ele, prev: nil, next: l.tail}
	} else {
		temp := &node{data: ele, prev: nil, next: l.head}
		l.head.prev = temp
		l.head = temp
	}
	l.len++
}

func (l *LinkedList) RemoveLast() {
	_ = l.RemoveByIndex(l.len - 1)
}

func (l *LinkedList) RemoveFirst() {
	_ = l.RemoveByIndex(0)
}

func (l *LinkedList) RemoveByIndex(index int) error {
	if index < 0 || index >= l.len {
		return errors.New("数组越界异常")
	}
	var curr int
	if index < l.len/2-1 {
		temp := l.head
		for curr < index {
			temp = temp.next
			curr++
		}
		if temp.prev == nil {
			next := temp.next
			temp.next = nil
			if next != nil {
				next.prev = nil
			}
			l.head = next
		} else {
			prev := temp.prev
			prev.next = temp.next
			if temp.next != nil {
				temp.next.prev = prev
			}
			temp.next = nil
			temp.prev = nil
		}
	} else {
		temp := l.tail
		index = l.len - index - 1
		for curr < index {
			temp = temp.prev
			curr++
		}
		if temp.next == nil {
			prev := temp.prev
			temp.prev = nil
			if prev != nil {
				prev.next = nil
			}
			l.tail = prev
		} else {
			next := temp.next
			next.prev = temp.prev
			if temp.prev != nil {
				temp.prev.next = next
			}
			temp.next = nil
			temp.prev = nil
		}
	}
	l.len--
	return nil
}

func (l *LinkedList) PollFirst() interface{} {
	temp := l.GetFirst()
	l.RemoveFirst()
	return temp
}

func (l *LinkedList) PollLast() interface{} {
	temp := l.GetLast()
	l.RemoveLast()
	return temp
}

func (l *LinkedList) GetLast() interface{} {
	data, _ := l.Get(l.len - 1)
	return data
}

func (l *LinkedList) GetFirst() interface{} {
	data, _ := l.Get(0)
	return data
}

func (l *LinkedList) Size() int {
	return l.len
}

func (l *LinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= l.len {
		return nil, errors.New("数组越界异常")
	}
	var curr int
	if index < l.len/2-1 {
		temp := l.head
		for curr < index {
			temp = temp.next
			curr++
		}
		return temp.data, nil
	} else {
		temp := l.tail
		index = l.len - index - 1
		for curr < index {
			temp = temp.prev
			curr++
		}
		return temp.data, nil
	}
}

func (l *LinkedList) Clear() {
	l.len = 0
	l.head = nil
	l.tail = nil
}

func (l *LinkedList) ForEach(each func(ele interface{})) {
	temp := l.head
	for temp != nil {
		each(temp.data)
		temp = temp.next
	}
}

func (l *LinkedList) AddAll(c Collection) {
	l.ForEach(func(ele interface{}) {
		l.Add(ele)
	})
}

func (l *LinkedList) RemoveAll(c Collection) {
	l.ForEach(func(ele interface{}) {
		l.Remove(ele)
	})
}

func (l *LinkedList) ToArray() []interface{} {
	data := make([]interface{}, 0, l.len)
	l.ForEach(func(ele interface{}) {
		data = append(data, ele)
	})
	return data
}
