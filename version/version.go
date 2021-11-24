package version

import (
	"fmt"
)

const (
	// ServiceName 服务名称
	ServiceName = "yCmdb"

	// Description 服务描述
	Description = "资源中心"
)

var (
	GIT_TAG    string
	BUILD_TIME string
	GIT_BRANCH string
	GIT_COMMIT string
	GO_VERSION string
)

// FullVersion show the version info
func FullVersion() string {
	version := fmt.Sprintf("GIT_TA: %s\nBuild Time: %s\nGit Branch: %s\nGit Commit: %s\nGo Version: %s\n", GIT_TAG, BUILD_TIME, GIT_BRANCH, GIT_COMMIT, GO_VERSION)

	return version
}

// Short 版本缩写
func Short() string {
	commit := ""
	if len(GIT_COMMIT) > 8 {
		commit = GIT_COMMIT[:8]
	}

	return fmt.Sprintf("%s[%s]", GIT_BRANCH, commit)
}
