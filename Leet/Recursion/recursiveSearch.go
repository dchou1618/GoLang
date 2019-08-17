// file recursiveSearch.go

package main

import (
  "fmt"
)

type Course struct {
    course string
    children []*Course
}

func courseSearch(rootCourse *Course, courseNumber string, currPath string) string {
    if (rootCourse != nil) { // if nil is included in the children courses
        if (rootCourse.course == courseNumber) {
            return currPath;
        } else {
            for _, child := range rootCourse.children {
                currPath += "/" + child.course;
                resultPath := courseSearch(child, courseNumber, currPath);
                if (resultPath != "No Course Found") {
                    return resultPath;
                } else {
                    currPath = currPath[0:len(currPath)-(len(child.course)+1)];
                }
            }
            return "No Course Found";
        }
    }
    return "No Course Found";
}

func getCourse(rootCourse *Course, courseNumber string) string {
    return courseSearch(rootCourse, courseNumber, rootCourse.course);
}


func main() {
    course1 := &Course{"99-307", nil};
    course2 := &Course{"99-308", nil};
    course3 := &Course{"15-122", nil};
    course4 := &Course{"15-150", nil};
    course5 := &Course{"15-213", nil};
    course6 := &Course{"15-110", nil};
    course7 := &Course{"15-112", nil};
    course8 := &Course{"42-101", nil};
    course9 := &Course{"42-201", nil};
    course10 := &Course{"18-100", nil};
    course11 := &Course{"18-202", nil};
    course12 := &Course{"18-213", nil};
    intro := &Course{"Intro", []*Course{course6, course7}};
    cs := &Course{"CS", []*Course{intro, course3, course4, course5}};
    scs := &Course{"SCS", []*Course{cs}};
    ece := &Course{"ECE", []*Course{course10, course11, course12}};
    bme := &Course{"BME", []*Course{course8, course9}};
    cit := &Course{"CIT", []*Course{ece, bme}};
    root := &Course{"CMU", []*Course{course1, course2, scs, cit}};
    fmt.Println(getCourse(root, "18-10"));
}
