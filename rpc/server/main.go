package main

import (
	"fmt"
	"win-his/server"
)

func main() {
	if err := server.Run(); err != nil {
		fmt.Println(err)
	}
}
