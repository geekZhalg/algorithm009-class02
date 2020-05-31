import (
    "fmt"
    "sort"
)
type bytes []byte
func (s bytes) Len() int {return len(s)}
func (s bytes) Swap(i, j int) {s[i], s[j] = s[j], s[i]} 
func (s bytes) Less(i, j int) bool {return s[i] < s[j]}
func groupAnagrams(strs []string) [][]string {
    group := make(map[string][]string)
    var result [][]string

    for _,s := range strs {
        chars := bytes(s)
        sort.Sort(chars)
        str := string(chars)
        group[str] = append(group[str], s)
    }

    for _,list := range group {
        result = append(result, list)
    }

    return result
}
