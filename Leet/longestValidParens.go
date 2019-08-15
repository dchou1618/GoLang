// file longestValidParens.go

package main

import (
  "fmt"
)
// uses idea of tracking current valid string from traversing the
// string forward and backwards
func longestValidParentheses(s string) int {
    var longest, open, closed int;
    for i := range s {
        if (s[i] == '(') {
            open++;
        } else if (s[i] == ')') {
            closed++;
        }
        if (open == closed) {
            if (2*open > longest) {
                longest = 2*open;
            }
        }
        if (closed > open) {
            open = 0;
            closed = 0;
        }
    }
    open, closed = 0,0;
    for j := len(s)-1; j >= 0; j-- {
        if (s[j] == ')') {
            closed++;
        } else if (s[j] == '(') {
            open++;
        }
        if (open == closed) {
            if (2*open > longest) {
                longest = 2*open;
            }
        }
        if (open > closed) {
            open, closed = 0,0;
        }

    }

    return longest;
}
