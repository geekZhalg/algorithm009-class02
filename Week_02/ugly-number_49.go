type ugHeap []int
func (h ugHeap) Len() int{
    return len(h)
}

func (h ugHeap) Less (i,j int) bool{
    return h[i] < h[j]
}

func (h ugHeap) Swap(i,j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *ugHeap) Push(n interface{}) {
    (*h) = append((*h), n.(int))
}

func (h* ugHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n - 1]
    (*h) = old[:n-1]

    return x
}

func nthUglyNumber(n int) int {
    hp := &ugHeap{2,3,5}
    base := []int{2,3,5}
    m := map[int]int{2:1,3:1,5:1}
    heap.Init(hp)

    var num = 1
    for i:= 1; i < n;i++{
        num = heap.Pop(hp).(int)
        for _,b := range base {
            k := num * b
            if _,ok := m[k]; !ok {
                m[k] = 1
                heap.Push(hp, k)
            }
        }
    }

    return num
}
