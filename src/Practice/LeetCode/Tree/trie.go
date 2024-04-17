type Trie struct {
    children map[string] *Trie
    isLast bool
}


func Constructor() Trie {
    trie := new(Trie)
    trie.children = make(map[string]*Trie)
    return *trie
}


func (this *Trie) Insert(word string)  {
    trie := this
    for pos := 0; pos < len(word); pos++ {
        ch := string(word[pos])
        if trie.children[ch] == nil {
            newTrie := new(Trie) // new branch
            newTrie.children = make(map[string]*Trie)
            trie.children[ch] = newTrie
            trie = newTrie
        } else {
            trie = trie.children[ch]
        }
    }
    trie.isLast = true
}


func (this *Trie) Search(word string) bool {
    trie := this
    for pos := 0; pos < len(word); pos++ {
        ch := string(word[pos])
        if trie.children[ch] == nil {
            return false
        }
        trie = trie.children[ch]
    }
    if trie == nil {
        return false
    }
    return trie.isLast == true
}


func (this *Trie) StartsWith(prefix string) bool {
    trie := this
    for pos := 0; pos < len(prefix); pos++ {
        ch := string(prefix[pos])
        if trie.children[ch] == nil {
            return false
        }
        trie = trie.children[ch]
    }
    return trie != nil
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
