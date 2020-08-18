// +build !windows

package restic

// isAllowedToSymlink is true if the process is allowed to create symlinks on
// windows
func isAllowedToSymlink() (bool, error) {
	return false, nil
}
