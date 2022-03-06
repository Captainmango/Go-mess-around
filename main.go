package main

import (
	"fmt"
    "main/fizzbuzz"
    "main/config"
)

var (
    myConfig config.IConfig = config.New("Fizz", "Buzz", 5, 8)
)

func main() {
    fmt.Printf("The config is now: %v \n", myConfig.GetConfig())
    fmt.Println("Printing fizzbuzz numbers")    
    fmt.Println("---------------------------")
    
	for i := 1; i <= 40; i++ {
        res, _ := fizzbuzz.GenerateFromNum(i, myConfig)
        fmt.Println(res)
    }

    fmt.Println("---------------------------")
    fmt.Println("Printing fizzbuzz numbers from slices")
    fmt.Println("---------------------------")
    res, _ := fizzbuzz.GenerateFromSlice([]int{1,2,3,4,5,6,8,40}, myConfig)
    fmt.Println(res)
    fmt.Println("---------------------------")

    myConfig.SetConfig("Harrold", "Farthing", 2, 7)
    fmt.Printf("The config is now: %v \n", myConfig.GetConfig())
    res2, _ := fizzbuzz.GenerateFromSlice([]int{1,2,3,4,5,6,7,8,14}, myConfig)
    fmt.Println(res2)
    fmt.Println("---------------------------")
}