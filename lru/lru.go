package lru

/*
 capcacity = 2
 lrucache.put(1, 1)
 lrucache.put(2, 2)
 lrucache.get(1) => return 1
 lrucache.put(3, 3) 
 lrucache.get(2) => return -1

get put 时间复杂度为O(1)
*/

import "fmt"

type LruNode struct {
	Key      int
	Value    int
	PreNode  *LruNode
	NextNode *LruNode
}

type LruCache struct {
	head      *LruNode
	tail      *LruNode
	LruMap    map[int]*LruNode
	used      int
	capcacity int 
}

func (this *LruCache) removeNode(node *LruNode) {
	node.PreNode.NextNode = node.NextNode
	node.NextNode.PreNode = node.PreNode

	delete(this.LruMap, node.Key)
	this.used--
}

func (this *LruCache) insetToHead(node *LruNode) {
	this.head.NextNode.PreNode = node
	node.NextNode = this.head.NextNode
	this.head.NextNode = node
	node.PreNode = this.head

	this.LruMap[node.Key] = node
	this.used++
}

func (this *LruCache) Get(k int)  int{
	node, ok := this.LruMap[k]
	if !ok {
		return -1
	} else {
		// 摘除
		this.removeNode(node)

		// 加入头节点
		this.insetToHead(node)
	}

	return node.Value
}

func (this *LruCache) Put(k int, v int)  {
	node := &LruNode{Key: k, Value: v}

	n, ok := this.LruMap[k]
	if ok {
		// 更新 lrumap， 插入头结点
		this.removeNode(n)
		this.insetToHead(node)
	} else {
		if this.used < this.capcacity {
			// 直接插入头结点
			this.insetToHead(node)
		} else {
			// 先删除尾巴节点，然后插入头结点
			this.removeNode(this.tail.PreNode)
			this.insetToHead(node)
	
		}
	}
}

func (this *LruCache) GetCacheDetail() {
	fmt.Println("=========================")

	for k, v := range(this.LruMap) {
		fmt.Println("The key/value is", k, v.Value)
	}

	fmt.Println(this.used)
	n := this.head.NextNode

	for n.NextNode != nil {
		fmt.Println(n)
		n = n.NextNode
	}
	fmt.Println("=========================")

}

func StartTest() {
	head := &LruNode{}
	tail := &LruNode{}
	head.PreNode = nil
	head.NextNode = tail
	tail.PreNode = head
	tail.NextNode = nil

	lrucacahe := &LruCache{head: head, tail: tail, LruMap: make(map[int]*LruNode), used: 0, capcacity: 3}

	lrucacahe.Put(1, 1)
	lrucacahe.Put(2, 2)
	lrucacahe.Put(3, 3)
	lrucacahe.GetCacheDetail()

	res := lrucacahe.Get(2)
	fmt.Println(res)
	lrucacahe.GetCacheDetail()

	lrucacahe.Put(4, 4)
	lrucacahe.GetCacheDetail()

	res = lrucacahe.Get(1)
	fmt.Println(res)
	lrucacahe.GetCacheDetail()

	lrucacahe.Put(4, 5)
	lrucacahe.GetCacheDetail()
}
