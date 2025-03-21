package main

import (
	"fmt"

	"github.com/orbulant/random-user-agent/internal/useragent"
)

func main() {
	ua := useragent.New().GetRandomAgent()
	fmt.Println(ua)
}
