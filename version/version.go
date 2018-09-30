package version

import (
	"runtime"
)

var (
	Version      = "1.5.3"
	UTCBuildTime = "undefined"
	GitCommit    = "undefined"
	OS           = runtime.GOOS
	Arch         = runtime.GOARCH
	GoVersion    = "undefined"
)
