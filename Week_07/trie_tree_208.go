type Trie struct {
    isEnd bool
    next [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    for i := 0; i < len(word); i++ {
        if this.next[word[i] - 'a'] == nil {
            this.next[word[i] - 'a'] = new(Trie)
        }

        this = this.next[word[i] - 'a']
    }

    this.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    for i := 0; i < len(word); i++ {
        if this.next[word[i] - 'a'] == nil {
            return false
        }

        this = this.next[word[i] - 'a']
    }

    return this.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
     for i := 0; i < len(prefix); i++ {
        if this.next[prefix[i] - 'a'] == nil {
            return false
        }

        this = this.next[prefix[i] - 'a']
    }

    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
