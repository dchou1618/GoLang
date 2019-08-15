// file stringPermutation.go
// name: Dylan
// id: dvchou

package main

import (
  "fmt"
)


func match(m1 map[byte]int, m2 map[byte]int) (bool) {
    for key, value := range m1 {
        if (value != m2[key]) {
            return false;
        }
    }
    return true;
}
func stringToMap(s string) (map[byte]int) {
    m := make(map[byte]int);
    for j := 0; j < len(s); j++ {
        m[s[j]] += 1;
    }
    return m;
}
func matchInclusion(s1 string, s2 string, m1 map[byte]int, i int) (bool) {
    for (i+len(s1) <= len(s2)) {
        if (match(m1,stringToMap(s2[i:i+len(s1)]))) {
            return true;
        } else {
            i++;
        }
    }
    return false;
}
func checkInclusion(s1 string, s2 string) bool {
    i := 0;
    m1 := stringToMap(s1);
    return matchInclusion(s1,s2,m1,i);
}


// other string manipulation
// objective is to memoize solution so dp[i][j] stores values of
// different lengths of
func isInterleave(s1 string, s2 string, s3 string) bool {
    if (len(s1) == 0 && len(s2) == 0 && len(s3) == 0) {
        return true;
    } else {
        if (len(s1) > 0 && len(s3) > 0 && s1[0] == s3[0]) {
            if (len(s2) > 0 && len(s3) > 0 && s2[0] == s3[0]) {
                return isInterleave(s1[1:], s2, s3[1:]) || isInterleave(s1, s2[1:], s3[1:]);
            } else {
                return isInterleave(s1[1:], s2, s3[1:]);
            }
        } else if (len(s2) > 0 && len(s3) > 0 && s2[0] == s3[0]) {
            return isInterleave(s1, s2[1:], s3[1:]);
        } else {
            return false;
        }
    }
}

func main() {
    
}
