// file subsequence.go

package main


import (
  "fmt"
)

func isSubsequence(s string, t string) bool {
    if (s == "") {
        return true;
    } else if (t == "") {
        return false;
    } else {
        if (s[0] == t[0]) {
            return isSubsequence(s[1:], t[1:]);
        } else {
            return isSubsequence(s, t[1:]);
        }
    }
}


// finding the k most frequent subsequences
func createMinHeap() []int {

}
