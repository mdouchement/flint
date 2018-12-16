package version

import (
	"runtime"
)

var (
	Version      = "1.6.0"
	UTCBuildTime = "undefined"
	GitCommit    = "undefined"
	OS           = runtime.GOOS
	Arch         = runtime.GOARCH
	GoVersion    = "undefined"
)
