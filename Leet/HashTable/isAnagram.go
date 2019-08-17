// file isAnagram.go

package main

import (
  "fmt"
)

func equalMaps(m1 map[byte]int, m2 map[byte]int) (bool) {
    for key, value := range m1 {
        if (value != m2[key]) {
            return false;
        }
    }
    for key, value := range m2 {
        if (value != m1[key]) {
            return false;
        }
    }
    return true;
}
func makeMap(str string) (map[byte]int) {
    m := make(map[byte]int);
    for i := 0; i < len(str); i++ {
        m[str[i]]++;
    }
    return m;
}
func isAnagram(s string, t string) bool {
    return equalMaps(makeMap(s),makeMap(t));
}

func main() {
  fmt.Println("______Testing isAnagram______");
  fmt.Println();
}
