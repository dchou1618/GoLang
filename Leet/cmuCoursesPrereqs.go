/* cmuCoursesPrereqs.go - employing means of topological sort to
   find ways to take all such courses in a specific major.

   Courses come from:
      http://coursecatalog.web.cmu.edu/schoolofcomputerscience/courses/
   Only courses in computer science department for now.
   Chose most popular/notable courses
*/

package main

import (
    "fmt"
    "os"
    "bufio"
)

func writeToFile(coursePath []string, path string) (error) {
    file, err := os.Open(path);
    if (err != nil) {
        return err;
    } else {
        write := bufio;
    }
}

type courseNode struct {
    courseName string;
    nextCourses []*courseNode; /* prerequisite for... */
}

type courseCatalog struct {
    courses []*courseNode;
}

func indegreeCourse() {
    
}

/* indegree nodes in graph (queue creation)- no prerequisites */
func _introCourses(catalog *courseCatalog) []string {
    introCourses := []string{};
    for _, course := range catalog.courses {
        if (course.nextCourses == nil) {
            introCourses = append(introCourses, course.courseNames);
        }
    }
    return introCourses;
}

func _prerequisiteCourses(numCourses int, catalog *courseCatalog) [][]string {
    prereqs := [][]string{};
    for _, course := range catalog.courses {

    }
    return prereqs;
}

func _CMUPrereqOrder() []string {

}


func main() {
    cs150 := &courseNode{courseName: "15-150: Functional Programming",
                         nextCourses: []*courseNode{}};
    cs122 := &courseNode{courseName: "15-122: Imperative Computation",
                         nextCourse: []*courseNode{}};
    cs121 := &courseNode{courseName: "15-121: Intro Data Structures",
                         nextCourses: []*courseNode{}};
    cs112 := &courseNode{courseName: "15-112: Programming Fundamentals",
                         nextCourses: []*courseNode{cs121, cs122, cs150}};
    cs110 := &courseNode{courseName: "15-110", nextCourses: nil};
    conceptsMath := &courseNode{courseName: "21-127",
                                nextCourses: []*courseNode{}};
    var catalog *courseCatalog;
    catalog.courses = []*courseNode{conceptsMath, cs110, cs112, cs121, cs122,
                                    cs150};
}
