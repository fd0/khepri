package dryrun

import (
	"bufio"
	"context"
	"io"

	"github.com/restic/restic/internal/debug"
	"github.com/restic/restic/internal/restic"
)

// Backend passes reads through to an underlying layer and accepts writes, but
// doesn't do anything. Also removes are ignored.
// So in fact, this backend silently ignores all operations that would modify
// the repo and does normal operations else.
// This is used for `backup --dry-run`.
type Backend struct {
	restic.Backend
}

// New returns a new backend that saves all data in a map in memory.
func New(be restic.Backend) *Backend {
	b := &Backend{Backend: be}
	debug.Log("created new dry backend")
	return b
}

// Save adds new Data to the backend.
func (be *Backend) Save(ctx context.Context, h restic.Handle, rd restic.RewindReader) error {
	if err := h.Valid(); err != nil {
		return err
	}

	// discard everything from rd, but count it
	maxInt := int(^uint(0) >> 1)
	bufRd := bufio.NewReader(rd)
	length, err := bufRd.Discard(maxInt)
	if err != io.EOF {
		return err
	}

	debug.Log("faked saving %v bytes at %v", length, h)

	return nil
}

// Remove deletes a file from the backend.
func (be *Backend) Remove(ctx context.Context, h restic.Handle) error {
	return nil
}

// Location returns the location of the backend.
func (be *Backend) Location() string {
	return "DRY:" + be.Backend.Location()
}

// Delete removes all data in the backend.
func (be *Backend) Delete(ctx context.Context) error {
	return nil
}
