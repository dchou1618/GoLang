// go file generate parentheses
package main

import (
  "fmt"
)

type Stack struct {
    vals []string
}

func (s Stack) push(str string) (Stack) {
    s.vals = append(s.vals,str);
    return s;
}
func (s Stack) pop() (Stack) {
    s.vals = s.vals[:len(s.vals)-1];
    return s;
}
func (s Stack) isEmpty() (bool) {
    return len(s.vals) == 0;
}
func (s Stack) top() (string) {
    if (!s.isEmpty()) {
        return s.vals[len(s.vals)-1];
    }
    return "";
}

func allParenthesesCombos(n int, nums *[]string, currParen string, currStack Stack) {
    if (len(currParen) == 2*n && currStack.isEmpty()) {
        *nums = append(*nums,currParen);
    } else if (len(currParen) > 2*n) {
        return;
    } else {
        if (currStack.isEmpty()) {
            allParenthesesCombos(n, nums, currParen+"(", currStack.push("("))
        } else {
            allParenthesesCombos(n, nums, currParen+")", currStack.pop());
            allParenthesesCombos(n, nums, currParen+"(", currStack.push("("));
        }
    }
}

func generateParenthesis(n int) []string {
    parens := []string{};
    currParen := "";
    currStack := Stack{[]string{}};
    allParenthesesCombos(n, &parens, currParen, currStack);
    return parens;
}

func main() {

}
