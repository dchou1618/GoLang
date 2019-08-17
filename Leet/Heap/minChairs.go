/* minChairs.go file - n guests invited to party
   kth guest attends at time S[k], leaving at E[k]
   - similar to question on interval list intersections
*/
package main

import (
    "fmt"
)

type minHeap struct {
    intervals []*Interval;
};

func (this *minHeap) size() int {
    return len(this.intervals);
}

func (this *minHeap) enqueue(interval *Interval) {
    if (this.size() == 0) {
        this.intervals = append(this.intervals, &Interval{0,0});
    }
    this.intervals = append(this.intervals, interval);
    var i int = len(this.intervals)-1;
    for (i/2 > 0) { 
        if (this.intervals[i].end < this.intervals[i/2].end) {
            var temp *Interval = this.intervals[i];
            this.intervals[i] = this.intervals[i/2];
            this.intervals[i/2] = temp;
            i /= 2;
        } else {
            break;
        }
    }
}

func (this *minHeap) dequeue() *Interval {
    if (this.size() > 0) {
      front := this.intervals[len(this.intervals)-1];
      this.intervals = this.intervals[:len(this.intervals)-1];
      if (this.size() == 1) {
          return front;
      }
      dequeued := this.intervals[1];
      this.intervals[1] = front;
      var i int = 1;
      for (2*i < len(this.intervals)) {
          if (2*i+1 < len(this.intervals)) {
              if (this.intervals[i].end <= this.intervals[2*i].end && this.intervals[i].end <= this.intervals[2*i+1].end) {
                  break;
              } else if (this.intervals[2*i].end < this.intervals[2*i+1].end) {
                  var temp *Interval = this.intervals[2*i+1];
                  this.intervals[2*i+1] = this.intervals[i];
                  this.intervals[i] = temp;
                  i = 2*i + 1;
              } else {
                  var temp *Interval = this.intervals[i];
                  this.intervals[i] = this.intervals[2*i];
                  this.intervals[2*i] = temp;
                  i *= 2;
              }
          } else {
              if (this.intervals[i].end <= this.intervals[2*i].end) {
                  break;
              } else {
                  var temp *Interval = this.intervals[2*i];
                  this.intervals[2*i] = this.intervals[i];
                  this.intervals[i] = temp;
              }
          }
      }
      return dequeued;
    } else {
        return nil;
    }
}

type Interval struct {
    start int;
    end int;
};

func (this *Interval) display() {
    fmt.Print("{", this.start, "<->", this.end, "}");
}

func binaryInsertion(interval *Interval, intervals []*Interval) int {
    lowerBound, upperBound := 0, len(intervals)-1;
    mid := (lowerBound+upperBound)/2;
    for (lowerBound <= upperBound) {
        if (intervals[mid].start < interval.start) {
            lowerBound = mid+1;
        } else if (intervals[mid].start > interval.start) {
            upperBound = mid-1;
        } else {
            return mid;
        }
        mid = (lowerBound+upperBound)/2;
    }
    return lowerBound;
}

/* sorts intervals by start time */
func sortIntervals(intervals []*Interval) []*Interval {
    sortedIntervals := []*Interval{};
    for _, interval := range intervals {
        var index int = binaryInsertion(interval, sortedIntervals);
        sortedIntervals = append(sortedIntervals[:index], append([]*Interval{interval}, sortedIntervals[index:]...)...);
    }
    return sortedIntervals;
}

func guestIntervals(S []int, E []int) []*Interval {
    var guests []*Interval;
    for i := 0; i < len(S); i++ {
        guests = append(guests, &Interval{start: S[i], end: E[i]});
    }
    return guests;
}

func _minChairsAlloc(S []int, E []int) int {
    if (len(S) != len(E)) {
        panic(1);
    }
    guests := guestIntervals(S,E);
    guests = sortIntervals(guests);
    priorityQueue := &minHeap{};
    priorityQueue.enqueue(guests[0]);
    for i := 1; i < len(guests); i++ {
        newInterval := guests[i];
        earliestArrival := priorityQueue.dequeue();
        // for _, interval := range priorityQueue.intervals {
        //     interval.display();
        // }
        if (newInterval.start >= earliestArrival.end) {
            earliestArrival.end = newInterval.end;
        } else {
            priorityQueue.enqueue(newInterval);
            // adds newInterval if intersects
        }
        priorityQueue.enqueue(earliestArrival);
        for _, interval := range priorityQueue.intervals {
            interval.display();
        }
        fmt.Println();
        /* otherwise, you would add onto the priorityqueue
           (ending length) being the number of intersecting intervals
        */
    }
    return priorityQueue.size()-1;
}


func main() {
    S1 := []int{1, 2, 6, 5, 3};
    E1 := []int{5, 5, 7, 6, 8};
    fmt.Println(_minChairsAlloc(S1,E1));
    S2 := []int{0, 5, 15};
    E2 := []int{30, 10, 20};
    fmt.Println(_minChairsAlloc(S2,E2));
}
