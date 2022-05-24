package main

import (
	"fmt"
	"sync"
)

// Node 通过链表存储二叉树节点信息
type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

// 队列实现
// 链表结点
type LinkNode struct {
	Data *Node
	Next *LinkNode
}

// 链表队列
type LinkQueue struct {
	Top  *LinkNode
	Size int
	Lock sync.Mutex
}

// 入队
func (q *LinkQueue) Add(t *Node) {
	q.Lock.Lock()
	defer q.Lock.Unlock()

	if q.Top == nil {
		q.Top = new(LinkNode)
		q.Top.Data = t
	} else {
		temp := new(LinkNode)
		temp.Data = t
		TNode := q.Top
		for TNode.Next != nil {
			TNode = TNode.Next
		}
		TNode.Next = temp
	}
	q.Size++
}

// 出队
func (q *LinkQueue) Remove() *Node {
	q.Lock.Lock()
	defer q.Lock.Unlock()

	if q.Size == 0 {
		panic("queue is empty!")
	}

	t := q.Top.Data
	q.Top = q.Top.Next
	q.Size--
	return t
}

// 是否为空
func (q *LinkQueue) IsEmpty() bool {
	return q.Size == 0
}

func (node *Node) GetData() string {
	return fmt.Sprintf("%v", node.Data)
}

// 前序遍历
func preOrderTraverse(treeNode *Node) {
	// 节点为空则退出当前递归
	if treeNode == nil {
		return
	}
	// 否则先打印当前节点值
	fmt.Printf("%s ", treeNode.GetData())
	// 然后对左子树和右子树做前序遍历
	preOrderTraverse(treeNode.Left)
	preOrderTraverse(treeNode.Right)
}

// 中序遍历
func midOrderTraverse(treeNode *Node) {
	// 节点为空则退出当前递归
	if treeNode == nil {
		return
	}
	// 否则先从左子树最左侧节点开始遍历
	midOrderTraverse(treeNode.Left)
	// 打印位于中间的根节点
	fmt.Printf("%s ", treeNode.GetData())
	// 最后按照和左子树一样的逻辑遍历右子树
	midOrderTraverse(treeNode.Right)
}

// 后序遍历
func postOrderTraverse(treeNode *Node) {
	// 节点为空则退出当前递归
	if treeNode == nil {
		return
	}
	// 否则先从左子树最左侧节点开始遍历
	postOrderTraverse(treeNode.Left)
	// 最后按照和左子树一样的逻辑遍历右子树
	postOrderTraverse(treeNode.Right)
	// 打印位于中间的根节点
	fmt.Printf("%s ", treeNode.GetData())
}

// 层级遍历
func levelOrderTraverse(treeNode *Node) {
	if treeNode == nil {
		return
	}

	q := new(LinkQueue)
	q.Add(treeNode)
	for !q.IsEmpty() {
		e := q.Remove()
		fmt.Print(e.Data, " ")

		if e.Left != nil {
			q.Add(e.Left)
		}
		if e.Right != nil {
			q.Add(e.Right)
		}
	}

}

/*
         0
		/ \
	   1   2
	  / \ / \
      3 4  5 6
*/
// 测试代码
func main() {
	// 初始化一个简单的二叉树
	node0 := NewNode(0) // 根节点
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3) // 根节点
	node4 := NewNode(4)
	node5 := NewNode(5)
	node6 := NewNode(6)
	node0.Left = node1
	node0.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	node2.Right = node6

	// 前序遍历这个二叉树
	fmt.Print("前序遍历: ")
	preOrderTraverse(node0)
	fmt.Println()
	// 中序遍历这个二叉树
	fmt.Print("中序遍历: ")
	midOrderTraverse(node0)
	fmt.Println()
	fmt.Print("后序遍历: ")
	postOrderTraverse(node0)
	fmt.Println()
	fmt.Print("层次遍历: ")
	levelOrderTraverse(node0)
	fmt.Println()
}
