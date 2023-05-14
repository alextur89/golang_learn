package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	reverse  map[*ListItem]Key
}

func (p *lruCache) Set(key Key, value interface{}) bool {
	i, has := p.items[key]
	if !has {
		newItem := p.queue.PushFront(value)
		p.items[key] = newItem
		p.reverse[newItem] = key

		if p.queue.Len() > p.capacity {
			back := p.queue.Back()
			delete(p.items, p.reverse[back])
			delete(p.reverse, back)
			p.queue.Remove(back)
		}
	} else {
		i.Value = value
		p.queue.MoveToFront(i)
	}
	return has
}

func (p *lruCache) Get(key Key) (interface{}, bool) {
	i, has := p.items[key]
	if has {
		p.queue.MoveToFront(i)
		return i.Value, true
	}
	return nil, false
}

func (p *lruCache) Clear() {
	for k, v := range p.items {
		delete(p.items, k)
		delete(p.reverse, v)
		p.queue.Remove(v)
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		reverse:  make(map[*ListItem]Key, capacity),
	}
}
