// file prisonCellularAutomata.go

package main

import (
  "fmt"
)

func prisonAfterNDays(cells []int, N int) []int {
    // cellular automata repeats after 14 times
    var newState []int;
    newState = append(newState, cells...);
    N = (N-1)%14+1;
    for (N > 0) {
        for i := range cells {
            if (i-1 < 0) {
                newState[i] = 0;
            } else if (i+1 >= len(cells)) {
                newState[i] = 0;
            } else {
                if ((cells[i-1] == 0 && cells[i+1] == 0) ||
                    (cells[i-1] == 1 && cells[i+1] == 1)) {
                    newState[i] = 1;
                } else {
                    newState[i] = 0;
                }
            }
        }
        cells = []int{};
        cells = append(cells, newState...);
        N--;
    }
    return newState;
}

func main() {
  
}
