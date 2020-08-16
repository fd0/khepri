package restic

import (
	"sync"

	"github.com/restic/restic/internal/debug"
	"golang.org/x/sys/windows"
)

var windowsAdmin struct {
	sync.Once
	isAdmin bool
	err     error
}

// isWindowsAdmin is true if the process is running with admin privileges
func isWindowsAdmin() (bool, error) {
	windowsAdmin.Do(func() {

		var sid *windows.SID

		err := windows.AllocateAndInitializeSid(
			&windows.SECURITY_NT_AUTHORITY,
			2,
			windows.SECURITY_BUILTIN_DOMAIN_RID,
			windows.DOMAIN_ALIAS_RID_ADMINS,
			0, 0, 0, 0, 0, 0,
			&sid,
		)

		if err != nil {
			debug.Log("AllocateAndInitializeSid() failed: %s", err)
			windowsAdmin.err = err
			return
		}

		defer windows.FreeSid(sid)

		token := windows.Token(0)
		windowsAdmin.isAdmin, err = token.IsMember(sid)

		if err != nil {
			debug.Log("token.IsMember() failed: %s", err)
			windowsAdmin.err = err
			return
		}

		debug.Log("isWindowsAdmin(): %v", windowsAdmin.isAdmin)
	})

	return windowsAdmin.isAdmin, windowsAdmin.err
}
