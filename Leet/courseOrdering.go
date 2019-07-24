// file courseOrdering

func inDegree(numCourses int, prerequisites [][]int) map[int]int {
    indegree := make(map[int]int);
    for _, course := range prerequisites {
        indegree[course[0]]++;
    }
    return indegree;
}

func neighboringCourses(numCourses int, prerequisites [][]int) [][]int{
    neighbors := make([][]int);
    for _, pair := range prerequisites {
        neighbors[pair[1]] = append(neighbors[pair[1]], pair[0]);
    }
    return neighbors;
}

func makeQueue(numCourses int, indegree map[int]int) []int {
    queue := []int{};
    for num := 0; num < numCourses; num++ { // included other courses
        // not among the prerequisites
        _,inclusive := indegree[num];
        if (!inclusive) {
            queue = append(queue, num);
        }
    }
    return queue;
}

func findOrder(numCourses int, prerequisites [][]int) []int {
    if (len(prerequisites) == 0) {
        result := []int{};
        for num := 0; num < numCourses; num++ {
            result = append(result, num);
        }
        return result;
    } else {
        courseOrdering := []int{};
        indegree := inDegree(numCourses, prerequisites);
        queue := makeQueue(numCourses, indegree);
        if (len(queue) == 0) {
            return []int{};
        } else {
            neighbors := neighboringCourses(numCourses, prerequisites);
            for len(queue) != 0 {
                firstCourse := queue[0];
                queue = queue[1:];
                courseOrdering = append(courseOrdering, firstCourse);
                for _, neighbor := range neighbors[firstCourse] {
                    indegree[neighbor]--;
                    if (indegree[neighbor] == 0) {
                        queue = append(queue, neighbor);
                    }
                }
            }
            if (len(courseOrdering) == numCourses) {
                return courseOrdering;
            } else {
                return []int{};
            }
        }
    }
}
