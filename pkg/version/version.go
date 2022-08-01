package version

import "fmt"

var (
	Version   = "0.0.1"
	GitCommit = "HEAD"
)

func CurrentVersion() string {
	return fmt.Sprintf("%s (%s)", Version, GitCommit)
}
