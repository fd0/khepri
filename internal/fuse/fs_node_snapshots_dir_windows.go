package fuse

import (
	"github.com/billziss-gh/cgofuse/fuse"
	"github.com/restic/restic/internal/debug"
	"golang.org/x/net/context"
)

// snapshotDirLatestName is a specially handled alias for latest snapshot
const snapshotDirLatestName = "latest"

// FsNodeSnapshotsDir is a virtual directory listing all available snapshots
// in the repository.
type FsNodeSnapshotsDir struct {
	ctx   context.Context
	root  *FsNodeRoot
	nodes map[string]*FsNodeSnapshotDir
}

var _ = FsNode(&FsNodeSnapshotsDir{})

// NewSnapshotsDir creates a new virtual directory,
func NewSnapshotsDir(ctx context.Context, root *FsNodeRoot) *FsNodeSnapshotsDir {
	return &FsNodeSnapshotsDir{
		ctx: ctx, root: root, nodes: make(map[string]*FsNodeSnapshotDir),
	}
}

// Readdir lists all items in the specified path. Results are returned
// through the given callback function.
func (self *FsNodeSnapshotsDir) Readdir(path []string, fill FsListItemCallback) {

	debug.Log("Readdir(%v)", path)

	if len(path) == 0 {

		fill(".", nil, 0)
		fill("..", nil, 0)

		self.root.snapshotManager.updateSnapshots()

		if self.root.snapshotManager.snapshotNameLatest != "" {
			fill(snapshotDirLatestName, &defaultDirectoryStat, 0)
		}

		for name, _ := range self.root.snapshotManager.snapshotByName {
			fill(name, &defaultDirectoryStat, 0)
		}
	} else {

		head := path[0]

		debug.Log("handle subtree %v", head)

		if head == snapshotDirLatestName {
			head = self.root.snapshotManager.snapshotNameLatest
		}

		if snapshot, ok := self.root.snapshotManager.snapshotByName[head]; ok {

			if _, contained := self.nodes[head]; !contained {

				node, err := NewFsNodeSnapshotDirFromSnapshot(self.ctx, self.root, snapshot)

				if err == nil {
					self.nodes[head] = node
				} else {
					debug.Log("Failed to create node for %v: %v", head, err.Error())
				}
			}

			if node, contained := self.nodes[head]; contained {
				node.Readdir(path[1:], fill)
			} else {
				debug.Log("ListDirectories error for %v", head)
			}

		} else {
			debug.Log("Snapshot not found: %v", head)
		}
	}
}

// GetAttributes fetches the attributes of the specified file or directory.
func (self *FsNodeSnapshotsDir) GetAttributes(path []string, stat *fuse.Stat_t) bool {

	debug.Log("GetAttributes(%v)", path)

	pathLength := len(path)

	if pathLength < 1 {
		*stat = defaultDirectoryStat
		return true
	} else {

		head := path[0]

		if pathLength == 1 {
			if head == snapshotDirLatestName && self.root.snapshotManager.snapshotNameLatest != "" {
				*stat = defaultDirectoryStat
				return true
			}

			if _, found := self.root.snapshotManager.snapshotByName[head]; found {
				*stat = defaultDirectoryStat
				return true
			}
		} else {

			if head == snapshotDirLatestName {
				head = self.root.snapshotManager.snapshotNameLatest
			}

			if snapshotDir, ok := self.nodes[head]; ok {
				return snapshotDir.GetAttributes(path[1:], stat)
			} else {
				return false
			}
		}
	}

	return false
}

// Open opens the file for the given path.
func (self *FsNodeSnapshotsDir) Open(path []string, flags int) (errc int, fh uint64) {

	pathLength := len(path)

	if pathLength < 1 {
		return -fuse.EISDIR, ^uint64(0)
	} else {

		head := path[0]

		if pathLength == 1 {
			if head == snapshotDirLatestName && self.root.snapshotManager.snapshotNameLatest != "" {
				return -fuse.EISDIR, ^uint64(0)
			}

			if _, found := self.root.snapshotManager.snapshotByName[head]; found {
				return -fuse.EISDIR, ^uint64(0)
			}
		} else {
			if head == snapshotDirLatestName {
				head = self.root.snapshotManager.snapshotNameLatest
			}

			if snapshotDir, ok := self.nodes[head]; ok {
				return snapshotDir.Open(path[1:], flags)
			}
		}
	}

	return -fuse.ENOENT, ^uint64(0)
}

// Read reads data to the given buffer from the specified file.
func (self *FsNodeSnapshotsDir) Read(path []string, buff []byte, ofst int64, fh uint64) (n int) {

	pathLength := len(path)

	if pathLength < 1 {
		return -fuse.EISDIR
	} else {

		head := path[0]

		if pathLength == 1 {
			if head == snapshotDirLatestName && self.root.snapshotManager.snapshotNameLatest != "" {
				return -fuse.EISDIR
			}

			if _, found := self.root.snapshotManager.snapshotByName[head]; found {
				return -fuse.EISDIR
			}
		} else {
			if head == snapshotDirLatestName {
				head = self.root.snapshotManager.snapshotNameLatest
			}

			if snapshotDir, ok := self.nodes[head]; ok {
				return snapshotDir.Read(path[1:], buff, ofst, fh)
			}
		}
	}

	return -fuse.ENOENT
}
