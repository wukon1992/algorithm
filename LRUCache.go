package main
// 缓存的大小有限，当缓存被用满时，哪些数据应该被清理出去，哪些数据应该被保留？这就需要缓存淘汰策略来决定。常见的策略有三种：先进先出策略FIFO（First In，First Out）、最少使用策略LFU（Least Frequently Used）、最近最少使用策略LRU（Least Recently Used）。
// - 向链表头部加入数据
// - 从链表尾部删除数据
// - 最近用到的数据节点，先删除再放到头部节点
// - get put addToHead removeNode removeTail moveToHead
import(
	"fmt"
)
//定义缓存结构
type LRUCache struct {
	size int //当前缓存的大小
	capacity int //缓存容器大小
	cache map[int]*LinkedNode //map缓存 
	head, tail *LinkedNode //虚拟头部尾部节点
}
// 定义个节点
type LinkedNode struct {
	key int
	value string
	prev, next *LinkedNode
}

//初始化一个节点
func initNode(key int,value string) *LinkedNode{
	return &LinkedNode{
		key: key,
		value: value,
	}
}

//初始化一个缓存体
func initCache(capacity int) *LRUCache{
	l := &LRUCache{
		size:0,
		capacity:capacity,
		cache:make(map[int]*LinkedNode),
		head:initNode(0,""),
		tail:initNode(0,""),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}
//获取一个节点
func (this *LRUCache)get(key int) string{
	if _, ok := this.cache[key]; !ok{
		return "不存在"
	}
	node:= this.cache[key]
	//值存在，将该节点放入头结点
	this.moveToHead(node)
	return node.value 
}

//将节点移动到头部
func(this *LRUCache)moveToHead(node *LinkedNode ){
	//移除节点，同时将节点移植到头部
	this.removeNode(node)
	this.addToHead(node)
}

//移除节点
func(this *LRUCache)removeNode(node *LinkedNode){
	node.prev.next = node.next
    node.next.prev = node.prev
}
//从头部添加结点
func(this *LRUCache)addToHead(node *LinkedNode){
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node

	node.prev = this.head
    node.next = this.head.next

    this.head.next.prev = node
    this.head.next = node
}
func (this *LRUCache)put(key int,value string){
	if _, ok := this.cache[key];!ok{
		node := initNode(key,value)
		this.cache[key]=node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity { //容量满了则删除最后尾节点
			removeNode := this.removeTail()
			delete(this.cache,removeNode.key)
			this.size--
		}
	}else{
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func(this *LRUCache)removeTail()*LinkedNode{
	node := this.tail.prev
	this.removeNode(node)
	return node
}

 func main() {
	lruC := initCache(2)
	fmt.Println(lruC.get(1))	
	lruC.put(1,"一")
	lruC.put(2,"贰")
	fmt.Println(lruC.get(1))
	fmt.Println(lruC.get(2))
	lruC.put(3,"叁")
	fmt.Println(lruC.get(1))
}