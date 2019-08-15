// file lexicographicPermutations.go
// Problem 24: Project Euler

package main

import (
    "fmt"
)


func lexPermutations(options []byte, count *int, lower int, result string) {
    if (len(options) == 0) {
        (*count)++;
        if (*count == 1000000) {
            fmt.Println(result);
            return;
        }
    } else {
        for i := lower; i < len(options); i++ {
            var val byte = options[i];
            options = append(options[:i], options[i+1:]...);
            lexPermutations(options, count, lower, result+string(val));
            options = append(options[:i], append([]byte{val}, options[i:]...)...);
        }
    }
}


func main() {
    count := 0;
    lexPermutations([]byte{'0','1','2','3','4','5','6','7','8','9'}, &count, 0, "");
}
