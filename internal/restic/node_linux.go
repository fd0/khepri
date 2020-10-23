package restic

import (
	"path/filepath"
	"syscall"

	"golang.org/x/sys/unix"

	"github.com/restic/restic/internal/errors"

	"github.com/restic/restic/internal/fs"
)

func (node Node) restoreSymlinkTimestamps(path string, utimes [2]syscall.Timespec) error {
	dir, err := fs.Open(filepath.Dir(path))
	if err != nil {
		return errors.Wrap(err, "Open")
	}
	defer dir.Close()

	times := []unix.Timespec{
		{Sec: utimes[0].Sec, Nsec: utimes[0].Nsec},
		{Sec: utimes[1].Sec, Nsec: utimes[1].Nsec},
	}

	// fd may be uintptr(0) for some filesystems
	// skip setting
	fd := int(dir.Fd())
	if fd == 0 {
		return nil
	}

	err = unix.UtimesNanoAt(fd, filepath.Base(path), times, unix.AT_SYMLINK_NOFOLLOW)

	if err != nil {
		return errors.Wrap(err, "UtimesNanoAt")
	}

	return nil
}

func (node Node) device() int {
	return int(node.Device)
}

func (s statUnix) atim() syscall.Timespec { return s.Atim }
func (s statUnix) mtim() syscall.Timespec { return s.Mtim }
func (s statUnix) ctim() syscall.Timespec { return s.Ctim }
