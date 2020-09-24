package collection

/**
* 作者：刘时明
* 时间：2019/8/15-17:56
* 作用：PriorityQueue优先队列
 */

// 基于最小二叉堆的双向优先队列
type PriorityDeque struct {
	list ArrayList
}

func (p *PriorityDeque) IsEmpty() bool {
	return p.list.IsEmpty()
}

func (p *PriorityDeque) Contains(ele interface{}) bool {
	for _, v := range p.list.data {
		if v == ele {
			return true
		}
	}
	return false
}

func (p *PriorityDeque) Clear() {
	p.list.Clear()
}

func (p *PriorityDeque) ToArray() []interface{} {
	return p.list.data
}

func (p *PriorityDeque) AddAll(c Collection) {
	p.ForEach(func(ele interface{}) {
		p.Add(ele)
	})
}

func (p *PriorityDeque) RemoveAll(c Collection) {
	p.ForEach(func(ele interface{}) {
		if c.Contains(ele) {
			p.Remove(ele)
		}
	})
}

func (p *PriorityDeque) Remove(ele interface{}) bool {
	index := p.IndexOf(ele)
	if index > -1 {
		_ = p.list.RemoveByIndex(index)
	}
	return false
}

func (p *PriorityDeque) Size() int {
	return p.list.Size()
}

func (p *PriorityDeque) Add(ele interface{}) bool {
	temp, ok := ele.(Comparable)
	if !ok {
		return false
	}
	p.list.Add(temp)
	p.siftUp(p.list.Size() - 1)
	return true
}

// AddFirst 优先队列只能Add
func (p *PriorityDeque) AddFirst(ele interface{}) {
	_ = p.Add(ele)
}

// AddLast 优先队列只能Add
func (p *PriorityDeque) AddLast(ele interface{}) {
	_ = p.Add(ele)
}

func (p *PriorityDeque) PollFirst() interface{} {
	temp := p.GetFirst()
	if temp != nil {
		p.swap(0, p.list.Size()-1)
		_ = p.list.RemoveByIndex(p.list.Size() - 1)
		p.siftDown(0)
	}
	return temp
}

func (p *PriorityDeque) RemoveFirst() {
	_ = p.PollFirst()
}

func (p *PriorityDeque) RemoveLast() {
	_ = p.PollLast()
}

func (p *PriorityDeque) PollLast() interface{} {
	temp := p.GetLast()
	if temp != nil {
		_ = p.list.RemoveByIndex(p.list.Size() - 1)
	}
	return temp
}

func (p *PriorityDeque) GetFirst() interface{} {
	if p.list.Size() > 0 {
		temp, _ := p.list.Get(0)
		return temp
	}
	return nil
}

func (p *PriorityDeque) GetLast() interface{} {
	if p.list.Size() > 0 {
		temp, _ := p.list.Get(p.list.Size() - 1)
		return temp
	}
	return nil
}

func (p *PriorityDeque) siftUp(index int) {
	for index > 0 && p.compareParent(index) {
		p.swap(index, parentIndex(index))
		index = parentIndex(index)
	}
}

func (p *PriorityDeque) siftDown(index int) {
	size := p.list.Size()
	for rightChildIndex(index) < size {
		temp := rightChildIndex(index)
		if leftChildIndex(index) < size && p.compareEle(temp+1, temp) {
			temp = leftChildIndex(index)
		}
		if p.compareEle(index, temp) {
			break
		}
		p.swap(index, temp)
		index = temp
	}
}

func (p *PriorityDeque) compareParent(index int) bool {
	return p.compareEle(index, parentIndex(index))
}

func (p *PriorityDeque) compareEle(i, j int) bool {
	e1, _ := p.list.Get(i)
	e2, _ := p.list.Get(j)
	c1 := e1.(Comparable)
	c2 := e2.(Comparable)
	return c1.Compare(c2) < 0
}

func (p *PriorityDeque) swap(i, j int) {
	p.list.data[i], p.list.data[j] = p.list.data[j], p.list.data[i]
}

func (p *PriorityDeque) ForEach(each func(e interface{})) {
	for _, v := range p.list.data {
		each(v)
	}
}

func (p *PriorityDeque) IndexOf(ele interface{}) int {
	for i, v := range p.list.data {
		if ele == v {
			return i
		}
	}
	return -1
}

func parentIndex(index int) int {
	if index == 0 {
		return -1
	}
	return (index - 1) / 2
}

func leftChildIndex(index int) int {
	return index*2 + 1
}

func rightChildIndex(index int) int {
	return index * 2
}
