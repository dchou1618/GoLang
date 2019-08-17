


package main

import (
  "fmt"
)



func corpFlightBookings(bookings [][]int, n int) []int {
    seatsBooked := make([]int, n+1);
    for _, label := range bookings {
        seatsBooked[label[0]-1] += label[2];
        seatsBooked[label[1]] -= label[2];
    }
    cumulativeSum := 0;
    for i, val := range seatsBooked {
        cumulativeSum += val;
        seatsBooked[i] = cumulativeSum;
    }
    return seatsBooked[:n];
}
