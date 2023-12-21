package main

import (
	"fmt"

	"github.com/greedyor/glua"
)

func main() {
	Data, err := glua.ExecToPath("./xxapi.lua")
	if err != nil {
		fmt.Println("ExecToPath error:", err)
		return
	}
	fmt.Println(Data)
}
