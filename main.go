package main

import (
	"fmt"
	"github.com/whitesource/ws-go-sdk/utils"
)

func main() {
	fmt.Println("Hello!!!")
	myConfs := utils.Configurations{
		Url:      "a",
		ApiKey:   "b",
		UserKey:  "c",
		Timeouts: nil,
	}
	fmt.Println(myConfs)
}
