package core

import (
	"context"
	"time"
)

// LotFileID it's unique lot file id.
type LotFileID int

type LotFile struct {
	ID LotFileID

	// Reference to Lot.
	LotID LotID

	// Name of uploaded file.
	Name string

	// Size of uploaded file in bytes.
	Size int

	// MIMEType of uploaded file.
	MIMEType string

	// Path of uploaded file.
	Path string

	// Time when file was uploaded.
	CreatedAt time.Time
}

type LotFileSlice []*LotFile

func (lfs LotFileSlice) FindByLotID(id LotID) LotFileSlice {
	result := make([]*LotFile, 0, len(lfs))
	for _, lf := range lfs {
		if lf.LotID == id {
			result = append(result, lf)
		}
	}
	return result
}

func NewLotFile(
	name string,
	size int,
	mimeType string,
	path string,
) *LotFile {
	return &LotFile{
		Name:      name,
		Size:      size,
		MIMEType:  mimeType,
		Path:      path,
		CreatedAt: time.Now(),
	}
}

var ErrLotFileNotFound = NewError("lot_file_not_found", "lot file not found")

type LotFileStoreQuery interface {
	ID(ids ...LotFileID) LotFileStoreQuery
	LotID(ids ...LotID) LotFileStoreQuery
	All(ctx context.Context) (LotFileSlice, error)
}

type LotFileStore interface {
	Add(ctx context.Context, lf *LotFile) error
	Update(ctx context.Context, lf *LotFile) error
	Query() LotFileStoreQuery
}
