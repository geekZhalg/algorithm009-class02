type node struct {
    key int
    cnt int
}

type nheap []node

func (h nheap) Len() int {
    return len(h)
}

func (h nheap) Less(i, j int) bool {
    return h[i].cnt < h[j].cnt
}

func (h nheap) Swap(i,j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *nheap) Push(v interface{}) {
    *h = append(*h, v.(node))
}

func (h *nheap) Pop()interface{} {
    n:= len(*h)
    x := (*h)[n - 1]
    (*h) = (*h)[:n-1]
    return x
}

func topKFrequent(nums []int, k int) []int {
    m := map[int]int{}
    for _,n := range nums{
        m[n]++
    }

    //for key,v := range m {
    //    fmt.Println(key,v)
    //}
    hp := &nheap{}
    heap.Init(hp)
    i := 0
    for key,v := range m {
        if i < k {
            heap.Push(hp, node{key, v})
            i++
        } else {
            if v  > (*hp)[0].cnt {
                heap.Pop(hp)
                heap.Push(hp, node{key, v})
            }
        }
        
    } 

    res := []int{}
    for i:=0; i < k;i++ {
        res = append(res, (*hp)[0].key)
        heap.Pop(hp)
    }

    return res
}
