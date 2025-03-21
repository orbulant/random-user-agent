# random-user-agent

A Go package that generates random, realistic user agent strings for common browsers.

## Installation

```sh
go get github.com/orbulant/random-user-agent
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/orbulant/random-user-agent/useragent"
)

func main() {
    ua := useragent.New()
    randomUserAgent := ua.GetRandomAgent()
    fmt.Println(randomUserAgent)
}
```
