/* file effSort.go covers efficiency and sorting methods in golang */


package main

import (
  "fmt"
  "sort"
  "time"
)

// can adjust the sorting algorithm to follow desc/asc order
// getting the duration of the sorting algorithm using package called
/* time. time.Duration are struct types that result in the time.Now() method */
func selectionSort(ascending bool, lst []int) (time.Duration,[]int){
  start := time.Now()
  for i := 0; i < len(lst); i++{
    index := i
    for j := i; j < len(lst); j++{
      if ascending{
        if lst[j] < lst[index]{
          index = j
        }
      } else {
        if lst[j] > lst[index]{
          index = j
        }
      }
    }
    lst[i],lst[index] = lst[index],lst[i]
  }
  end := time.Now()
  return end.Sub(start),lst
}



// func mergeLsts(L []int, i int, j int, last int){
//   index1 := i
//   index2 := j
//   length := last-i
//   auxLst := [length]int
//
//   for iter := 0; iter < length; iter++{
//
//   }
// }
//
// func mergeSort(ascending bool, lst []int) ([]int,time.Duration,int){
//
//   if ascending{
//
//   }
// }

/*
To sort specific ages, you override the sort interface methods.
The interface has Len() Less(i, j int) and Swap (i, j int) as methods
*/
type Person struct{
  Name string
  Age int
  Birthday string
}

type ByName []Person /* will organize individuals based on age */

func (b ByName) Len() (int){return len(b)}
func (b ByName) Less(i, j int) (bool){return b[i].Birthday < b[j].Birthday}
func (b ByName) Swap(i, j int) {b[j],b[i] = b[i],b[j]}


// Sorting by a key




func main(){
  people := []Person{
    {"Dan",18,"10/1/2000"},
    {"Frank",21,"10/1/1997"},
    {"John",24,"10/1/1994"},
    {"Arnold",28,"10/1/1990"}}
  sort.Sort(ByName(people))
  fmt.Println(people)
  for person := 0; person < len(people);person++{
    fmt.Println(people[person].Birthday)
  }
  fmt.Println(selectionSort(false,[]int{2,6,1,34,54,123,10,2,-2,3,4,21}))
}
