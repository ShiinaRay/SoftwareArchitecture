package linktable

// Node 定义双向链表节点
type Node[T any] struct {
    Value    T
    Prev, Next *Node[T]
}

// LinkTable 管理头尾和长度
type LinkTable[T any] struct {
    head, tail *Node[T]
    Len        int
}

// New 创建空表
func New[T any]() *LinkTable[T] {
    return &LinkTable[T]{}
}

// PushBack 在尾部插入
func (lt *LinkTable[T]) PushBack(val T) {
    node := &Node[T]{Value: val}
    if lt.tail == nil {
        lt.head = node
        lt.tail = node
    } else {
        lt.tail.Next = node
        node.Prev = lt.tail
        lt.tail = node
    }
    lt.Len++
}

// PopFront 从头部弹出
func (lt *LinkTable[T]) PopFront() (T, bool) {
    if lt.head == nil {
        var zero T
        return zero, false
    }
    node := lt.head
    lt.head = node.Next
    if lt.head != nil {
        lt.head.Prev = nil
    } else {
        lt.tail = nil
    }
    lt.Len--
    return node.Value, true
}

// Iter 遍历所有元素
func (lt *LinkTable[T]) Iter(f func(val T)) {
    for cur := lt.head; cur != nil; cur = cur.Next {
        f(cur.Value)
    }
}
