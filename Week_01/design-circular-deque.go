// @Title  Week_01
// @Description  双端队列
// @Author  alan  2020/9/12 16:13
// @Update  alan  2020/9/12 16:13
package Week_01

type node struct {
	Val  int
	Prev *node
	Next *node
}
type MyCircularDeque struct {
	head     *node // 首部节点
	last     *node // 尾部节点
	capacity int   // 双向队列容量
	length   int   // 当前队列长度
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	// 初始化head、last节点
	head := &node{
		Val: -1,
	}
	last := &node{
		Val: -1,
	}
	// 初始化链表节点指向
	head.Next = last
	last.Prev = head
	return MyCircularDeque{
		head:     head,
		last:     last,
		capacity: k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	// 如果队列已满则返回false
	if this.IsFull() {
		return false
	}

	// 获取head节点的next节点
	headNextNode := this.head.Next
	insertNode := &node{
		Prev: this.head,
		Next: headNextNode,
		Val:  value,
	}
	// 调整插入节点的首尾指向
	this.head.Next = insertNode
	headNextNode.Prev = insertNode
	this.length++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	// 如果队列已满则返回false
	if this.IsFull() {
		return false
	}

	lastPrevNode := this.last.Prev
	insertNode := &node{
		Prev: lastPrevNode,
		Next: this.last,
		Val:  value,
	}
	this.last.Prev = insertNode
	lastPrevNode.Next = insertNode
	this.length++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	deleteNode := this.head.Next     // 获取要删除的node
	this.head.Next = deleteNode.Next // 调整要删除node的父子指针的前后指向
	deleteNode.Next.Prev = this.head
	this.length--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	deleteNode := this.last.Prev     // 获取要删除的node
	this.last.Prev = deleteNode.Prev // 调整要删除node的父子指针的前后指向
	deleteNode.Prev.Next = this.last
	this.length--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	return this.head.Next.Val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	return this.last.Prev.Val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return  this.length == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.length == this.capacity
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
