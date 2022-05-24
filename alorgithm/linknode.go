package main

import (
	"fmt"
)

// 单链表
// 定义节点
type SingleNode struct {
	Value int
	Next  *SingleNode
}

type DoubleNode struct {
	Value    int
	Previous *DoubleNode
	Next     *DoubleNode
}

// 循环链表
type RingNode struct {
	next, prev *RingNode   // 前驱和后驱节点
	Value      interface{} // 数据
}

// 添加节点
func (s *SingleNode) SingleAddNode(v int) int {
	if v == s.Value {
		fmt.Println("节点已存在:", v)
		return -1
	}

	// 如果当前节点下一个节点为空
	if s.Next == nil {
		s.Next = &SingleNode{v, nil}
		return -2
	}

	// 如果当前节点下一个节点不为空
	//return t.SingleAddNode(t.Next, v)
	return s.Next.SingleAddNode(v)
}

// 遍历链表
func (s *SingleNode) SingleTraverse() {
	if s == nil {
		fmt.Println("-> 空链表!")
		return
	}

	for s != nil {
		fmt.Printf("%d -> ", s.Value)
		s = s.Next
	}

	fmt.Println()
}

// 查找节点
func (s *SingleNode) SingleLookupNode(v int) bool {
	if v == s.Value {
		return true
	}

	if s.Next == nil {
		return false
	}

	return s.Next.SingleLookupNode(v)
}

// 获取链表长度
func (s *SingleNode) SingleSize() int {
	if s == nil {
		fmt.Println("-> 空链表!")
		return 0
	}

	i := 0
	for s != nil {
		i++
		s = s.Next
	}
	return i
}

///////////////////////////////////////////////////////////////
// 定义节点

// 添加节点
func (d *DoubleNode) DoubleAddNode(v int) int {
	if v == d.Value {
		fmt.Println("节点已存在:", v)
		return -1
	}

	// 如果当前节点下一个节点为空
	if d.Next == nil {
		// 与单链表不同的是每个节点还要维护前驱节点指针
		temp := d
		d.Next = &DoubleNode{v, temp, nil}
		return -2
	}

	// 如果当前节点下一个节点不为空
	//return DoubleAddNode(t.Next, v)
	return d.Next.DoubleAddNode(v)
}

// 遍历链表
func (d *DoubleNode) DoubleTraverse() {
	if d == nil {
		fmt.Println("-> 空链表!")
		return
	}

	for d != nil {
		fmt.Printf("%d -> ", d.Value)
		d = d.Next
	}

	fmt.Println()
}

// 反向遍历链表
func (d *DoubleNode) DoubleReverse() {
	if d == nil {
		fmt.Println("-> 空链表!")
		return
	}

	temp := d
	for d != nil {
		temp = d
		d = d.Next
	}

	for temp.Previous != nil {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Previous
	}

	fmt.Printf("%d -> ", temp.Value)
	fmt.Println()
}

// 获取链表长度
func (d *DoubleNode) DoubleSize() int {
	if d == nil {
		fmt.Println("-> 空链表!")
		return 0
	}

	n := 0
	for d != nil {
		n++
		d = d.Next
	}

	return n
}

// 查找节点
func (d *DoubleNode) DoubleLookupNode(v int) bool {
	if v == d.Value {
		return true
	}

	if d.Next == nil {
		return false
	}

	return d.Next.DoubleLookupNode(v)
}

func (r *RingNode) init() *RingNode {
	r.next = r
	r.prev = r
	return r
}

// 创建N个节点的循环链表
func New(n int) *RingNode {
	if n <= 0 {
		return nil
	}
	r := new(RingNode)
	p := r
	for i := 1; i < n; i++ {
		p.next = &RingNode{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

func (r *RingNode) Next() *RingNode {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// 获取上一个节点
func (r *RingNode) Prev() *RingNode {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

func (r *RingNode) Move(n int) *RingNode {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}
func (r *RingNode) Link(s *RingNode) *RingNode {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

// 入口函数
func main() {
	fmt.Println("***********单链表************")
	s := SingleNode{Value: 1}
	// 遍历链表
	s.SingleTraverse()
	//SingleTraverse(SingleHead)
	// 添加节点
	s.SingleAddNode(1)
	s.SingleAddNode(-1)
	// 再次遍历
	s.SingleTraverse()
	// 添加更多节点
	s.SingleAddNode(10)
	s.SingleAddNode(5)
	s.SingleAddNode(45)
	// 添加已存在节点
	s.SingleAddNode(5)
	// 再次遍历
	s.SingleTraverse()

	// 查找已存在节点
	if s.SingleLookupNode(5) {
		fmt.Println("该节点已存在!")
	} else {
		fmt.Println("该节点不存在!")
	}

	// 查找不存在节点
	if s.SingleLookupNode(-100) {
		fmt.Println("该节点已存在!")
	} else {
		fmt.Println("该节点不存在!")
	}

	// 双链表
	fmt.Println("***********双链表************")
	d := DoubleNode{Value: 1}
	// 遍历链表
	d.DoubleTraverse()
	// 新增节点
	d.DoubleAddNode(1)
	// 再次遍历
	d.DoubleTraverse()
	// 继续添加节点
	d.DoubleAddNode(10)
	d.DoubleAddNode(5)
	d.DoubleAddNode(100)
	// 再次遍历
	d.DoubleTraverse()
	// 添加已存在节点
	d.DoubleAddNode(100)
	fmt.Println("链表长度:", d.DoubleSize())
	// 再次遍历
	d.DoubleTraverse()
	// 反向遍历
	d.DoubleReverse()
	// 查找已存在节点
	if d.DoubleLookupNode(5) {
		fmt.Println("该节点已存在!")
	} else {
		fmt.Println("该节点不存在!")
	}
	// 循环链表
	// 第一个节点
	fmt.Println("***********循环链表************")
	r := &RingNode{Value: -1}

	// 链接新的五个节点
	r.Link(&RingNode{Value: 1})
	r.Link(&RingNode{Value: 2})
	r.Link(&RingNode{Value: 3})
	r.Link(&RingNode{Value: 4})

	node := r
	for {
		// 打印节点值
		fmt.Println(node.Value)

		// 移到下一个节点
		node = node.Next()

		//  如果节点回到了起点，结束
		if node == r {
			return
		}
	}

}
