package useragent

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/orbulant/random-user-agent/lib"
)

type UserAgent struct {
	randomiser *rand.Rand
}

func New() *UserAgent {
	// Creates random source once, but with current time as seed
	return &UserAgent{
		randomiser: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (ua *UserAgent) GetRandomAgent() string {
	// Get random platform
	platformIdx := ua.randomiser.Intn(len(lib.ValidPlatforms))
	platform := lib.ValidPlatforms[platformIdx]

	// Get random browser for this platform
	browserIdx := ua.randomiser.Intn(len(platform.Browsers))
	browser := platform.Browsers[browserIdx]

	// Get random version for this browser
	versionIdx := ua.randomiser.Intn(len(browser.Versions))
	browserVersion := browser.Versions[versionIdx]

	switch browser.Name {
	case "Chrome":
		return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s",
			platform.Name, lib.WebkitVersion, browserVersion, lib.WebkitVersion)
	case "Firefox":
		return fmt.Sprintf("Mozilla/5.0 (%s) Gecko/20100101 Firefox/%s",
			platform.Name, browserVersion)
	case "Safari":
		return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/%s (KHTML, like Gecko) Version/%s Safari/%s",
			platform.Name, lib.WebkitVersion, browserVersion, lib.WebkitVersion)
	default:
		return ""
	}
}
