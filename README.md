# random-user-agent

A Go package that generates random, realistic user agent strings for common browsers.

**This is useful for bypassing API restrictions that block requests with the same or missing `User-Agent` header in the HTTP request**

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
