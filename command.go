//go:build unix

package gorscript

import "regexp"

const cmdName = "Rscript"

var cmdRegex = regexp.MustCompile(`^Rscript$`)
