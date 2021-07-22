package main

import (
	"fmt"
	"sync"
)

func main() {
	var users sync.Map
	users.Store(10, "kk")
	users.Store(20, "郭聪")

	if value, ok := users.Load(10); ok {
		fmt.Println(value.(string))
	}

	if value, ok := users.Load(30); ok {
		fmt.Println(value)

	}

	if value, ok := users.Load(20); ok {
		fmt.Println(value)
	}

}
