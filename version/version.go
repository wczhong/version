package version

import (
	"fmt"
)

// 版本信息，利用Makefile编译时自动注入值
var (
	GitTag    string // git tag
	BuildTime string // build time
)

// String 返回完整的版本信息
func String() string {
	return fmt.Sprintf("Git Tag:%s, Build Time:%s", GitTag, BuildTime)
}
