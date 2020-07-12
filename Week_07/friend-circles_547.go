type set struct {
    p []int
    count int
}

func Constructor(n int) set {
    this := set{}
    this.p = make([]int, n)

    for i:=0; i < n; i++ {
        this.p[i] = i
    }
    this.count = n
    return this
}

func (this *set) find(i int)int {
    root := i
    for this.p[root] != root {
        root = this.p[root]
    }

    for this.p[i] != root {
        tmp := i
        i = this.p[i]
        this.p[tmp] = root
    }

    return root
}

func (this *set) union(a, b int) {
    la := this.find(a)
    lb := this.find(b)

    if la != lb {
        this.p[lb] = la
        this.count--
    }
}

func findCircleNum(M [][]int) int {
    bset := Constructor(len(M))

    for i := 0; i < len(M); i++ {
        for j := 0; j < len(M); j++ {
            if M[i][j] == 1 {
                bset.union(i, j)
            }
        }
    }

    return bset.count
}
