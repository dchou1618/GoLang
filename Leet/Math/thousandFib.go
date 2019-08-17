// file thousandFib.go

package main

import (
    "fmt"
    "strconv"
)

func thousandFib(first int, second int, digits int) int {
    for (len(strconv.Itoa(second)) < digits) {
        var newSecond int = first+second;
        first = second;
        second = newSecond;
    }
    return second;
}

func main() {
    fmt.Println(thousandFib(1, 1, 1000));
}
