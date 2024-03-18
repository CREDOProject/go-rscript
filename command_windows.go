//go:build windows

package gorscript

import "regexp"

const cmdName = "Rscript.exe"

var cmdRegex = regexp.MustCompile(`^Rscript(?:\.exe)$`)
