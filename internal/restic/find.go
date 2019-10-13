package restic

import "context"

// FindUsedBlobs traverses the tree ID and adds all seen blobs (trees and data
// blobs) to the set blobs. The tree blobs in the `seen` BlobSet will not be visited
// again.
func FindUsedBlobs(ctx context.Context, repo Repository, treeID ID, blobs BlobSet, seen BlobSet) error {
	blobs.Insert(NewBlobHandle(treeID, TreeBlob))

	tree, err := repo.LoadTree(ctx, treeID)
	if err != nil {
		return err
	}

	for _, node := range tree.Nodes {
		switch node.Type {
		case "file":
			for _, blob := range node.Content {
				blobs.Insert(NewBlobHandle(blob, DataBlob))
			}
		case "dir":
			subtreeID := *node.Subtree
			h := NewBlobHandle(subtreeID, TreeBlob)
			if seen.Has(h) {
				continue
			}

			seen.Insert(h)

			err := FindUsedBlobs(ctx, repo, subtreeID, blobs, seen)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
