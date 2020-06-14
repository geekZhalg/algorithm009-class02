func isEndWordExit(endWord string, wordList[]string) bool {
    for _,str := range wordList {
        if endWord == str {
            return true
        }
    }

    return false
}

func ladderLength(beginWord string, endWord string, wordList []string) int {

    if isEndWordExit(endWord, wordList) != true {
        return 0
    }

    queue := []string{}
    othQueue := []string{}
    queue = append(queue, beginWord)
    othQueue = append(othQueue, endWord)
    step := 1
    vist := map[string]bool{beginWord:true}
    othVist := map[string]bool{endWord:true}
    /* build map */
    wordbank := map[string][]string{}
    for _,word := range wordList{
        for j := 0; j < len(word); j++ {
            key := word[0:j] + "*" + word[j+1:]
            wordbank[key] = append(wordbank[key], word)
        }
    }

    for len(queue) > 0 {
        size := len(queue)
        step++

        //fmt.Println(queue)
        for i:= 0; i < size;i++ {
            word := queue[0]
            queue = queue[1:]

            for j:=0; j < len(word); j++ {
                key:=word[0:j] + "*" + word[j+1:]
                //fmt.Println(key)
                if strlist,ok := wordbank[key];ok == true {
                    for _,str := range strlist {
                        if othVist[str] == true {
                            return step
                        }

                        if vist[str] != true {
                            vist[str] = true
                            queue = append(queue, str)
                        }
                    }
                }
            }
        }

        if len(queue) > len(othQueue) {
            queue, othQueue = othQueue, queue
            vist, othVist = othVist, vist
        }
    }

    return 0
}
