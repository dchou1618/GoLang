// file excelSheetColumnTitle.go

package main

import (
  "fmt"
  "math"
)

func convertToTitle(n int) string {
    var column string = "";
    for n > 0 {
        if (n%26 == 0) {
            column = "Z" + column;
        } else {
            if (n%26 > 0) {
                column = (string(int('A') + (n%26-1))) + column;
            }
        }
        if (n%26 == 0) {
            n -= 26;
        } else {
            n -= n%26;
        }
        n /= 26;
    }
    return column;
}

func titleToNumber(s string) int {
    var columnNumber int;
    currPower := 0;
    for pos := len(s)-1; pos >= 0; pos-- {
        columnNumber += (int(s[pos])-int('A')+1)*int(math.Pow(26,float64(currPower)));
        currPower++;
    }
    return columnNumber;
}


func main() {

}
