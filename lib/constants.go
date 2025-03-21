package lib

// Platform represents an operating system and its compatible browsers
type Platform struct {
	Name     string
	Browsers []Browser
}

// Browser represents a browser and its versions
type Browser struct {
	Name      string
	Versions  []string
	HasWebkit bool
}

// ValidPlatforms contains all valid platform-browser combinations
var ValidPlatforms = []Platform{
	{
		Name: "Windows NT 10.0; Win64; x64",
		Browsers: []Browser{
			{
				Name:      "Chrome",
				Versions:  []string{"89.0.4389.82", "90.0.4430.85", "91.0.4472.124"},
				HasWebkit: true,
			},
			{
				Name:      "Firefox",
				Versions:  []string{"87.0", "88.0", "89.0", "90.0"},
				HasWebkit: false,
			},
		},
	},
	{
		Name: "Macintosh; Intel Mac OS X 10_15_7",
		Browsers: []Browser{
			{
				Name:      "Safari",
				Versions:  []string{"14.0.3", "14.1", "14.1.1", "15.0"},
				HasWebkit: true,
			},
			{
				Name:      "Chrome",
				Versions:  []string{"89.0.4389.82", "90.0.4430.85"},
				HasWebkit: true,
			},
		},
	},
	{
		Name: "X11; Linux x86_64",
		Browsers: []Browser{
			{
				Name:      "Firefox",
				Versions:  []string{"87.0", "88.0", "89.0"},
				HasWebkit: false,
			},
			{
				Name:      "Chrome",
				Versions:  []string{"89.0.4389.82", "90.0.4430.85"},
				HasWebkit: true,
			},
		},
	},
}

// WebkitVersion is the standard WebKit version string
const WebkitVersion = "537.36"
