package b2_test

import (
	"bytes"
	"math/rand"
	"os"
	"time"

	"golang.org/x/net/context"

	"restic"
	"restic/backend/b2"
	"restic/backend/test"
	"restic/errors"
)

//go:generate go run ../test/generate_backend_tests.go

func init() {
	if os.Getenv("B2_ACCOUNT_ID") == "" || os.Getenv("B2_ACCOUNT_KEY") == "" {
		SkipMessage = "No B2 credentials found. Skipping B2 backend tests."
		return
	}

	cfg := b2.Config{
		AccountID: os.Getenv("B2_ACCOUNT_ID"),
		Key:       os.Getenv("B2_ACCOUNT_KEY"),
		Bucket:    generateBucketName(),
		Prefix:    "test",
	}

	test.CreateFn = func() (restic.Backend, error) {
		be, err := b2.Open(cfg)
		if err != nil {
			return nil, err
		}

		exists, err := be.Test(restic.ConfigFile, "")
		if err != nil {
			return nil, err
		}

		if exists {
			return nil, errors.New("config already exists")
		}

		return be, nil
	}

	test.OpenFn = func() (restic.Backend, error) {
		return b2.Open(cfg)
	}

	// Remove the test bucket as clean up.
	test.CleanupFn = func() error {
		be, err := b2.Open(cfg)
		if err != nil {
			return err
		}

		ctx := context.Background()
		b2be := be.(*b2.B2)
		return b2be.Bucket().Delete(ctx)
	}
}

// Generates a random bucket name starting with "restic-test-".
func generateBucketName() string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	const lenChars = len(chars)
	const lenBucket = 16
	var bucket bytes.Buffer
	bucket.WriteString("restic-test-")
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	for i := 0; i < lenBucket; i++ {
		bucket.WriteByte(chars[r.Intn(lenChars)])
	}
	return bucket.String()
}
