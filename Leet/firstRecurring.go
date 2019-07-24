// file firstRecurring.go

package main

import (
    "fmt"
)

func firstRecurring(s string) (string){
    // queue
    stack := []string{};
    for i := 0; i < len(s); i++ {
        if (len(stack) == 0) {
            stack = append(stack,string(s[i]));
        } else {
            if (stack[0] == string(s[i])) {
                return string(s[i]);
            } else {
                stack = append(stack,string(s[i]));
            }
        }
    }
    return "";
}

func main() {
    s := "BCABA";
    fmt.Println(firstRecurring(s));
}
