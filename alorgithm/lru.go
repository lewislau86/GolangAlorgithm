package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	list     *list.List
}

func (this *LRUCache) Get(item any) any {
	if elem, ok := this.Find(item); ok {
		this.list.MoveToFront(elem)
		return elem.Value
	}
	return nil
}
func (this *LRUCache) Find(item any) (*list.Element, bool) {
	for elem := this.list.Front(); elem != nil; elem = elem.Next() {
		if elem.Value == item {
			return elem, true
		}
	}
	return nil, false
}

func (this *LRUCache) Show() {
	for e := this.list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func (this *LRUCache) Put(item any) {
	if elem, ok := this.Find(item); ok {
		this.list.MoveToFront(elem)
		elem.Value = item
	} else {
		if this.list.Len() >= this.capacity {
			//delete(this.cache, this.list.Back().Value.(Pair).key)
			this.list.Remove(this.list.Back())
		}
		this.list.PushFront(item)
		//this.cache[item] = this.list.Front()
	}
}

func main() {
	var cache = LRUCache{
		capacity: 3,
		list:     list.New(),
	}
	cache.Put(1)
	cache.Put("aaa")
	cache.Put(3)
	cache.Show()
	fmt.Println("*******")
	cache.Put("bbb")
	cache.Show()
	/*
		cache.Put(2, 2)
		fmt.Println(cache.Get(2))
		cache.Put(3, 3)
		fmt.Println(cache.Get(1))
		cache.Put(4, 4)
		cache.Put(1, 1)
		fmt.Println(cache.Get(1))
		fmt.Println(cache.Get(2))
	*/
}
