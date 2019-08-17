/* file leafDistribution.go aims to construct and query a decision tree
*/

package main

import (
    "fmt"
    "math"
)


type Employee struct {
    Name string
    Age int
    Experience int
    Salary float64
    Department string
}

type decisionNode struct {
    e *Employee;
    left *decisionNode;
    right *decisionNode;
};

func sortEmployees(vals []interface{}) []interface{} {
    newVals := []interface{}; // doesn't make sense to split on name
    for _, elem := range vals {
        var indexToInsert int = -1;
        low, up := 0, len(newVals)-1;
        mid := (low+up)/2;
        for (low <= up) {
            if (newVals[mid] == elem) {
                indexToInsert = mid;
                break;
            } else if (newVals[mid] > elem) {
                up = mid-1;
            } else {
                low = mid+1;
            }
            mid = (low+up)/2;
        }
        if (indexToInsert < 0) {
            indexToInsert = low;
        }
        singleton := []interface{};
        singleton = append(singleton, elem);
        newVals = append(newVals[:indexToInsert], append(singleton, newVals[indexToInsert:]...)...);
    }
    return newVals;
}
// used by ID3 decision trees, the attribute with the largest information gain
// is chosen to split upon
func highestInformationGain(employeeMap map[string][]interface{}) float64 {
    var lowestEntropy float64 = math.MaxInt64;
    // categorical variables
    for _, value := range employeeMap {

    }
}

func getEntropy(attributes []interface{}) float64 {
    employees = sortEmployees(employees, attribute);
    var entropy float64;

}

// improvement in split is from summation of square errors between node and
// child nodes after split (y_i-c)^2
func splitDecision(attributes []interface{}, quantitative bool) {
    if (quantitative) {
        for _, val := range attributes {

        }
    } else {
        var entropy float64;
        types := make(map[string]int);
        total := 0;
        for _, val := range attributes {
            types[val]++;
            total++;
        }
    }
}


func employeeMap(employees []*Employee) map[string][]interface{} {
    mappedEmployees := make(map[string][]interface{});
    for _, employee := range employees {
        mappedEmployees["Name"] = append(mappedEmployees["Name"], employee.Name);
        mappedEmployees["Department"] = append(mappedEmployees["Department"], employee.Department);
        mappedEmployees["Age"] = append(mappedEmployees["Age"], employee.Age);
        mappedEmployees["Experience"] = append(mappedEmployees["Experience"], employee.Experience);
        mappedEmployees["Salary"] = append(mappedEmployees["Salary"], employee.Salary);
    }
    return mappedEmployees;
}

func constructDecisionTree(employees []*Employee) *decisionNode {
    mapEmployees := employeeMap(employees);
    root := &decisionNode{nil,nil,nil};

}

func main() {
    employee1 := &Employee{"Joe", 21, 3, 70000, "Administration"};
    employee2 := &Employee{"Bo", 25, 10, 120000, "IT"};
    employee3 := &Employee{"Rufus", 29, 1, 90000, "Sales"};
    employee4 := &Employee{"Monroe", 19, 2, 60000, "Production"};
    employees := []*Employee{employee1, employee2, employee3, employee4};
    fmt.Println(constructDecisionTree(employees));
}
