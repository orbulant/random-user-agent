package main

import (
	"fmt"

	"github.com/orbulant/random-user-agent/useragent"
)

func main() {
	ua := useragent.New()

	// Generate 5 different user agents
	for i := 0; i < 5; i++ {
		fmt.Println(ua.GetRandomAgent())
	}
}
