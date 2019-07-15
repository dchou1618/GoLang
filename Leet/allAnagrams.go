// file allAnagrams.go
package main

import (
  "fmt"
)

// Problem 438: Runtime: 4760 MS :( (5.70%), Memory: 7.9 MB (90.38%)

func equalMaps(m1 map[byte]int, m2 map[byte]int) (bool) {
    for key,value := range m1 {
        if (value != m2[key]) {
            return false;
        }
    }
    return true;
}

func makeMap(s string) (map[byte]int) {
    m := make(map[byte]int);
    for i := 0; i < len(s); i++ {
        m[s[i]] += 1;
    }
    return m;
}

func findAnagrams(s string, p string) []int {
    i := 0;
    anagramIndices := []int{};
    pMap := makeMap(p);
    for (i+len(p) <= len(s)) {
        if (equalMaps(pMap,makeMap(s[i:i+len(p)]))) {
            anagramIndices = append(anagramIndices, i);
        }
        i += 1;
    }
    return anagramIndices;
}


func main() {

}
