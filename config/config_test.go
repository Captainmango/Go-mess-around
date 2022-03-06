package config

import (
    "fmt"
    "testing"
)

type TestData struct {
    description string
    data interface{}
    result interface{}
}

// Mock the config object and stub responses in the table. 
// Should be able to include SetConfig() test in here too

var (
    myConfig IConfig = New("Fizz", "Buzz", 5, 8)
)

func TestConfigStructMethods (t *testing.T) {
    myTestData := []TestData{
        {
            "It can retrieve the config",
            myConfig.GetConfig(),
            myConfig,
        },
        {
            "It can get trigger element one",
            myConfig.GetNumOne(),
            5,
        },
        {
            "It can get trigger element two",
            myConfig.GetNumTwo(),
            8,
        },
        {
            "It can get replacement element one",
            myConfig.GetWordOne(),
           "Fizz",
        },
        {
            "It can get replacement element two",
            myConfig.GetWordTwo(),
           "Buzz",
        },      
    }

    for _, scenario := range myTestData {
        result := scenario.data

        if result != scenario.result {
            t.Errorf("Result does not match scenario. Got %v expected %v", result, scenario.result)
        } else {
            fmt.Println(scenario.description)
        }
    }
}

func TestSetConfig(t *testing.T) {
    dummyConf := New("Test", "Dummy", 2, 3)
    result := dummyConf.SetConfig("Fez", "Koz", 5, 8)

    if result != dummyConf {
        t.Errorf("Result has not changed. Expected %v to be %v", dummyConf, result)
    } else {
        fmt.Println("It can set config values after being initialised")
    }
}
