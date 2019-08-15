// file BFS wordLadder.go

package main

import (
  "fmt"
)


func oneDiff(s1 string, s2 string) (bool) {
    if (len(s1) != len(s2)) {
        return false;
    }
    count := 0;
    for i := 0; i < len(s1); i++ {
        if (s1[i] != s2[i]) {
            if (count == 1) {
                return false;
            }
            count++;
        }
    }
    return count == 1;
}

func neighboringWords(beginWord string, endWord string, wordList []string, inLst *bool) map[string][]string {
    neighbors := make(map[string][]string);
    currLst := append(wordList, append([]string{beginWord}, endWord)...);
    for i := range currLst {
        for j := range currLst {
            if (oneDiff(currLst[i], currLst[j])) {
                neighbors[currLst[i]] = append(neighbors[currLst[i]], currLst[j]);
            }
        }
    }
    for i := range wordList {
        if (wordList[i] == endWord) {
            *inLst = true;
        }
    }
    return neighbors;
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
    inLst := false;
    wordNeighbors := neighboringWords(beginWord, endWord, wordList, &inLst);
    if (!inLst) {
        return 0;
    }
    queue := [][]string{[]string{beginWord, "1"}};
    seen := make(map[string]bool);
    for (len(queue) != 0) {
        top := queue[0];
        queue = queue[1:];
        for _, oneDiffWord := range wordNeighbors[top[0]] {
            if (oneDiffWord == endWord) {
                val, _ := strconv.Atoi(top[1]);
                wordNeighbors = nil;
                queue = nil;
                seen = nil;
                return val+1;
            }
            if (!seen[oneDiffWord] && oneDiff(top[0], oneDiffWord)) {
                num, _ := strconv.Atoi(top[1]);
                str := strconv.Itoa(num+1);
                queue = append(queue, []string{oneDiffWord, str});
                seen[top[0]] = true;
            }
        }
    }
    wordNeighbors = nil;
    queue = nil;
    seen = nil;
    return 0;
}

func main() {
  
}
