// file zigZagConversion.go

package main

import (
  "fmt"
)



// rather than perceiving movement as 2d, consider linear up/down
func zigzag(s string, numRows int) map[int]string {
    lettersPerRow := make(map[int]string);
    pos := 0;
    currRow := 1;
    var increasing bool = true;
    for pos < len(s) {
        lettersPerRow[currRow] += s[pos:pos+1];
        if (currRow == numRows) {
            increasing = false;
        }
        if (currRow == 1) {
            increasing = true;
        }
        if (increasing) {
            currRow++;
        } else {
            currRow--;
        }
        pos++;
    }
    return lettersPerRow;
}

// down and diagonal up-right
func convert(s string, numRows int) string {
    zigzagConversion := "";
    zigMap := zigzag(s, numRows);
    keys := []int{};
    for key,_ := range zigMap {
        keys = append(keys,key);
    }
    sort.Ints(keys); // having sorted the keys
    for _,key := range keys {
        zigzagConversion += zigMap[key];
    }
    zigMap = nil;
    keys = nil;
    return zigzagConversion;
}


func main() {

}
