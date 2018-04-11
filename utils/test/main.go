package main

import (
	"fmt"

	log "github.com/gonethopper/libs/logs"
	"github.com/gonethopper/libs/utils"
)

func main() {
	fmt.Println("hello utils")
	defer utils.PrintPanicStack()

	vals := utils.Rand(10, 100)
	for _, v := range vals {
		log.Info("normal values is :[%d]", v)
	}

	vals2 := utils.UniqRand(10, 100)
	for _, v := range vals2 {
		log.Info("uniq values is :[%d]", v)
	}

	uuid, err := utils.GenUUID()
	if err != nil {
		log.Info("get uuid error :", err)
	}
	log.Info("get uuid %s", uuid)
}
