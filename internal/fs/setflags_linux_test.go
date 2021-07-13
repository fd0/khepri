package fs

import (
	"io"
	"os"
	"testing"
	"time"

	rtest "github.com/restic/restic/internal/test"

	"golang.org/x/sys/unix"
)

func TestNoatime(t *testing.T) {
	tmp, cleanup := TestTempFile(t, "restic-test-noatime")
	defer cleanup()
	f := tmp.(*os.File)

	// Only run this test on common filesystems that support O_NOATIME.
	// On others, we may not get an error.
	if !supportsNoatime(t, f) {
		t.Log("temp directory may not support O_NOATIME, skipping")
		t.Skip()
	}
	// From this point on, we own the file, so we should not get EPERM.

	_, err := io.WriteString(f, "Hello!")
	rtest.OK(t, err)
	_, err = f.Seek(0, io.SeekStart)
	rtest.OK(t, err)

	getAtime := func() time.Time {
		info, err := f.Stat()
		rtest.OK(t, err)
		return ExtendedStat(info).AccessTime
	}

	atime := getAtime()

	err = setFlags(f)
	rtest.OK(t, err)

	_, err = f.Read(make([]byte, 1))
	rtest.OK(t, err)
	rtest.Equals(t, atime, getAtime())
}

func supportsNoatime(t *testing.T, f *os.File) bool {
	var fsinfo unix.Statfs_t
	err := unix.Fstatfs(int(f.Fd()), &fsinfo)
	rtest.OK(t, err)

	return fsinfo.Type == unix.BTRFS_SUPER_MAGIC ||
		fsinfo.Type == unix.EXT2_SUPER_MAGIC ||
		fsinfo.Type == unix.EXT3_SUPER_MAGIC ||
		fsinfo.Type == unix.EXT4_SUPER_MAGIC ||
		fsinfo.Type == unix.TMPFS_MAGIC
}
