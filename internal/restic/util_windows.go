package restic

import (
	"github.com/restic/restic/internal/debug"
	"golang.org/x/sys/windows"
)

// IsRunningAsAdminOnWindows is true if the process is running with admin privileges
var IsRunningAsAdminOnWindows = func() bool {
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
		return false
	}

	defer windows.FreeSid(sid)

	token := windows.Token(0)
	isMemberOfAdminGroup, err := token.IsMember(sid)

	if err != nil {
		debug.Log("token.IsMember() failed: %s", err)
		return false
	}

	debug.Log("isRunningAsAdmin(): %v", isMemberOfAdminGroup)

	return isMemberOfAdminGroup
}()
