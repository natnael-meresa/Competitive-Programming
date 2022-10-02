package main

import "fmt"

// slow as f**k they say

// type LFUCache struct {
// 	Cap      int
// 	HashMap  map[int]Data
// 	LastTime int
// }

// type Data struct {
// 	Val   int
// 	Cnt   int
// 	Wtime int
// }

// func Constructor(capacity int) LFUCache {
// 	return LFUCache{
// 		Cap:      capacity,
// 		HashMap:  make(map[int]Data),
// 		LastTime: 0,
// 	}
// }

// func (this *LFUCache) Get(key int) int {
// 	val, ok := this.HashMap[key]
// 	if !ok {
// 		return -1
// 	}
// 	this.HashMap[key] = Data{
// 		Val:   this.HashMap[key].Val,
// 		Cnt:   this.HashMap[key].Cnt + 1,
// 		Wtime: this.LastTime,
// 	}
// 	this.LastTime++

// 	return val.Val
// }

// func (this *LFUCache) Put(key int, value int) {
// 	if this.Cap == 0 {
// 		return
// 	}
// 	_, ok := this.HashMap[key]
// 	if ok {
// 		this.HashMap[key] = Data{
// 			Val:   value,
// 			Cnt:   this.HashMap[key].Cnt + 1,
// 			Wtime: this.LastTime,
// 		}
// 		this.LastTime++
// 	} else {
// 		if len(this.HashMap) >= this.Cap {
// 			lsdKey := -1
// 			leastCnt := 1
// 			for k, v := range this.HashMap {
// 				if lsdKey == -1 {
// 					lsdKey = k
// 					leastCnt = v.Cnt
// 					continue
// 				}
// 				if v.Cnt <= leastCnt {
// 					if v.Cnt == leastCnt && lsdKey != -1 {
// 						if v.Wtime < this.HashMap[lsdKey].Wtime {
// 							lsdKey = k
// 							leastCnt = v.Cnt
// 						}
// 					} else {
// 						lsdKey = k
// 						leastCnt = v.Cnt
// 					}
// 				}
// 			}
// 			fmt.Println("lsdkey", lsdKey)
// 			delete(this.HashMap, lsdKey)
// 		}
// 		this.HashMap[key] = Data{
// 			Val:   value,
// 			Cnt:   1,
// 			Wtime: this.LastTime,
// 		}
// 		this.LastTime++
// 	}

// }

type LFUCache struct {
	CacheList map[int]*Node
	FrqList   map[int]*FreqList
	minFreq   int
	cap       int
}

type FreqList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	Val  int
	Key  int
	freq int
	next *Node
	prev *Node
}

func (f *FreqList) popNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	f.size--
}

func (f *FreqList) AddNode(node *Node) {
	node.prev = f.tail.prev
	node.next = f.tail
	f.tail.prev.next = node
	f.tail.prev = node
	f.size++
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:       capacity,
		minFreq:   0,
		CacheList: make(map[int]*Node),
		FrqList:   make(map[int]*FreqList),
	}
}

func (this *LFUCache) Get(key int) int {
	currNode, ok := this.CacheList[key]
	if !ok {
		return -1
	}
	this.Update(currNode)

	return currNode.Val
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}

	currNode, ok := this.CacheList[key]
	if ok {
		currNode.Val = value
		this.Update(currNode)
		return
	}
	if len(this.CacheList) >= this.cap {
		leastList := this.FrqList[this.minFreq]
		leastNode := leastList.head.next
		leastList.popNode(leastNode)
		delete(this.CacheList, leastNode.Key)
	}
	this.minFreq = 1
	newNode := &Node{
		Key:  key,
		Val:  value,
		freq: 1,
	}
	this.CacheList[key] = newNode

	freqList, ok := this.FrqList[1]
	if !ok {
		head := &Node{}
		tail := &Node{}
		head.next = tail
		tail.prev = head
		freqList = &FreqList{
			size: 0,
			head: head,
			tail: tail,
		}
	}
	freqList.AddNode(newNode)
	this.FrqList[1] = freqList
}

func (this *LFUCache) Update(currentNode *Node) {
	freq := currentNode.freq
	this.FrqList[freq].popNode(currentNode)

	currentFrq, ok := this.FrqList[freq+1]
	if !ok {
		head := &Node{}
		tail := &Node{}
		head.next = tail
		tail.prev = head
		currentFrq = &FreqList{
			head: head,
			tail: tail,
		}
	}
	currentNode.freq += 1
	currentFrq.AddNode(currentNode)
	this.FrqList[freq+1] = currentFrq

	if this.FrqList[freq].size == 0 && this.minFreq == currentNode.freq-1 {
		this.minFreq++
	}

}
func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	param_1 := obj.Get(1)
	fmt.Println(">>1", param_1)

	obj.Put(3, 3)
	param_2 := obj.Get(2)
	fmt.Println(">>2", param_2)
	param_3 := obj.Get(3)
	fmt.Println(">>3", param_3)

	// obj.Put(4, 4)
	// param_1 = obj.Get(1)
	// fmt.Println(">>1", param_1)

	// param_3 = obj.Get(3)
	// fmt.Println(">>3", param_3)

	// param_4 := obj.Get(4)
	// fmt.Println(">>4", param_4)

}
