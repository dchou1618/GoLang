// file minWindow.go implements sliding window algorithm

package main

import (
    "fmt"
)

/* Smallest window of string containing all characters in the string
   Problem can be reduced to finding smallest window in string that contains
   all characters of another string */
func minWindowContainAll(s string) string {
    // maintain distinct characters
    numChars := 0;
    seen := make(map[byte]bool);
    for i := range s {
        if (!seen[s[i]]){
            seen[s[i]] = true;
            numChars++;
        }
    }
    start, minStart, windowSize := 0, -1, len(s);
    counts := make(map[byte]int);
    matched := 0;
    for i := 0; i < len(s); i++ {
        counts[s[i]]++;
        if (counts[s[i]] == 1) {
            matched++;
        }
        if (matched == numChars) { // all chars included
            for (counts[s[start]] > 1) {
                if (counts[s[start]] > 1) {
                  counts[s[start]]--;
                }
                start++;
            }
            // new minWindow
            if (i-start+1 < windowSize) {
                windowSize = i-start+1;
                minStart = start;
            }
        }
    }
    return s[minStart:minStart+windowSize];
}

/* MinWindowSubstring
Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
*/
/* minimum window substring - O(n);
What really is the sliding window algorithm? (Turns nested -> single loops)

*/
func minWindow(s string, t string) string {
    // obtain distinct characters in t
    if (len(t) == 0) {
        return s;
    }
    tCounts := make(map[string]int);
    for j := 0; j < len(t); j++ {
        tCounts[t[j:j+1]]++;
    }
    startIndex, minIndex, minWindow := 0, -1, len(s)+1;
    counts := make(map[string]int);
    match := 0;
    for i := 0; i < len(s); i++ {
        counts[s[i:i+1]]++;
        if (counts[s[i:i+1]] <= tCounts[s[i:i+1]] && tCounts[s[i:i+1]] != 0) {
            match++;
        }
        if (match == len(t)) {
            for (counts[s[startIndex:startIndex+1]] > tCounts[s[startIndex:startIndex+1]] || tCounts[s[startIndex:startIndex+1]] == 0) {
                if (counts[s[startIndex:startIndex+1]] > tCounts[s[startIndex:startIndex+1]]) {
                    counts[s[startIndex:startIndex+1]]--;
                }
                startIndex++;
            }
            if (i-startIndex+1 < minWindow) {
                minWindow = i-startIndex+1;
                minIndex = startIndex;
            }
        }
    }
    if (minIndex == -1) {
        return "";
    } else {
        return s[minIndex:minIndex+minWindow];
    }
}



func main() {
    fmt.Println(minWindowContainAll("aabcbcdbca"));
}
