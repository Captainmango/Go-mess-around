package fizzbuzz

import (
    "strconv"
    "main/config"
)

// Accepts an int and config and returns a string.
// FizzBuzz!!! if num is multiple of 8 and 5
// Buzz! if num is multiple of 8
// Fizz! if num is multiple of 5
func GenerateFromNum(num int, myConfig config.IConfig) (string, error) {
    if num % myConfig.GetNumOne() == 0 && num % myConfig.GetNumTwo() == 0 {
        return myConfig.GetWordOne() + myConfig.GetWordTwo() + "!!!", nil
    }

    if num % myConfig.GetNumTwo() == 0 {
        return myConfig.GetWordTwo() + "!", nil
    }

    if num % myConfig.GetNumOne() == 0 {
        return myConfig.GetWordOne() + "!", nil
    }
    
    return strconv.Itoa(num), nil
}

// Accepts a slice of ints converts to a fizzbuzz string 
// and returns a slice of the values
// Uses GenerateFromNum to compute the string values
func GenerateFromSlice(slice []int, myConfig config.IConfig) ([]string, error) {
    var output []string
    
    for i := 0; i < len(slice); i++ {
        item, err := GenerateFromNum(slice[i], myConfig)
        
        if err != nil {
            return []string{}, err
        }
        
        output = append(output, item)
    }

    return output, nil
}