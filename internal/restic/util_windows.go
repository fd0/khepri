package restic

import (
	"errors"
	"sync"
	"syscall"
	"unsafe"

	"github.com/restic/restic/internal/debug"
	"golang.org/x/sys/windows"
)

var windowsPrivileges struct {
	sync.Once
	isAllowedToSymlink bool
	err                error
}

// isAllowedToSymlink is true if the process is allowed to create symlinks on
// windows
func isAllowedToSymlink() (bool, error) {
	windowsPrivileges.Do(func() {

		// Lookup LUID of SeCreateSymbolicLinkPrivilege which is necessary
		// to create symlinks on windows
		symbolicLinkPrivlege, err := syscall.UTF16FromString(
			"SeCreateSymbolicLinkPrivilege",
		)

		if err != nil {
			debug.Log("UTF16FromString() failed: %s", err)
			windowsPrivileges.err = err
			return
		}

		var luidSymbolicLinkPrivilege windows.LUID
		err = windows.LookupPrivilegeValue(
			nil, &symbolicLinkPrivlege[0], &luidSymbolicLinkPrivilege,
		)

		if err != nil {
			debug.Log("LookupPrivilegeValue() failed: %s", err)
			windowsPrivileges.err = err
			return
		}

		// Get the current process token which is later used to get privilege
		// information for the current process.
		var processToken windows.Token
		err = windows.OpenProcessToken(
			windows.CurrentProcess(),
			windows.TOKEN_READ|windows.TOKEN_WRITE,
			&processToken,
		)

		if err != nil {
			windowsPrivileges.err = err
			debug.Log("OpenProcessToken() failed: %s", err)
			return
		}

		defer processToken.Close()

		// Inital call to GetTokenInformation() is only used to get the
		// necessary buffer size for fetching actual data.
		var requiredBufferSize uint32
		windows.GetTokenInformation(
			processToken, windows.TokenPrivileges,
			nil, 0, &requiredBufferSize,
		)

		if requiredBufferSize == 0 ||
			requiredBufferSize < uint32(unsafe.Sizeof(windows.Tokenprivileges{}.PrivilegeCount)) {
			windowsPrivileges.err = errors.New(
				"GetTokenInformation() failed to get buffer size",
			)
			debug.Log("%s", windowsPrivileges.err)
			return
		}

		// Allocate a buffer for the privilege token information and fetch
		// information using a second call to GetTokenInformation() to actually
		// return data
		buffer := make([]byte, requiredBufferSize)
		var bytesWritten uint32

		err = windows.GetTokenInformation(
			processToken, windows.TokenPrivileges, &buffer[0],
			uint32(len(buffer)), &bytesWritten,
		)

		if err != nil {
			windowsPrivileges.err = err
			debug.Log("GetTokenInformation() failed: %s", windowsPrivileges.err)
			return
		}

		if bytesWritten != requiredBufferSize {
			windowsPrivileges.err = errors.New(
				"GetTokenInformation() failed to returned complete data",
			)
			debug.Log("%s", windowsPrivileges.err)
			return
		}

		// Iterate over returned items and check if the LUID of one of the
		// privileges is the one we are looking for.
		tokenPrivileges := (*windows.Tokenprivileges)(unsafe.Pointer(&buffer[0]))

		for i := uint32(0); i < tokenPrivileges.PrivilegeCount; i++ {

			item := (*windows.LUIDAndAttributes)(unsafe.Pointer(
				uintptr(unsafe.Pointer(&tokenPrivileges.Privileges[0])) +
					unsafe.Sizeof(tokenPrivileges.Privileges[0])*uintptr(i)))

			if item.Luid == luidSymbolicLinkPrivilege {
				// If the correct LUID is found this process has the privileges
				// to create symbolic links.
				windowsPrivileges.isAllowedToSymlink = true
				windowsPrivileges.err = nil
				break
			}
		}

		debug.Log("isAllowedToSymlink(): %v", windowsPrivileges.isAllowedToSymlink)
	})

	return windowsPrivileges.isAllowedToSymlink, windowsPrivileges.err
}
