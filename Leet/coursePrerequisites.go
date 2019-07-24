// file using

package main

import (
  "fmt"
)

// generate indegree mapping runtime(53%), memory (93%)
func getEdges(prerequisites [][]int) map[int]int {
    result := make(map[int]int);
    for _, coursePair := range prerequisites {
        result[coursePair[1]]++;
        _, included := result[coursePair[0]];
        if (!included) {
            result[coursePair[0]] = 0;
        }
    }
    return result;
}

func inDegreeQueue(vertexDegree map[int]int) []int {
    queue := []int{};
    for key ,val := range vertexDegree {
        if (val == 0) {
            queue = append(queue, key);
        }
    }
    return queue;
}

func generateAdjacencyLst(numCourses int, prerequisites [][]int) [][]int {
    neighbors := make([][]int, numCourses);
    for _, pair := range prerequisites {
        neighbors[pair[0]] = append(neighbors[pair[0]], pair[1]);
    }
    return neighbors;
}

// begin somewhere in which there are no prerequisites
func canFinish(numCourses int, prerequisites [][]int) bool {
    if (numCourses == 0 || len(prerequisites) == 0) {
        return true;
    } else {
        indegree := getEdges(prerequisites);
        queue := inDegreeQueue(indegree);
        if (len(queue) == 0) {
            return false;
        } else {
            courseOrder := []int{};
            neighbors := generateAdjacencyLst(numCourses, prerequisites);
            for (len(queue) != 0) { // repeat while queue is not empty
                topNode := queue[0]; // dequeue
                queue = queue[1:];
                courseOrder = append(courseOrder, topNode);
                for _, neighbor := range neighbors[topNode] {
                    indegree[neighbor]--;
                    if (indegree[neighbor] == 0) {
                        queue = append(queue, neighbor);
                    }
                }
            }
            // fmt.Println(courseOrder);
            for _, val := range indegree { // if all are indegree values
                if (val != 0) {
                    return false;
                }
            }
            return true;
        }
    }
}
/*
Test
21
[[0,1],[0,2],[0,3],[4,0],[4,5],[4,3],[4,6],[6,7],[8,4],[4,9],[8,10],[11,10],
[11,4],[4,12],[12,13],[12,14],[4,15],[15,16],[15,17],[15,18],[15,19],[15,20]]
*/
