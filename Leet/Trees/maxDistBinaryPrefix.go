// file maxDistBinaryPrefix.go
// needs WORK
package main

import (
    "fmt"
    "strconv"
)

// intention is to get shortest prefixes
type TrieNode struct {
    prefix string
    postfixs []string
    maxPostSum int
    inDictionary bool
    binaryNodes [2]*TrieNode
}
// https://leetcode.com/discuss/interview-question/350363/Google-or-OA-2018-or-Max-Distance
func constructTrie(binaryStrs []string) (*TrieNode, int) {
    root := &TrieNode{postfixs: []string{}};
    root.binaryNodes[0] = &TrieNode{postfixs: []string{}};
    root.binaryNodes[1] = &TrieNode{postfixs: []string{}};
    var temp *TrieNode = root;
    longestPostSum := 0;
    for _, str := range binaryStrs {
        temp = root;
        depth := 0;
        maxDepth := -1;
        for i, _ := range str {
            val, _ := strconv.Atoi(str[i:i+1]);
            //fmt.Println(val, temp.binaryNodes[val].postfixs);
                // reaching in the loop implies that
            if (temp.binaryNodes[val] != nil) {
                for _, postfix := range temp.binaryNodes[val].postfixs {
                    if (depth > maxDepth) {
                        temp.maxPostSum = len(str[i+1:]) + len(postfix);
                        longestPostSum = temp.maxPostSum;
                        maxDepth = depth;
                    }
                    if (depth == maxDepth) {
                        if (len(str[i+1:]) + len(postfix) > temp.maxPostSum) {
                            temp.maxPostSum = len(str[i+1:]) + len(postfix);
                            if (temp.maxPostSum > longestPostSum) {
                                longestPostSum = temp.maxPostSum;
                            }
                        }
                    }
                }
            }
            temp.binaryNodes[val] = &TrieNode{prefix: str[:i+1],
                             postfixs: append(temp.postfixs, str[i+1:])};
            temp.binaryNodes[val].inDictionary = true;
            fmt.Println(temp.postfixs, longestPostSum);
            temp = temp.binaryNodes[val];
            depth++;
        }
    }
    return root, longestPostSum;
}

func maxDistPrefix(binaryStrs []string) int {
    _, longestPostSum := constructTrie(binaryStrs);
    return longestPostSum;
}

func main() {
    binaryStrs := []string{"1011000", "1011110"};
    fmt.Println(maxDistPrefix(binaryStrs));
}
