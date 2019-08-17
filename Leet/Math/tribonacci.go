// file tribonacci.go

package main

import (
  "fmt"
)


func tribonacci(n int) int {
    if (n == 0) {
        return 0;
    } else if (n == 1 || n == 2) {
        return 1;
    } else {
        zeroth := 0;
        first := 1;
        second := 1;
        num := 2;
        for (num < n) {
            tempSecond := second;
            second = second+first+zeroth;
            zeroth = first;
            first = tempSecond;
            num++;
        }
        return second
    }
}
