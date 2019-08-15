// file minIndexSumTwoLsts.go

package main

import (
  "fmt"
)

func getMinSum(nums []int) (int) {
    if (len(nums) == 0) {
        return 0;
    }
    min := nums[0];
    for _, elem := range nums {
        if (elem < min) {
            min = elem;
        }
    }
    return min;
}

func getValues(m map[string]int) []int {
    result := []int{};
    for _, value := range m {
        result = append(result, value);
    }
    return result;
}

func findRestaurant(list1 []string, list2 []string) []string {
    // used var = nil to free up memory - golang has garbage collection
    restaurants := make(map[string]int);
    finalRestaurants := make(map[string]int);
    for i, restaurant := range list1 {restaurants[restaurant] += i;}
    for j, restaurant := range list2 {
        _, included := restaurants[restaurant];
        if (!included) {
            continue;
        } else {
            finalRestaurants[restaurant] = restaurants[restaurant] + j;
        }
    }
    restaurants = nil;
    values := getValues(finalRestaurants);
    min := getMinSum(values);
    result := []string{};
    values = nil;
    for key, _ := range finalRestaurants {
        if (finalRestaurants[key] == min) {
            result = append(result, key);
        }
    }
    finalRestaurants = nil;
    return result;
}
