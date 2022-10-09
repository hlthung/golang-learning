//go:build tools
// +build tools

// More Info:
// - https://www.jvt.me/posts/2022/06/15/go-tools-dependency-management/
// - https://github.com/golang/go/issues/25922
// - https://marcofranssen.nl/manage-go-tools-via-go-modules/
package tools

import (
	_ "golang.org/x/tools/cmd/stringer"
)
