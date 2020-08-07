// +build !windows

package restic

// isWindowsAdmin is true if the process is running with admin privileges
func isWindowsAdmin() (bool, error) {
	return false, nil
}
