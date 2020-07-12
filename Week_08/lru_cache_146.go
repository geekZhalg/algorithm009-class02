type LRUCache struct {
    maps map[int]*Dlist
    head *Dlist
    tail *Dlist
    cap int
    size int
}

type Dlist struct {
    key int
    val int
    prev *Dlist
    next *Dlist
}

func (this *LRUCache)DlistNodeAddHead(node *Dlist) {
    node.next = this.head.next
    node.prev = this.head
    this.head.next.prev = node
    this.head.next = node
}

func (this *LRUCache)DlistRmNode(node *Dlist) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func Constructor(capacity int) LRUCache {
    cache := LRUCache {
        maps:make(map[int]*Dlist),
        head:&Dlist{},
        tail:&Dlist{},
        cap:capacity,
        size:0,
    }

    cache.head.next = cache.tail
    cache.head.prev = cache.tail
    cache.tail.next = cache.head
    cache.tail.prev = cache.head

    return cache
}


func (this *LRUCache) Get(key int) int {
    if _,ok := this.maps[key]; !ok {
        return -1
    }

    cur := this.maps[key]
    this.DlistRmNode(cur)
    this.DlistNodeAddHead(cur)

    return cur.val
}   


func (this *LRUCache) Put(key int, value int)  {
    if _,ok := this.maps[key]; !ok {
        //fmt.Println(this.size, this.cap)
        node := &Dlist{key:key, val:value}
        this.DlistNodeAddHead(node)
        this.maps[key] = node
        this.size++
        
        if this.size > this.cap {
            last := this.tail.prev
            this.DlistRmNode(last)
            delete(this.maps, last.key)
            this.size--
        }
    } else {
        node := this.maps[key]
        node.val = value
        this.DlistRmNode(node)
        this.DlistNodeAddHead(node)
    }
}
