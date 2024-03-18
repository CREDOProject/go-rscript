package gorscript

import (
	"errors"

	"github.com/CREDOProject/sharedutils/files"
	"github.com/CREDOProject/sharedutils/shell"
)

var execCommander = shell.New

// DetectRscriptBinary returns the Rscript binary path in the system.
func DetectRscriptBinary() (string, error) {
	return execCommander().LookPath(cmdName)
}

// RscriptBinaryFrom returns the Rscript binary from a specified path.
func RscriptBinaryFrom(path string) (string, error) {
	execs, err := files.ExecsInPath(path, func(s string) bool {
		return cmdRegex.MatchString(s)
	})
	if len(execs) < 1 {
		return "", errors.New("No rscript found.")
	}
	return execs[0], err
}
