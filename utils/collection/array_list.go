package collection

import "errors"

/**
* 作者：刘时明
* 时间：2019/8/14-19:41
* 作用：ArrayList实现类
 */

// 基于数组的线性表
type ArrayList struct {
	data []interface{}
}

func (a *ArrayList) IsEmpty() bool {
	return len(a.data) == 0
}

func (a *ArrayList) Contains(ele interface{}) bool {
	for _, v := range a.data {
		if v == ele {
			return true
		}
	}
	return false
}

func (a *ArrayList) Add(ele interface{}) bool {
	a.AddLast(ele)
	return true
}

func (a *ArrayList) Clear() {
	a.data = a.data[0:0]
}

func (a *ArrayList) Remove(ele interface{}) bool {
	index := a.IndexOf(ele)
	if index < 0 {
		return false
	} else {
		_ = a.RemoveByIndex(index)
		return true
	}
}

func (a *ArrayList) Size() int {
	return len(a.data)
}

func (a *ArrayList) AddByIndex(index int, ele interface{}) error {
	if index < 0 || index > len(a.data) {
		return errors.New("索引越界")
	}
	if index == 0 {
		temp := []interface{}{ele}
		a.data = append(temp, a.data...)
	} else if index == len(a.data) {
		a.data = append(a.data, ele)
	} else {
		left := a.data[0:index]
		right := a.data[index:len(a.data)]
		temp := a.data[index]
		left = append(left, ele)
		a.data = append(left, right...)
		a.data[index+1] = temp
	}
	return nil
}

func (a *ArrayList) AddLast(ele interface{}) {
	_ = a.AddByIndex(len(a.data), ele)
}

func (a *ArrayList) AddFirst(ele interface{}) {
	_ = a.AddByIndex(0, ele)
}

func (a *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index > len(a.data)-1 {
		return nil, errors.New("索引越界")
	}
	return a.data[index], nil
}

func (a *ArrayList) GetLast() interface{} {
	data, _ := a.Get(len(a.data) - 1)
	return data
}

func (a *ArrayList) GetFirst() interface{} {
	data, _ := a.Get(0)
	return data
}

func (a *ArrayList) RemoveByIndex(index int) error {
	if index < 0 || index > len(a.data)-1 {
		return errors.New("索引越界")
	}
	if len(a.data) > 0 {
		left := a.data[0:index]
		right := a.data[index+1 : len(a.data)]
		a.data = append(left, right...)
	}
	return nil
}

func (a *ArrayList) IndexOf(ele interface{}) int {
	for i, v := range a.data {
		if v == ele {
			return i
		}
	}
	return -1
}

func (a *ArrayList) LastIndexOf(ele interface{}) int {
	for i := len(a.data) - 1; i >= 0; i-- {
		if ele == a.data[i] {
			return i
		}
	}
	return -1
}

func (a *ArrayList) ForEach(each func(ele interface{})) {
	for _, v := range a.data {
		each(v)
	}
}

func (a *ArrayList) AddAll(c Collection) {
	temp := c.ToArray()
	a.data = append(a.data, temp...)
}

func (a *ArrayList) RemoveAll(c Collection) {
	c.ForEach(func(ele interface{}) {
		a.Remove(ele)
	})
}

func (a *ArrayList) ToArray() []interface{} {
	return a.data
}
