package fizzbuzz

import (
    "fmt"
    "testing"
    "main/config"
)

type TestData struct {
    description string
    data interface{}
    result interface{}
}

func TestGenerateFromNum(t *testing.T) {
    myConfig := config.New("Fizz", "Buzz", 5, 8)
    myTestData := []TestData{
        {
            "It can handle number one condition triggering",
            5,
            "Fizz!",
        },
        {
            "It can handle number two condition triggering",
            8,
            "Buzz!",
        },
        {
            "It can handle combined condition triggering",
            40,
            "FizzBuzz!!!",
        },
        {
            "It can handle numbers without triggering a condition",
            3,
            "3",
        },
    }

    for _, scenario := range myTestData {
        result, _ := GenerateFromNum(scenario.data.(int), myConfig)

        if result != scenario.result {
            t.Errorf("GenerateFromNum(%d, %v) did not return expected result. Expected %v. Got %v",scenario.data, myConfig, scenario.result, result)
        } else {
            fmt.Println(scenario.description)
        }
    }
}

func TestGenerateFromSlice(t *testing.T){
    myConfig := config.New("Fizz", "Buzz", 2, 3)
    myTestData := []TestData{
        {
            "It can handle a slice of non triggering numbers correctly",
            []int{1,5,7},
            []string{"1", "5", "7"},
        },
        {
            "It can handle a slice of triggering numbers correctly",
            []int{2,3,6},
            []string{"Fizz!", "Buzz!", "FizzBuzz!!!"},
        },
        {
            "It handles a mix of triggering and non triggering numbers correctly",
            []int{2,3,7},
            []string{"Fizz!", "Buzz!", "7"},
        },
    }
    
    for _, scenario := range myTestData {
        result, _ := GenerateFromSlice(scenario.data.([]int), myConfig)

        for idx, item := range result {
            if item != scenario.result.([]string)[idx] {
                t.Errorf("Data at index %v in result is not equal to %v.", idx, scenario.result.([]string)[idx])
            }
        }
        
        fmt.Println(scenario.description)
    }
}