package main

import (
	"fmt"

	"github.com/gonethopper/libs/config"
)

//TConf test sub config
type TConf struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	EnableUDP bool   `yaml:"enableudp"`
}

//AKConf test config
type AKConf struct {
	Test *TConf `yaml:"test"`
}

func main() {
	fmt.Println("Hello Lib")
	c := new(AKConf)
	c.Test = new(TConf)
	err := config.ParseConfig(c, "test", "../conf")
	if err != nil {
		fmt.Println("read config error")
		return
	}

	fmt.Printf("read config %v ", c.Test)

}
