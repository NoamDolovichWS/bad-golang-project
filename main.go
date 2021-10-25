package main

import (
	"fmt"
	"github.com/smallnest/ringbuffer"
	"github.com/whitesource/ws-go-sdk/utils"
	"golang.org/x/text/transform"
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
	var buff ringbuffer.RingBuffer
	buff.WriteString("hey")
	fmt.Println(buff.ReadByte())
	n := transform.Chain()
	fmt.Println(n)
}
