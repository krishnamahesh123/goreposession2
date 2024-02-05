package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Data       int
	KeyPointer *list.Element
}
type LRUcache struct {
	Queue    *list.List
	Items    map[int]*Node
	Capacity int
}

func main() {
	lru := LRUcache{Queue: list.New(), Items: make(map[int]*Node), Capacity: 2}
	fmt.Println(lru)
	lru.put(10, 40)
	lru.put(20, 60)
	lru.put(30, 70)
	fmt.Println(lru.Get(10))
	fmt.Println(lru.Get(20))
}
func (l LRUcache) put(key int, value int) {
	if item, ok := l.Items[key]; !ok {
		if l.Capacity == len(l.Items) {
			back := l.Queue.Back()
			l.Queue.Remove(back)
			delete(l.Items, back.Value.(int))
		}
		l.Items[key] = &Node{Data: value, KeyPointer: l.Queue.PushFront(key)}
	} else {
		item.Data = value
		l.Items[key] = item
		l.Queue.MoveToFront(item.KeyPointer)
	}
	fmt.Println(l)
}
func (l *LRUcache) Get(key int) int {
	if item, ok := l.Items[key]; ok {
		l.Queue.MoveToFront(item.KeyPointer)
		return item.Data
	}
	return -1
}
