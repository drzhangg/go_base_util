package main

type Node struct {
	Key, Value int
	//双向链表
	Pre, Next *Node
}

type LruCache struct {
	//现有元素数量
	Size int
	//lru容量
	Capacity int
	//map存储链表节点，查询速度O(1)
	HashMap map[int]*Node
	//哨兵头尾，方便节点插入初始位置
	Dummy, Tail *Node
}

//初始化新节点
func InitNode(key, value int) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}

//初始化lru
func (m *LruCache) InitLruCache(capacity int) LruCache {
	l := LruCache{
		Capacity: capacity,
		HashMap:  make(map[int]*Node, capacity),
		Dummy:    InitNode(0, 0),
		Tail:     InitNode(0, 0),
	}
	//头尾连接
	l.Dummy.Next = l.Tail
	l.Tail.Pre = l.Dummy
	return l
}

// 两个基础函数，增加到头，移除节点，方便使用

//增加到头部
func (m *LruCache) AddToHead(n *Node) {
	//双向链表，两把都增加
	n.Next = m.Dummy.Next
	n.Pre = m.Dummy
	m.Dummy.Next.Pre = n
	m.Dummy.Next = n
}

//移除节点
func (m *LruCache) Remove(n *Node) {
	n.Pre.Next = n.Next
	n.Next.Pre = n.Pre
}

// 节点移动，删除节点，移动到头部
func (m *LruCache) MoveHead(n *Node) {
	m.MoveHead(n)
	m.AddToHead(n)
}

//超过容量，删除最后一个
//其实是删除tail的pre，返回node，方便移除map中的数据
func (m *LruCache) DelTail() *Node {
	n := m.Tail.Pre
	m.Remove(n)
	return n
}
