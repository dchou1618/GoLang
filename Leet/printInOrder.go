/* file printInOrder.go employs threads and go routines to
   run multiple tasks in PROGRESS simultaneously - concurrency
   concurrency ≠ parallelism
   Concurrency: About scale - more workers -> run faster
   Blocking lead to deadlocks -> ruins concurrency
*/

package main

import (
    "fmt"
    "time"
    "math/rand"
    "sync"
)

type printOrder struct {
    order map[string]int;
    wg sync.WaitGroup;
};

func Constructor(magnitudes map[string]int) printOrder{
    this := new(printOrder);
    this.order = magnitudes;
    return *this;
}

func (this *printOrder) first() {
    //defer this.wg.Done();
    this.wg.Add(1);
    go func(){
        time.Sleep(time.Duration(this.order["first"])*time.Millisecond);
        fmt.Print("first");
        this.wg.Done();
    }()
}

func (this *printOrder) second() {
    //defer this.wg.Done();
    this.wg.Add(1);
    go func(){
        time.Sleep(time.Duration(this.order["second"])*time.Millisecond);
        fmt.Print("second");
        this.wg.Done();
    }()
}
func (this *printOrder) third() {
    //defer this.wg.Done();
    this.wg.Add(1);
    go func(){
        time.Sleep(time.Duration(this.order["third"])*time.Millisecond);
        fmt.Print("third");
        this.wg.Done();
    }()
    //this.wg.Wait();
}

/* Goroutines */
// (1)

func seq(first int, second int, iters int) {
    if (iters > 0) {
        fmt.Print(first, " ");
    }
    for i := 0; i < iters; i++ {
        fmt.Print(second, " ");
        temp := first+second;
        first = second;
        second = temp;
        // can allow for time delay between goroutines
        waitTime := time.Duration(rand.Intn(250));
        time.Sleep(time.Millisecond*waitTime);
    }
}

/* Channels can allow go routines to synchronize with each other */

/* Task 1: Design concurrent function that accepts integer values and
   outputs the values in ascending order (Sleep Sort)
*/

func intPing(num chan int, sentNums int) {
    for i := 0; i < sentNums; i++ {
        rand.Seed(time.Now().UnixNano());
        num <- rand.Intn(100);
    }
}

func orderByTime(nums chan int, sentNums int) {
    var wg sync.WaitGroup;
    var sortedOrder []int;
    for i := 0; i < sentNums; i++ {
        wg.Add(1);
        go func() {
            defer wg.Done(); // after operations done running
            val := <- nums;
            time.Sleep(time.Duration(val*3)*time.Millisecond);
            sortedOrder = append(sortedOrder, val);
            fmt.Print(val, " ");
        }()
    }
    wg.Wait();
    fmt.Println(sortedOrder);
}


/* Alternately output foobar */

type FooBar struct {
    iterations int;
    wg sync.WaitGroup;
};

func ConstructFooBar(n int) FooBar {
    this := new(FooBar);
    this.iterations = n;
    return *this;
}

func (this *FooBar) foo() {
    this.wg.Add(2);
    go func() {
        for i := 0; i < this.iterations; i++ {
            fmt.Print("foo");
            time.Sleep(time.Duration(1)*time.Millisecond);
        }
    }();
}

func (this *FooBar) bar() {
  go func() {
      for i := 0; i < this.iterations; i++ {
          fmt.Print("bar");
          time.Sleep(time.Duration(1)*time.Millisecond);
      }
      this.wg.Wait();
  }();
}

/* with go routines running in the main, the main function executes
   for a shorter amount of time than other routines.
   Using a wait group
*/

/*
    Select shapes data stream flow
    Fan-out: can send values in channel into other output channels
*/
func main() {
    // for i := 2; i < 16; i += 2 {
    //     go seq(1, 1, i);
    // }
    // var input string;
    // fmt.Scanln(&input);

    // pOrder := Constructor(map[string]int{"first":1,"second":2,"third":3});
    // pOrder.third();
    // pOrder.second();
    // pOrder.first();
    // pOrder.second();
    // pOrder.wg.Wait();
    newFoo := ConstructFooBar(3);
    newFoo.foo();
    newFoo.bar();
    // var n chan int = make(chan int);
    // go intPing(n, 10);
    // orderByTime(n, 10);
    // close(n);
    /* if orderByTime is a goroutine, it will run in the
       background and
    */
    var input string;
    fmt.Scanln(&input);
}
