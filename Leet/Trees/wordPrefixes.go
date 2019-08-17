// file wordPrefixes.go
// implements storage of letters and isWord booleans to determine
// if a prefix is in a dictionary.

package main

import (
    "fmt"
)
// using a trie - prefix tree
// contains 26 entries for children Trie nodes of all letters from alphabet
type Trie struct {
    letters string
    children [26]*Trie
    isWord bool
}

func constructTrie(words []string) *Trie {
    var root *Trie = &Trie{letters: "", isWord: true};
    var currNode *Trie = root;
    for _, word := range words {
        currNode = root;
        for i := 0; i < len(word); i++ {
            if (currNode.children[int(word[i])-'a'] == nil) {
                currNode.children[int(word[i])-'a'] = &Trie{};
            }
            currNode = currNode.children[int(word[i])-'a'];
        }
        currNode.isWord = true;
        currNode.letters = word;
    }
    return root;
}

func findLongest(root *Trie, seen map[string]bool, longest *string) {
    if (root != nil) {
        if (root.isWord) {
            for i := 0; i < 26; i++ {
                if (seen[root.letters+string('a'+i)]) {
                    if (len(root.letters+string('a'+i)) > len(*longest)) {
                        *longest = root.letters+string('a'+i);
                    } else if (len(root.letters+string('a'+i)) == len(*longest)) {
                        if (root.letters+string('a'+i) < *longest) {
                            *longest = root.letters+string('a'+i);
                        }
                    }
                    findLongest(root.children[i], seen, longest)
                }
            }
        }
    }
}

func longestWord(words []string) string {
    longest := "";
    seen := make(map[string]bool);
    for _, val := range words {
        seen[val] = true;
    }
    var root *Trie = constructTrie(words);
    findLongest(root, seen, &longest);
    return longest;
}
