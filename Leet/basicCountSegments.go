// file basicCountSegments.go is practice with basic algorithms

package main

import (
  "fmt"
  "unicode"
)

func countSegments(s string) int {
    var segments int;
    var newWord bool = false;
    for i := range s {
        if (!newWord && !unicode.IsSpace(rune(s[i]))) {
            newWord = true;
            segments++;
        }
        if (newWord && unicode.IsSpace(rune(s[i]))) {
            newWord = false;
        }
    }
    return segments;
}
