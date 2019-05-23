/* file basics.go covers variables, types and data type manipulation */

package main

import (
  "fmt"
  "strconv"
  "math"
)

// isPrime function

func isPrime(n int) (bool){
  if n < 2{
    return false
  } else if n == 2{
    return true
  } else if n%2 == 0{
    return false
  }
  for factor := 3; factor < n/2; factor+=2{
    if n%factor == 0{
      return false
    }
  }
  return true
}

// rounding number num to n places
func nRound(n float64, num float64) (float64){
  return math.Ceil(num*math.Pow(10,n))/math.Pow(10,n)
}

func convertToInches(quantity float64, unit string) (float64,bool){
  units := map[string]float64{"centimeter":1/2.54,"kilometer":39370.079,
           "meter":39.37,"millimeter":25.4,"micrometer":25400,
           "nanometer":2.54*(10^7),"mile":63360,"yard":36,"foot":12,
           "nautical mile":72913.386} // map maps string keys to float values
  for measure := range units{
    if unit == measure{
      return nRound(1,quantity*units[measure]), true
    }
  }
  return 0.0, false
}



func isCircularPrime(n int) (bool){
  num := strconv.Itoa(n)
  for i := 0;i<len(num);i++{
    newNum := string(num[i:])+string(num[:i])
    rotation, err := strconv.Atoi(newNum)
    if err == nil{
      if !isPrime(rotation){
        return false
      }
    } else {
      return false
    }
  }
  return true
}

func nthCircularPrime(n int) (int){
  guess,found := 2,0;
  for found != n{
    guess += 1
    if isCircularPrime(guess){
      found += 1
    }
  }
  return guess
}

func test() (int){
  v := 0
  for _,i := range []int{1,2,3,4}{
    v = i // reassigns v to point to i (latest value in loop)
  }
  return v
}

func main(){
  fmt.Println(convertToInches(3.4,"meter"))
  // fmt.Println(nthCircularPrime(30))
}
