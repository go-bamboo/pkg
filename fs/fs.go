package fs

import "context"

type FileStorage interface {
	UploadBytes(ctx context.Context, fileName string, data []byte) (string, error)
}
