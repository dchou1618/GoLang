// file storingResults.go

package main

import (
  "fmt"

)

func fibMemo(n int, memo map[int]int) int {
  _, included := memo[n];
  if (included) {
    return memo[n];
  }
  f := 0;
  if (n <= 2) {
    f = 1;
  } else {
    f = fibMemo(n-1,memo) + fibMemo(n-2, memo);
  }
  memo[n] = f;
  return f;
}

func fib(n int) int {
  memo := make(map[int]int);
  return fibMemo(n, memo);
}

func main() {
  fmt.Println(fib(1000)); // overflows if above 2^31
}
