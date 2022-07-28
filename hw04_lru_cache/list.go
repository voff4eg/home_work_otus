package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

func (li *ListItem) Remove() {
	li.Value = nil
	li.Next = nil
	li.Prev = nil
}

type list struct {
	len  int
	head *ListItem
	last *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	v1 := &ListItem{
		Value: v,
	}

	if l.head == nil {
		initHeadAndLast(v1, l)
	} else {
		head := l.head
		l.head = v1
		l.head.Next = head
		head.Prev = l.head
	}
	l.len++

	return l.head
}

func (l *list) PushBack(v interface{}) *ListItem {
	v1 := &ListItem{
		Value: v,
	}

	if l.last == nil {
		initHeadAndLast(v1, l)
	} else {
		last := l.last
		v1.Prev = last
		last.Next = v1
		l.last = v1
	}
	l.len++

	return l.last
}

func initHeadAndLast(i *ListItem, l *list) {
	head := i
	last := i
	head.Next = last
	last.Prev = head
	l.head = head
	l.last = last
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	if l.head == i {
		l.head = i.Next
	}

	if l.last == i {
		l.last = i.Prev
	}

	if l.len != 0 {
		l.len--
	}
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	head := l.head
	l.head = i
	l.head.Next = head
	head.Prev = l.head
	l.len++
}

func NewList() List {
	return new(list)
}
