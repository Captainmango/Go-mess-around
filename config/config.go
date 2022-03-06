package config

type configObject struct {
    wordOne string
    wordTwo string
    numOne int
    numTwo int
}

type IConfig interface{
    GetWordOne() string
    GetWordTwo() string
    GetNumOne() int
    GetNumTwo() int
    GetConfig() *configObject
    SetConfig(string, string, int, int) *configObject
}

// Constructor for the config object
func New (wordOne string, wordTwo string, numOne int, numTwo int) *configObject {
    myConfig := configObject{wordOne, wordTwo, numOne, numTwo}

    return &myConfig
}

func (c *configObject) GetConfig () *configObject {
    return c
}

func (c *configObject) SetConfig (wordOne string, wordTwo string, numOne int, numTwo int) *configObject {
    c.numOne = numOne
    c.numTwo = numTwo
    c.wordOne = wordOne
    c.wordTwo = wordTwo

    return c
}

func (c *configObject) GetNumOne () int {
    return c.GetConfig().numOne
}

func (c *configObject) GetNumTwo () int {
    return c.GetConfig().numTwo
}

func (c *configObject) GetWordOne () string {
    return c.GetConfig().wordOne
}

func (c *configObject) GetWordTwo () string {
    return c.GetConfig().wordTwo
}