package collection

/**
* 作者：刘时明
* 时间：2019/8/14-19:40
* 作用：集合接口定义
 */

// 集合顶层接口
type Collection interface {
	// 获取集合元素数量
	Size() int
	// 集合是否为空
	IsEmpty() bool
	// 查找元素
	Contains(ele interface{}) bool
	// 删除指定元素
	Remove(ele interface{}) bool
	// 添加元素
	Add(ele interface{}) bool
	// 清空集合
	Clear()
	// 添加集合
	AddAll(c Collection)
	// 删除集合
	RemoveAll(c Collection)
	// 集合遍历器
	ForEach(each func(ele interface{}))
	// 转换为切片
	ToArray() []interface{}
}

// 针对树设计的Map集合顶层接口
type Map interface {
	// 获取集合元素数量
	Size() int
	// 集合是否为空
	IsEmpty() bool
	// 根据Key查找元素
	ContainsKey(key Comparable) bool
	// 根据Value查找元素
	ContainsValue(value interface{}) bool
	// 添加元素
	Put(key Comparable, value interface{}) bool
	// 获取元素
	Get(key Comparable) interface{}
	// 删除元素
	Remove(key Comparable)
	// 清空集合
	Clear()
	// 集合遍历器
	ForEach(each func(key Comparable, value interface{}))
}

// 有序Map
type NavigableMap interface {
	Map
	// 获取首个Key
	FirstKey() Comparable
	// 获取末尾Key
	LastKey() Comparable
	// 获取首个Key对应的值
	FirstValue() interface{}
	// 获取末尾Key对应的值
	LastValue() interface{}
}

// 比较器接口
type Comparable interface {
	Compare(o interface{}) int
}

// 线性表
type List interface {
	Collection
	// 指定位置插入
	AddByIndex(index int, ele interface{}) error
	// 获取指定位置元素
	Get(index int) (interface{}, error)
	// 指定位置删除
	RemoveByIndex(index int) error
	// 查找元素并获取下标
	IndexOf(ele interface{}) int
	// 从尾部查找元素并获取下标
	LastIndexOf(ele interface{}) int
}

// 哈希表
type Set interface {
	Collection
}

// 双向队列
type Deque interface {
	Collection
	// 获取首部位置元素
	GetFirst() interface{}
	// 获取尾部位置元素
	GetLast() interface{}
	// 首部位置入队
	AddFirst(ele interface{})
	// 尾部位置入队
	AddLast(ele interface{})
	// 删除首部位置元素
	RemoveFirst()
	// 删除尾部位置元素
	RemoveLast()
	// 首部位置出队
	PollFirst() interface{}
	// 尾部位置出队
	PollLast() interface{}
}

// 阻塞队列
type BlockingDeque interface {
	Deque
	// 入队，如果队列已满则线程阻塞
	Put(ele interface{})
	// 出队，如果队列为空则线程阻塞
	Take() interface{}
}
