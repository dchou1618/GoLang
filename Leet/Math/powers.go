// file powers.go

package main

import (
  "fmt"
)

func isPowerOfK(n int, k int) bool {
  logEquality := math.Pow(k,math.Floor(math.Log(float64(n))/math.Log(float64(k)))) == float64(n);
  return n != 0 && logEquality;
}

func main() {

}
