package api

type Node struct {
	Data interface{} //节点存储的数据
	Next *Node       //指向下一个节点的指针
}

type LinkedList struct {
	Head *Node //指向第一个节点的指针
}

// 添加一个节点
func (l *LinkedList) AddNode(data interface{}) {
	newNode := &Node{Data: data}

	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

// 遍历链表并执行函数
func (l *LinkedList) Traverse(fn func(interface{})) {
	current := l.Head
	for current != nil {
		fn(current.Data)
		current = current.Next
	}
}

// 删除链表中的节点
func (l *LinkedList) RemoveNode(target interface{}) {
	if l.Head == nil {
		return
	}

	if l.Head.Data == target {
		l.Head = l.Head.Next
		return
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Data == target {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}
