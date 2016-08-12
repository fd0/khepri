// +build !windows

package local

import (
	"os"
	"restic/patchedos"
)

// set file to readonly
func setNewFileMode(f string, fi os.FileInfo) error {
	return patchedos.Chmod(f, fi.Mode()&os.FileMode(^uint32(0222)))
}
