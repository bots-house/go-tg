package storage

import (
	"context"
	"io"
)

// Storage of files.
type Storage interface {
	// Add file to storage.
	// Assign it random name.
	Add(ctx context.Context, dir string, body io.Reader, ext string) (string, error)

	// Add file to storage from remote URL
	AddByURL(ctx context.Context, dir string, url string) (string, error)

	// Delete file from storage by path
	// Delete(ctx context.Context, path string) error

	// Returns full path to file
	PublicURL(path string) string
}
