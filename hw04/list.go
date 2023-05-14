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

type list struct {
	size  int
	front *ListItem
	back  *ListItem
}

func (p *list) Back() *ListItem {
	return p.back
}

func (p *list) Front() *ListItem {
	return p.front
}

func (p *list) Len() int {
	return p.size
}

func (p *list) PushFront(v interface{}) *ListItem {
	newElem := &ListItem{Value: v, Next: p.front}
	if p.size > 0 {
		p.front.Prev = newElem
	} else {
		p.back = newElem
	}
	p.front = newElem
	p.size++
	return newElem
}

func (p *list) PushBack(v interface{}) *ListItem {
	newElem := &ListItem{Value: v, Prev: p.back}
	if p.size > 0 {
		p.back.Next = newElem
	} else {
		p.front = newElem
	}
	p.back = newElem
	p.size++
	return newElem
}

func (p *list) cut(i *ListItem) {
	if i != p.front && i != p.back {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	} else if i == p.back {
		i.Prev.Next = nil
		p.back = i.Prev
		if p.size == 1 {
			p.front = p.back
		}
	} else if i == p.front {
		i.Next.Prev = nil
		p.front = i.Next
		if p.size == 1 {
			p.back = p.front
		}
	}
}

func (p *list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	p.size--
	if p.size == 0 {
		p.front = nil
		p.back = nil
		return
	}

	p.cut(i)
}

func (p *list) MoveToFront(i *ListItem) {
	p.cut(i)
	i.Prev = nil
	if p.size > 0 {
		p.front.Prev = i
		i.Next = p.front
	} else {
		p.back = i
	}
	p.front = i
}

func NewList() List {
	return new(list)
}
