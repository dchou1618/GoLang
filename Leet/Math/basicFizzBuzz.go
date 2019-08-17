
package main

import (
  "fmt"
)

// outputs fizz, buzz for the first n integers >= 1
// or output
func fizzBuzz(n int, fizz int, buzz int) {
    if (n <= 0) {
        fmt.Println("n must be larger than 0");
        return;
    }
    fizzed := make(map[int]bool);
    buzzed := make(map[int]bool); // poor memory
    for i := 1; i <= n/fizz; i++ {
      fizzed[i*fizz] = true;
    }
    for j := 1; j <= n/buzz; j++ {
      buzzed[j*buzz] = true;
    }
    for i := 1; i <= n; i++ {
        if (fizzed[i] && buzzed[i]) {
            fmt.Println("fizzbuzz");
        } else if (fizzed[i]) {
            fmt.Println("fizz");
        } else if (buzzed[i]) {
            fmt.Println("buzz");
        } else {
            fmt.Println(i);
        }
    }
}


func main() {
  fizzBuzz(100,3,5);
}
