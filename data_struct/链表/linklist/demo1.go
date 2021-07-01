package main

import (
	"fmt"
	"strings"
)

type Object interface{}

type Node struct {
	Data Object //定义数据域
	Next *Node
}

type List struct {
	headNode *Node
}

//判断链表是否为空
func (l *List) IsEmpty() bool {
	if l.headNode == nil {
		return true
	}
	return false
}

// 获取链表长度
func (l *List) Length() int {
	count := 0
	cur := l.headNode
	for cur != nil {
		cur = cur.Next
		count++
	}
	return count
}

// 从链表头部添加元素
func (l *List) Add(data Object) *Node {
	node := &Node{Data: data}
	node.Next = l.headNode
	l.headNode = node
	return node
}

// 从链表尾部添加元素
func (l *List) Append(data Object) {
	node := &Node{Data: data}
	if l.IsEmpty() { //如果链表为空，则直接将新元素作为头结点
		l.headNode = node
	} else {
		cur := l.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

// 在链表指定位置添加元素
func (l *List) Insert(index int, data Object) {
	if index < 0 { //如果index<0，则进行头部插入
		l.Add(data)
	} else if index > l.Length() { //如果index大于链表长度，则进行尾部插入
		l.Append(data)
	} else {
		pre := l.headNode
		count := 0
		for count < index-1 {
			pre = pre.Next
			count++
		}
		node := &Node{Data: data}
		node.Next = pre.Next
		pre.Next = node
	}
}

// 删除指定元素
func (l *List) Remove(data Object) {
	pre := l.headNode
	if pre.Data == data{
		l.headNode = pre.Next
	}else {
		for pre.Next != nil {
			if pre.Next.Data == data{
				pre.Next = pre.Next.Next
			}else {
				pre = pre.Next
			}
		}
	}
}

func main() {
	s := "2021-01-03 12:12:32"
	arr := strings.Split(s," ")
	arr1 := strings.Split(arr[0],"-")
	arr2 := strings.Split(arr[1],":")
	var s1 string
	for _,v := range arr1{
		s1 += v
	}

	var s2 string
	for _,v := range arr2{
		s2 += v
	}
	fmt.Println(s1+"."+s2)
}
